package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFiles_UpdateEveryColumn - создаёт 1 файл в папке grpc_client
func CreateFiles_UpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPC_Client := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPC_Client := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()

	FilenameTemplateGRPC_ClientFunc := DirTemplatesGRPC_Client + config.Settings.GRPC_CLIENT_TABLE_UPDATE_FUNC_FILENAME
	DirReadyTable := DirReadyGRPC_Client
	FilenameReadyGRPC_ClientUpdate := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateGRPC_ClientFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPC_ClientFunc, " error: ", err)
	}
	TextGRPC_Client_UpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateCrud := DirTemplatesGRPC_Client + config.Settings.TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextGRPC_Client := string(bytes)
	TextGRPC_Client = TextGRPC_Client + "\n"
	TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//заменим имя пакета на новое
	TextGRPC_Client = create_files.Replace_PackageName(TextGRPC_Client, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPC_Client = create_files.Delete_TemplateRepositoryImports(TextGRPC_Client)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ProtoURL)

		//GRPCClientURL := create_files.Find_GRPClientURL()
		//TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCClientURL)

		//grpc client func
		GRPCClientFuncURL := create_files.Find_GRPCClient_func_URL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCClientFuncURL)

		GRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCConstantsURL)

		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPC_NRPC_URL)

		if config.Settings.NEED_CREATE_NRPC == true {
			NRPCClientURL := create_files.Find_NRPClientURL()
			TextGRPC_Client = create_files.AddImport(TextGRPC_Client, NRPCClientURL)
		}

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ModelTableURL)

		//TextGRPC_Client = create_files.Convert_RequestIdToAlias(TextGRPC_Client, Table1)
	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextGRPC_Client_UpdateFunc, Table1)
	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	//NRPC
	if config.Settings.NEED_CREATE_NRPC == true {
		//уберём "//"
		TextGRPC_Client = Replace_NRPC_CLIENT(TextGRPC_Client)
	}

	//ModelName := Table1.NameGo
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPC_Client = TextGRPC_Client + TextUpdateEveryColumn

	TextGRPC_Client = config.Settings.TEXT_MODULE_GENERATED + TextGRPC_Client

	//удаление пустого импорта
	TextGRPC_Client = create_files.Delete_EmptyImport(TextGRPC_Client)
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPC_Client = create_files.CheckAndAdd_ImportTime_FromText(TextGRPC_Client)
		TextGRPC_Client = create_files.CheckAndAdd_ImportTimestamp_FromText(TextGRPC_Client)
		TextGRPC_Client = create_files.CheckAndAdd_ImportAlias(TextGRPC_Client)

		//добавим импорт uuid
		TextGRPC_Client = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPC_Client)

	}

	//удаление пустых строк
	TextGRPC_Client = create_files.Delete_EmptyLines(TextGRPC_Client)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPC_ClientUpdate, []byte(TextGRPC_Client), config.Settings.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextGRPC_ClientUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumn1(TextGRPC_ClientUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextGRPC_ClientUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPC_ClientUpdateFunc

	ModelNameTranslit := Table1.NameGo_translit
	ColumnName := Column1.NameGo
	ColumnNameTranslit := Column1.NameGo_translit
	FuncName := "Update_" + ColumnNameTranslit
	_, TextRequestFieldName, _, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "Request.")
	TextRequest := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)

	//замена RequestId{}
	Otvet = ReplaceText_RequestID_and_Column(Otvet, Table1, Column1)
	Otvet = create_files.ReplaceText_RequestID_PrimaryKey(Otvet, Table1)

	//замена ID на PrimaryKey
	Otvet = Replace_PrimaryKeyM_ID(Otvet, Table1)

	//
	ColumnNameGolang := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")

	_, IDTypeGo := create_files.Find_PrimaryKeyNameTypeGo(Table1)

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Update", ModelNameTranslit+"_"+FuncName)
	Otvet = strings.ReplaceAll(Otvet, " Update ", " "+FuncName+" ")
	Otvet = strings.ReplaceAll(Otvet, " Update(", " "+FuncName+"(")

	Otvet = create_files.Replace_TemplateModel_to_Model(Otvet, Table1.NameGo)
	Otvet = create_files.Replace_TemplateTableName_to_TableName(Otvet, Table1.Name)

	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "m.ColumnName", ColumnNameGolang)
	Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", " "+IDTypeGo+"(m.ID)")
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.FieldName", "Request."+TextRequestFieldName)

	Otvet = Replace_PrimaryKeyRequest_ID(Otvet, Table1)
	Otvet = Replace_PrimaryKeyOtvetID(Otvet, Table1)

	return Otvet
}

