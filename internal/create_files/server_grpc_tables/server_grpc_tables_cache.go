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

// CreateFiles_Cache - создаёт 1 файл в папке grpc_server
func CreateFiles_Cache(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesGRPCServer + config.Settings.SERVER_GRPC_TABLE_CACHE_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyCache := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_cache.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, DBConstantsURL)

		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

	}

	//TextGRPCServer = create_files.Replace_IntFromProtoRequest(TextGRPCServer, Table1)

	//замена RequestId{}
	TextGRPCServer = ReplaceText_RequestID_PrimaryKey_ManyPK(TextGRPCServer, Table1)

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "grpc_proto.", TextProto+".")

	//TextGRPCServer = create_files.Replace_PrimaryKeyOtvetID(TextGRPCServer, Table1)

	TextGRPCServer = Replace_PrimaryKeyM_ID(TextGRPCServer, Table1)

	if Table1.PrimaryKeyColumnsCount == 1 {
		ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
		TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "ReplaceManyID(ID)", ColumnPK.NameGo)
		//ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
	} else {
		TextIDMany := "ReplaceManyID(ID)"
		//TextIDMany = create_files.Replace_IDtoID_Many(TextIDMany, Table1)
		TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "ReplaceManyID(ID)", TextIDMany)
	}
	TextGRPCServer = Replace_IDtoID_Many(TextGRPCServer, Table1)

	//добавим импорт uuid
	TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	//добавим импорт alias
	TextGRPCServer = create_files.CheckAndAdd_ImportAlias(TextGRPCServer)

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	//ModelName := Table1.NameGo
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//	TextGRPCServer = create_files.Convert_RequestIdToAlias(TextGRPCServer, Table1)
	//}

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_Cache_Test - создаёт 1 файл в папке grpc_server
func CreateFiles_Cache_Test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateCache := DirTemplatesGRPCServer + config.Settings.SERVER_GRPC_TABLE_CACHE_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyCache := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_cache_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCache)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCache, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudStarterURL := create_files.Find_CrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ConstantsURL := create_files.Find_ConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudFuncURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.Replace_Connect_WithApplicationName(TextGRPCServer)

		//Postgres_ID_Test = ID Minimum
		TextGRPCServer = Replace_Model_ID_Test(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = ReplaceText_RequestID_PrimaryKey(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

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

	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//	TextGRPCServer = create_files.Convert_RequestIdToAlias(TextGRPCServer, Table1)
	//}

	//import timestamppb.New()
	TextGRPCServer = create_files.CheckAndAdd_ImportTimestamp_FromText(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//SkipNow()
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyCache, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}

// ReplaceText_RequestID_PrimaryKey_ManyPK - заменяет RequestId{} на RequestString{}
func ReplaceText_RequestID_PrimaryKey_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextRequestID := create_files.FindText_ProtobufRequest_ManyPK(Table1)
	TextProto := create_files.TextProto()

	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
	Otvet = strings.ReplaceAll(Otvet, "*grpc_proto.RequestId", "*"+TextProto+"."+TextRequestID)

	return Otvet
}

// Replace_IDtoID_Many - заменяет int64(ID) на ID, и остальные PrimaryKey
func Replace_IDtoID_Many(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextNames, _, _ := create_files.FindText_IDMany(Table1)

	Otvet = strings.ReplaceAll(Otvet, "ReplaceManyID(ID)", TextNames)
	//Otvet = strings.ReplaceAll(Otvet, "int64(ID)", TextProtoNames)
	//Otvet = strings.ReplaceAll(Otvet, "(ID int64", "("+TextNamesTypes)
	//Otvet = strings.ReplaceAll(Otvet, "(ID)", "("+TextNames+")")
	//Otvet = strings.ReplaceAll(Otvet, ", ID)", ", "+TextNames+")")
	//Otvet = strings.ReplaceAll(Otvet, ", ID int64)", ", "+TextNamesTypes+")")

	return Otvet
}
