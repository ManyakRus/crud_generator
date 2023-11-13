package main

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/load_configs"
	"github.com/ManyakRus/crud_generator/internal/logic"
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadEnv()
	config.FillSettings()
	config.FillFlags()

	load_configs.LoadConfigsAll()

	postgres_gorm.StartDB()
	postgres_gorm.GetConnection().Logger.LogMode(1)

	folders.CreateAllFolders()

	log.Info("postgres host: ", postgres_gorm.Settings.DB_HOST)
	err := logic.StartFillAll()
	if err != nil {
		println(constants.TEXT_HELP)
	}

}