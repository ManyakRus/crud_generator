package config

import (
	"log"
	"os"
	"strings"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	INCLUDE_TABLES                  string
	EXCLUDE_TABLES                  string
	TEMPLATE_FOLDERNAME             string
	TEMPLATE_FOLDERNAME_MODEL       string
	TEMPLATE_FOLDERNAME_DB          string
	TEMPLATE_FOLDERNAME_GRPC        string
	TEMPLATE_FOLDERNAME_GRPC_SERVER string
	TEMPLATE_FOLDERNAME_GRPC_CLIENT string
	TEMPLATE_FOLDERNAME_NRPC        string
	TEMPLATE_FOLDERNAME_NRPC_SERVER string
	TEMPLATE_FOLDERNAME_NRPC_CLIENT string
	TEMPLATE_SERVICE_NAME           string
	NEED_CRUD                       bool
	NEED_GRPC                       bool
	NEED_NRPC                       bool
	SERVICE_NAME                    string
	TEXT_TEMPLATE_MODEL             string
	TEXT_TEMPLATE_TABLENAME         string
	HAS_IS_DELETED                  bool
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.INCLUDE_TABLES = os.Getenv("INCLUDE_TABLES")
	Settings.EXCLUDE_TABLES = os.Getenv("EXCLUDE_TABLES")
	Settings.TEMPLATE_FOLDERNAME = os.Getenv("TEMPLATE_FOLDERNAME")
	Settings.TEMPLATE_FOLDERNAME_MODEL = os.Getenv("TEMPLATE_FOLDERNAME_MODEL")
	Settings.TEMPLATE_FOLDERNAME_DB = os.Getenv("TEMPLATE_FOLDERNAME_DB")
	Settings.TEMPLATE_FOLDERNAME_GRPC = os.Getenv("TEMPLATE_FOLDERNAME_GRPC")
	Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER = os.Getenv("TEMPLATE_FOLDERNAME_GRPC_SERVER")
	Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT = os.Getenv("TEMPLATE_FOLDERNAME_GRPC_CLIENT")
	Settings.TEMPLATE_FOLDERNAME_NRPC = os.Getenv("TEMPLATE_FOLDERNAME_NRPC")
	Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER = os.Getenv("TEMPLATE_FOLDERNAME_NRPC_SERVER")
	Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT = os.Getenv("TEMPLATE_FOLDERNAME_NRPC_CLIENT")
	Settings.TEXT_TEMPLATE_MODEL = os.Getenv("TEXT_TEMPLATE_MODEL")
	Settings.TEXT_TEMPLATE_TABLENAME = os.Getenv("TEXT_TEMPLATE_TABLENAME")
	Settings.TEMPLATE_SERVICE_NAME = os.Getenv("TEMPLATE_SERVICE_NAME")
	sHAS_IS_DELETED := os.Getenv("HAS_IS_DELETED")

	sNEED_CRUD := os.Getenv("NEED_CRUD")
	Settings.NEED_CRUD = BoolFromString(sNEED_CRUD)

	sNEED_GRPC := os.Getenv("NEED_GRPC")
	Settings.NEED_GRPC = BoolFromString(sNEED_GRPC)

	sNEED_NRPC := os.Getenv("NEED_NRPC")
	Settings.NEED_NRPC = BoolFromString(sNEED_NRPC)

	Settings.SERVICE_NAME = os.Getenv("SERVICE_NAME")

	if Settings.TEMPLATE_FOLDERNAME == "" {
		log.Panic("Need fill TEMPLATE_FOLDERNAME")
	}

	if Settings.TEXT_TEMPLATE_MODEL == "" {
		Settings.TEXT_TEMPLATE_MODEL = "LawsuitStatusType"
	}

	if Settings.TEXT_TEMPLATE_TABLENAME == "" {
		Settings.TEXT_TEMPLATE_TABLENAME = "lawsuit_status_types"
	}

	HAS_IS_DELETED := BoolFromString(sHAS_IS_DELETED)
	Settings.HAS_IS_DELETED = HAS_IS_DELETED
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
