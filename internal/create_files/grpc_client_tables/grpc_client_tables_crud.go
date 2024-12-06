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

// CreateFiles - создаёт 1 файл в папке grpc_client
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + "grpc_client_table.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + ".go"

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

		//model
		RepositoryModelURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryModelURL)

		////grpc client
		//RepositoryGRPCClientlURL := create_files.Find_GRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCClientlURL)

		//grpc client func
		GRPCClientFuncURL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClientFuncURL)

		//nrpc client
		if config.Settings.NEED_CREATE_NRPC == true {
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

	TextGRPCClient = Replace_RequestExtID(TextGRPCClient, Table1)

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

// CreateFiles_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_Test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + "grpc_client_table_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + "_test.go"

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

		GRPClientURL := create_files.Find_GRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		ModelTableName := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableName)

		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

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
	TextGRPCClient = Replace_RequestExtID(TextGRPCClient, Table1)

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

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete ")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore ")

	return Otvet
}

// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
func DeleteFuncDeleteCtx(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete_ctx ")

	return Otvet
}

// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
func DeleteFuncRestoreCtx(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore_ctx ")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Find_ByExtID ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestDelete")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestRestore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestFindByExtID")

	return Otvet
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

// Replace_Postgres_ID_Test - заменяет текст "const Postgres_ID_Test = 0" на нужные ИД, для много колонок PrimaryKey
func Replace_Postgres_ID_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	if len(MassPK) == 0 {
		return Otvet
	}

	//заменим const Postgres_ID_Test = 0
	TextFind := "const Postgres_ID_Test = 0\n"
	TextNew := ""
	for _, PrimaryKey1 := range MassPK {
		TextNew = TextNew + create_files.Replace_Postgres_ID_Test1(TextFind, Table1, PrimaryKey1)
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим Otvet.ID = Postgres_ID_Test
	TextFind = "\tOtvet.ID = Postgres_ID_Test\n"
	TextNew = ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := create_files.FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "Otvet."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим m.ID = Postgres_ID_Test
	TextFind = "\tm.ID = Postgres_ID_Test\n"
	TextNew = ""
	for _, PrimaryKey1 := range MassPK {
		Text1 := create_files.FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "m."+PrimaryKey1.NameGo)
		TextNew = TextNew + "\t" + Text1 + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	////заменим m1.ID = Postgres_ID_Test
	//TextFind = "\tm1.ID = Postgres_ID_Test\n"
	//TextNew = ""
	//for _, PrimaryKey1 := range MassPK {
	//	Text1 := create_files.FindText_VariableEqual_ColumnName_Test(PrimaryKey1, "m1."+PrimaryKey1.NameGo)
	//	TextNew = TextNew + "\t" + Text1 + "\n"
	//}
	//Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим ReadFromCache(Postgres_ID_Test)
	TextFind = "ReadFromCache(Postgres_ID_Test)"
	TextNew = "ReadFromCache("
	Comma := ""
	for _, PrimaryKey1 := range MassPK {
		Name := create_files.FindText_ColumnNameTest(PrimaryKey1)
		TextNew = TextNew + Comma + Name
		Comma = ", "
	}
	TextNew = TextNew + ")"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	////заменим ненужные Otvet.ID на Otvet.Name
	PrimaryKey1 := MassPK[0]
	//Otvet = strings.ReplaceAll(Otvet, " Otvet.ID ", " Otvet."+PrimaryKey1.NameGo+" ")
	//Otvet = strings.ReplaceAll(Otvet, " Otvet.ID)", " Otvet."+PrimaryKey1.NameGo+")")
	Name := create_files.FindText_ColumnNameTest(PrimaryKey1)
	Otvet = strings.ReplaceAll(Otvet, "Postgres_ID_Test", Name)

	return Otvet
}

// Replace_OtvetIDEqual1 - заменяет Otvet.ID = -1
func Replace_OtvetIDEqual1(Text string, Table1 *types.Table) string {
	Otvet := Text

	PrimaryKeyColumns := create_files.Find_PrimaryKeyColumns(Table1)
	if len(PrimaryKeyColumns) == 0 {
		return Otvet
	}

	TextFind := "\tOtvet.ID = -1\n"
	TextNew := ""
	for _, ColumnPK1 := range PrimaryKeyColumns {
		Value := create_files.Find_NegativeValue(ColumnPK1.TypeGo)
		TextNew = TextNew + "\tOtvet." + ColumnPK1.NameGo + " = " + Value + "\n"
	}

	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}

// Replace_OtvetIDEqual0 - заменяет Otvet.ID = -1
func Replace_OtvetIDEqual0(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	if len(ColumnsPK) == 0 {
		return Otvet
	}

	//
	TextFind := "\tOtvet.ID = 0\n"
	TextNew := ""
	for _, ColumnPK1 := range ColumnsPK {
		Value := create_files.FindText_DefaultValue(ColumnPK1.TypeGo)
		TextNew = TextNew + "\tOtvet." + ColumnPK1.NameGo + " = " + Value + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//
	TextFind = " Otvet.ID == 0"
	for _, ColumnPK1 := range ColumnsPK {
		Value := create_files.FindText_DefaultValue(ColumnPK1.TypeGo)
		TextNew = " Otvet." + ColumnPK1.NameGo + " == " + Value + ""
		Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)
		break
	}

	return Otvet
}
