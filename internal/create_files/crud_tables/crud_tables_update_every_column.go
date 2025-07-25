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
	"strconv"
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

	FilenameTemplateCrudFunc := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_UPDATE_MANY_FUNC_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCrudUpdateFunc := DirReadyTable + config.Settings.PREFIX_CRUD + TableName + "_update_many.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateCrudFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunc, " error: ", err)
	}
	TextCrudUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_UPDATE_MANY_FILENAME
	bytes, err = micro.ReadFile_Linux_Windows(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)
	TextCrud = TextCrud + "\n"

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextCrud = CreateFiles_UpdateEveryColumn1(TextCrud, Table1)

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextCrudUpdateFunc, Table1)

	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	TextCrud = TextCrud + TextUpdateEveryColumn

	TextCrud = create_files.Replace_TemplateModel_to_Model(TextCrud, Table1.NameGo)
	TextCrud = create_files.Replace_TemplateTableName_to_TableName(TextCrud, Table1.Name)
	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, DBConstantsURL)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		//TextCrud = create_files.CheckAndAdd_ImportGorm_FromText(TextCrud)
		//TextCrud = create_files.Convert_RequestIdToAlias(TextCrud, Table1)
		//добавим импорт uuid

		//postgres_func
		TextCrud = create_files.CheckAndAdd_Import(TextCrud, "github.com/ManyakRus/starter/postgres_func")

	}

	//кэш
	if config.Settings.NEED_CREATE_CACHE_API == true {
		const TEXT_CACHE_REMOVE = "cache.Remove("
		TextCrud = strings.ReplaceAll(TextCrud, `//`+TEXT_CACHE_REMOVE, TEXT_CACHE_REMOVE)
	}

	//переименование функций
	//TextCrud = RenameFunctions(TextCrud, Table1)

	//заменяет "m.ID" на название колонки PrimaryKey
	//TextCrud = Replace_PrimaryKeyOtvetID(TextCrud, Table1)

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

