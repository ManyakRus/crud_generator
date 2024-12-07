package object_tables

import (
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
			log.Warn("CreateFiles() table: ", Table1.Name, ", warning: ", err1)
			continue
		}

		//model
		err = CreateFiles(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, ", error: ", err)
			return err
		}

		//crud
		err = CreateFiles_crud(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, ", error: ", err)
			return err
		}

	}

	return err
}
