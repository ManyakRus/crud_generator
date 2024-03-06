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
	INCLUDE_TABLES                                   string
	EXCLUDE_TABLES                                   string
	TEMPLATE_FOLDERNAME                              string
	TEMPLATE_FOLDERNAME_MODEL                        string
	TEMPLATE_FOLDERNAME_DB                           string
	TEMPLATE_FOLDERNAME_CRUD                         string
	TEMPLATE_FOLDERNAME_TABLES                       string
	TEMPLATE_FOLDERNAME_GRPC                         string
	TEMPLATE_FOLDERNAME_GRPC_PROTO                   string
	TEMPLATE_FOLDERNAME_GRPC_SERVER                  string
	TEMPLATE_FOLDERNAME_GRPC_CLIENT                  string
	TEMPLATE_FOLDERNAME_NRPC                         string
	TEMPLATE_FOLDERNAME_NRPC_SERVER                  string
	TEMPLATE_FOLDERNAME_NRPC_CLIENT                  string
	TEMPLATE_FOLDERNAME_GRPC_NRPC                    string
	TEMPLATE_SERVICE_NAME                            string
	TEMPLATE_FOLDERNAME_CRUD_STARTER                 string
	TEMPLATE_FOLDERNAME_ALIAS                        string
	TEMPLATE_FOLDER_CRUD_FUNCTIONS                   string
	TEMPLATE_FOLDERNAME_MAIN                         string
	TEMPLATE_REPOSITORY_URL                          string
	TEMPLATE_EXTERNAL_PROTO_FILENAME                 string
	TEMPLATES_CRUD_FILENAME                          string
	TEMPLATES_CRUD_TEST_FILENAME                     string
	TEMPLATES_ALIAS_FILENAME                         string
	TEMPLATES_CRUD_TABLE_UPDATE_FUNC_FILENAME        string
	TEMPLATES_CRUD_TABLE_UPDATE_FUNC_TEST_FILENAME   string
	TEMPLATES_CONVERT_ID_FILENAME                    string
	TEMPLATES_CRUD_FUNCTIONS_RENAME_FILENAME         string
	TEMPLATES_MODEL_CRUD_DELETE_FUNCTIONS_FILENAME   string
	TEMPLATES_NAME_REPLACE_FILENAME                  string
	TEMPLATES_NULLABLE_FILENAME                      string
	TEMPLATES_CRUD_TABLE_UPDATE_FILENAME             string
	TEMPLATES_CRUD_TABLE_UPDATE_TEST_FILENAME        string
	TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_FILENAME      string
	TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_TEST_FILENAME string
	TEMPLATES_GRPC_SERVER_TABLE_UPDATE_FILENAME      string
	TEMPLATES_GRPC_SERVER_TABLE_UPDATE_TEST_FILENAME string
	NEED_CREATE_MODEL_STRUCT                         bool
	NEED_CREATE_MODEL_CRUD                           bool
	NEED_CREATE_DB                                   bool
	NEED_CREATE_DB_TEST                              bool
	NEED_CREATE_DB_TABLES                            bool
	NEED_CREATE_GRPC                                 bool
	NEED_CREATE_GRPC_SERVER_TEST                     bool
	NEED_CREATE_GRPC_CLIENT_TEST                     bool
	NEED_CREATE_NRPC                                 bool
	NEED_CREATE_NRPC_SERVER_TEST                     bool
	NEED_CREATE_NRPC_CLIENT_TEST                     bool
	NEED_CREATE_MANUAL_FILES                         bool
	NEED_CREATE_CACHE_API                            bool
	NEED_CREATE_CACHE_FILES                          bool
	NEED_CREATE_CACHE_TEST_FILES                     bool
	SERVICE_NAME                                     string
	SERVICE_REPOSITORY_URL                           string
	TEXT_TEMPLATE_MODEL                              string
	TEXT_TEMPLATE_TABLENAME                          string
	HAS_IS_DELETED                                   bool
	READY_FOLDERNAME                                 string
	TEXT_DB_MODIFIED_AT                              string
	TEXT_DB_CREATED_AT                               string
	TEXT_DB_IS_DELETED                               string
	USE_DEFAULT_TEMPLATE                             bool
	PREFIX_SERVER_GRPC                               string
	PREFIX_CLIENT_GRPC                               string
	PREFIX_CRUD                                      string
	PREFIX_MODEL                                     string
	PREFIX_TABLE                                     string
	COMMENT_MODEL_STRUCT                             string
	TEXT_MODULE_GENERATED                            string
	READY_ALIAS_FILENAME                             string
	NEED_CREATE_UPDATE_EVERY_COLUMN                  bool
}