// CreateFiles_UpdateEveryColumn_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_UpdateEveryColumn_Test(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPC_Client := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPC_Client := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()

	FilenameTemplateGRPC_ClientFunc := DirTemplatesGRPC_Client + config.Settings.GRPC_CLIENT_TABLE_UPDATE_FUNC_TEST_FILENAME
	DirReadyTable := DirReadyGRPC_Client + "tests" + micro.SeparatorFile()
	FilenameReadyGRPC_ClientUpdate := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_update_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateGRPC_ClientFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPC_ClientFunc, " error: ", err)
	}
	TextGRPC_Client_UpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateCrud := DirTemplatesGRPC_Client + config.Settings.TEMPLATES_GRPC_CLIENT_TABLE_UPDATE_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextGRPC_Client := string(bytes)
	TextGRPC_Client = TextGRPC_Client + "\n"
	TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//заменим имя пакета на новое
	TextGRPC_Client = create_files.Replace_PackageName(TextGRPC_Client, DirReadyTable)

	//	TextGRPC_Client := "package " + config.Settings.PREFIX_CLIENT_GRPC + TableName + "\n\n"
	//	TextGRPC_Client = TextGRPC_Client + `import (
	//	"testing"
	//	"github.com/ManyakRus/starter/config_main"
	//)
	//
	//`

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPC_Client = create_files.Delete_TemplateRepositoryImports(TextGRPC_Client)

		GRPCClientURL := create_files.Find_GRPClientURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCClientURL)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ModelTableURL)

		//ConstantsURL := create_files.Find_ConstantsURL()
		//TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ConstantsURL)

		//TextGRPC_Client = create_files.Convert_RequestIdToAlias(TextGRPC_Client, Table1)

		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPClientTableURL)

		//GRPClientFuncURL := create_files.Find_GRPCClient_func_URL()
		//TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPClientFuncURL)

		TextGRPC_Client = Replace_Postgres_ID_Test(TextGRPC_Client, Table1)

		TextGRPC_Client = Replace_PrimaryKeyOtvetID(TextGRPC_Client, Table1)

		//замена m.ID = Postgres_ID_Test
		TextGRPC_Client = Replace_PrimaryKeyM_ID(TextGRPC_Client, Table1)

	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumnTest(TextGRPC_Client_UpdateFunc, Table1)
	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	//ModelName := Table1.NameGo
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPC_Client = TextGRPC_Client + TextUpdateEveryColumn

	//добавим импорт uuid
	TextGRPC_Client = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPC_Client)

	//
	TextGRPC_Client = config.Settings.TEXT_MODULE_GENERATED + TextGRPC_Client

	//TextGRPC_Client = create_files.ReplaceText_RequestID_and_Column(TextGRPC_Client, Table1)

	//SkipNow() если нет строк в БД
	TextGRPC_Client = create_files.AddSkipNow(TextGRPC_Client, Table1)

	//удаление пустого импорта
	TextGRPC_Client = create_files.Delete_EmptyImport(TextGRPC_Client)

	//удаление пустых строк
	TextGRPC_Client = create_files.Delete_EmptyLines(TextGRPC_Client)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPC_ClientUpdate, []byte(TextGRPC_Client), config.Settings.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumnTest - возвращает текст для всех таблиц
func FindTextUpdateEveryColumnTest(TextGRPC_ClientUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumnTest1(TextGRPC_ClientUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumnTest1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumnTest1(TextGRPC_ClientUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPC_ClientUpdateFunc

	Otvet = Replace_Postgres_ID_Test(Otvet, Table1)

	Otvet = ReplaceText_RequestID_and_Column(Otvet, Table1, Column1)
	Otvet = Replace_PrimaryKeyM_ID(Otvet, Table1)
	Otvet = Replace_PrimaryKeyOtvetID(Otvet, Table1)

	//ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	ColumnNameTranslit := Column1.NameGo_translit
	FuncName := "Update_" + ColumnNameTranslit
	TextRequest, TextRequestFieldName, _, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "Request.")
	DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

	Otvet = strings.ReplaceAll(Otvet, "TestCrud_GRPC_Update(", "TestCrud_GRPC_"+FuncName+"(")

	Otvet = create_files.Replace_TemplateModel_to_Model(Otvet, Table1.NameGo)
	Otvet = create_files.Replace_TemplateTableName_to_TableName(Otvet, Table1.Name)

	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNameTranslit", ColumnNameTranslit)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	Otvet = strings.ReplaceAll(Otvet, "TestRead(", "Test"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, "error: ID =0", "error: "+ColumnName+" ="+DefaultValue)
	Otvet = strings.ReplaceAll(Otvet, "_Update(", "_"+FuncName+"(")

	return Otvet
}

// ReplaceText_RequestID_and_Column - заменяет RequestId{} на RequestString{}
func ReplaceText_RequestID_and_Column(Text string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := Text

	//TypeGo := Column1.TypeGo

	TextRequestID, _, _, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "Request")
	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
	//Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextID)

	return Otvet
}
