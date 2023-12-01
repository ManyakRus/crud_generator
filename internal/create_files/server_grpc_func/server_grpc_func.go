package server_grpc_func

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

	err = CreateServerGRPCFunc()
	if err != nil {
		log.Error("CreateServerGRPCFunc() error: ", err)
		return err
	}

	return err
}

// CreateServerGRPCFunc - создаёт 1 файл в папке server_grpc
func CreateServerGRPCFunc() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesServerGRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyServerGRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	FilenameReadyServerGRPCFunc := DirReadyServerGRPC + constants.SERVER_GRPC_FUNC_FILENAME
	FilenameTemplateServerGRPCFunc := DirTemplatesServerGRPC + constants.SERVER_GRPC_FUNC_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyServerGRPC)

	bytes, err := os.ReadFile(FilenameTemplateServerGRPCFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateServerGRPCFunc, " error: ", err)
	}
	TextGRPCFunc := string(bytes)

	//заменим имя пакета на новое
	TextGRPCFunc = create_files.ReplacePackageName(TextGRPCFunc, DirReadyServerGRPC)

	////заменим имя сервиса на новое
	//ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	//ServiceName := config.Settings.SERVICE_NAME
	//TextGRPCFunc = strings.ReplaceAll(TextGRPCFunc, ServiceNameTemplate, ServiceName)
	//TextGRPCFunc = strings.ReplaceAll(TextGRPCFunc, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))
	//
	////заменим имя сервиса на новое с CamelCase
	//ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	//ServiceName = create_files.FormatName(ServiceName)
	//TextGRPCFunc = strings.ReplaceAll(TextGRPCFunc, ServiceNameTemplate, ServiceName)
	//TextGRPCFunc = strings.ReplaceAll(TextGRPCFunc, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCFunc = create_files.DeleteTemplateRepositoryImports(TextGRPCFunc)

		ConstantsURL := create_files.FindGRPCConstantsURL()
		TextGRPCFunc = create_files.AddImport(TextGRPCFunc, ConstantsURL)
	}

	//удаление пустого импорта
	TextGRPCFunc = create_files.DeleteEmptyImport(TextGRPCFunc)

	//запись файла
	err = os.WriteFile(FilenameReadyServerGRPCFunc, []byte(TextGRPCFunc), constants.FILE_PERMISSIONS)

	return err
}
