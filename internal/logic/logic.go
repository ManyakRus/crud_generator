package logic

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files/alias"
	"github.com/ManyakRus/crud_generator/internal/create_files/constants_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_starter_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/entities_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/env_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/generation_code_sh"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/main_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/makefile"
	"github.com/ManyakRus/crud_generator/internal/create_files/nrpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/protobuf"
	"github.com/ManyakRus/crud_generator/internal/create_files/readme_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_func"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_nrpc_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/tables_tables"
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
		log.Error("FillMapTable()error: len(MapAll) == 0")
		return err
	}

	//копируем все файлы
	dir := micro.ProgramDir_bin()
	err = folders.CopyAllFiles_Exclude_(dir+config.Settings.TEMPLATE_FOLDERNAME, dir+config.Settings.READY_FOLDERNAME)
	if err != nil {
		//log.Error("folders.CopyAllFiles_Exclude_() error: ", err)
		return err
	}

	//модель
	err = entities_tables.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("model.CreateAllFiles() error: ", err)
		return err
	}

	//db crud
	err = crud_tables.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("db_crud_tables.CreateAllFiles() error: ", err)
		return err
	}

	//grpc_server
	err = server_grpc_tables.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("grpc_server.CreateAllFiles() error: ", err)
		return err
	}

	//grpc_client tables
	err = grpc_client_tables.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("grpc_client_tables.CreateAllFiles() error: ", err)
		return err
	}

	//grpc_client
	err = grpc_client.CreateAllFiles()
	if err != nil {
		//log.Error("grpc_client.CreateAllFiles() error: ", err)
		return err
	}

	//nrpc_client
	err = nrpc_client.CreateAllFiles()
	if err != nil {
		//log.Error("nrpc_client.CreateAllFiles() error: ", err)
		return err
	}

	////nrpc_client tables
	//err = nrpc_client_tables.CreateAllFiles(MapAll)
	//if err != nil {
	//	//log.Error("nrpc_client_tables.CreateAllFiles() error: ", err)
	//	return err
	//}

	//grpc_proto
	err = protobuf.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("protobuf.CreateAllFiles() error: ", err)
		return err
	}

	//crud_starter
	err = crud_starter.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("crud_starter.CreateAllFiles() error: ", err)
		return err
	}

	//main file
	err = main_file.CreateAllFiles()
	if err != nil {
		//log.Error("main_file.CreateAllFiles() error: ", err)
		return err
	}

	//server_grpc_starter
	err = server_grpc_starter.CreateAllFiles()
	if err != nil {
		//log.Error("server_grpc_starter.CreateAllFiles() error: ", err)
		return err
	}

	//server_nrpc_starter
	err = server_nrpc_starter.CreateAllFiles()
	if err != nil {
		//log.Error("server_nrpc_starter.CreateAllFiles() error: ", err)
		return err
	}

	//server_grpc_func
	err = server_grpc_func.CreateAllFiles()
	if err != nil {
		//log.Error("server_grpc_func.CreateAllFiles() error: ", err)
		return err
	}

	//makefile
	err = makefile.CreateAllFiles()
	if err != nil {
		//log.Error("makefile.CreateAllFiles() error: ", err)
		return err
	}

	//generation_code.sh
	err = generation_code_sh.CreateAllFiles()
	if err != nil {
		//log.Error("generation_code_sh.CreateAllFiles() error: ", err)
		return err
	}

	//tables
	err = tables_tables.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("db_tables.CreateAllFiles() error: ", err)
		return err
	}

	//tables
	err = crud_starter_tables.CreateAllFiles(MapAll)
	if err != nil {
		//log.Error("db_tables.CreateAllFiles() error: ", err)
		return err
	}

	//env
	err = env_file.CreateAllFiles()
	if err != nil {
		//log.Error("env_file.CreateAllFiles() error: ", err)
		return err
	}

	//alias
	err = alias.CreateAllFiles()
	if err != nil {
		//log.Error("env_file.CreateAllFiles() error: ", err)
		return err
	}

	//readme
	err = readme_file.CreateAllFiles()
	if err != nil {
		//log.Error("env_file.CreateAllFiles() error: ", err)
		return err
	}

	//constants
	err = constants_file.CreateAllFiles()
	if err != nil {
		//log.Error("env_file.CreateAllFiles() error: ", err)
		return err
	}

	return err
}
