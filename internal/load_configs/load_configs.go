package load_configs

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
)

// LoadMappingsAll - загружает маппинг ТипБД = ТипGolang, из файла .json
func LoadMappingsAll() {
	dir := micro.ProgramDir_bin()
	Filename := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + "configs" + micro.SeparatorFile() + "mapping.json"
	err := dbmeta.LoadMappings(Filename, false)
	if err != nil {
		log.Panic("LoadMappings() error: ", err)
	}
}