func Getenv(Name string, IsRequired bool) string {
	TextError := "Need fill OS environment variable: "
	Otvet := os.Getenv(Name)
	if IsRequired == true && Otvet == "" {
		log.Error(TextError + Name)
	}

	return Otvet
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	//dir := micro.ProgramDir_bin()

	Settings = SettingsINI{}

	//Заполнение переменных окружения
	Name := ""
	s := ""
	//
	Name = "INCLUDE_TABLES"
	s = Getenv(Name, false)
	Settings.INCLUDE_TABLES = s

	//
	Name = "EXCLUDE_TABLES"
	s = Getenv(Name, false)
	Settings.EXCLUDE_TABLES = s

	//
	Name = "TEMPLATE_FOLDERNAME"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME = s

	//
	Name = "TEMPLATE_FOLDERNAME_MODEL"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_MODEL = s

	//
	Name = "TEMPLATE_FOLDERNAME_CRUD"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_CRUD = s

	//
	Name = "TEMPLATE_FOLDERNAME_GRPC_PROTO"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO = s

	//
	Name = "TEMPLATE_FOLDERNAME_GRPC_SERVER"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER = s

	//
	Name = "TEMPLATE_FOLDERNAME_GRPC_CLIENT"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT = s

	//
	Name = "TEMPLATE_FOLDERNAME_NRPC"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_NRPC = s

	//
	Name = "TEMPLATE_FOLDERNAME_NRPC_SERVER"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER = s

	//
	Name = "TEMPLATE_FOLDERNAME_NRPC_CLIENT"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT = s

	//
	Name = "TEXT_TEMPLATE_MODEL"
	s = Getenv(Name, true)
	Settings.TEXT_TEMPLATE_MODEL = s

	//
	Name = "TEXT_TEMPLATE_TABLENAME"
	s = Getenv(Name, true)
	Settings.TEXT_TEMPLATE_TABLENAME = s

	//
	Name = "TEMPLATE_SERVICE_NAME"
	s = Getenv(Name, true)
	Settings.TEMPLATE_SERVICE_NAME = s

	//
	Name = "TEMPLATE_FOLDERNAME_CRUD_STARTER"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER = s

	//
	Name = "SERVICE_REPOSITORY_URL"
	s = Getenv(Name, true)
	s = strings.ToLower(s)
	Settings.SERVICE_REPOSITORY_URL = s

	//
	Name = "TEMPLATE_FOLDERNAME_ALIAS"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_ALIAS = s

	//
	Name = "PREFIX_SERVER_GRPC"
	s = Getenv(Name, true)
	Settings.PREFIX_SERVER_GRPC = s

	//
	Name = "HAS_IS_DELETED"
	s = Getenv(Name, true)
	Settings.HAS_IS_DELETED = BoolFromString(s)

	//
	Name = "NEED_CREATE_DB"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_DB = BoolFromString(s)

	//
	Name = "NEED_CREATE_DB_TABLES"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_DB_TABLES = BoolFromString(s)

	//
	Name = "NEED_CREATE_GRPC"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_GRPC = BoolFromString(s)

	//
	Name = "NEED_CREATE_NRPC"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_NRPC = BoolFromString(s)

	//
	Name = "NEED_CREATE_MODEL_STRUCT"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_MODEL_STRUCT = BoolFromString(s)

	//
	Name = "NEED_CREATE_MODEL_CRUD"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_MODEL_CRUD = BoolFromString(s)

	//
	Name = "SERVICE_NAME"
	s = Getenv(Name, true)
	Settings.SERVICE_NAME = s

	//
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

	//
	Name = "USE_DEFAULT_TEMPLATE"
	s = Getenv(Name, true)
	Settings.USE_DEFAULT_TEMPLATE = BoolFromString(s)

	//
	Name = "HAS_IS_DELETED"
	s = Getenv(Name, true)
	Settings.HAS_IS_DELETED = BoolFromString(s)

	//
	Name = "COMMENT_MODEL_STRUCT"
	s = Getenv(Name, true)
	Settings.COMMENT_MODEL_STRUCT = s

	//
	Name = "TEXT_MODULE_GENERATED"
	s = Getenv(Name, true)
	Settings.TEXT_MODULE_GENERATED = s

	//
	Name = "NEED_CREATE_DB_TEST"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_DB_TEST = BoolFromString(s)

	//
	Name = "NEED_CREATE_GRPC_SERVER_TEST"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_GRPC_SERVER_TEST = BoolFromString(s)

	//
	Name = "NEED_CREATE_GRPC_CLIENT_TEST"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_GRPC_CLIENT_TEST = BoolFromString(s)

	//
	Name = "NEED_CREATE_NRPC_SERVER_TEST"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_NRPC_SERVER_TEST = BoolFromString(s)

	//
	Name = "NEED_CREATE_NRPC_CLIENT_TEST"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_NRPC_CLIENT_TEST = BoolFromString(s)

	//
	Name = "TEMPLATE_FOLDERNAME_MAIN"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_MAIN = s

	//
	Name = "TEMPLATE_REPOSITORY_URL"
	s = Getenv(Name, true)
	Settings.TEMPLATE_REPOSITORY_URL = s

	//
	Name = "TEMPLATE_FOLDERNAME_GRPC"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_GRPC = s

	//
	Name = "PREFIX_CRUD"
	s = Getenv(Name, true)
	Settings.PREFIX_CRUD = s

	//
	Name = "PREFIX_TABLE"
	s = Getenv(Name, true)
	Settings.PREFIX_TABLE = s

	//
	Name = "TEMPLATE_FOLDERNAME_TABLES"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_TABLES = s

	//
	Name = "PREFIX_MODEL"
	s = Getenv(Name, false)
	Settings.PREFIX_MODEL = s

	//
	Name = "TEMPLATE_FOLDERNAME_DB"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_DB = s

	//
	Name = "TEMPLATE_FOLDER_CRUD_FUNCTIONS"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDER_CRUD_FUNCTIONS = s

	//
	Name = "TEMPLATE_EXTERNAL_PROTO_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATE_EXTERNAL_PROTO_FILENAME = s

	//
	Name = "TEMPLATE_FOLDERNAME_GRPC_NRPC"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_GRPC_NRPC = s

	//
	Name = "NEED_CREATE_MANUAL_FILES"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_MANUAL_FILES = BoolFromString(s)

	//
	Name = "TEMPLATES_CRUD_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TEST_FILENAME = s

	//
	Name = "TEMPLATES_ALIAS_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_ALIAS_FILENAME = s

	//
	Name = "READY_ALIAS_FILENAME"
	s = Getenv(Name, true)
	Settings.READY_ALIAS_FILENAME = s

	//
	Name = "NEED_CREATE_UPDATE_EVERY_COLUMN"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_UPDATE_EVERY_COLUMN = BoolFromString(s)

	//
	Name = "PREFIX_CLIENT_GRPC"
	s = Getenv(Name, true)
	Settings.PREFIX_CLIENT_GRPC = s

	//
	Name = "TEMPLATES_CRUD_TABLE_UPDATE_FUNC_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_UPDATE_FUNC_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_UPDATE_FUNC_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_UPDATE_FUNC_TEST_FILENAME = s

	//
	Name = "TEMPLATES_CONVERT_ID_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CONVERT_ID_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_FUNCTIONS_RENAME_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_FUNCTIONS_RENAME_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_CRUD_DELETE_FUNCTIONS_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_CRUD_DELETE_FUNCTIONS_FILENAME = s

	//
	Name = "TEMPLATES_NAME_REPLACE_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_NAME_REPLACE_FILENAME = s

	//
	Name = "TEMPLATES_NULLABLE_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_NULLABLE_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_UPDATE_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_UPDATE_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_UPDATE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_UPDATE_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_TABLE_UPDATE_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_TABLE_UPDATE_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_TABLE_UPDATE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_TABLE_UPDATE_TEST_FILENAME = s

	//
	Name = "NEED_CREATE_CACHE_API"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_CACHE_API = BoolFromString(s)

	//
	Name = "NEED_CREATE_CACHE_FILES"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_CACHE_FILES = BoolFromString(s)

	//
	Name = "NEED_CREATE_CACHE_TEST_FILES"
	s = Getenv(Name, true)
	Settings.NEED_CREATE_CACHE_TEST_FILES = BoolFromString(s)

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
