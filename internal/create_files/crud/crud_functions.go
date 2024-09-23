package crud

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"

	"os"
	"strings"
)

// CreateCrud - создаёт 1 файл в папке grpc_client
func CreateCrud() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATES_CRUD_FUNC_FOLDERNAME + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATES_CRUD_FUNC_FOLDERNAME + micro.SeparatorFile()
	FilenameReadyCrud := DirReadyCrud + create_files.Delete_LastUnderline(config.Settings.TEMPLATES_CRUD_FUNC_TEST_FILENAME)
	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_FUNC_TEST_FILENAME

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyCrud)

	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyCrud)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		//constants db
		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, DBConstantsURL)

	}

	//добавим импорт uuid
	TextCrud = create_files.CheckAndAdd_ImportUUID_FromText(TextCrud)

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextCrud = strings.ReplaceAll(TextCrud, ServiceNameTemplate, ServiceName)
	TextCrud = strings.ReplaceAll(TextCrud, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextCrud = strings.ReplaceAll(TextCrud, ServiceNameTemplate, ServiceName)
	TextCrud = strings.ReplaceAll(TextCrud, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrud, []byte(TextCrud), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateCrudTest - создаёт 1 файл в папке grpc_client
func CreateCrudTest() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATES_CRUD_FUNC_FOLDERNAME + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATES_CRUD_FUNC_FOLDERNAME + micro.SeparatorFile()
	FilenameReadyCrud := DirReadyCrud + create_files.Delete_LastUnderline(config.Settings.TEMPLATES_CRUD_FUNC_TEST_FILENAME)
	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_FUNC_TEST_FILENAME

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyCrud)

	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyCrud)

	//добавим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, DBConstantsURL)

	}

	//заменим имя сервиса на новое
	ServiceNameTemplate := config.Settings.TEMPLATE_SERVICE_NAME
	ServiceName := config.Settings.SERVICE_NAME
	TextCrud = strings.ReplaceAll(TextCrud, ServiceNameTemplate, ServiceName)
	TextCrud = strings.ReplaceAll(TextCrud, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//заменим имя сервиса на новое с CamelCase
	ServiceNameTemplate = create_files.FormatName(ServiceNameTemplate)
	ServiceName = create_files.FormatName(ServiceName)
	TextCrud = strings.ReplaceAll(TextCrud, ServiceNameTemplate, ServiceName)
	TextCrud = strings.ReplaceAll(TextCrud, micro.StringFromUpperCase(ServiceNameTemplate), micro.StringFromUpperCase(ServiceName))

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrud, []byte(TextCrud), config.Settings.FILE_PERMISSIONS)

	return err
}
