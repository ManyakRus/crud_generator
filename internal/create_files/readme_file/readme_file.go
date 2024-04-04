package readme_file

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

// CreateAllFiles - создаёт файл makefile в корне проекта
func CreateAllFiles() error {
	var err error

	err = CreateReadmeMD()
	if err != nil {
		log.Error("CreateReadmeMD() error: ", err)
		return err
	}

	err = CreateReadmeRus()
	if err != nil {
		log.Error("CreateReadmeRus() error: ", err)
		return err
	}

	return err
}

// CreateReadmeMD - создаёт 1 файл в корне проекта
func CreateReadmeMD() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesReadmeFile := DirTemplates
	DirReadyReadmeFile := DirReady
	filename := create_files.DeleteLastUnderline(config.Settings.TEMPLATES_README_MD_FILENAME)
	FilenameTemplateReadmeFile := DirTemplatesReadmeFile + config.Settings.TEMPLATES_README_MD_FILENAME
	FilenameReadyReadmeFile := DirReadyReadmeFile + filename

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyReadmeFile)

	bytes, err := os.ReadFile(FilenameTemplateReadmeFile)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateReadmeFile, " error: ", err)
	}
	TextReadmeFile := string(bytes)

	//заменим URL
	CrudStarterURLOld := "gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/crud_starter"
	CrudStarterURL := create_files.FindCrudStarterURL()
	TextReadmeFile = strings.ReplaceAll(TextReadmeFile, CrudStarterURLOld, CrudStarterURL)

	//запись файла
	err = os.WriteFile(FilenameReadyReadmeFile, []byte(TextReadmeFile), constants.FILE_PERMISSIONS)

	return err
}

// CreateReadmeRus - создаёт 1 файл в корне проекта
func CreateReadmeRus() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesReadmeFile := DirTemplates
	DirReadyReadmeFile := DirReady
	filename := create_files.DeleteLastUnderline(config.Settings.TEMPLATES_README_RUS_FILENAME)
	FilenameTemplateReadmeFile := DirTemplatesReadmeFile + config.Settings.TEMPLATES_README_RUS_FILENAME
	FilenameReadyReadmeFile := DirReadyReadmeFile + filename

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyReadmeFile)

	bytes, err := os.ReadFile(FilenameTemplateReadmeFile)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateReadmeFile, " error: ", err)
	}
	TextReadmeFile := string(bytes)

	//заменим URL
	CrudStarterURLOld := "gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/crud_starter"
	CrudStarterURL := create_files.FindCrudStarterURL()
	TextReadmeFile = strings.ReplaceAll(TextReadmeFile, CrudStarterURLOld, CrudStarterURL)

	//запись файла
	err = os.WriteFile(FilenameReadyReadmeFile, []byte(TextReadmeFile), constants.FILE_PERMISSIONS)

	return err
}
