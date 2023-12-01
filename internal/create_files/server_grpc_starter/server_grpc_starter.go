package server_grpc_starter

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
	DirTemplatesServerGRPC := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyServerGRPC := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	FilenameReadyMain := DirReadyServerGRPC + constants.SERVER_GRPC_STARTER_FILENAME
	FilenameTemplateMain := DirTemplatesServerGRPC + constants.SERVER_GRPC_STARTER_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyServerGRPC)

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextGRPCStarter := string(bytes)

	//заменим имя пакета на новое
	TextGRPCStarter = create_files.ReplacePackageName(TextGRPCStarter, DirReadyServerGRPC)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCStarter = create_files.DeleteTemplateRepositoryImports(TextGRPCStarter)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCStarter = create_files.AddImport(TextGRPCStarter, ProtoURL)
	}

	//найдём текст после конца импортов
	TextAfterImport := ""
	pos1 := strings.Index(TextGRPCStarter, "\n)")
	if pos1 >= 0 {
		TextAfterImport = TextGRPCStarter[pos1+2:]
	}

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextAfterImport = strings.ReplaceAll(TextAfterImport, ServiceNameTemplate, ServiceName)
	TextAfterImport = strings.ReplaceAll(TextAfterImport, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextAfterImport = strings.ReplaceAll(TextAfterImport, ServiceNameTemplate, ServiceName)
	TextAfterImport = strings.ReplaceAll(TextAfterImport, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))
	TextGRPCStarter = TextGRPCStarter[:pos1+2] + TextAfterImport

	//удаление пустого импорта
	TextGRPCStarter = create_files.DeleteEmptyImport(TextGRPCStarter)

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextGRPCStarter), constants.FILE_PERMISSIONS)

	return err
}
