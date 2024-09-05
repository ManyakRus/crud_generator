package entities_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
)

// CreateAllFiles - создаёт все файлы в папке model
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	//для каждой таблицы
	for _, Table1 := range MapAll {
		//проверка имени таблицы "DELETED_"
		err1 := create_files.IsGoodTableNamePrefix(Table1)
		if err1 != nil {
			log.Warn("CreateFiles() table: ", Table1.Name, " warning: ", err)
			continue
		}

		//создание файлов
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			err = CreateFilesUpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}

			//
			err = CreateFilesFindBy(Table1)
			if err != nil {
				log.Error("CreateFilesFindBy() table: ", Table1.Name, " error: ", err)
				return err
			}

			//
			err = CreateFilesFindMassBy(Table1)
			if err != nil {
				log.Error("CreateFilesFindMassBy() table: ", Table1.Name, " error: ", err)
				return err
			}

		}
	}

	return err
}
