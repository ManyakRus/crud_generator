package logic

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/db_crud_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/db_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/env_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/generation_code_sh"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_server"
	"github.com/ManyakRus/crud_generator/internal/create_files/main_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/makefile"
	"github.com/ManyakRus/crud_generator/internal/create_files/model"
	"github.com/ManyakRus/crud_generator/internal/create_files/nrpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/nrpc_client_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/protobuf"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_func"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_nrpc_starter"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/postgres"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
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

	//копируем все файлы
	dir := micro.ProgramDir_bin()
	err = folders.CopyAllFiles_Exclude_(dir+config.Settings.TEMPLATE_FOLDERNAME, dir+config.Settings.READY_FOLDERNAME)
	if err != nil {
		return err
	}

	//модель
	err = model.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//db crud
	err = db_crud_tables.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_server
	err = grpc_server.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_client tables
	err = grpc_client_tables.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_client
	err = grpc_client.CreateAllFiles()
	if err != nil {
		return err
	}

	//nrpc_client
	err = nrpc_client.CreateAllFiles()
	if err != nil {
		return err
	}

	//grpc_client
	err = nrpc_client_tables.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//grpc_proto
	err = protobuf.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//crud_starter
	err = crud_starter.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//main file
	err = main_file.CreateAllFiles()
	if err != nil {
		return err
	}

	//server_grpc_starter
	err = server_grpc_starter.CreateAllFiles()
	if err != nil {
		return err
	}

	//server_nrpc_starter
	err = server_nrpc_starter.CreateAllFiles()
	if err != nil {
		return err
	}

	//server_grpc_func
	err = server_grpc_func.CreateAllFiles()
	if err != nil {
		return err
	}

	//makefile
	err = makefile.CreateAllFiles()
	if err != nil {
		return err
	}

	//generation_code.sh
	err = generation_code_sh.CreateAllFiles()
	if err != nil {
		return err
	}

	//tables
	err = db_tables.CreateAllFiles(MapAll)
	if err != nil {
		return err
	}

	//env
	err = env_file.CreateAllFiles()
	if err != nil {
		return err
	}

	return err
}
