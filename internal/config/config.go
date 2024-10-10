package config

import (
	"github.com/ManyakRus/crud_generator/internal/constants"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"io/fs"
	"os"
	"strconv"
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
	TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC             string
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
	TEMPLATES_NAME_PRIMARYKEYS_FILENAME              string
	TEMPLATES_NULLABLE_FILENAME                      string
	TEMPLATES_CRUD_TABLE_UPDATE_FILENAME             string
	TEMPLATES_CRUD_TABLE_UPDATE_TEST_FILENAME        string
	TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_FILENAME      string
	TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_TEST_FILENAME string
	TEMPLATES_GRPC_SERVER_TABLE_UPDATE_FILENAME      string
	TEMPLATES_GRPC_SERVER_TABLE_UPDATE_TEST_FILENAME string
	TEMPLATES_GRPC_CLIENT_TABLES_CACHE_FILENAME      string
	TEMPLATES_GRPC_CLIENT_TABLES_CACHE_TEST_FILENAME string
	TEMPLATE_GRPC_CLIENT_FUNC_FILENAME               string
	TEMPLATE_GRPC_CLIENT_FUNC_TEST_FILENAME          string
	TEMPLATES_README_MD_FILENAME                     string
	TEMPLATES_README_RUS_FILENAME                    string
	TEMPLATE_FOLDERNAME_CONSTANTS                    string
	TEMPLATES_CONSTANTS_FILENAME                     string
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
	CACHE_ELEMENTS_COUNT_MAX                         int64
	SINGULAR_TABLE_NAMES                             bool

	//---FINDBY---
	TEMPLATES_FINDBY_FILENAME                                  string
	TEMPLATES_CRUD_TABLE_FINDBY_FILENAME                       string
	TEMPLATES_CRUD_TABLE_FINDBY_TEST_FILENAME                  string
	TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_FILENAME              string
	TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_TEST_FILENAME         string
	TEMPLATES_GRPC_SERVER_FINDBY_FILENAME                      string
	TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_FILENAME             string
	TEMPLATES_GRPC_SERVER_FINDBY_TEST_FILENAME                 string
	TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_TEST_FILENAME        string
	TEMPLATES_MODEL_FINDBY_FILENAME                            string
	TEMPLATES_MODEL_FINDBY_FUNCTION_FILENAME                   string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FILENAME               string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_TEST_FILENAME          string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_FILENAME      string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_TEST_FILENAME string

	//---FINDMASSBY---
	TEMPLATES_FINDMASSBY_FILENAME                                  string
	TEMPLATES_CRUD_TABLE_FINDMASSBY_FILENAME                       string
	TEMPLATES_CRUD_TABLE_FINDMASSBY_TEST_FILENAME                  string
	TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_FILENAME              string
	TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_TEST_FILENAME         string
	TEMPLATES_GRPC_SERVER_FINDMASSBY_FILENAME                      string
	TEMPLATES_GRPC_SERVER_FINDMASSBY_FUNCTION_FILENAME             string
	TEMPLATES_GRPC_SERVER_FINDMASSBY_TEST_FILENAME                 string
	TEMPLATES_GRPC_SERVER_FINDMASSBY_FUNCTION_TEST_FILENAME        string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FILENAME               string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_TEST_FILENAME          string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_FILENAME      string
	TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_TEST_FILENAME string
	TEMPLATES_MODEL_FINDMASSBY_FILENAME                            string
	TEMPLATES_MODEL_FINDMASSBY_FUNCTION_FILENAME                   string

	//---READALL---
	TEMPLATES_READALL_FILENAME                                  string
	TEMPLATES_CRUD_TABLE_READALL_FILENAME                       string
	TEMPLATES_CRUD_TABLE_READALL_TEST_FILENAME                  string
	TEMPLATES_CRUD_TABLE_READALL_FUNCTION_FILENAME              string
	TEMPLATES_CRUD_TABLE_READALL_FUNCTION_TEST_FILENAME         string
	TEMPLATES_GRPC_SERVER_READALL_FILENAME                      string
	TEMPLATES_GRPC_SERVER_READALL_FUNCTION_FILENAME             string
	TEMPLATES_GRPC_SERVER_READALL_TEST_FILENAME                 string
	TEMPLATES_GRPC_SERVER_READALL_FUNCTION_TEST_FILENAME        string
	TEMPLATES_GRPC_CLIENT_TABLES_READALL_FILENAME               string
	TEMPLATES_GRPC_CLIENT_TABLES_READALL_TEST_FILENAME          string
	TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_FILENAME      string
	TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_TEST_FILENAME string
	TEMPLATES_MODEL_READALL_FILENAME                            string
	TEMPLATES_MODEL_READALL_FUNCTION_FILENAME                   string

	TEMPLATES_CRUD_FUNC_FOLDERNAME    string
	TEMPLATES_CRUD_FUNC_FILENAME      string
	TEMPLATES_CRUD_FUNC_TEST_FILENAME string

	//
	FILE_PERMISSIONS fs.FileMode //= 0666

	GENERATION_PROTO_FILENAME string

	GRPC_CLIENT_FILENAME      string
	GRPC_CLIENT_TEST_FILENAME string

	NRPC_CLIENT_FILENAME      string
	NRPC_CLIENT_TEST_FILENAME string

	NRPC_CLIENT_TABLE_FILENAME      string
	NRPC_CLIENT_TABLE_TEST_FILENAME string

	SERVER_GRPC_STARTER_FILENAME string
	SERVER_GRPC_FUNC_FILENAME    string

	MAKEFILE_FILENAME string
	ENV_FILENAME      string

	STARTER_TABLES_FILENAME             string
	STARTER_TABLES_TEST_FILENAME        string
	STARTER_TABLES_MANUAL_FILENAME      string
	STARTER_TABLES_TEST_MANUAL_FILENAME string
	STARTER_TABLES_PREFIX               string
	CRUD_TABLES_FREFIX                  string

	MODEL_TABLE_MANUAL_FILENAME string
	MODEL_TABLE_UPDATE_FILENAME string

	SERVER_GRPC_TABLE_UPDATE_FUNC_FILENAME      string
	SERVER_GRPC_TABLE_UPDATE_FUNC_TEST_FILENAME string

	GRPC_CLIENT_TABLE_UPDATE_FUNC_FILENAME      string
	GRPC_CLIENT_TABLE_UPDATE_FUNC_TEST_FILENAME string

	CRUD_TABLES_CACHE_FILENAME      string
	CRUD_TABLES_CACHE_TEST_FILENAME string

	SERVER_GRPC_TABLE_CACHE_FILENAME      string
	SERVER_GRPC_TABLE_CACHE_TEST_FILENAME string

	TEXT_READALL string

	NEED_USE_DB_VIEWS bool
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
	s = Getenv(Name, false)
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

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_CACHE_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_CACHE_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_CACHE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_CACHE_TEST_FILENAME = s

	//
	Name = "CACHE_ELEMENTS_COUNT_MAX"
	s = Getenv(Name, true)
	i, err := micro.Int64FromString(s)
	if err != nil {
		log.Error("CACHE_ELEMENTS_COUNT_MAX: ", s, " Int64FromString() error: ", err)
	}
	Settings.CACHE_ELEMENTS_COUNT_MAX = i

	//
	Name = "TEMPLATES_README_MD_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_README_MD_FILENAME = s

	//
	Name = "TEMPLATES_README_RUS_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_README_RUS_FILENAME = s

	//
	Name = "TEMPLATE_FOLDERNAME_CONSTANTS"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_CONSTANTS = s

	//
	Name = "TEMPLATES_CONSTANTS_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CONSTANTS_FILENAME = s

	//
	Name = "SINGULAR_TABLE_NAMES"
	s = Getenv(Name, true)
	Settings.SINGULAR_TABLE_NAMES = BoolFromString(s)

	//
	Name = "TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC"
	s = Getenv(Name, true)
	Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC = s

	//
	Name = "TEMPLATE_GRPC_CLIENT_FUNC_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATE_GRPC_CLIENT_FUNC_FILENAME = s

	//
	Name = "TEMPLATE_GRPC_CLIENT_FUNC_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATE_GRPC_CLIENT_FUNC_TEST_FILENAME = s

	//
	Name = "TEMPLATES_FINDBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_FINDBY_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDBY_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDBY_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDBY_TEST_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDBY_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDBY_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDBY_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_TEST_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_FINDBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_FINDBY_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_FINDBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_FINDBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_TEST_FILENAME = s

	//-----------------FINDMASSBY---------------------------
	//
	Name = "TEMPLATES_FINDMASSBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_FINDMASSBY_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDMASSBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDMASSBY_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDMASSBY_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDMASSBY_TEST_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDMASSBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDMASSBY_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDMASSBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDMASSBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDMASSBY_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDMASSBY_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_FINDMASSBY_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_FINDMASSBY_FUNCTION_TEST_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_FINDMASSBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_FINDMASSBY_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_FINDMASSBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_FINDMASSBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_TEST_FILENAME = s

	//-----------------READALL---------------------------
	//
	Name = "TEMPLATES_READALL_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_READALL_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_READALL_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_READALL_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_READALL_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_READALL_TEST_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_READALL_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_READALL_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_TABLE_READALL_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_TABLE_READALL_FUNCTION_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_READALL_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_READALL_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_READALL_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_READALL_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_READALL_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_READALL_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_SERVER_READALL_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_SERVER_READALL_FUNCTION_TEST_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_READALL_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_READALL_FILENAME = s

	//
	Name = "TEMPLATES_MODEL_READALL_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_MODEL_READALL_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_READALL_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_READALL_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_TEST_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_FILENAME = s

	//
	Name = "TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_TEST_FILENAME = s

	//----------------CRUD_FUNC----------------------------
	//
	Name = "TEMPLATES_CRUD_FUNC_FOLDERNAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_FUNC_FOLDERNAME = s

	//
	Name = "TEMPLATES_CRUD_FUNC_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_FUNC_FILENAME = s

	//
	Name = "TEMPLATES_CRUD_FUNC_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_CRUD_FUNC_TEST_FILENAME = s

	//--------------------------------------------

	//
	Name = "FILE_PERMISSIONS"
	s = Getenv(Name, true)
	x, err := strconv.Atoi(s)
	if err != nil {
		x = 0666
		log.Error("FILE_PERMISSIONS error: ", err)
	}
	Settings.FILE_PERMISSIONS = fs.FileMode(x)

	//
	Name = "GENERATION_PROTO_FILENAME"
	s = Getenv(Name, true)
	Settings.GENERATION_PROTO_FILENAME = s

	//
	Name = "GRPC_CLIENT_FILENAME"
	s = Getenv(Name, true)
	Settings.GRPC_CLIENT_FILENAME = s

	//
	Name = "GRPC_CLIENT_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.GRPC_CLIENT_TEST_FILENAME = s

	//
	Name = "NRPC_CLIENT_FILENAME"
	s = Getenv(Name, true)
	Settings.NRPC_CLIENT_FILENAME = s

	//
	Name = "NRPC_CLIENT_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.NRPC_CLIENT_TEST_FILENAME = s

	//
	Name = "NRPC_CLIENT_TABLE_FILENAME"
	s = Getenv(Name, true)
	Settings.NRPC_CLIENT_TABLE_FILENAME = s

	//
	Name = "NRPC_CLIENT_TABLE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.NRPC_CLIENT_TABLE_TEST_FILENAME = s

	//
	Name = "SERVER_GRPC_STARTER_FILENAME"
	s = Getenv(Name, true)
	Settings.SERVER_GRPC_STARTER_FILENAME = s

	//
	Name = "SERVER_GRPC_FUNC_FILENAME"
	s = Getenv(Name, true)
	Settings.SERVER_GRPC_FUNC_FILENAME = s

	//
	Name = "MAKEFILE_FILENAME"
	s = Getenv(Name, true)
	Settings.MAKEFILE_FILENAME = s

	//
	Name = "ENV_FILENAME"
	s = Getenv(Name, true)
	Settings.ENV_FILENAME = s

	//
	Name = "STARTER_TABLES_FILENAME"
	s = Getenv(Name, true)
	Settings.STARTER_TABLES_FILENAME = s

	//
	Name = "STARTER_TABLES_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.STARTER_TABLES_TEST_FILENAME = s

	//
	Name = "STARTER_TABLES_MANUAL_FILENAME"
	s = Getenv(Name, true)
	Settings.STARTER_TABLES_MANUAL_FILENAME = s

	//
	Name = "STARTER_TABLES_TEST_MANUAL_FILENAME"
	s = Getenv(Name, true)
	Settings.STARTER_TABLES_TEST_MANUAL_FILENAME = s

	//
	Name = "STARTER_TABLES_PREFIX"
	s = Getenv(Name, true)
	Settings.STARTER_TABLES_PREFIX = s

	//
	Name = "CRUD_TABLES_FREFIX"
	s = Getenv(Name, true)
	Settings.CRUD_TABLES_FREFIX = s

	//
	Name = "MODEL_TABLE_MANUAL_FILENAME"
	s = Getenv(Name, true)
	Settings.MODEL_TABLE_MANUAL_FILENAME = s

	//
	Name = "MODEL_TABLE_UPDATE_FILENAME"
	s = Getenv(Name, true)
	Settings.MODEL_TABLE_UPDATE_FILENAME = s

	//
	Name = "SERVER_GRPC_TABLE_UPDATE_FUNC_FILENAME"
	s = Getenv(Name, true)
	Settings.SERVER_GRPC_TABLE_UPDATE_FUNC_FILENAME = s

	//
	Name = "SERVER_GRPC_TABLE_UPDATE_FUNC_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.SERVER_GRPC_TABLE_UPDATE_FUNC_TEST_FILENAME = s

	//
	Name = "GRPC_CLIENT_TABLE_UPDATE_FUNC_FILENAME"
	s = Getenv(Name, true)
	Settings.GRPC_CLIENT_TABLE_UPDATE_FUNC_FILENAME = s

	//
	Name = "GRPC_CLIENT_TABLE_UPDATE_FUNC_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.GRPC_CLIENT_TABLE_UPDATE_FUNC_TEST_FILENAME = s

	//
	Name = "CRUD_TABLES_CACHE_FILENAME"
	s = Getenv(Name, true)
	Settings.CRUD_TABLES_CACHE_FILENAME = s

	//
	Name = "CRUD_TABLES_CACHE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.CRUD_TABLES_CACHE_TEST_FILENAME = s

	//
	Name = "SERVER_GRPC_TABLE_CACHE_FILENAME"
	s = Getenv(Name, true)
	Settings.SERVER_GRPC_TABLE_CACHE_FILENAME = s

	//
	Name = "SERVER_GRPC_TABLE_CACHE_TEST_FILENAME"
	s = Getenv(Name, true)
	Settings.SERVER_GRPC_TABLE_CACHE_TEST_FILENAME = s

	//
	Name = "TEXT_READALL"
	s = Getenv(Name, true)
	Settings.TEXT_READALL = s

	//
	Name = "NEED_USE_DB_VIEWS"
	s = Getenv(Name, true)
	Settings.NEED_USE_DB_VIEWS = BoolFromString(s)

	//
	Name = "TEMPLATES_NAME_PRIMARYKEYS_FILENAME"
	s = Getenv(Name, true)
	Settings.TEMPLATES_NAME_PRIMARYKEYS_FILENAME = s

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
