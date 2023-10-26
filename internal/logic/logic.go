package logic

import (
	"github.com/ManyakRus/crud_generator/internal/create_files/db"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_proto"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_server"
	"github.com/ManyakRus/crud_generator/internal/create_files/model"
	"github.com/ManyakRus/crud_generator/internal/create_files/nrpc_client"
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

	//модель
	err = model.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//db crud
	err = db.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_server
	err = grpc_server.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_client
	err = grpc_client.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_client
	err = nrpc_client.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_proto
	err = grpc_proto.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

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
