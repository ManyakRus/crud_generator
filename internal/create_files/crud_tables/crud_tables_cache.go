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
	"strings"
)

// CreateFilesCache - создаёт 1 файл "*_cache.go" в папке crud
func CreateFilesCache(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesCrud + constants.CRUD_TABLES_CACHE_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCache := DirReadyTable + "crud_" + TableName + "_cache.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextCache := string(bytes)

	//заменим имя пакета на новое
	TextCache = create_files.ReplacePackageName(TextCache, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCache = create_files.DeleteTemplateRepositoryImports(TextCache)

		DBConstantsURL := create_files.FindDBConstantsURL()
		TextCache = create_files.AddImport(TextCache, DBConstantsURL)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextCache = create_files.AddImport(TextCache, ModelTableURL)

		//TextCache = create_files.ConvertRequestIdToAlias(TextCache, Table1)
	}

	//alias
	TextCache = create_files.ReplaceIDToAlias_OtvetID(TextCache, Table1)

	//const CACHE_SIZE = 1000
	CACHE_ELEMENTS_COUNT_MAX := config.Settings.CACHE_ELEMENTS_COUNT_MAX
	Count_Now := Table1.RowsCount
	CACHE_ELEMENTS_COUNT := micro.MinInt64(Count_Now, CACHE_ELEMENTS_COUNT_MAX)
	sCACHE_ELEMENTS_COUNT := micro.StringFromInt64(CACHE_ELEMENTS_COUNT)
	TextCache = create_files.FillVariable(TextCache, constants.TEXT_CACHE_SIZE_1000, sCACHE_ELEMENTS_COUNT)

	ColumnPK := create_files.FindPrimaryKeyColumn(Table1)

	//тип ID кэша
	if Table1.PrimaryKeyColumnsCount == 1 {
		_, ColumnTypeGo := create_files.FindPrimaryKeyNameTypeGo(Table1)
		TextCache = strings.ReplaceAll(TextCache, "LRU[int64", "LRU["+ColumnTypeGo)
		TextCache = strings.ReplaceAll(TextCache, "ID int64", ColumnPK.NameGo+" "+ColumnTypeGo)
		TextCache = create_files.ReplacePrimaryKeyOtvetID(TextCache, Table1)
		TextCache = strings.ReplaceAll(TextCache, "int64(ID)", ColumnPK.NameGo)
		TextCache = strings.ReplaceAll(TextCache, ", ID)", ", "+ColumnPK.NameGo+")")
	} else {
		TextCache = strings.ReplaceAll(TextCache, "LRU[int64", "LRU[string")
		TextCache = create_files.ReplacePrimaryKeyOtvetID_Many(TextCache, Table1)
		TextIDMany := "(ID)"
		TextIDMany = create_files.ReplaceIDtoID_Many(TextIDMany, Table1)
		TextCache = strings.ReplaceAll(TextCache, "int64(ID)", "("+Table1.Name+".StringIdentifier"+TextIDMany+")")
		TextCache = create_files.ReplaceIDtoID_Many(TextCache, Table1)
	}

	//uuid
	TextCache = create_files.CheckAndAddImportUUID_FromText(TextCache)

	//замена слов
	ModelName := Table1.NameGo
	TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextCache = config.Settings.TEXT_MODULE_GENERATED + TextCache

	//удаление пустого импорта
	TextCache = create_files.DeleteEmptyImport(TextCache)

	//удаление пустых строк
	TextCache = create_files.DeleteEmptyLines(TextCache)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextCache), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesCacheTest - создаёт 1 файл "*_cache_test.go" в папке crud
func CreateFilesCacheTest(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesCrud + constants.CRUD_TABLES_CACHE_TEST_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCache := DirReadyTable + "crud_" + TableName + "_cache_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextCache := string(bytes)

	//заменим имя пакета на новое
	TextCache = create_files.ReplacePackageName(TextCache, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCache = create_files.DeleteTemplateRepositoryImports(TextCache)

		//DBConstantsURL := create_files.FindDBConstantsURL()
		//TextCache = create_files.AddImport(TextCache, DBConstantsURL)
		//
		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextCache = create_files.AddImport(TextCache, ModelTableURL)

		//TextCache = create_files.ConvertRequestIdToAlias(TextCache, Table1)
		ConstantsURL := create_files.FindConstantsURL()
		TextCache = create_files.AddImport(TextCache, ConstantsURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextCache = create_files.ReplaceConnect_WithApplicationName(TextCache)

		//тип ID кэша
		if Table1.PrimaryKeyColumnsCount == 1 {
			//Postgres_ID_Test = ID Minimum
			TextCache = create_files.Replace_Postgres_ID_Test(TextCache, Table1)

		} else {
			TextIDMany := "(ID)"
			TextIDMany = create_files.ReplaceIDtoID_Many(TextIDMany, Table1)
			//TextCache = strings.ReplaceAll(TextCache, "(ID)", "("+Table1.Name+".StringIdentifier"+TextIDMany+")")
			//TextCache = create_files.ReplaceIDtoID_Many(TextCache, Table1)
			//TextIDMany := create_files.FindTextNameTest_ManyPK(Table1)
			//TextCache = strings.ReplaceAll(TextCache, "ReadFromCache(Postgres_ID_Test)", "ReadFromCache("+TextIDMany+")")
			//TextCache = create_files.Replace_Postgres_ID_Test(TextCache, Table1)
			EntityURL := create_files.FindModelTableURL(Table1.Name)
			TextCache = create_files.AddImport(TextCache, EntityURL)

			TextCache = create_files.Replace_Postgres_ID_Test_ManyPK(TextCache, Table1)
		}
	}

	//замена слов
	ModelName := Table1.NameGo
	TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextCache = config.Settings.TEXT_MODULE_GENERATED + TextCache

	//удаление пустого импорта
	TextCache = create_files.DeleteEmptyImport(TextCache)

	//удаление пустых строк
	TextCache = create_files.DeleteEmptyLines(TextCache)

	//SkipNow() если нет строк в БД
	TextCache = create_files.AddSkipNow(TextCache, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextCache), constants.FILE_PERMISSIONS)

	return err
}
