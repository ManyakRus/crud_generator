package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_FindModelBy - создаёт 1 файл в папке grpc_client
func CreateFiles_FindModelBy(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	if len(types.MassFindModelBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMODELBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_findmodelby.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMODELBY_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCClientFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ProtoURL)

		GRPCClient_func_URL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClient_func_URL)

		GRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCConstantsURL)

		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		NRPC_Client_URL := create_files.Find_NRPC_Client_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, NRPC_Client_URL)

	}

	//создание функций
	TextClientGRPCFunc := CreateFiles_FindModelByTable(MapAll, Table1, &TextGRPCClient, TextTemplatedFunction)
	if TextClientGRPCFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextClientGRPCFunc

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//uuid
	TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	//alias
	TextGRPCClient = create_files.CheckAndAdd_ImportAlias(TextGRPCClient)

	//time
	TextGRPCClient = create_files.CheckAndAdd_ImportTime_FromText(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCClient), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindModelByTable - создаёт текст всех функций
func CreateFiles_FindModelByTable(MapAll map[string]*types.Table, Table1 *types.Table, TextGRPCClient *string, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}
		Otvet1 := CreateFiles_FindModelByTable1(MapAll, Table1, TextGRPCClient, TextTemplateFunction, TableColumns1.Column)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindModelByTable1 - создаёт текст всех функций
func CreateFiles_FindModelByTable1(MapAll map[string]*types.Table, Table1 *types.Table, TextGRPCClient *string, TextTemplateFunction string, Column1 *types.Column) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	//TextAssign := ""

	////
	//TextFind := "\t" + `Request.RequestFieldName = m.FieldName` + "\n"
	Underline := ""
	Plus := ""
	//RequestName := "Request_"
	RequestName := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)
	//TextRequest := create_files.Find_RequestFieldName(Table1, Column1)
	//ValueM := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "m")
	//TextAssign = TextAssign + "\tRequest." + TextRequest + " = " + ValueM + "\n"
	ColumnNameTranslit := Column1.NameGo
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + ColumnNameTranslit
	FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo

	//ProtoTypeName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
	//RequestName = RequestName + Underline + ProtoTypeName

	Underline = "_"
	Plus = "+"

	Otvet = strings.ReplaceAll(Otvet, "RequestName", RequestName)
	//Otvet = strings.ReplaceAll(Otvet, TextFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)

	//
	ForeignTableName := Column1.TableKey
	ForeignTable, ok := MapAll[ForeignTableName]
	if ok == false {
		log.Panic("Table not found: ", ForeignTableName)
	}

	//
	TextFindModelBy := "Find" + ForeignTable.NameGo_translit + "By"
	Otvet = strings.ReplaceAll(Otvet, "FindModelBy", TextFindModelBy)

	//
	TextForeignPackage := ForeignTable.Name
	Otvet = strings.ReplaceAll(Otvet, "foreign_package", TextForeignPackage)

	//
	TextForeignModel := ForeignTable.NameGo
	Otvet = strings.ReplaceAll(Otvet, "ForeignModel", TextForeignModel)

	//
	TextRequestColumnName := ""
	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	MassPK_and_Column := create_files.AppendColumn(MassPK, Column1)
	for _, ColumnPK1 := range MassPK_and_Column {
		RequestColumnName := create_files.Convert_GolangVariableToProtobufVariable(Table1, ColumnPK1, "m")
		//TextIDEqual = TextIDEqual + "\t" + ColumnPK1.NameGo + " := m." + ColumnPK1.NameGo + "\n"
		TextRequest := create_files.Find_RequestFieldName(Table1, ColumnPK1)
		TextRequestColumnName = TextRequestColumnName + "\t" + "Request." + TextRequest + " = " + RequestColumnName + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, "\tRequest.RequestFieldName = m.FieldName", TextRequestColumnName)

	//добавим URL
	ModelTableURL := create_files.Find_ModelTableURL(ForeignTableName)
	*TextGRPCClient = create_files.AddImport(*TextGRPCClient, ModelTableURL)

	return Otvet
}

// CreateFiles_FindModelBy_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_FindModelBy_Test(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	if len(types.MassFindModelBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMODELBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_findmodelby_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMODELBY_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCClientFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableURL)

		//GRPCClient_func_URL := create_files.Find_GRPCClient_func_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClient_func_URL)

		GRPClientURL := create_files.Find_GRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

		CrudFuncURL := create_files.Find_CrudFuncURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, CrudFuncURL)

	}

	//создание функций
	TextGRPCClientFunc := CreateFiles_FindModelBy_Test_Table(MapAll, Table1, TextTemplatedFunction)
	if TextGRPCClientFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextGRPCClientFunc

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//uuid
	TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	//alias
	TextGRPCClient = create_files.CheckAndAdd_ImportAlias(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCClient), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindModelBy_Test_Table - создаёт текст всех функций
func CreateFiles_FindModelBy_Test_Table(MapAll map[string]*types.Table, Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}
		Otvet1 := CreateFiles_FindModelBy_Test_Table1(MapAll, Table1, TextTemplateFunction, TableColumns1.Column)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindModelBy_Test_Table1 - создаёт текст всех функций
func CreateFiles_FindModelBy_Test_Table1(MapAll map[string]*types.Table, Table1 *types.Table, TextTemplateFunction string, Column1 *types.Column) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""

	//
	TextAssignFind := "\t" + `Otvet.FieldName = 0` + "\n"
	TextAssign := ""
	TextFieldName_TEST := ""

	//MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

	Underline := ""
	Comma := ""
	//for _, ColumnName1 := range MassColumnsString {
	//	Column1, ok := Table1.MapColumns[ColumnName1]
	//	if ok == false {
	//		log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
	//	}
	DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)
	//RequestFieldName := create_files.Find_RequestFieldName_FromMass(Column1, MassColumns)
	//TextAssign = TextAssign + "\t" + `Otvet.` + Column1.NameGo + ` = ` + DefaultValue + "\n"
	ColumnNameTranslit := Column1.NameGo
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + ColumnNameTranslit
	FieldNamesWithComma = FieldNamesWithComma + Comma + Column1.NameGo
	TextFieldName_TEST = TextFieldName_TEST + Comma + DefaultValue

	TextRequestColumnName := ""
	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	MassPK_and_Column := create_files.AppendColumn(MassPK, Column1)
	for _, ColumnPK1 := range MassPK_and_Column {
		Value := ""
		if ColumnPK1.IsPrimaryKey == true {
			Value = create_files.FindText_ColumnNameTest(ColumnPK1)
		} else {
			Value = create_files.FindText_DefaultValue(ColumnPK1.TypeGo)
		}
		TextAssign = TextAssign + "\t" + `Otvet.` + ColumnPK1.NameGo + ` = ` + Value + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, "\tRequest.RequestFieldName = m.FieldName", TextRequestColumnName)

	Underline = "_"
	Comma = ", "
	//}
	Otvet = strings.ReplaceAll(Otvet, TextAssignFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldName_TEST)

	//
	ForeignTableName := Column1.TableKey
	ForeignTable, ok := MapAll[ForeignTableName]
	if ok == false {
		log.Panic("Table not found: ", ForeignTableName)
	}

	//
	TextFindModelBy := "Find" + ForeignTable.NameGo_translit + "By"
	Otvet = strings.ReplaceAll(Otvet, "FindModelBy", TextFindModelBy)

	return Otvet
}
