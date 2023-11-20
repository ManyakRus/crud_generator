package grpc_client

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles() error {
	var err error

	if config.Settings.NEED_CREATE_GRPC == false {
		return err
	}

	err = CreateGRPCClient()
	if err != nil {
		log.Error("CreateGRPCClient() error: ", err)
		return err
	}

	err = CreateGRPCClientTest()
	if err != nil {
		log.Error("CreateGRPCClientTest() error: ", err)
		return err
	}

	return err
}

// CreateGRPCClient - создаёт 1 файл в папке grpc_client
func CreateGRPCClient() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesClientGRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyClientGRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyMain := DirReadyClientGRPC + constants.GRPC_CLIENT_FILENAME
	FilenameTemplateMain := DirTemplatesClientGRPC + constants.GRPC_CLIENT_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyClientGRPC)

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextGRPCClient, DirReadyClientGRPC)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ProtoURL)

		DBConstantsURL := create_files.FindDBConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//заменим имя сервиса на новое
		ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
		ServiceName := config.Settings.SERVICE_NAME
		TextGRPCClient = strings.ReplaceAll(TextGRPCClient, ServiceNameTemplate, ServiceName)
		TextGRPCClient = strings.ReplaceAll(TextGRPCClient, strings.ToUpper(ServiceNameTemplate), strings.ToUpper(ServiceName))
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

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

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
	FilenameReadyMain := DirReadyClientGRPC + constants.GRPC_CLIENT_TEST_FILENAME
	FilenameTemplateMain := DirTemplatesClientGRPC + constants.GRPC_CLIENT_TEST_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyClientGRPC)

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextGRPCClient, DirReadyClientGRPC)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		DBConstantsURL := create_files.FindDBConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//заменим имя сервиса на новое
		ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
		ServiceName := config.Settings.SERVICE_NAME
		TextGRPCClient = strings.ReplaceAll(TextGRPCClient, ServiceNameTemplate, ServiceName)
		TextGRPCClient = strings.ReplaceAll(TextGRPCClient, strings.ToUpper(ServiceNameTemplate), strings.ToUpper(ServiceName))

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

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}
