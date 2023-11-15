package main_file

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

	err = CreateFileMain()
	if err != nil {
		log.Error("CreateFileMain() error: ", err)
		return err
	}

	return err
}

// CreateFileMain - создаёт 1 файл в папке grpc
func CreateFileMain() error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesMain := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MAIN + micro.SeparatorFile()
	DirReadyMain := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MAIN + micro.SeparatorFile() + config.Settings.SERVICE_NAME + micro.SeparatorFile()
	FilenameReadyMain := DirReadyMain + "main.go"
	FilenameTemplateMain := DirTemplatesMain + "main.go_"

	bytes, err := os.ReadFile(FilenameTemplateMain)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateMain, " error: ", err)
	}
	TextMain := string(bytes)

	//создадим папку ready
	folders.CreateFolder(DirReadyMain)

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//GRPC
		if config.Settings.NEED_CREATE_GRPC == true {
			TextImport := create_files.FindGRPCServerlURL()
			TextMain = create_files.AddImport(TextMain, TextImport)
		} else {
			TextMain = strings.ReplaceAll(TextMain, "\n\tserver_grpc.Start()", "")
		}

		//NRPC
		if config.Settings.NEED_CREATE_NRPC == true {
			TextImport := create_files.FindNRPCServerlURL()
			TextMain = create_files.AddImport(TextMain, TextImport)
		} else {
			TextMain = strings.ReplaceAll(TextMain, "\n\tserver_nrpc.Start()", "")
		}
	}

	//замена импортов на новые URL
	TextMain = create_files.ReplaceServiceURLImports(TextMain)

	//запись файла
	err = os.WriteFile(FilenameReadyMain, []byte(TextMain), constants.FILE_PERMISSIONS)

	return err
}
