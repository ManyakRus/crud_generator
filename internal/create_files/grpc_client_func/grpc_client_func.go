package grpc_client_func

import (
	"github.com/ManyakRus/crud_generator/internal/config"
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

	err = CreateFileGRPCClientFunc()
	if err != nil {
		log.Error("CreateFileGRPCClientFunc() error: ", err)
		return err
	}

	err = CreateFileGRPCClientFuncTest()
	if err != nil {
		log.Error("CreateFileGRPCClientFuncTest() error: ", err)
		return err
	}

	return err
}

// CreateFileGRPCClientFunc - создаёт 1 файл в папке grpc_client_func
func CreateFileGRPCClientFunc() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClientFunc := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC + micro.SeparatorFile()
	DirReadyGRPCClientFunc := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC + micro.SeparatorFile()
	FilenameReadyGRPCClientFunc := DirReadyGRPCClientFunc + config.Settings.TEMPLATE_GRPC_CLIENT_FUNC_FILENAME
	FilenameReadyGRPCClientFunc = create_files.Delete_LastUnderline(FilenameReadyGRPCClientFunc)
	FilenameTemplateGRPCClientFunc := DirTemplatesGRPCClientFunc + config.Settings.TEMPLATE_GRPC_CLIENT_FUNC_FILENAME

	ok, err := micro.FileExists(FilenameTemplateGRPCClientFunc)
	if err != nil {
		log.Panic("FileExists() ", FilenameTemplateGRPCClientFunc, " error: ", err)
	}
	if ok == false {
		log.Debug("FileExists() ", FilenameTemplateGRPCClientFunc, " = false")
		return err
	}

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyGRPCClientFunc)

	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateGRPCClientFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunc, " error: ", err)
	}
	TextGRPCClientFunc := string(bytes)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClientFunc = create_files.Delete_TemplateRepositoryImports(TextGRPCClientFunc)

		//
		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCClientFunc = create_files.AddImport(TextGRPCClientFunc, ProtoURL)

		//
		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextGRPCClientFunc = create_files.AddImport(TextGRPCClientFunc, DBConstantsURL)

		//
		GRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClientFunc = create_files.AddImport(TextGRPCClientFunc, GRPCConstantsURL)

	}

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCClientFunc = strings.ReplaceAll(TextGRPCClientFunc, "grpc_proto.", TextProto+".")

	//заменим имя сервиса на новое
	TextGRPCClientFunc = create_files.Replace_ServiceName(TextGRPCClientFunc)
	//ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	//ServiceName := config.Settings.SERVICE_NAME
	//TextAfterImport = strings.ReplaceAll(TextAfterImport, ServiceNameTemplate, ServiceName)
	//TextAfterImport = strings.ReplaceAll(TextAfterImport, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	TextGRPCClientFunc = create_files.Replace_ServiceName_CamelCase(TextGRPCClientFunc)

	////заменим название сервиса
	//ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	//ServiceNameNew := config.Settings.SERVICE_NAME
	//TextGRPCClientFunc = strings.ReplaceAll(TextGRPCClientFunc, ServiceNameTemplate, ServiceNameNew)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClientFunc, []byte(TextGRPCClientFunc), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFileGRPCClientFuncTest - создаёт 1 файл в папке grpc_client_func
func CreateFileGRPCClientFuncTest() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClientFunc := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC + micro.SeparatorFile()
	DirReadyGRPCClientFunc := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC + micro.SeparatorFile()
	FilenameReadyGRPCClientFunc := DirReadyGRPCClientFunc + config.Settings.TEMPLATE_GRPC_CLIENT_FUNC_TEST_FILENAME
	FilenameReadyGRPCClientFunc = create_files.Delete_LastUnderline(FilenameReadyGRPCClientFunc)
	FilenameTemplateGRPCClientFunc := DirTemplatesGRPCClientFunc + config.Settings.TEMPLATE_GRPC_CLIENT_FUNC_TEST_FILENAME

	ok, err := micro.FileExists(FilenameTemplateGRPCClientFunc)
	if err != nil {
		log.Panic("FileExists() ", FilenameTemplateGRPCClientFunc, " error: ", err)
	}
	if ok == false {
		log.Debug("FileExists() ", FilenameTemplateGRPCClientFunc, " = false")
		return err
	}

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyGRPCClientFunc)

	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateGRPCClientFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunc, " error: ", err)
	}
	TextGRPCClientFunc := string(bytes)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClientFunc = create_files.Delete_TemplateRepositoryImports(TextGRPCClientFunc)

		//
		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextGRPCClientFunc = create_files.AddImport(TextGRPCClientFunc, DBConstantsURL)

	}

	//заменим название сервиса
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceNameNew := config.Settings.SERVICE_NAME
	TextGRPCClientFunc = strings.ReplaceAll(TextGRPCClientFunc, ServiceNameTemplate, ServiceNameNew)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClientFunc, []byte(TextGRPCClientFunc), config.Settings.FILE_PERMISSIONS)

	return err
}
