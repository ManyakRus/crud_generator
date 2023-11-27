package nrpc_client

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

	err = CreateNRPCClient()
	if err != nil {
		log.Error("CreateNRPCClient() error: ", err)
		return err
	}

	err = CreateNRPCClientTest()
	if err != nil {
		log.Error("CreateNRPCClientTest() error: ", err)
		return err
	}

	return err
}

// CreateNRPCClient - создаёт 1 файл в папке grpc_client
func CreateNRPCClient() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesClientNRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	DirReadyClientNRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyNRPC := DirReadyClientNRPC + constants.NRPC_CLIENT_FILENAME
	FilenameTemplateNRPC := DirTemplatesClientNRPC + constants.NRPC_CLIENT_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyClientNRPC)

	bytes, err := os.ReadFile(FilenameTemplateNRPC)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateNRPC, " error: ", err)
	}
	TextNRPCClient := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextNRPCClient, DirReadyClientNRPC)

	//добавим комментарий
	TextNRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextNRPCClient

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextNRPCClient = create_files.DeleteTemplateRepositoryImports(TextNRPCClient)

		ProtoURL := create_files.FindProtoURL()
		TextNRPCClient = create_files.AddImport(TextNRPCClient, ProtoURL)

		DBConstantsURL := create_files.FindDBConstantsURL()
		TextNRPCClient = create_files.AddImport(TextNRPCClient, DBConstantsURL)
	}

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, ServiceNameTemplate, ServiceName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, ServiceNameTemplate, ServiceName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//удаление пустого импорта
	TextNRPCClient = create_files.DeleteEmptyImport(TextNRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyNRPC, []byte(TextNRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateNRPCClientTest - создаёт 1 файл в папке grpc_client
func CreateNRPCClientTest() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesClientNRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	DirReadyClientNRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyNRPC := DirReadyClientNRPC + constants.NRPC_CLIENT_TEST_FILENAME
	FilenameTemplateNRPC := DirTemplatesClientNRPC + constants.NRPC_CLIENT_TEST_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyClientNRPC)

	bytes, err := os.ReadFile(FilenameTemplateNRPC)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateNRPC, " error: ", err)
	}
	TextNRPCClient := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextNRPCClient, DirReadyClientNRPC)

	//добавим комментарий
	TextNRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextNRPCClient

	//добавим импорты
	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//	TextNRPCClient = create_files.DeleteTemplateRepositoryImports(TextNRPCClient)
	//
	//	DBConstantsURL := create_files.FindDBConstantsURL()
	//	TextNRPCClient = create_files.AddImport(TextNRPCClient, DBConstantsURL)
	//}

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, ServiceNameTemplate, ServiceName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, ServiceNameTemplate, ServiceName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//удаление пустого импорта
	TextNRPCClient = create_files.DeleteEmptyImport(TextNRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyNRPC, []byte(TextNRPCClient), constants.FILE_PERMISSIONS)

	return err
}
