package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFilesUpdateEveryColumn - создаёт 1 файл в папке crud
func CreateFilesUpdateEveryColumn(Table1 *types.Table) error {
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
		TextCrud = strings.ReplaceAll(TextCrud, `//`+constants.TEXT_CACHE_REMOVE, constants.TEXT_CACHE_REMOVE)
	}

	//переименование функций
	//TextCrud = RenameFunctions(TextCrud, Table1)

	//заменяет "m.ID" на название колонки PrimaryKey
	TextCrud = create_files.Replace_PrimaryKeyM_ID(TextCrud, Table1)

	//добавим импорт uuid
	TextCrud = create_files.CheckAndAdd_ImportUUID_FromText(TextCrud)

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//импорт "fmt"
	TextCrud = create_files.CheckAndAdd_ImportFmt(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.Delete_EmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudUpdateFunc, []byte(TextCrud), constants.FILE_PERMISSIONS)

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

	Otvet = create_files.Replace_PrimaryKeyOtvetID_ManyPK1(Otvet, Table1, "m")

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

	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNamePK", ColumnPK.NameGo)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNameField", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Model.ID", "Model."+ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	//Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	//TextIntFromAlias := create_files.ConvertFromAlias(Table1, Column1, "m")
	//DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

	//TextEqual0 := create_files.FindText_Equal0(Column1)
	//Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(m.ID) == 0", TextIntFromAlias+TextEqual0)
	//Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(m.ID)", TextIntFromAlias)

	//внешние ключи заменяем 0 на null
	TextEqualEmpty := create_files.FindText_EqualEmpty(Column1, "Value")
	Otvet = strings.ReplaceAll(Otvet, "Value == 0", TextEqualEmpty)
	//if Column1.IsNullable == true && (Column1.TableKey != "" || Column1.TypeGo == "time.Time") {
	//	Otvet = strings.ReplaceAll(Otvet, "0==1 && ", "")
	//}

	return Otvet
}

// CreateFilesUpdateEveryColumnTest - создаёт 1 файл в папке grpc_client
func CreateFilesUpdateEveryColumnTest(Table1 *types.Table) error {
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

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)
	TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		//Postgres_ID_Test = ID Minimum
		TextCrud = create_files.Replace_Postgres_ID_Test(TextCrud, Table1)

		TextCrud = create_files.Replace_PrimaryKeyM_ID(TextCrud, Table1)
		TextCrud = create_files.Replace_PrimaryKeyOtvetID(TextCrud, Table1)
		//TextCrud = create_files.Convert_RequestIdToAlias(TextCrud, Table1)

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

	TextCrud = create_files.CheckAndAdd_ImportFmt(TextCrud)

	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//SkipNow() если нет строк в БД
	TextCrud = create_files.AddSkipNow(TextCrud, Table1)

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.Delete_EmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudUpdate, []byte(TextCrud), constants.FILE_PERMISSIONS)

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
	Otvet = create_files.Replace_Postgres_ID_Test(Otvet, Table1)

	Otvet = create_files.Replace_PrimaryKeyM_ID(Otvet, Table1)
	Otvet = create_files.Replace_PrimaryKeyOtvetID(Otvet, Table1)

	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	//Otvet = strings.ReplaceAll(Otvet, "Otvet.Name", "Otvet."+ColumnName)
	//Otvet = strings.ReplaceAll(Otvet, "Postgres_ID_Test", DefaultValue)
	Otvet = strings.ReplaceAll(Otvet, "TestUpdate(", "Test"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, ".Update(", "."+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, " DefaultValue", " "+DefaultValue)

	return Otvet
}
