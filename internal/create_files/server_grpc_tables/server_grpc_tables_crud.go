package server_grpc_tables

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

// CreateFiles - создаёт 1 файл в папке grpc_server
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "server_grpc.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + ".go"

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

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

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
	TextGRPCServer = Replace_RequestExtID(TextGRPCServer, Table1)

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

// CreateFiles_Test - создаёт 1 файл в папке grpc_server
func CreateFiles_Test(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "server_grpc_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_test.go"

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

// DeleteFunc_Delete - удаляет функцию Delete()
func DeleteFunc_Delete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// "+ModelName+"_Delete ")

	return Otvet
}

// DeleteFunc_Restore - удаляет функцию Restore()
func DeleteFunc_Restore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore ")

	return Otvet
}

//// DeleteFunc_DeleteCtx - удаляет функцию Delete_ctx()
//func DeleteFunc_DeleteCtx(Text, ModelName string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Delete_ctx ")
//
//	return Otvet
//}
//
//// DeleteFunc_RestoreCtx - удаляет функцию Restore_ctx()
//func DeleteFunc_RestoreCtx(Text, ModelName string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore_ctx ")
//
//	return Otvet
//}

// DeleteFunc_Find_byExtID - удаляет функцию Find_ByExtID()
func DeleteFunc_Find_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_FindByExtID ")

	return Otvet
}

// DeleteFunc_Test_Delete - удаляет функцию Delete()
func DeleteFunc_Test_Delete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_Delete")

	return Otvet
}

// DeleteFunc_Test_Restore - удаляет функцию Restore()
func DeleteFunc_Test_Restore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"Restore")

	return Otvet
}

// DeleteFunc_Find_byExtID - удаляет функцию Find_ByExtID()
func DeleteFunc_Test_Find_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_FindByExtID")

	return Otvet
}

// Replace_PrimaryKey_ID - заменяет "m.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyM_ID(Text string, Table1 *types.Table) string {
	Otvet := Text

	VariableName := "m"

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"

		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)

	//заменим ID := Request.ID
	Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextIDRequestID)

	return Otvet
}

//// Replace_PrimaryKeyOtvet_ID - заменяет "m.ID" на название колонки PrimaryKey
//func Replace_PrimaryKeyOtvet_ID(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	VariableName := "Otvet"
//
//	//сортировка по названию таблиц
//	keys := make([]string, 0, len(Table1.MapColumns))
//	for k := range Table1.MapColumns {
//		keys = append(keys, k)
//	}
//	sort.Strings(keys)
//
//	TextOtvetIDAliasID := ""
//	TextIfMId := ""
//	TextIfMIdNot0 := ""
//	TextM2ID := ""
//	TextIDRequestID := ""
//	TextOtvetIDID := ""
//	TextRequestIDmID := ""
//	TextRequestIDInt64ID := ""
//	TextOtvetIDmID := ""
//	TextMID0 := ""
//	TextOR := ""
//	for _, key1 := range keys {
//		Column1, _ := Table1.MapColumns[key1]
//		if Column1.IsPrimaryKey != true {
//			continue
//		}
//		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
//		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
//		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
//		if GolangCode == "" {
//			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
//		} else {
//			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
//		}
//		TextM := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
//		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
//		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
//		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
//		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"
//
//		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)
//
//		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
//		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
//		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue
//
//		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
//		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
//		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
//		TextOR = " || "
//	}
//
//	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)
//
//	return Otvet
//}

// ReplaceText_RequestID_PrimaryKey - заменяет RequestId{} на RequestString{}
func ReplaceText_RequestID_PrimaryKey(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextRequestID := create_files.FindText_ProtobufRequest_ManyPK(Table1)
	TextProto := create_files.TextProto()

	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
	Otvet = strings.ReplaceAll(Otvet, "*grpc_proto.RequestId", "*"+TextProto+"."+TextRequestID)

	return Otvet
}

// Replace_ModelIDEqual1 - заменяет Otvet.ID = -1
func Replace_ModelIDEqual1(Text string, Table1 *types.Table) string {
	Otvet := Text

	Otvet = Replace_ModelIDEqual1_ManyPK(Otvet, Table1)

	return Otvet
}

// Replace_ModelIDEqual1_ManyPK - заменяет m.ID = -1
func Replace_ModelIDEqual1_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextFind := "\tm.ID = -1\n"
	TextNew := ""
	MassPrimaryKey := create_files.Find_PrimaryKeyColumns(Table1)
	for _, Column1 := range MassPrimaryKey {
		Value := create_files.Find_NegativeValue(Column1.TypeGo)
		TextNew = TextNew + "\tm." + Column1.NameGo + " = " + Value + "\n"
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

	////
	//TextFind := "\tOtvet.ID = 0\n"
	//TextNew := ""
	//for _, ColumnPK1 := range ColumnsPK {
	//	Value := create_files.FindText_DefaultValue(ColumnPK1.TypeGo)
	//	TextNew = TextNew + "\tOtvet." + ColumnPK1.NameGo + " = " + Value + "\n"
	//}
	//Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//
	TextFind := " Otvet.ID == 0"
	for _, ColumnPK1 := range ColumnsPK {
		Value := create_files.FindText_DefaultValue(ColumnPK1.TypeGo)
		TextNew := " Otvet." + ColumnPK1.NameGo + " == " + Value + ""
		Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)
		break
	}

	return Otvet
}

// Replace_RequestExtID - заменяет RequestExtID{} на Request_Int64_String{}
func Replace_RequestExtID(TextGRPCServer string, Table1 *types.Table) string {
	Otvet := TextGRPCServer

	//если нет таких колонок - ничего не делаем
	if create_files.Has_Column_ExtID_ConnectionID(Table1) == false {
		return Otvet
	}

	//если обе колонки Int64 - ничего не делаем
	if create_files.Has_Column_ExtID_ConnectionID_Int64(Table1) == true {
		return Otvet
	}

	//
	ColumnExtID := create_files.FindColumn_ExtID(Table1)
	if ColumnExtID == nil {
		return Otvet
	}

	//
	if ColumnExtID.TypeGo != "string" {
		return Otvet
	}

	//
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestExtID", "grpc_proto.RequestExtIDString")

	return Otvet
}
