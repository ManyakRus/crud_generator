package server_grpc_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFilesUpdateEveryColumn - создаёт 1 файл в папке grpc_server
func CreateFilesUpdateEveryColumn(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServerFunc := DirTemplatesGRPCServer + constants.SERVER_GRPC_TABLE_UPDATE_FUNC_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServerUpdateFunc := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateGRPCServerFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunc, " error: ", err)
	}
	TextGRPCServerUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateGRPCServerUpdate := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_TABLE_UPDATE_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerUpdate)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerUpdate, " error: ", err)
	}
	TextGRPCServer := string(bytes)
	TextGRPCServer = TextGRPCServer + "\n"

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

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

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextGRPCServerUpdateFunc, Table1)
	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	TextGRPCServer = TextGRPCServer + TextUpdateEveryColumn

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Convert_RequestIdToAlias(TextGRPCServer, Table1)
		TextGRPCServer = create_files.CheckAndAdd_ImportAlias(TextGRPCServer)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceText_RequestID_PrimaryKey(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	}

	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.Delete_EmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServerUpdateFunc, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextGRPCServerUpdateFunc string, Table1 *types.Table) string {
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

		TextColumn1 := FindTextUpdateEveryColumn1(TextGRPCServerUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextGRPCServerUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPCServerUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	_, _, TextRequestFieldGolang, TextGolangLine := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "Request.")
	//if Table1.PrimaryKeyColumnsCount > 1 {
	TextRequest := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)
	//ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
	IsPrimaryKey := create_files.IsPrimaryKeyColumn(Table1, Column1)

	//замена ID на PrimaryKey
	Otvet = create_files.Replace_PrimaryKeyOtvetID(Otvet, Table1)
	Otvet = create_files.Replace_PrimaryKeyM_ID(Otvet, Table1)

	//ColumnNameGolang := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "m")

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Update", ModelName+"_"+FuncName)

	Otvet = create_files.Replace_TemplateModel_to_Model(Otvet, Table1.NameGo)
	Otvet = create_files.Replace_TemplateTableName_to_TableName(Otvet, Table1.Name)

	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	if IsPrimaryKey == true {
		Otvet = strings.ReplaceAll(Otvet, "\tColumnName := Request.FieldName\n", "")
	} else if TextGolangLine != "" {
		Otvet = strings.ReplaceAll(Otvet, "ColumnName := Request.FieldName", TextGolangLine)
	}
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "Request.FieldName", TextRequestFieldGolang)
	Otvet = strings.ReplaceAll(Otvet, "Model.ColumnName", "Model."+ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Model.Update()", "Model."+FuncName+"()")

	Otvet = ReplaceIDRequestID_1PK(Otvet, Table1)

	return Otvet
}

// ReplaceIDRequestID_1PK - замена "ID := Request.ID"
func ReplaceIDRequestID_1PK(Text string, Table1 *types.Table) string {
	Otvet := Text

	//замена ID := Request.ID
	ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
	_, _, RequestFieldPK, GolangLinePK := create_files.FindText_ProtobufRequest_ID_Type(Table1, ColumnPK, "Request.")

	if GolangLinePK != "" {
		Otvet = strings.ReplaceAll(Otvet, "ID := Request.ID", GolangLinePK)
	} else {
		Otvet = strings.ReplaceAll(Otvet, "ID := Request.ID", ColumnPK.NameGo+" := "+RequestFieldPK)
	}

	return Otvet
}

// CreateFilesUpdateEveryColumnTest - создаёт 1 файл в папке grpc_server
func CreateFilesUpdateEveryColumnTest(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServerFunc := DirTemplatesGRPCServer + constants.SERVER_GRPC_TABLE_UPDATE_FUNC_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServerUpdate := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_update_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	//читаем шаблон файла, только функции
	bytes, err := os.ReadFile(FilenameTemplateGRPCServerFunc)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunc, " error: ", err)
	}
	TextGRPCServerUpdateFunc := string(bytes)

	//читаем шаблон файла, без функций
	FilenameTemplateGRPCServerUpdate := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_TABLE_UPDATE_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerUpdate)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerUpdate, " error: ", err)
	}
	TextGRPCServer := string(bytes)
	TextGRPCServer = TextGRPCServer + "\n"

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		//ModelURL := create_files.Find_ModelURL()
		//TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelURL)

		CrudStarterURL := create_files.Find_CrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ConstantsURL := create_files.Find_ConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.Replace_Connect_WithApplicationName(TextGRPCServer)

		//TextGRPCServer = create_files.Convert_RequestIdToAlias(TextGRPCServer, Table1)
	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumnTest(TextGRPCServerUpdateFunc, Table1)

	//Postgres_ID_Test = ID Minimum
	TextGRPCServer = Replace_Model_ID_Test(TextGRPCServer, Table1)

	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//// пустой файл не нужен
	//if TextUpdateEveryColumn == "" {
	//	return err
	//}

	TextGRPCServer = TextGRPCServer + TextUpdateEveryColumn

	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена RequestId{}
	TextGRPCServer = create_files.ReplaceText_RequestID_PrimaryKey(TextGRPCServer, Table1)

	//добавим импорт uuid
	TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	//Import Timestamp
	TextGRPCServer = create_files.CheckAndAdd_ImportTimestamp_FromText(TextGRPCServer)

	//SkipNow() если нет строк в БД
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.Delete_EmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServerUpdate, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumnTest - возвращает текст для всех таблиц
func FindTextUpdateEveryColumnTest(TextGRPCServerUpdateFunc string, Table1 *types.Table) string {
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

		TextColumn1 := FindTextUpdateEveryColumnTest1(TextGRPCServerUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumnTest1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumnTest1(TextGRPCServerUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPCServerUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	_, TextRequestField, TextRequestFieldGolang, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "Request2.")
	TextModelColumnName := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "m")
	TextRequestID := create_files.FindText_ProtobufRequest_ManyPK(Table1)

	//Postgres_ID_Test = ID Minimum
	Otvet = Replace_Model_ID_Test(Otvet, Table1)

	//if TextGolangLine != "" {
	//	Otvet = strings.ReplaceAll(Otvet, "value := Request.FieldName", TextGolangLine)
	//}

	//if Table1.PrimaryKeyColumnsCount == 1 {
	//} else {
	TextRequestString := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId{}", "grpc_proto."+TextRequestID+"{}")
	//}

	Otvet = strings.ReplaceAll(Otvet, "Request.ColumnName", TextRequestFieldGolang)
	Otvet = strings.ReplaceAll(Otvet, "Request2.ColumnName", "Request2."+TextRequestField)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestString", "grpc_proto."+TextRequestString)
	Otvet = strings.ReplaceAll(Otvet, "m.ColumnName", TextModelColumnName)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Update(", ModelName+"_"+FuncName+"(")

	Otvet = create_files.Replace_TemplateModel_to_Model(Otvet, Table1.NameGo)
	Otvet = create_files.Replace_TemplateTableName_to_TableName(Otvet, Table1.Name)

	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	return Otvet
}
