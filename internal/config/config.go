package config

import (
	"github.com/ManyakRus/crud_generator/internal/constants"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	INCLUDE_TABLES                   string
	EXCLUDE_TABLES                   string
	TEMPLATE_FOLDERNAME              string
	TEMPLATE_FOLDERNAME_MODEL        string
	TEMPLATE_FOLDERNAME_DB           string
	TEMPLATE_FOLDERNAME_CRUD         string
	TEMPLATE_FOLDERNAME_TABLES       string
	TEMPLATE_FOLDERNAME_GRPC         string
	TEMPLATE_FOLDERNAME_GRPC_PROTO   string
	TEMPLATE_FOLDERNAME_GRPC_SERVER  string
	TEMPLATE_FOLDERNAME_GRPC_CLIENT  string
	TEMPLATE_FOLDERNAME_NRPC         string
	TEMPLATE_FOLDERNAME_NRPC_SERVER  string
	TEMPLATE_FOLDERNAME_NRPC_CLIENT  string
	TEMPLATE_SERVICE_NAME            string
	TEMPLATE_FOLDERNAME_CRUD_STARTER string
	TEMPLATE_FOLDERNAME_ALIAS        string
	TEMPLATE_FOLDER_CRUD_FUNCTIONS   string
	NEED_CREATE_MODEL_STRUCT         bool
	NEED_CREATE_MODEL_CRUD           bool
	NEED_CREATE_DB                   bool
	NEED_CREATE_DB_TEST              bool
	NEED_CREATE_GRPC                 bool
	NEED_CREATE_GRPC_SERVER_TEST     bool
	NEED_CREATE_GRPC_CLIENT_TEST     bool
	NEED_CREATE_NRPC                 bool
	NEED_CREATE_NRPC_SERVER_TEST     bool
	NEED_CREATE_NRPC_CLIENT_TEST     bool
	SERVICE_NAME                     string
	SERVICE_REPOSITORY_URL           string
	TEXT_TEMPLATE_MODEL              string
	TEXT_TEMPLATE_TABLENAME          string
	HAS_IS_DELETED                   bool
	READY_FOLDERNAME                 string
	TEXT_DB_MODIFIED_AT              string
	TEXT_DB_CREATED_AT               string
	TEXT_DB_IS_DELETED               string
	USE_DEFAULT_TEMPLATE             bool
	PREFIX_SERVER_GRPC               string
	PREFIX_CRUD                      string
	PREFIX_MODEL                     string
	COMMENT_MODEL_STRUCT             string
	TEXT_MODULE_GENERATED            string
	TEMPLATE_FOLDERNAME_MAIN         string
	TEMPLATE_REPOSITORY_URL          string
	PREFIX_TABLE                     string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	//dir := micro.ProgramDir_bin()

	Settings = SettingsINI{}
	Settings.INCLUDE_TABLES = os.Getenv("INCLUDE_TABLES")
	Settings.EXCLUDE_TABLES = os.Getenv("EXCLUDE_TABLES")
	Settings.TEMPLATE_FOLDERNAME = os.Getenv("TEMPLATE_FOLDERNAME")
	Settings.TEMPLATE_FOLDERNAME_MODEL = os.Getenv("TEMPLATE_FOLDERNAME_MODEL")
	Settings.TEMPLATE_FOLDERNAME_CRUD = os.Getenv("TEMPLATE_FOLDERNAME_CRUD")
	Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO = os.Getenv("TEMPLATE_FOLDERNAME_GRPC_PROTO")
	Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER = os.Getenv("TEMPLATE_FOLDERNAME_GRPC_SERVER")
	Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT = os.Getenv("TEMPLATE_FOLDERNAME_GRPC_CLIENT")
	Settings.TEMPLATE_FOLDERNAME_NRPC = os.Getenv("TEMPLATE_FOLDERNAME_NRPC")
	Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER = os.Getenv("TEMPLATE_FOLDERNAME_NRPC_SERVER")
	Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT = os.Getenv("TEMPLATE_FOLDERNAME_NRPC_CLIENT")
	Settings.TEXT_TEMPLATE_MODEL = os.Getenv("TEXT_TEMPLATE_MODEL")
	Settings.TEXT_TEMPLATE_TABLENAME = os.Getenv("TEXT_TEMPLATE_TABLENAME")
	Settings.TEMPLATE_SERVICE_NAME = os.Getenv("TEMPLATE_SERVICE_NAME")
	Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER = os.Getenv("TEMPLATE_FOLDERNAME_CRUD_STARTER")

	SERVICE_REPOSITORY_URL := strings.ToLower(os.Getenv("SERVICE_REPOSITORY_URL"))
	Settings.SERVICE_REPOSITORY_URL = SERVICE_REPOSITORY_URL

	Settings.TEMPLATE_FOLDERNAME_ALIAS = os.Getenv("TEMPLATE_FOLDERNAME_ALIAS")
	Settings.PREFIX_SERVER_GRPC = os.Getenv("PREFIX_SERVER_GRPC")

	sHAS_IS_DELETED := os.Getenv("HAS_IS_DELETED")

	sNEED_CREATE_DB := os.Getenv("NEED_CREATE_DB")
	Settings.NEED_CREATE_DB = BoolFromString(sNEED_CREATE_DB)

	sNEED_GRPC := os.Getenv("NEED_CREATE_GRPC")
	Settings.NEED_CREATE_GRPC = BoolFromString(sNEED_GRPC)

	sNEED_NRPC := os.Getenv("NEED_CREATE_NRPC")
	Settings.NEED_CREATE_NRPC = BoolFromString(sNEED_NRPC)

	sNEED_CREATE_MODEL_STRUCT := os.Getenv("NEED_CREATE_MODEL_STRUCT")
	Settings.NEED_CREATE_MODEL_STRUCT = BoolFromString(sNEED_CREATE_MODEL_STRUCT)

	sNEED_CREATE_MODEL_CRUD := os.Getenv("NEED_CREATE_MODEL_CRUD")
	Settings.NEED_CREATE_MODEL_CRUD = BoolFromString(sNEED_CREATE_MODEL_CRUD)

	Settings.SERVICE_NAME = os.Getenv("SERVICE_NAME")
	Settings.READY_FOLDERNAME = strings.ToLower(Settings.SERVICE_NAME)

	if Settings.TEMPLATE_FOLDERNAME == "" {
		log.Panic("Need fill TEMPLATE_FOLDERNAME")
	}

	if Settings.TEXT_TEMPLATE_MODEL == "" {
		Settings.TEXT_TEMPLATE_MODEL = "LawsuitStatusType"
	}

	if Settings.TEXT_TEMPLATE_TABLENAME == "" {
		Settings.TEXT_TEMPLATE_TABLENAME = "lawsuit_status_types"
	}

	sUSE_DEFAULT_TEMPLATE := os.Getenv("USE_DEFAULT_TEMPLATE")
	Settings.USE_DEFAULT_TEMPLATE = BoolFromString(sUSE_DEFAULT_TEMPLATE)

	HAS_IS_DELETED := BoolFromString(sHAS_IS_DELETED)
	Settings.HAS_IS_DELETED = HAS_IS_DELETED
	Settings.COMMENT_MODEL_STRUCT = os.Getenv("COMMENT_MODEL_STRUCT")

	Settings.TEXT_MODULE_GENERATED = os.Getenv("TEXT_MODULE_GENERATED")

	sNEED_CREATE_DB_TEST := os.Getenv("NEED_CREATE_DB_TEST")
	NEED_CREATE_DB_TEST := BoolFromString(sNEED_CREATE_DB_TEST)
	Settings.NEED_CREATE_DB_TEST = NEED_CREATE_DB_TEST

	sNEED_CREATE_GRPC_SERVER_TEST := os.Getenv("NEED_CREATE_GRPC_SERVER_TEST")
	NEED_CREATE_GRPC_SERVER_TEST := BoolFromString(sNEED_CREATE_GRPC_SERVER_TEST)
	Settings.NEED_CREATE_GRPC_SERVER_TEST = NEED_CREATE_GRPC_SERVER_TEST

	sNEED_CREATE_GRPC_CLIENT_TEST := os.Getenv("NEED_CREATE_GRPC_CLIENT_TEST")
	NEED_CREATE_GRPC_CLIENT_TEST := BoolFromString(sNEED_CREATE_GRPC_CLIENT_TEST)
	Settings.NEED_CREATE_GRPC_CLIENT_TEST = NEED_CREATE_GRPC_CLIENT_TEST

	sNEED_CREATE_NRPC_SERVER_TEST := os.Getenv("NEED_CREATE_NRPC_SERVER_TEST")
	NEED_CREATE_NRPC_TEST := BoolFromString(sNEED_CREATE_NRPC_SERVER_TEST)
	Settings.NEED_CREATE_NRPC_SERVER_TEST = NEED_CREATE_NRPC_TEST

	sNEED_CREATE_NRPC_CLIENT_TEST := os.Getenv("NEED_CREATE_NRPC_CLIENT_TEST")
	NEED_CREATE_NRPC_CLIENT_TEST := BoolFromString(sNEED_CREATE_NRPC_CLIENT_TEST)
	Settings.NEED_CREATE_NRPC_CLIENT_TEST = NEED_CREATE_NRPC_CLIENT_TEST

	Settings.TEMPLATE_FOLDERNAME_MAIN = os.Getenv("TEMPLATE_FOLDERNAME_MAIN")
	Settings.TEMPLATE_REPOSITORY_URL = os.Getenv("TEMPLATE_REPOSITORY_URL")
	Settings.TEMPLATE_FOLDERNAME_GRPC = os.Getenv("TEMPLATE_FOLDERNAME_GRPC")
	Settings.PREFIX_CRUD = os.Getenv("PREFIX_CRUD")
	Settings.PREFIX_TABLE = os.Getenv("PREFIX_TABLE")
	Settings.TEMPLATE_FOLDERNAME_TABLES = os.Getenv("TEMPLATE_FOLDERNAME_TABLES")
	Settings.PREFIX_MODEL = os.Getenv("PREFIX_MODEL")
	Settings.TEMPLATE_FOLDERNAME_DB = os.Getenv("TEMPLATE_FOLDERNAME_DB")
	Settings.TEMPLATE_FOLDER_CRUD_FUNCTIONS = os.Getenv("TEMPLATE_FOLDER_CRUD_FUNCTIONS")

}

