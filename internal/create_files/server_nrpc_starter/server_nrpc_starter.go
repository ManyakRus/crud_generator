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

	err = CreateFile_ServerGRPCStarter()
	if err != nil {
		log.Error("CreateFile_ServerGRPCStarter() error: ", err)
		return err
	}

	return err
}

// CreateFile_ServerGRPCStarter - создаёт 1 файл в папке server_grpc
func CreateFile_ServerGRPCStarter() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesServerNRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER + micro.SeparatorFile()
	DirReadyServerNRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER + micro.SeparatorFile()
	FilenameReadyMain := DirReadyServerNRPC + "server_nrpc_starter.go"
	FilenameTemplateMain := DirTemplatesServerNRPC + "server_nrpc_starter.go_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyServerNRPC)

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextNRPCStarter := string(bytes)

	//заменим имя пакета на новое
	TextNRPCStarter = create_files.Replace_PackageName(TextNRPCStarter, DirReadyServerNRPC)

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextNRPCStarter = create_files.Delete_TemplateRepositoryImports(TextNRPCStarter)

		//Proto URL
		RepositoryProtoURL := create_files.Find_ProtoURL()
		TextNRPCStarter = create_files.AddImport(TextNRPCStarter, RepositoryProtoURL)

		//server grpc URL
		RepositoryServerGRPCURL := create_files.Find_GRPCServerURL()
		TextNRPCStarter = create_files.AddImport(TextNRPCStarter, RepositoryServerGRPCURL)
	}

	//найдём текст после конца импортов
	TextAfterImport := ""
	pos1 := strings.Index(TextNRPCStarter, "\n)")
	if pos1 >= 0 {
		TextAfterImport = TextNRPCStarter[pos1+2:]
	}

	//заменим название сервиса
	ServiceName := config.Settings.SERVICE_NAME
	ServiceNameProto := micro.StringFromUpperCase(ServiceName)
	TEMPLATE_SERVICE_NAME := config.Settings.TEMPLATE_SERVICE_NAME
	TextAfterImport = strings.ReplaceAll(TextAfterImport, TEMPLATE_SERVICE_NAME, ServiceNameProto)
	//заменим ещё раз с большой буквы
	TEMPLATE_SERVICE_NAME = micro.StringFromUpperCase(TEMPLATE_SERVICE_NAME)
	TextAfterImport = strings.ReplaceAll(TextAfterImport, TEMPLATE_SERVICE_NAME, ServiceNameProto)
	TextNRPCStarter = TextNRPCStarter[:pos1+2] + TextAfterImport

	//удаление пустого импорта
	TextNRPCStarter = create_files.Delete_EmptyImport(TextNRPCStarter)

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextNRPCStarter), constants.FILE_PERMISSIONS)

	return err
}
