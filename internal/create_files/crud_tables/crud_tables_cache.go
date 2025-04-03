package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_Cache - создаёт 1 файл "*_cache.go" в папке crud
func CreateFiles_Cache(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesCrud + config.Settings.CRUD_TABLES_CACHE_FILENAME
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
	TextCache = create_files.Replace_PackageName(TextCache, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCache = create_files.Delete_TemplateRepositoryImports(TextCache)

		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextCache = create_files.AddImport(TextCache, DBConstantsURL)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCache = create_files.AddImport(TextCache, ModelTableURL)

		//TextCache = create_files.Convert_RequestIdToAlias(TextCache, Table1)
	}

	//alias
	TextCache = Replace_IDToAlias_OtvetID(TextCache, Table1)

	//const CACHE_SIZE = 1000
	CACHE_ELEMENTS_COUNT_MAX := config.Settings.CACHE_ELEMENTS_COUNT_MAX
	Count_Now := Table1.RowsCount
	CACHE_ELEMENTS_COUNT := micro.MinInt64(Count_Now, CACHE_ELEMENTS_COUNT_MAX)
	CACHE_ELEMENTS_COUNT = micro.MaxInt64(CACHE_ELEMENTS_COUNT, 1)
	sCACHE_ELEMENTS_COUNT := micro.StringFromInt64(CACHE_ELEMENTS_COUNT)
	TextCache = create_files.FillVariable(TextCache, "CACHE_SIZE", sCACHE_ELEMENTS_COUNT)

	ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)

	//тип ID кэша
	if Table1.PrimaryKeyColumnsCount == 1 {
		_, ColumnTypeGo := create_files.Find_PrimaryKeyNameTypeGo(Table1)
		TextCache = strings.ReplaceAll(TextCache, "LRU[int64", "LRU["+ColumnTypeGo)
		TextCache = strings.ReplaceAll(TextCache, "ID int64", ColumnPK.NameGo+" "+ColumnTypeGo)
		TextCache = create_files.Replace_PrimaryKeyOtvetID(TextCache, Table1)
		TextCache = strings.ReplaceAll(TextCache, "int64(ID)", ColumnPK.NameGo)
		TextCache = strings.ReplaceAll(TextCache, ", ID)", ", "+ColumnPK.NameGo+")")
	} else {
		TextCache = strings.ReplaceAll(TextCache, "LRU[int64", "LRU[string")
		TextCache = create_files.Replace_PrimaryKeyOtvetID_Many(TextCache, Table1)
		TextIDMany := "(ID)"
		TextIDMany = Replace_IDtoID(TextIDMany, Table1)
		TextCache = strings.ReplaceAll(TextCache, "int64(ID)", "("+Table1.Name+".StringIdentifier"+TextIDMany+")")
		TextCache = Replace_IDtoID(TextCache, Table1)
	}

	//uuid
	TextCache = create_files.CheckAndAdd_ImportUUID_FromText(TextCache)

	//замена слов
	TextCache = create_files.Replace_TemplateModel_to_Model(TextCache, Table1.NameGo)
	TextCache = create_files.Replace_TemplateTableName_to_TableName(TextCache, Table1.Name)
	TextCache = create_files.AddText_ModuleGenerated(TextCache)

	//ModelName := Table1.NameGo
	//TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextCache = config.Settings.TEXT_MODULE_GENERATED + TextCache

	//удаление пустого импорта
	TextCache = create_files.Delete_EmptyImport(TextCache)

	//удаление пустых строк
	TextCache = create_files.Delete_EmptyLines(TextCache)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextCache), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_Cache_Test - создаёт 1 файл "*_cache_test.go" в папке crud
