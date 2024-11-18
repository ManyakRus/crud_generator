package grpc_client

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"

	"os"
	"strings"
)

// CreateGRPCClient - создаёт 1 файл в папке grpc_client
func CreateGRPCClient() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesClientGRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyClientGRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyMain := DirReadyClientGRPC + config.Settings.GRPC_CLIENT_FILENAME
	FilenameTemplateMain := DirTemplatesClientGRPC + config.Settings.GRPC_CLIENT_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyClientGRPC)

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyClientGRPC)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//grpc_proto
		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ProtoURL)

		//constants db
		//DBConstantsURL := create_files.Find_DBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//grpc_nrpc
		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		//nrpc_client
		if config.Settings.NEED_CREATE_NRPC == false {
			NRPC_CLIENT_URL := create_files.Find_NRPC_Client_URL()
			TextGRPCClient = create_files.AddImport(TextGRPCClient, NRPC_CLIENT_URL)
		}

		//grpc_client_func
		GRPC_CLIENT_FUNC_URL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_CLIENT_FUNC_URL)

		//CRUD_STARTER_URL
		CRUD_STARTER_URL := create_files.Find_CrudStarterURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, CRUD_STARTER_URL)

		//constants GRPC
		RepositoryGRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

	}

	//добавим импорт uuid
	TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, ServiceNameTemplate, ServiceName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, ServiceNameTemplate, ServiceName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextGRPCClient), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateGRPCClientTest - создаёт 1 файл в папке grpc_client
func CreateGRPCClientTest() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesClientGRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyClientGRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyMain := DirReadyClientGRPC + config.Settings.GRPC_CLIENT_TEST_FILENAME
	FilenameTemplateMain := DirTemplatesClientGRPC + config.Settings.GRPC_CLIENT_TEST_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyClientGRPC)

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyClientGRPC)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//DBConstantsURL := create_files.Find_DBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

	}

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, ServiceNameTemplate, ServiceName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, ServiceNameTemplate, ServiceName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextGRPCClient), config.Settings.FILE_PERMISSIONS)

	return err
}
