package crud_starter_tables

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

// CreateAllFiles - создаёт все файлы в папке db
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы starter
		if config.Settings.NEED_CREATE_DB == true {
			err = CreateFiles(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//тестовые файлы starter
		if config.Settings.NEED_CREATE_DB_TEST == true {
			err = CreateFiles_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		// создание файла manual
		if config.Settings.NEED_CREATE_MANUAL_FILES == true {
			err = CreateFiles_manual(Table1)
			if err != nil {
				log.Error("CreateFilesModel_manual() table: ", Table1.Name, " error: ", err)
				return err
			}
			err = CreateFiles_manual_test(Table1)
			if err != nil {
				log.Error("CreateFiles_manual_test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}
	}

	return err
}

// CreateFiles - создаёт 1 файл в папке db
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + ".go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//заменим имя пакета на новое
	TextDB = create_files.Replace_PackageName(TextDB, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.Delete_TemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

	}

	//создание текста
	TextDB = create_files.Replace_TemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_TemplateTableName_to_TableName(TextDB, Table1.Name)
	TextDB = create_files.AddText_ModuleGenerated(TextDB)

	//ModelName := Table1.NameGo
	//TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//замена импортов на новые URL
	TextDB = create_files.Replace_RepositoryImportsURL(TextDB)

	//удаление пустого импорта
	TextDB = create_files.Delete_EmptyImport(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_Test - создаёт 1 файл в папке db
func CreateFiles_Test(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_TEST_FILENAME
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//заменим имя пакета на новое
	TextDB = create_files.Replace_PackageName(TextDB, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.Delete_TemplateRepositoryImports(TextDB)

		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, CrudTableURL)

	}

	//создание текста
	TextDB = create_files.Replace_TemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_TemplateTableName_to_TableName(TextDB, Table1.Name)
	TextDB = create_files.AddText_ModuleGenerated(TextDB)

	//ModelName := Table1.NameGo
	//TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//замена импортов на новые URL
	TextDB = create_files.Replace_RepositoryImportsURL(TextDB)

	//удаление пустого импорта
	TextDB = create_files.Delete_EmptyImport(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_manual - создаёт 1 файл в папке crud_starter
func CreateFiles_manual(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile() + "starter_tables" + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_MANUAL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyManual := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_manual.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextManual := string(bytes)

	//заменим имя пакета на новое
	TextManual = create_files.Replace_PackageName(TextManual, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextManual = create_files.Delete_TemplateRepositoryImports(TextManual)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextManual = create_files.AddImport(TextManual, ModelTableURL)

	}

	//создание текста
	TextManual = create_files.Replace_TemplateModel_to_Model(TextManual, Table1.NameGo)
	TextManual = create_files.Replace_TemplateTableName_to_TableName(TextManual, Table1.Name)

	//ModelName := Table1.NameGo
	//TextManual = strings.ReplaceAll(TextManual, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextManual = strings.ReplaceAll(TextManual, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextManual = TextManual

	//замена импортов на новые URL
	TextManual = create_files.Replace_RepositoryImportsURL(TextManual)

	//удаление пустого импорта
	TextManual = create_files.Delete_EmptyImport(TextManual)

	//запись файла
	err = os.WriteFile(FilenameReadyManual, []byte(TextManual), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_manual_test - создаёт 1 файл в папке crud_starter
func CreateFiles_manual_test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile() + "starter_tables" + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_TEST_MANUAL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyManual := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_manual_test.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextManual := string(bytes)

	//заменим имя пакета на новое
	TextManual = create_files.Replace_PackageName(TextManual, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextManual = create_files.Delete_TemplateRepositoryImports(TextManual)

		//
		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextManual = create_files.AddImport(TextManual, ModelTableURL)

		//
		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextManual = create_files.AddImport(TextManual, CrudTableURL)

	}

	//создание текста
	TextManual = create_files.Replace_TemplateModel_to_Model(TextManual, Table1.NameGo)
	TextManual = create_files.Replace_TemplateTableName_to_TableName(TextManual, Table1.Name)

	//ModelName := Table1.NameGo
	//TextManual = strings.ReplaceAll(TextManual, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextManual = strings.ReplaceAll(TextManual, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextManual = TextManual

	//замена импортов на новые URL
	TextManual = create_files.Replace_RepositoryImportsURL(TextManual)

	//удаление пустого импорта
	TextManual = create_files.Delete_EmptyImport(TextManual)

	//запись файла
	err = os.WriteFile(FilenameReadyManual, []byte(TextManual), config.Settings.FILE_PERMISSIONS)

	return err
}
