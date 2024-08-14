package load_configs

import (
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
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
	Load_TEXT_DB_CREATED_AT()
	Load_TEXT_DB_IS_DELETED()
	LoadConvertID()
	LoadMapModelCrudDeleteFunctions()
	LoadCrudFunctionsRename()
	LoadFindBy()
	LoadFindMassBy()
}

// LoadMappings - загружает маппинг ТипБД = ТипGolang, из файла .json
func LoadMappings() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + "mapping.json"
	err := dbmeta.LoadMappings(FileName, false)
	if err != nil {
		log.Panic("LoadMappings() error: ", err)
	}
}

// LoadNameReplace - загружает маппинг ТипБД = ТипGolang, из файла .json
func LoadNameReplace() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_NAME_REPLACE_FILENAME

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
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_NULLABLE_FILENAME

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
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_CONVERT_ID_FILENAME

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

// LoadMapModelCrudDeleteFunctions - загружает map ИмяТаблицыPostgres:ИмяФункцииGolang
func LoadMapModelCrudDeleteFunctions() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_MODEL_CRUD_DELETE_FUNCTIONS_FILENAME

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	//var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &types.MapModelCrudDeleteFunctions)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}

// Load_TEXT_DB_CREATED_AT - загружает текст created_at.go_
func Load_TEXT_DB_CREATED_AT() {
	DirTemplatesDB := create_files.Find_Template_DB_Foldername()
	FileName := DirTemplatesDB + "created_at.go_"

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Error(TextError)
	}

	config.Settings.TEXT_DB_CREATED_AT = string(bytes)

}

// LoadCrudFunctionsRename - загружает маппинг ИмяТаблицы:{old:"",new:""}
func LoadCrudFunctionsRename() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_CRUD_FUNCTIONS_RENAME_FILENAME

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	//var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &types.MapRenameFunctions)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}

// LoadFindBy - загружает из файла .json список функций FindBy которые надо создать
func LoadFindBy() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_FINDBY_FILENAME

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	err = json.Unmarshal(bytes, &types.MapFindBy)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}

// LoadFindMassBy - загружает из файла .json список функций FindMassBy которые надо создать
func LoadFindMassBy() {
	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + config.Settings.TEMPLATES_FINDMASSBY_FILENAME

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Panic(TextError)
	}

	//json в map
	err = json.Unmarshal(bytes, &types.MapFindMassBy)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}
