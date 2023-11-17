package server_nrpc_starter

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
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

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextNRPCStarter, DirReadyServerNRPC)

	//создадим папку ready
	folders.CreateFolder(DirReadyServerNRPC)

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextNRPCStarter = create_files.DeleteTemplateRepositoryImports(TextNRPCStarter)

		//Proto URL
		RepositoryProtoURL := create_files.FindGRPCProtoURL()
		TextNRPCStarter = create_files.AddImport(TextNRPCStarter, RepositoryProtoURL)

		//server grpc URL
		RepositoryServerGRPCURL := create_files.FindGRPCServerURL()
		TextNRPCStarter = create_files.AddImport(TextNRPCStarter, RepositoryServerGRPCURL)
	}

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextNRPCStarter), constants.FILE_PERMISSIONS)

	return err
}