// CreateFiles_UpdateEveryColumn1 - заполняет Text
func CreateFiles_UpdateEveryColumn1(Text string, Table1 *types.Table) string {
	Otvet := Text

	//Primary key
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	ReplacePKFieldsWithComma := ""
	ReplaceID0 := ""
	Comma := ""
	TextAnd := ""
	for _, Column1 := range ColumnsPK {
		ReplacePKFieldsWithComma = ReplacePKFieldsWithComma + Comma + "m." + Column1.NameGo
		TextEmpty := create_files.FindText_EqualEmpty(Column1, "m."+Column1.NameGo)
		ReplaceID0 = ReplaceID0 + TextAnd + TextEmpty
		Comma = ", "
		TextAnd = " && "
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldsWithComma", ReplacePKFieldsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceID0", ReplaceID0)

	Otvet = ReplaceCacheRemove(Otvet, Table1)

	return Otvet

}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextCrudUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//найдём новый текст для каждой таблицы
	MassColumns := micro.MassFrom_Map(Table1.MapColumns)
	for _, Column1 := range MassColumns {
		//кроме ненужных колонок
		if create_files.Is_Need_Сolumn(Column1) == false {
			continue
		}
		//if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
		//	continue
		//}

		TextColumn1 := FindTextUpdateEveryColumn1(TextCrudUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextCrudUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextCrudUpdateFunc

	//Primary key
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	//TableAlias := create_files.Find_TableAlias(Table1)
	ReplacePKFieldsWithComma := ""
	ReplacePKFieldNamesFormat := ""
	ReplaceID0 := ""
	ReplaceWhereID := ""
	Comma := ""
	TextAnd := ""
	for i, Column1 := range ColumnsPK {
		//sNumber := strconv.Itoa(i + 1)
		ReplacePKFieldsWithComma = ReplacePKFieldsWithComma + Comma + "m." + Column1.NameGo
		ReplacePKFieldNamesFormat = ReplacePKFieldNamesFormat + Comma + Column1.NameGo + ": %v"
		TextEmpty := create_files.FindText_EqualEmpty(Column1, "m."+Column1.NameGo)
		ReplaceID0 = ReplaceID0 + TextAnd + TextEmpty
		sNumber := strconv.Itoa(i + 2) // 1 - уже занят для Column1
		ReplaceWhereID = ReplaceWhereID + "\tand " + `"` + Column1.Name + `"` + " = $" + sNumber + "\n"
		//ReplaceWhereID = ReplaceWhereID + "\tand " + Column1.Name + " = $" + sNumber + "\n"

		Comma = ", "
		TextAnd = " && "
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldsWithComma", ReplacePKFieldsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldNamesFormat", ReplacePKFieldNamesFormat)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceID0", ReplaceID0)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceWhereID", ReplaceWhereID)

	//
	TextAlias := create_files.ConvertFromAlias(Table1, Column1, "m."+Column1.NameGo)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceValueFromAlias(m.ID)", TextAlias)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceFieldName", Column1.NameGo)
	//Otvet = strings.ReplaceAll(Otvet, "ReplaceTableName", Column1.Name)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTableName", Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceColumnNameEqualDollarComma", `"`+Column1.Name+`"`+" = $1")

	ReplaceValueEqual := "Value := m." + Column1.NameGo
	if Column1.IsNullable == true {
		TextValue := create_files.FindText_NullValue(Column1.TypeGo, "m."+Column1.NameGo)
		ReplaceValueEqual = "Value := " + TextValue
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceValueEqual", ReplaceValueEqual)

	Otvet = ReplaceCacheRemove(Otvet, Table1)

	//ModelName := Table1.NameGo
	//ColumnName := Column1.NameGo
	//FuncName := "Update_" + ColumnName
	//TextRequest, TextRequestFieldName := create_files.FindText_ProtobufRequest(Table1)
	//
	////ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
	//
	//Otvet = ReplaceCacheRemove(Otvet, Table1)
	//
	//Otvet = Replace_PrimaryKeyOtvetID(Otvet, Table1)
	//
	////запись null в nullable колонки
	//if Column1.IsNullable == true{
	//} else {
	//}
	//
	////заменяем Read_ctx()
	//Otvet = strings.ReplaceAll(Otvet, " Read_ctx ", " "+FuncName+"_ctx ")
	//Otvet = strings.ReplaceAll(Otvet, " Read_ctx(", " "+FuncName+"_ctx(")
	//Otvet = strings.ReplaceAll(Otvet, ".Read_ctx(", "."+FuncName+"_ctx(")
	//
	////заменяем Read()
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Read", ModelName+"_"+FuncName)
	//Otvet = strings.ReplaceAll(Otvet, " Read ", " "+FuncName+" ")
	//Otvet = strings.ReplaceAll(Otvet, " Read(", " "+FuncName+"(")
	//Otvet = strings.ReplaceAll(Otvet, `"Read()`, `"`+FuncName+"()")
	//
	//Otvet = create_files.Replace_TemplateModel_to_Model(Otvet, Table1.NameGo)
	//Otvet = create_files.Replace_TemplateTableName_to_TableName(Otvet, Table1.Name)
	//
	//TextProto := create_files.TextProto()
	//Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", TextProto+"."+TextRequest)
	//Otvet = strings.ReplaceAll(Otvet, "ColumnNamePK", ColumnPK.NameGo)
	//Otvet = strings.ReplaceAll(Otvet, "ColumnNameField", ColumnName)
	//Otvet = strings.ReplaceAll(Otvet, "Model.ID", "Model."+ColumnName)
	//Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	//
	////внешние ключи заменяем 0 на null
	//TextEqualEmpty := create_files.FindText_EqualEmpty(Column1, "Value")
	//Otvet = strings.ReplaceAll(Otvet, "Value == 0", TextEqualEmpty)

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

	FilenameTemplateCrudFunc := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_UPDATE_MANY_FUNC_TEST_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCrudUpdate := DirReadyTable + config.Settings.PREFIX_CRUD + TableName + "_update_many_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateCrudFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunc, " error: ", err)
	}
	TextCrudUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_UPDATE_MANY_TEST_FILENAME
	bytes, err = micro.ReadFile_Linux_Windows(FilenameTemplateCrud)
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

		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextCrud = create_files.AddImport(TextCrud, CrudFuncURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextCrud = create_files.Replace_Connect_WithApplicationName(TextCrud)

	}

	//ReplacePKColumnName
	ColumnsPK := create_files.Find_PrimaryKeyColumn(Table1)
	TextCrud = strings.ReplaceAll(TextCrud, "ReplacePKColumnName", ColumnsPK.NameGo)

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
		if create_files.Is_Need_Сolumn(Column1) == false {
			continue
		}
		//if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
		//	continue
		//}

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

	TextProto := create_files.TextProto()
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", TextProto+"."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceFieldName", ColumnName)
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
		TextM := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
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

// Replace_Postgres_ID_Update - заменяет текст "const Postgres_ID = 0" на нужные ИД, для много колонок PrimaryKey
