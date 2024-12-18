package logic

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/create_files/alias"
	"github.com/ManyakRus/crud_generator/internal/create_files/constants_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_readobject_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_starter_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/crud_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/entities_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/env_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/generation_code_sh"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client_func"
	"github.com/ManyakRus/crud_generator/internal/create_files/grpc_client_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/main_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/makefile"
	"github.com/ManyakRus/crud_generator/internal/create_files/nrpc_client"
	"github.com/ManyakRus/crud_generator/internal/create_files/object_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/protobuf"
	"github.com/ManyakRus/crud_generator/internal/create_files/readme_file"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_func"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_grpc_tables"
	"github.com/ManyakRus/crud_generator/internal/create_files/server_nrpc_starter"
	"github.com/ManyakRus/crud_generator/internal/create_files/tables_tables"
	"github.com/ManyakRus/crud_generator/internal/database"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/load_configs"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
)

//var MassTable []types.Table

func StartFillAll() error {
	var err error

	//заполним MapAll
	SettingsFillFromDatabase := types.SettingsFillFromDatabase{}
	SettingsFillFromDatabase.INCLUDE_TABLES = config.Settings.INCLUDE_TABLES
	SettingsFillFromDatabase.EXCLUDE_TABLES = config.Settings.EXCLUDE_TABLES
	SettingsFillFromDatabase.NEED_USE_DB_VIEWS = config.Settings.NEED_USE_DB_VIEWS
	SettingsFillFromDatabase.SINGULAR_TABLE_NAMES = config.Settings.SINGULAR_TABLE_NAMES
	SettingsFillFromDatabase.MapDBTypes = dbmeta.GetMappings()
	MapAll, err := database.FillMapTable(SettingsFillFromDatabase)
	if err != nil {
		log.Error("FillMapTable() error: ", err)
		return err
	}

	if len(MapAll) == 0 {
		log.Error("FillMapTable()error: len(MapAll) == 0")
		return err
	}

	//загрузим FindModelBy .json
	load_configs.LoadFindModelBy(MapAll)

	//
	FillGlobalVariables(MapAll)

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

	//grpc_client_func
	err = grpc_client_func.CreateAllFiles()
	if err != nil {
		//log.Error("env_file.CreateAllFiles() error: ", err)
		return err
	}

	//crud
	err = crud.CreateAllFiles()
	if err != nil {
		//log.Error("env_file.CreateAllFiles() error: ", err)
		return err
	}

	//objects
	if config.Settings.NEED_CREATE_READOBJECT == true {
		err = object_tables.CreateAllFiles(MapAll)
		if err != nil {
			//log.Error("env_file.CreateAllFiles() error: ", err)
			return err
		}

		err = crud_readobject_tables.CreateAllFiles(MapAll)
		if err != nil {
			//log.Error("env_file.CreateAllFiles() error: ", err)
			return err
		}

	}

	return err
}

// FillGlobalVariables - заполняет глобальные переменные
func FillGlobalVariables(MapAll map[string]*types.Table) {

	// MassFindBy
	Mass1 := create_files.FindMass_TableColumns(MapAll, types.MassFindBy_String)
	types.MassFindBy = Mass1

	// MassFindMassBy
	Mass1 = create_files.FindMass_TableColumns(MapAll, types.MassFindMassBy_String)
	types.MassFindMassBy = Mass1

	// ReadAll
	types.MapReadAll = load_configs.LoadReadAll(MapAll)

}
