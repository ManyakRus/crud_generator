package server_nrpc_starter

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

	err = CreateServerGRPCStarter()
	if err != nil {
		log.Error("CreateServerGRPCStarter() error: ", err)
		return err
	}

	return err
}

// CreateServerGRPCStarter - создаёт 1 файл в папке server_grpc
func CreateServerGRPCStarter() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesServerNRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER + micro.SeparatorFile()
	DirReadyServerNRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER + micro.SeparatorFile()
	FilenameReadyMain := DirReadyServerNRPC + "server_nrpc_starter.go"
	FilenameTemplateMain := DirTemplatesServerNRPC + "server_nrpc_starter.go_"

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextNRPCStarter := string(bytes)

	//создадим папку ready
	folders.CreateFolder(DirReadyServerNRPC)

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//заменим имя сервиса на новое
		ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
		ServiceName := config.Settings.SERVICE_NAME
		TextNRPCStarter = strings.ReplaceAll(TextNRPCStarter, ServiceNameTemplate, ServiceName)
		TextNRPCStarter = strings.ReplaceAll(TextNRPCStarter, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

		TextNRPCStarter = create_files.DeleteTemplateRepositoryImports(TextNRPCStarter)

		ProtoURL := create_files.FindProtoURL() + "/grpc_proto"
		TextNRPCStarter = create_files.AddImport(TextNRPCStarter, ProtoURL)

		GRPCServer_URL := create_files.FindGRPCServerlURL()
		TextNRPCStarter = create_files.AddImport(TextNRPCStarter, GRPCServer_URL)
	}

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextNRPCStarter), constants.FILE_PERMISSIONS)

	return err
}
