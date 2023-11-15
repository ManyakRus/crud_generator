package makefile

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateAllFiles - создаёт файл makefile в корне проекта
func CreateAllFiles() error {
	var err error

	err = CreateMakefile()
	if err != nil {
		log.Error("CreateMakefile() error: ", err)
		return err
	}

	return err
}

// CreateMakefile - создаёт 1 файл в папке grpc
func CreateMakefile() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesMakefile := DirTemplates
	DirReadyMakefile := DirReady
	FilenameReadyMakefile := DirReadyMakefile + "Makefile"
	FilenameTemplateMakefile := DirTemplatesMakefile + "Makefile_"

	bytes, err := os.ReadFile(FilenameTemplateMakefile)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMakefile, " error: ", err)
	}
	TextMakefile := string(bytes)

	//создадим папку ready
	folders.CreateFolder(DirReadyMakefile)

	//ReplaceAll
	TextMakefile = strings.ReplaceAll(TextMakefile, config.Settings.TEMPLATE_SERVICE_NAME, strings.ToLower(config.Settings.SERVICE_NAME))

	//запись файла
	err = os.WriteFile(FilenameReadyMakefile, []byte(TextMakefile), constants.FILE_PERMISSIONS)

	return err
}
