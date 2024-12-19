package entities_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
)

// CreateAllFiles - создаёт все файлы в папке model
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	//для каждой таблицы
	MassAll := micro.MassFrom_Map(MapAll)
	for _, Table1 := range MassAll {
		//проверка имени таблицы "DELETED_"
		err1 := create_files.IsGood_TableName(Table1)
		if err1 != nil {
			log.Warn("CreateFiles() table: ", Table1.Name, " warning: ", err1)
			continue
		}

		//создание файлов
		err = CreateFiles(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			err = CreateFiles_UpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//
		err = CreateFiles_FindBy(Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//FindMassBy
		err = CreateFiles_FindMassBy(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//ReadAll
		err = CreateFiles_ReadAll(Table1)
		if err != nil {
			log.Error("CreateFiles_ReadAll() table: ", Table1.Name, " error: ", err)
			return err
		}

		//FindModelBy
		err = CreateFiles_FindModelBy(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_FindModelBy() table: ", Table1.Name, " error: ", err)
			return err
		}

	}

	return err
}
