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

// CreateFiles_ReadObject - создаёт 1 файл в папке db
func CreateFiles_ReadObject(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_READOBJECT_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_readobject.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

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

		ObjectTableURL := create_files.Find_ObjectTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ObjectTableURL)

	}

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextDB = strings.ReplaceAll(TextDB, "grpc_proto.", TextProto+".")

	//создание текста
	TextDB = create_files.Replace_ObjectTemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_ObjectTemplateTableName_to_TableName(TextDB, Table1.Name)
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

// CreateFiles_ReadObject_Test - создаёт 1 файл в папке db
func CreateFiles_ReadObject_Test(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.STARTER_TABLES_READOBJECT_TEST_FILENAME
	DirReadyTable := DirReadyDB + config.Settings.STARTER_TABLES_PREFIX + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.STARTER_TABLES_PREFIX + TableName + "_readobject_test.go"

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

		CrudObjectTableURL := create_files.Find_CrudObjectTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, CrudObjectTableURL)
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
