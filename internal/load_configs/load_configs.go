package load_configs

import (
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
)

func LoadConfigsAll() {
	LoadMappings()
	LoadNameReplace()
	LoadNullable()
	Load_TEXT_DB_MODIFIED_AT()
	Load_TEXT_DB_IS_DELETED()
	LoadConvertID()
}

// LoadMappings - загружает маппинг ТипБД = ТипGolang, из файла .json
func LoadMappings() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + "configs" + micro.SeparatorFile() + "mapping.json"
	err := dbmeta.LoadMappings(FileName, false)
	if err != nil {
		log.Panic("LoadMappings() error: ", err)
	}
}

// LoadNameReplace - загружает маппинг ТипБД = ТипGolang, из файла .json
func LoadNameReplace() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + "configs" + micro.SeparatorFile() + "name_replace.json"

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	//var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &types.MapReplaceName)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}

// LoadNullable - загружает список полей которые могут быть null
func LoadNullable() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + "configs" + micro.SeparatorFile() + "nullable.json"

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	//var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &types.MapNullableFileds)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}

// Load_TEXT_DB_MODIFIED_AT - загружает текст DB_MODIFIED_AT
func Load_TEXT_DB_MODIFIED_AT() {
	DirTemplatesDB := create_files.Find_Template_DB_Foldername()
	FileName := DirTemplatesDB + "modified_at.go_"

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Error(TextError)
	}

	config.Settings.TEXT_DB_MODIFIED_AT = string(bytes)

}

// Load_TEXT_DB_IS_DELETED - загружает текст DB_IS_DELETED
func Load_TEXT_DB_IS_DELETED() {
	DirTemplatesDB := create_files.Find_Template_DB_Foldername()
	FileName := DirTemplatesDB + "is_deleted.go_"

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Error(TextError)
	}

	config.Settings.TEXT_DB_IS_DELETED = string(bytes)

}

// LoadConvertID - загружает map: ИмяТаблицы:Тип
func LoadConvertID() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + "configs" + micro.SeparatorFile() + "convert_id.json"

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	//var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &types.MapConvertID)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}
