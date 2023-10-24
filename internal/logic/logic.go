package logic

import (
	"github.com/ManyakRus/crud_generator/internal/create_files/model"
	"github.com/ManyakRus/crud_generator/internal/postgres"
	"github.com/ManyakRus/starter/log"
)

//var MassTable []types.Table

func StartFillAll() error {
	var err error

	//заполним MapAll
	MapAll, err := postgres.FillMapTable()
	if err != nil {
		log.Error("FillMapTable() error: ", err)
		return err
	}

	if len(MapAll) == 0 {
		return err
	}

	err = model.CreateModelFiles(MapAll)

	return err
}

//// MassFromMapColumns - возвращает Slice из Map
//func MassFromMapColumns(MapColumns map[string]types.Column) []types.Column {
//	Otvet := make([]types.Column, 0)
//
//	for _, v := range MapColumns {
//		Otvet = append(Otvet, v)
//	}
//
//	sort.Slice(Otvet[:], func(i, j int) bool {
//		return Otvet[i].OrderNumber < Otvet[j].OrderNumber
//	})
//
//	return Otvet
//}
