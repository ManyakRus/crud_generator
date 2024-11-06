package crud_starter_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"

	"os"
	"strings"
)

// CreateFiles_manual - создаёт 1 файл в папке crud_starter
func CreateFiles_manual(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_MANUAL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyManual := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_manual.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//
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
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_TEST_MANUAL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyManual := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_manual_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//
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
