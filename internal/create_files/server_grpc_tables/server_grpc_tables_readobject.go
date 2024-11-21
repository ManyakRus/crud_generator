package server_grpc_tables

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

// CreateFiles_ReadObject - создаёт 1 файл в папке grpc_server
func CreateFiles_ReadObject(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_READOBJECT_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_readobject" + ".go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		//
		ObjectTableURL := create_files.Find_ObjectTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ObjectTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudObjectTableURL := create_files.Find_CrudObjectTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudObjectTableURL)

		if Table1.PrimaryKeyColumnsCount == 1 {
			TextGRPCServer = Replace_IDRequestID_1PK(TextGRPCServer, Table1)
		}

		//замена "m.ID = AliasFromInt(ID)"
		TextGRPCServer = Replace_PrimaryKeyM_ID(TextGRPCServer, Table1)

		//замена "ID := Request.ID"
		//TextGRPCServer = create_files.Replace_PrimaryKeyOtvetID(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = ReplaceText_RequestID_PrimaryKey(TextGRPCServer, Table1)

		//замена int64(ID) на ID
		//TextGRPCServer = create_files.Replace_IDtoID(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

		//удалим лишние функции
		TextGRPCServer = DeleteFunc_Delete(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFunc_Restore(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFunc_Find_byExtID(TextGRPCServer, Table1)
	}

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "grpc_proto.", TextProto+".")

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	//ModelName := Table1.NameGo
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = Convert_RequestIdToAlias(TextGRPCServer, Table1)
	}

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_ReadObject_Test - создаёт 1 файл в папке grpc_server
func CreateFiles_ReadObject_Test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_READOBJECT_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_readobject" + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//удалим лишние функции
		TextGRPCServer = DeleteFunc_Test_Delete(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFunc_Test_Restore(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFunc_Test_Find_byExtID(TextGRPCServer, Table1)

		//добавим импорты
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		//object
		ObjectTableURL := create_files.Find_ObjectTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ObjectTableURL)

		//
		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		//
		CrudStarterURL := create_files.Find_CrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		//
		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudFuncURL)

		//
		ConstantsURL := create_files.Find_ConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		//
		TextGRPCServer = create_files.CheckAndAdd_Import(TextGRPCServer, "encoding/json")

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.Replace_Connect_WithApplicationName(TextGRPCServer)

		if Table1.PrimaryKeyColumnsCount > 1 {
		}

		//Postgres_ID_Test = ID Minimum
		TextGRPCServer = Replace_Model_ID_Test(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = ReplaceText_RequestID_PrimaryKey(TextGRPCServer, Table1)

		//замена Otvet.ID = -1
		TextGRPCServer = Replace_ModelIDEqual1(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

		//
		TextGRPCServer = Replace_OtvetIDEqual0(TextGRPCServer, Table1)
	}

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "grpc_proto.", TextProto+".")

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	// замена ID на PrimaryKey
	//TextGRPCServer = create_files.Replace_PrimaryKeyOtvetID(TextGRPCServer, Table1)
	//TextGRPCServer = create_files.Replace_PrimaryKeyM_ID(TextGRPCServer, Table1)

	//SkipNow()
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//замена импортов на новые URL
	//TextGRPCServer = create_files.Replace_RepositoryImportsURL(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}
