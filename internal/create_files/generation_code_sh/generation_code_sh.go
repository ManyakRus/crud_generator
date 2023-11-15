package generation_code_sh

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	if config.Settings.NEED_CREATE_GRPC == false {
		return err
	}

	err = CreateFileGenerationCodeSh()
	if err != nil {
		log.Error("CreateFileGenerationCodeSh() error: ", err)
		return err
	}

	return err
}

// CreateFileGenerationCodeSh - создаёт 1 файл в папке grpc proto
func CreateFileGenerationCodeSh() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesProto := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile()
	DirReadyProto := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile()
	FilenameReadyProto := DirReadyProto + "generation_code.sh"
	FilenameTemplateProto := DirTemplatesProto + "generation_code.sh_"

	bytes, err := os.ReadFile(FilenameTemplateProto)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateProto, " error: ", err)
	}
	TextGenerationCode := string(bytes)

	//создадим папку ready proto
	folders.CreateFolder(DirReadyProto)

	//replace
	TextGenerationCode = strings.ReplaceAll(TextGenerationCode, config.Settings.TEMPLATE_SERVICE_NAME, config.Settings.SERVICE_NAME)

	//запись файла
	err = os.WriteFile(FilenameReadyProto, []byte(TextGenerationCode), 0777)

	return err
}
