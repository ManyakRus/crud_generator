package alias

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"

	"os"
)

// CreateAllFiles - создаёт файл makefile в корне проекта
func CreateAllFiles() error {
	var err error

	err = CreateAlias()
	if err != nil {
		log.Error("CreateAlias() error: ", err)
		return err
	}

	return err
}

// CreateAlias - создаёт 1 файл в папке grpc
func CreateAlias() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesAlias := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_ALIAS + micro.SeparatorFile()
	DirReadyAlias := DirReady + config.Settings.TEMPLATE_FOLDERNAME_ALIAS + micro.SeparatorFile()
	FilenameTemplateAlias := DirTemplatesAlias + config.Settings.TEMPLATES_ALIAS_FILENAME
	FilenameReadyAlias := DirReadyAlias + config.Settings.READY_ALIAS_FILENAME

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyAlias)

	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateAlias)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateAlias, " error: ", err)
	}
	TextAlias := string(bytes)

	//запись файла в bin
	err = os.WriteFile(FilenameReadyAlias, []byte(TextAlias), config.Settings.FILE_PERMISSIONS)

	return err
}
