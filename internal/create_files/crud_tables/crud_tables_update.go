package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFiles_UpdateEveryColumn - создаёт 1 файл в папке crud
func CreateFiles_UpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCrudFunc := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_UPDATE_FUNC_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCrudUpdateFunc := DirReadyTable + config.Settings.PREFIX_CRUD + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateCrudFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunc, " error: ", err)
	}
	TextCrudUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_UPDATE_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)
	TextCrud = TextCrud + "\n"

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)
	TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextCrudUpdateFunc, Table1)

	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	TextCrud = TextCrud + TextUpdateEveryColumn
	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, DBConstantsURL)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		TextCrud = create_files.CheckAndAdd_ImportGorm_FromText(TextCrud)
		//TextCrud = create_files.Convert_RequestIdToAlias(TextCrud, Table1)
		//добавим импорт uuid
	}

	//кэш
	if config.Settings.NEED_CREATE_CACHE_API == true {
		const TEXT_CACHE_REMOVE = "cache.Remove("
		TextCrud = strings.ReplaceAll(TextCrud, `//`+TEXT_CACHE_REMOVE, TEXT_CACHE_REMOVE)
	}

	//переименование функций
	//TextCrud = RenameFunctions(TextCrud, Table1)

	//заменяет "m.ID" на название колонки PrimaryKey
	TextCrud = Replace_PrimaryKeyOtvetID(TextCrud, Table1)

	//добавим импорт uuid
	TextCrud = create_files.CheckAndAdd_ImportUUID_FromText(TextCrud)

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//импорт "fmt"
	TextCrud = create_files.CheckAndAdd_ImportFmt(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.Delete_EmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudUpdateFunc, []byte(TextCrud), config.Settings.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextCrudUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextUpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumn1(TextCrudUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextCrudUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextCrudUpdateFunc

	ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	TextRequest, TextRequestFieldName := create_files.FindText_ProtobufRequest(Table1)

	//ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)

	Otvet = ReplaceCacheRemove(Otvet, Table1)

	Otvet = Replace_PrimaryKeyOtvetID(Otvet, Table1)

	//запись null в nullable колонки
	if Column1.IsNullable == true && (Column1.TableKey != "" || Column1.TypeGo == "time.Time") {
	} else {
		TextFind := `	if Value == 0 {
		tx = db.Model(&m).Update("ColumnNameField", gorm.Expr("NULL"))
	} else {
		tx = db.Model(&m).Update("ColumnNameField", Value)
	}`
		TextReplace := `	tx = db.Model(&m).Update("ColumnNameField", Value)`
		Otvet = strings.ReplaceAll(Otvet, TextFind, TextReplace)
	}

	//заменяем Read_ctx()
	Otvet = strings.ReplaceAll(Otvet, " Read_ctx ", " "+FuncName+"_ctx ")
	Otvet = strings.ReplaceAll(Otvet, " Read_ctx(", " "+FuncName+"_ctx(")
	Otvet = strings.ReplaceAll(Otvet, ".Read_ctx(", "."+FuncName+"_ctx(")

	//заменяем Read()
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Read", ModelName+"_"+FuncName)
	Otvet = strings.ReplaceAll(Otvet, " Read ", " "+FuncName+" ")
	Otvet = strings.ReplaceAll(Otvet, " Read(", " "+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, `"Read()`, `"`+FuncName+"()")

	Otvet = create_files.Replace_TemplateModel_to_Model(Otvet, Table1.NameGo)
	Otvet = create_files.Replace_TemplateTableName_to_TableName(Otvet, Table1.Name)

	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNamePK", ColumnPK.NameGo)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNameField", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Model.ID", "Model."+ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)

	//внешние ключи заменяем 0 на null
	TextEqualEmpty := create_files.FindText_EqualEmpty(Column1, "Value")
	Otvet = strings.ReplaceAll(Otvet, "Value == 0", TextEqualEmpty)

	return Otvet
}

// CreateFiles_UpdateEveryColumn_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_UpdateEveryColumn_Test(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCrudFunc := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_UPDATE_FUNC_TEST_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCrudUpdate := DirReadyTable + config.Settings.PREFIX_CRUD + TableName + "_update_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateCrudFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunc, " error: ", err)
	}
	TextCrudUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_UPDATE_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)
	TextCrud = TextCrud + "\n"

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		TextCrud = Replace_Postgres_ID_Update_Test(TextCrud, Table1)

		//TextCrud = create_files.Replace_PrimaryKeyM_ID(TextCrud, Table1)
		//TextCrud = create_files.Replace_PrimaryKeyOtvetID(TextCrud, Table1)

		ConstantsURL := create_files.Find_ConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, ConstantsURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextCrud = create_files.Replace_Connect_WithApplicationName(TextCrud)

	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumnTest(TextCrudUpdateFunc, Table1)
	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	TextCrud = TextCrud + TextUpdateEveryColumn

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)
	TextCrud = create_files.Replace_TemplateModel_to_Model(TextCrud, Table1.NameGo)
	TextCrud = create_files.Replace_TemplateTableName_to_TableName(TextCrud, Table1.Name)

	//
	TextCrud = create_files.CheckAndAdd_ImportFmt(TextCrud)

	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//SkipNow() если нет строк в БД
	TextCrud = create_files.AddSkipNow(TextCrud, Table1)

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.Delete_EmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudUpdate, []byte(TextCrud), config.Settings.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumnTest - возвращает текст для всех таблиц
func FindTextUpdateEveryColumnTest(TextCrudUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextUpdateEveryColumnTest() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumnTest1(TextCrudUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumnTest1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumnTest1(TextCrudUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextCrudUpdateFunc

	//ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	TextRequest, TextRequestFieldName := create_files.FindText_ProtobufRequest(Table1)
	DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

	//Postgres_ID_Test = ID Minimum
	Otvet = Replace_Postgres_ID_Update_Test(Otvet, Table1)

	//Otvet = create_files.Replace_PrimaryKeyM_ID(Otvet, Table1)
	//Otvet = create_files.Replace_PrimaryKeyOtvetID(Otvet, Table1)

	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	Otvet = strings.ReplaceAll(Otvet, "TestUpdate(", "Test"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, ".Update(", "."+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, " DefaultValue", " "+DefaultValue)

	return Otvet
}

// Replace_PrimaryKeyOtvetID - заменяет "Otvet.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyOtvetID(Text string, Table1 *types.Table) string {
	Otvet := Text

	VariableName := "m"

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"

		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)
	////Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDID)
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = ProtoFromInt(m.ID)", TextRequestIDmID)
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = int64(ID)", TextRequestIDInt64ID)
	//Otvet = strings.ReplaceAll(Otvet, "\tOtvet.ID = "+VariableName+".ID\n", TextOtvetIDmID)
	Otvet = strings.ReplaceAll(Otvet, " IntFromAlias("+VariableName+".ID) == 0", TextMID0)
	//Otvet = strings.ReplaceAll(Otvet, "\tm2.ID = int64(m.ID)", TextM2ID)
	//Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", TextIfMId)
	//Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", TextIfMIdNot0)

	////заменим ID := Request.ID
	//Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextIDRequestID)

	return Otvet
}

// Replace_Postgres_ID_Update_Test - заменяет текст "const Postgres_ID_Test = 0" на нужные ИД, для много колонок PrimaryKey
func Replace_Postgres_ID_Update_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	if len(MassPK) == 0 {
		return Otvet
	}

	//заменим Otvet.ID = Postgres_ID_Test
	TextFind := "\tOtvet.ID = Postgres_ID_Test\n"
	TextNew := ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := create_files.FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "Otvet."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим m.ID = Postgres_ID_Test
	TextFind = "\tm.ID = Postgres_ID_Test\n"
	TextNew = ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := create_files.FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "m."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}
