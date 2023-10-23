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
	INCLUDE_TABLES            string
	EXCLUDE_TABLES            string
	TEMPLATE_FOLDERNAME       string
	TEMPLATE_FOLDERNAME_MODEL string
	TEMPLATE_FOLDERNAME_DB    string
	TEMPLATE_FOLDERNAME_GRPC  string
	TEMPLATE_FOLDERNAME_NRPC  string
	NEED_CRUD                 bool
	NEED_GRPC                 bool
	NEED_NRPC                 bool
	SERVICE_NAME              string
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
	Settings.TEMPLATE_FOLDERNAME_NRPC = os.Getenv("TEMPLATE_FOLDERNAME_NRPC")

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
	//
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
