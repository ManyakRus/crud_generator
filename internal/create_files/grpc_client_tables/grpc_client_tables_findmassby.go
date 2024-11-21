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

// CreateFiles_FindMassBy - создаёт 1 файл в папке grpc_client
func CreateFiles_FindMassBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindMassBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_findmassby.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_FILENAME
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

		//NRPC
		if config.Settings.NEED_CREATE_NRPC == true {
			//
			NRPC_Client_URL := create_files.Find_NRPC_Client_URL()
			TextGRPCClient = create_files.AddImport(TextGRPCClient, NRPC_Client_URL)
		}

	}

	//создание функций
	TextGRPCClientFunc := CreateFiles_FindMassBy_Table(Table1, TextTemplatedFunction)
	if TextGRPCClientFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextGRPCClientFunc

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, "grpc_proto.", TextProto+".")

	//NRPC
	if config.Settings.NEED_CREATE_NRPC == true {
		//уберём "//"
		TextGRPCClient = Replace_NRPC_CLIENT(TextGRPCClient)
	}

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

// CreateFiles_FindMassBy_Table - создаёт текст всех функций
func CreateFiles_FindMassBy_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindMassBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFiles_FindMassBy_Table1(Table1, TextTemplateFunction, TableColumns1.ColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindMassBy_Table1 - создаёт текст всех функций
func CreateFiles_FindMassBy_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumnsString []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	TextAssign := ""

	MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

	//
	TextFind := "\t" + `Request.RequestFieldName = m.FieldName` + "\n"
	Underline := ""
	Plus := ""
	RequestName := "Request_"
	for _, ColumnName1 := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		TextRequest := create_files.Find_RequestFieldName_FromMass(Column1, MassColumns)
		ColumnName := Column1.NameGo
		ColumnNameTranslit := Column1.NameGo_translit
		TextAssign = TextAssign + "\tRequest." + TextRequest + " = m." + ColumnName + "\n"
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + ColumnNameTranslit
		FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo

		ProtoTypeName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
		RequestName = RequestName + Underline + ProtoTypeName

		Underline = "_"
		Plus = "+"
	}

	//функция ReadAll()
	if len(MassColumnsString) == 0 {
		FuncName := config.Settings.TEXT_READALL
		Otvet = strings.ReplaceAll(Otvet, "FindMassBy_FieldNamesWithUnderline", FuncName)
		Otvet = strings.ReplaceAll(Otvet, ".RequestName", ".Request_Empty")
		Otvet = strings.ReplaceAll(Otvet, "(m *lawsuit_status_types.LawsuitStatusType)", "()")
	}

	//
	Otvet = strings.ReplaceAll(Otvet, "RequestName", RequestName)
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)

	return Otvet
}

// CreateFiles_FindMassBy_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_FindMassBy_Test(Table1 *types.Table) error {
	var err error

	if len(types.MassFindMassBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_findmassby_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDMASSBY_FUNCTION_TEST_FILENAME
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

		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, CrudFuncURL)
	}

	//создание функций
	TextGRPCClientFunc := CreateFiles_FindMassBy_Test_Table(Table1, TextTemplatedFunction)
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

// CreateFiles_FindMassBy_Test_Table - создаёт текст всех функций
func CreateFiles_FindMassBy_Test_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindMassBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFiles_FindMassBy_Test_Table1(Table1, TextTemplateFunction, TableColumns1.ColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindMassBy_Test_Table1 - создаёт текст всех функций
func CreateFiles_FindMassBy_Test_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumnsString []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""

	//
	TextAssignFind := "\t" + `Model1.FieldName = 0` + "\n"
	TextAssign := ""
	TextFieldName_TEST := ""

	//MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

	Underline := ""
	Comma := ""
	for _, ColumnName1 := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)
		//RequestFieldName := create_files.Find_RequestFieldName_FromMass(Column1, MassColumns)
		TextAssign = TextAssign + "\t" + `Model1.` + Column1.NameGo + ` = ` + DefaultValue + "\n"
		ColumnNameTranslit := Column1.NameGo_translit
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + ColumnNameTranslit
		FieldNamesWithComma = FieldNamesWithComma + Comma + Column1.NameGo
		TextFieldName_TEST = TextFieldName_TEST + Comma + DefaultValue

		Underline = "_"
		Comma = ", "
	}

	//функция ReadAll()
	if len(MassColumnsString) == 0 {
		FuncName := config.Settings.TEXT_READALL
		Otvet = strings.ReplaceAll(Otvet, "FindMassBy_FieldNamesWithUnderline", FuncName)
		Otvet = strings.ReplaceAll(Otvet, ".RequestName", ".Request_Empty")
		Otvet = strings.ReplaceAll(Otvet, "(&Model1)", "()")
	}

	//
	Otvet = strings.ReplaceAll(Otvet, TextAssignFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldName_TEST)

	return Otvet
}
