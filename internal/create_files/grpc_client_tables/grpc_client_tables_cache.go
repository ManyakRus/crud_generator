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
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//ModelName := Table1.NameGo
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//proto
		RepositoryGRPCProtoURL := create_files.Find_ProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)

		//model
		RepositoryModelURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryModelURL)

		//grpc client
		//RepositoryGRPCClientlURL := create_files.Find_GRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCClientlURL)

		//grpc client func
		GRPCClientFuncURL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClientFuncURL)

		//nrpc client
		RepositoryNRPCClientlURL := create_files.Find_NRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)

		//constants GRPC
		RepositoryGRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		//DBConstantsURL := create_files.Find_DBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//grpc_nrpc
		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		//замена Request.ID = Int64(ID)
		TextGRPCClient = create_files.Replace_PrimaryKeyM_ID(TextGRPCClient, Table1)

		//замена RequestId{}
		TextGRPCClient = create_files.ReplaceText_RequestID_PrimaryKey(TextGRPCClient, Table1)

		//замена int64(ID) на ID
		TextGRPCClient = create_files.Replace_IDtoID(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	}

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

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
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		GRPClientURL := create_files.Find_GRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		ModelTableName := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableName)

		TextGRPCClient = create_files.Replace_Postgres_ID_Test(TextGRPCClient, Table1)
		////proto
		//RepositoryGRPCProtoURL := create_files.Find_ProtoURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)
		//
		////nrpc client
		//RepositoryNRPCClientlURL := create_files.Find_NRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)
		//
		////grpc_nrpc
		//GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)
		//
		////constants GRPC
		//RepositoryGRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

		//GRPClientFuncURL := create_files.Find_GRPCClient_func_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientFuncURL)

		// замена ID на PrimaryKey
		TextGRPCClient = create_files.Replace_PrimaryKeyOtvetID(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)
	}

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//ModelName := Table1.NameGo
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//SkipNow()
	TextGRPCClient = create_files.AddSkipNow(TextGRPCClient, Table1)

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}