func CreateFiles_Cache_Test(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesCrud + config.Settings.CRUD_TABLES_CACHE_TEST_FILENAME
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
	TextCache = create_files.Replace_PackageName(TextCache, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCache = create_files.Delete_TemplateRepositoryImports(TextCache)

		//DBConstantsURL := create_files.Find_DBConstantsURL()
		//TextCache = create_files.AddImport(TextCache, DBConstantsURL)
		//
		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCache = create_files.AddImport(TextCache, ModelTableURL)

		//TextCache = create_files.Convert_RequestIdToAlias(TextCache, Table1)
		ConstantsURL := create_files.Find_ConstantsURL()
		TextCache = create_files.AddImport(TextCache, ConstantsURL)

		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextCache = create_files.AddImport(TextCache, CrudFuncURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextCache = create_files.Replace_Connect_WithApplicationName(TextCache)

		//тип ID кэша
		if Table1.PrimaryKeyColumnsCount == 1 {
			//Postgres_ID_Test = ID Minimum
			TextCache = Replace_Postgres_ID_Test(TextCache, Table1)

		} else {
			//TextIDMany := "(ID)"
			//TextIDMany = Replace_IDtoID_Many(TextIDMany, Table1)
			EntityURL := create_files.Find_ModelTableURL(Table1.Name)
			TextCache = create_files.AddImport(TextCache, EntityURL)

			TextCache = Replace_Postgres_ID_Test(TextCache, Table1)
		}
	}

	//замена слов
	TextCache = create_files.Replace_TemplateModel_to_Model(TextCache, Table1.NameGo)
	TextCache = create_files.Replace_TemplateTableName_to_TableName(TextCache, Table1.Name)
	TextCache = create_files.AddText_ModuleGenerated(TextCache)

	//ModelName := Table1.NameGo
	//TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextCache = strings.ReplaceAll(TextCache, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextCache = config.Settings.TEXT_MODULE_GENERATED + TextCache

	//удаление пустого импорта
	TextCache = create_files.Delete_EmptyImport(TextCache)

	//удаление пустых строк
	TextCache = create_files.Delete_EmptyLines(TextCache)

	//SkipNow() если нет строк в БД
	TextCache = create_files.AddSkipNow(TextCache, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextCache), config.Settings.FILE_PERMISSIONS)

	return err
}

// Replace_IDToAlias_OtvetID - заменяет "Otvet.ID = ID" на "Otvet.ID = alias.Name(ID)"
func Replace_IDToAlias_OtvetID(Text string, Table1 *types.Table) string {
	Otvet := Text

	TableName := Table1.Name
	IDName, _ := create_files.Find_PrimaryKeyNameType(Table1)
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == false {
		return Otvet
	}

	if TextConvert[:6] != "alias." {
		return Otvet
	}

	TextFrom := "Otvet.ID = AliasFromInt(ID)"
	TextTo := TextFrom
	TextTo = strings.ReplaceAll(TextFrom, " AliasFromInt(ID)", " "+TextConvert+"(ID)")

	Otvet = strings.ReplaceAll(Otvet, TextFrom, TextTo)
	//URL := FindURL_Alias()
	//if URL == "" {
	//	return Otvet
	//}

	Otvet = create_files.CheckAndAdd_ImportAlias(Otvet)

	return Otvet
}

// Replace_Postgres_ID_Test - заменяет текст "const Postgres_ID_Test = 0" на нужные ИД, для много колонок PrimaryKey
func Replace_Postgres_ID_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	if len(MassPK) == 0 {
		return Otvet
	}

	//заменим m1.ID = Postgres_ID_Test
	TextFind := "\tm1.ID = Postgres_ID_Test\n"
	TextNew := ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := create_files.FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "m1."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим ReadFromCache(Postgres_ID_Test)
	TextFind = "ReadFromCache(Postgres_ID_Test)"
	TextNew = "ReadFromCache("
	Comma := ""
	for _, PrimaryKey1 := range MassPK {
		Name := create_files.FindText_ColumnNameTest(PrimaryKey1)
		TextNew = TextNew + Comma + Name
		Comma = ", "
	}
	TextNew = TextNew + ")"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	////заменим ненужные Otvet.ID на Otvet.Name
	//PrimaryKey1 := MassPK[0]
	//Name := create_files.FindText_ColumnNameTest(PrimaryKey1)
	//Otvet = strings.ReplaceAll(Otvet, "Postgres_ID_Test", Name)

	return Otvet
}

// Replace_IDtoID_Many - заменяет int64(ID) на ID, и остальные PrimaryKey
func Replace_IDtoID(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextNames, TextNamesTypes, TextProtoNames := create_files.FindText_IDMany(Table1)

	Otvet = strings.ReplaceAll(Otvet, "ReplaceManyID(ID)", TextNames)
	Otvet = strings.ReplaceAll(Otvet, "int64(ID)", TextProtoNames)
	Otvet = strings.ReplaceAll(Otvet, "(ID int64", "("+TextNamesTypes)
	Otvet = strings.ReplaceAll(Otvet, "(ID)", "("+TextNames+")")
	Otvet = strings.ReplaceAll(Otvet, ", ID)", ", "+TextNames+")")
	Otvet = strings.ReplaceAll(Otvet, ", ID int64)", ", "+TextNamesTypes+")")

	return Otvet
}
