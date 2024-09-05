package grpc_client_tables

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

// CreateFiles_GRPC_Client_Cache - создаёт 1 файл в папке grpc_client
func CreateFiles_GRPC_Client_Cache(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCache := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyCache := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesCache + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_CACHE_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCache + "grpc_" + TableName + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + "_cache.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.ReplacePackageName(TextGRPCClient, DirReadyTable)

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//TextGRPCClient = create_files.ReplaceServiceURLImports(TextGRPCClient)
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		//proto
		RepositoryGRPCProtoURL := create_files.FindProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)

		//model
		RepositoryModelURL := create_files.FindModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryModelURL)

		//grpc client
		//RepositoryGRPCClientlURL := create_files.FindGRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCClientlURL)

		//grpc client func
		GRPCClientFuncURL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClientFuncURL)

		//nrpc client
		RepositoryNRPCClientlURL := create_files.FindNRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)

		//constants GRPC
		RepositoryGRPCConstantsURL := create_files.FindGRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		//DBConstantsURL := create_files.FindDBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//grpc_nrpc
		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		//замена Request.ID = Int64(ID)
		TextGRPCClient = create_files.ReplacePrimaryKeyM_ID(TextGRPCClient, Table1)

		//замена RequestId{}
		TextGRPCClient = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCClient, Table1)

		//замена int64(ID) на ID
		TextGRPCClient = create_files.ReplaceIDtoID(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAddImportUUID_FromText(TextGRPCClient)

	}

	//удаление пустого импорта
	TextGRPCClient = create_files.DeleteEmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.DeleteEmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClient, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateFiles_GRPC_Client_Cache_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_GRPC_Client_Cache_Test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCache := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyCache := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesCache + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_CACHE_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCache + "grpc_" + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReadyCache := DirReadyTable + "grpc_" + TableName + "_cache_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.ReplacePackageName(TextGRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		GRPClientURL := create_files.FindGRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		ModelTableName := create_files.FindModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableName)

		TextGRPCClient = create_files.Replace_Postgres_ID_Test(TextGRPCClient, Table1)
		////proto
		//RepositoryGRPCProtoURL := create_files.FindProtoURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)
		//
		////nrpc client
		//RepositoryNRPCClientlURL := create_files.FindNRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)
		//
		////grpc_nrpc
		//GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)
		//
		////constants GRPC
		//RepositoryGRPCConstantsURL := create_files.FindGRPCConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		GRPClientTableURL := create_files.FindGRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

		//GRPClientFuncURL := create_files.Find_GRPCClient_func_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientFuncURL)

		// замена ID на PrimaryKey
		TextGRPCClient = create_files.ReplacePrimaryKeyOtvetID(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAddImportUUID_FromText(TextGRPCClient)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//SkipNow()
	TextGRPCClient = create_files.AddSkipNow(TextGRPCClient, Table1)

	//замена импортов на новые URL
	TextGRPCClient = create_files.ReplaceServiceURLImports(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.DeleteEmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.DeleteEmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}
