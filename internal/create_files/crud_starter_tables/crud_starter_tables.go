package crud_starter_tables

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

// CreateAllFiles - создаёт все файлы в папке db
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
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
			err = CreateTestFiles(Table1)
			if err != nil {
				log.Error("CreateTestFiles() table: ", Table1.Name, " error: ", err)
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
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile() + "starter_tables" + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + constants.STARTER_TABLES_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + constants.STARTER_TABLES_PREFIX + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + constants.STARTER_TABLES_PREFIX + TableName + ".go"

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
	TextDB = create_files.ReplacePackageName(TextDB, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.DeleteTemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

	}

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

	//удаление пустого импорта
	TextDB = create_files.DeleteEmptyImport(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// CreateTestFiles - создаёт 1 файл в папке db
func CreateTestFiles(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile() + "starter_tables" + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + constants.STARTER_TABLES_FILENAME_TEST
	DirReadyTable := DirReadyDB + constants.STARTER_TABLES_PREFIX + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + constants.STARTER_TABLES_PREFIX + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//заменим имя пакета на новое
	TextDB = create_files.ReplacePackageName(TextDB, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.DeleteTemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

	}

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

	//удаление пустого импорта
	TextDB = create_files.DeleteEmptyImport(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}
