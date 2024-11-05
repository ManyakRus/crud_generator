package crud_object_tables

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
)

// CreateAllFiles - создаёт все файлы в папке objects
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

		//файлы crud objects
		err = CreateFiles_ReadObject(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_ReadObject() table: ", Table1.Name, " error: ", err)
			return err
		}

		//файлы crud objects test
		err = CreateFiles_ReadObject_Test(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_ReadObject() table: ", Table1.Name, " error: ", err)
			return err
		}
	}
	return err
}
