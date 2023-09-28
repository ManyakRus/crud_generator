package logic

import (
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/postgres"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
)

//var MassTable []types.Table

func StartFillAll() bool {
	Otvet := false

	//заполним MapAll
	MapAll, err := postgres.FillMapTable()
	if err != nil {
		log.Error("FillMapTable() error: ", err)
		return Otvet
	}

	if len(MapAll) > 0 {
		Otvet = true
	}

	if Otvet == false {
		println("warning: Empty file not saved !")
		return Otvet
	}

	err = CreateModelFiles(MapAll)

	return Otvet
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

func CreateModelFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, table1 := range MapAll {
		err = CreateModelFiles1(table1)
		if err != nil {
			return err
		}
	}

	return err
}

func CreateModelFiles1(table1 *types.Table) error {
	var err error

	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()

	return err
}
