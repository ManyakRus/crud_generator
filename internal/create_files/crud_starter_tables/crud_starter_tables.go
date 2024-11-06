package crud_starter_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
)

// CreateAllFiles - создаёт все файлы в папке db
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	MassAll := micro.MassFrom_Map(MapAll)
	for _, Table1 := range MassAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы starter
		if config.Settings.NEED_CREATE_DB == true {
			err = CreateFiles(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//тестовые файлы starter
		if config.Settings.NEED_CREATE_DB_TEST == true {
			err = CreateFiles_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		// создание файла manual
		if config.Settings.NEED_CREATE_MANUAL_FILES == true {
			err = CreateFiles_manual(Table1)
			if err != nil {
				log.Error("CreateFilesModel_manual() table: ", Table1.Name, " error: ", err)
				return err
			}
			err = CreateFiles_manual_test(Table1)
			if err != nil {
				log.Error("CreateFiles_manual_test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//ReadObject
		if config.Settings.NEED_CREATE_READOBJECT == true {
			err = CreateFiles_ReadObject(Table1)
			if err != nil {
				log.Error("CreateFiles_ReadObject() table: ", Table1.Name, " error: ", err)
				return err
			}
			err = CreateFiles_ReadObject_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_ReadObject_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}
	}

	return err
}
