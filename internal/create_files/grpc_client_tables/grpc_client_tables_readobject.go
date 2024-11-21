package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_ReadObject - создаёт 1 файл в папке grpc_client
func CreateFiles_ReadObject(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_READOBJECT_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + "_readobject.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//создание текста
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//proto
		RepositoryGRPCProtoURL := create_files.Find_ProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)

		////model
		//RepositoryModelURL := create_files.Find_ModelTableURL(TableName)
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryModelURL)

		//object
		ObjectModelURL := create_files.Find_ObjectTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ObjectModelURL)

		////grpc client
		//RepositoryGRPCClientlURL := create_files.Find_GRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCClientlURL)

		//grpc client func
		GRPCClientFuncURL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClientFuncURL)

		//nrpc client
		if config.Settings.NEED_CREATE_NRPC == true {
			//
			RepositoryNRPCClientlURL := create_files.Find_NRPClientURL()
			TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)

			//уберём "//"
			TextGRPCClient = Replace_NRPC_CLIENT(TextGRPCClient)
		}

		//constants GRPC
		RepositoryGRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		//DBConstantsURL := create_files.Find_DBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//grpc_nrpc
		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		//замена ID на PrimaryKey
		TextGRPCClient = Replace_PrimaryKeyRequest_ID(TextGRPCClient, Table1)
		TextGRPCClient = Replace_PrimaryKeyOtvetID(TextGRPCClient, Table1)

		//замена ID на PrimaryKey
		TextGRPCClient = Replace_PrimaryKeyM_ID(TextGRPCClient, Table1)

		//замена RequestId{}
		TextGRPCClient = ReplaceText_RequestID_PrimaryKey_ManyPK(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	}

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, "grpc_proto.", TextProto+".")

	//удалим лишние функции
	TextGRPCClient = create_files.DeleteFunc_Delete(TextGRPCClient, Table1)
	TextGRPCClient = create_files.DeleteFunc_Restore(TextGRPCClient, Table1)
	TextGRPCClient = create_files.DeleteFunc_Find_byExtID(TextGRPCClient, Table1)

	//замена имени таблицы
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//ModelName := Table1.NameGo
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClient, []byte(TextGRPCClient), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_ReadObject_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_ReadObject_Test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_READOBJECT_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + "_readobject_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//
		GRPClientURL := create_files.Find_GRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		////
		//ModelTableName := create_files.Find_ModelTableURL(TableName)
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableName)

		//object
		ObjectTableURL := create_files.Find_ObjectTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ObjectTableURL)

		//
		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

		//
		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, CrudFuncURL)

		//GRPClientFuncURL := create_files.Find_GRPCClient_func_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientFuncURL)

		//Postgres_ID_Test = ID Minimum
		TextGRPCClient = Replace_Postgres_ID_Test(TextGRPCClient, Table1)

		//замена Otvet.ID = -1
		TextGRPCClient = Replace_OtvetIDEqual1(TextGRPCClient, Table1)

		//замена Otvet.ID = 0
		TextGRPCClient = Replace_OtvetIDEqual0(TextGRPCClient, Table1)

		//замена ID на PrimaryKey
		TextGRPCClient = Replace_PrimaryKeyOtvetID(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	}

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	ModelName := Table1.NameGo
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	if config.Settings.HAS_IS_DELETED == true {
		TextGRPCClient = DeleteFuncTestDelete(TextGRPCClient, ModelName, Table1)
		TextGRPCClient = DeleteFuncTestRestore(TextGRPCClient, ModelName, Table1)
	}
	TextGRPCClient = DeleteFuncTestFind_byExtID(TextGRPCClient, ModelName, Table1)

	//SkipNow()
	TextGRPCClient = create_files.AddSkipNow(TextGRPCClient, Table1)

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClient, []byte(TextGRPCClient), config.Settings.FILE_PERMISSIONS)

	return err
}
