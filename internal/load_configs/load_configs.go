package load_configs

import (
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
)

func LoadConfigsAll() {
	LoadMappings()
	LoadNameReplace()
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