// CurrentDirectory - возвращает текущую директорию ОС
func CurrentDirectory() string {
	Otvet, err := os.Getwd()
	if err != nil {
		//log.Println(err)
	}

	return Otvet
}

// FillFlags - заполняет параметры из командной строки
func FillFlags() {
	Args := os.Args[1:]
	if len(Args) > 1 {
		return
	}

}

// BoolFromString - возвращает true если строка = true, или =1
func BoolFromString(s string) bool {
	Otvet := false

	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	s = strings.ToLower(s)

	if s == "true" || s == "1" {
		Otvet = true
	}

	return Otvet
}

func LoadSettingsTxt() {
	var err error

	Fill_TEMPLATES_FOLDER_NAME()

	DirBin := micro.ProgramDir_bin()
	Dir := DirBin + constants.TEMPLATES_FOLDER_NAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile()
	FilenameEnv := Dir + ".env"
	err = ConfigMain.LoadEnv_from_file_err(FilenameEnv)
	if err == nil {
		return
	}

	FilenameSettings := Dir + "settings.txt"
	err = ConfigMain.LoadEnv_from_file_err(FilenameSettings)
	if err != nil {
		log.Panic("LoadSettingsTxt() filename: ", FilenameSettings, " error: ", err)
	}

}

// Fill_TEMPLATES_FOLDER_NAME - заполняет переменную TEMPLATES_FOLDER_NAME = "templates_main" или "templates"
func Fill_TEMPLATES_FOLDER_NAME() {
	DirBin := micro.ProgramDir_bin()
	FileName := DirBin + "templates_main"
	ok, err := micro.FileExists(FileName)
	if err != nil {
		log.Panic("FileExists() ", FileName, " error: ", err)
	}
	if ok == true {
		constants.TEMPLATES_FOLDER_NAME = "templates_main"
	}
	log.Info("TEMPLATES_FOLDER_NAME = ", constants.TEMPLATES_FOLDER_NAME)
}
