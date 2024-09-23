package env_file

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"

	"os"
	"strings"
)

// CreateAllFiles - создаёт файл makefile в корне проекта
func CreateAllFiles() error {
	var err error

	err = CreateENV()
	if err != nil {
		log.Error("CreateENV() error: ", err)
		return err
	}

	return err
}

// CreateENV - создаёт 1 файл в папке grpc
func CreateENV() error {
	var err error

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesMakefile := DirTemplates
	DirReadyENV := DirReady
	FilenameReadyENV := DirReadyENV + config.Settings.ENV_FILENAME
	FilenameTemplateENV := DirTemplatesMakefile + config.Settings.ENV_FILENAME + "_"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyENV)

	//не стираем файл .env
	ok, err := micro.FileExists(FilenameReadyENV)
	if ok == true {
		return err
	}

	bytes, err := os.ReadFile(FilenameTemplateENV)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateENV, " error: ", err)
	}
	TextMakefile := string(bytes)

	//ReplaceAll
	SERVICE_NAME := config.Settings.SERVICE_NAME
	TEMPLATE_SERVICE_NAME := config.Settings.TEMPLATE_SERVICE_NAME
	TextMakefile = strings.ReplaceAll(TextMakefile, strings.ToUpper(TEMPLATE_SERVICE_NAME), strings.ToUpper(SERVICE_NAME))

	//DB_HOST
	VariableName := "DB_HOST"
	Value := postgres_gorm.Settings.DB_HOST
	TextMakefile = ReplaceVariable(TextMakefile, VariableName, Value)

	//DB_PORT
	VariableName = "DB_PORT"
	Value = postgres_gorm.Settings.DB_PORT
	TextMakefile = ReplaceVariable(TextMakefile, VariableName, Value)

	//DB_SCHEMA
	VariableName = "DB_SCHEME"
	Value = postgres_gorm.Settings.DB_SCHEMA
	TextMakefile = ReplaceVariable(TextMakefile, VariableName, Value)

	//DB_NAME
	VariableName = "DB_NAME"
	Value = postgres_gorm.Settings.DB_NAME
	TextMakefile = ReplaceVariable(TextMakefile, VariableName, Value)

	//DB_USER
	VariableName = "DB_USER"
	Value = postgres_gorm.Settings.DB_USER
	TextMakefile = ReplaceVariable(TextMakefile, VariableName, Value)

	//DB_PASSWORD
	VariableName = "DB_PASSWORD"
	Value = postgres_gorm.Settings.DB_PASSWORD
	TextMakefile = ReplaceVariable(TextMakefile, VariableName, Value)

	//запись файла в корень проекта
	err = os.WriteFile(FilenameReadyENV, []byte(TextMakefile), config.Settings.FILE_PERMISSIONS)

	//запись файла в bin
	Dir := DirReady + "bin" + micro.SeparatorFile()
	folders.CreateFolder(Dir)
	FilenameReadyENV = Dir + config.Settings.ENV_FILENAME
	err = os.WriteFile(FilenameReadyENV, []byte(TextMakefile), config.Settings.FILE_PERMISSIONS)

	return err
}

// ReplaceVariable - заменяет переменную в тексте
func ReplaceVariable(Text, VariableName, Value string) string {
	Otvet := Text

	pos1 := strings.Index(Otvet, "\n"+VariableName)
	if pos1 < 0 {
		//нет такого
		Otvet = Otvet + "\n" + VariableName + " = " + Value
		return Otvet
	}

	s2 := Text[pos1+1:]
	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1+1] + VariableName + "=" + Value + Otvet[pos1+1+posEnd:]

	return Otvet
}
