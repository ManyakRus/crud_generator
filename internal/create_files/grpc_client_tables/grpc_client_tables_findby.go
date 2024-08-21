package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFilesFindBy - создаёт 1 файл в папке grpc_client
func CreateFilesFindBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_findby.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCClientFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.ReplacePackageName(TextGRPCClient, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ProtoURL)

		GRPCClient_func_URL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClient_func_URL)

		GRPCConstantsURL := create_files.FindGRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCConstantsURL)

		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		NRPC_Client_URL := create_files.FindNRPC_Client_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, NRPC_Client_URL)

	}

	//создание функций
	TextClientGRPCFunc := CreateFilesFindByTable(Table1, TextTemplatedFunction)
	if TextClientGRPCFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextClientGRPCFunc

	//создание текста
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//замена импортов на новые URL
	TextGRPCClient = create_files.ReplaceServiceURLImports(TextGRPCClient)

	//uuid
	TextGRPCClient = create_files.CheckAndAddImportUUID_FromText(TextGRPCClient)

	//alias
	TextGRPCClient = create_files.CheckAndAddImportAlias(TextGRPCClient)

	//time
	TextGRPCClient = create_files.CheckAndAddImportTime_FromText(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.DeleteEmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.DeleteEmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesFindByTable - создаёт текст всех функций
func CreateFilesFindByTable(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFilesFindByTable1(Table1, TextTemplateFunction, TableColumns1.MassColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFilesFindByTable1 - создаёт текст всех функций
func CreateFilesFindByTable1(Table1 *types.Table, TextTemplateFunction string, MassColumnsString []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	TextAssign := ""

	MassColumns := create_files.FindMassColumns_from_MassColumnsString(Table1, MassColumnsString)

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
		TextRequest := create_files.FindRequestFieldName_FromMass(Column1, MassColumns)
		ValueM := create_files.ConvertGolangVariableToProtobufVariable(Table1, Column1, "m")
		TextAssign = TextAssign + "\tRequest." + TextRequest + " = " + ValueM + "\n"
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
		FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo

		ProtoTypeName := create_files.ConvertGolangTypeNameToProtobufFieldName(Column1.TypeGo)
		RequestName = RequestName + Underline + ProtoTypeName

		Underline = "_"
		Plus = "+"
	}

	Otvet = strings.ReplaceAll(Otvet, "RequestName", RequestName)
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)

	return Otvet
}

// CreateFilesFindByTest - создаёт 1 файл в папке grpc_client
func CreateFilesFindByTest(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_findby_test.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_FINDBY_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCClientFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.ReplacePackageName(TextGRPCClient, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableURL)

		GRPCClient_func_URL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClient_func_URL)

		GRPClientURL := create_files.FindGRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		GRPClientTableURL := create_files.FindGRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)
	}

	//создание функций
	TextGRPCClientFunc := CreateFilesFindByTestTable(Table1, TextTemplatedFunction)
	if TextGRPCClientFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextGRPCClientFunc

	//создание текста
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//замена импортов на новые URL
	TextGRPCClient = create_files.ReplaceServiceURLImports(TextGRPCClient)

	//uuid
	TextGRPCClient = create_files.CheckAndAddImportUUID_FromText(TextGRPCClient)

	//alias
	TextGRPCClient = create_files.CheckAndAddImportAlias(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.DeleteEmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.DeleteEmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesFindByTestTable - создаёт текст всех функций
func CreateFilesFindByTestTable(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFilesFindByTestTable1(Table1, TextTemplateFunction, TableColumns1.MassColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFilesFindByTestTable1 - создаёт текст всех функций
func CreateFilesFindByTestTable1(Table1 *types.Table, TextTemplateFunction string, MassColumnsString []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""

	//
	TextAssignFind := "\t" + `Otvet.FieldName = 0` + "\n"
	TextAssign := ""
	TextFieldName_TEST := ""

	//MassColumns := create_files.FindMassColumns_from_MassColumnsString(Table1, MassColumnsString)

	Underline := ""
	Comma := ""
	for _, ColumnName1 := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		DefaultValue := create_files.FindTextDefaultValue(Column1.TypeGo)
		//RequestFieldName := create_files.FindRequestFieldName_FromMass(Column1, MassColumns)
		TextAssign = TextAssign + "\t" + `Otvet.` + Column1.NameGo + ` = ` + DefaultValue + "\n"
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
		FieldNamesWithComma = FieldNamesWithComma + Comma + Column1.NameGo
		TextFieldName_TEST = TextFieldName_TEST + Comma + DefaultValue

		Underline = "_"
		Comma = ", "
	}
	Otvet = strings.ReplaceAll(Otvet, TextAssignFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldName_TEST)

	return Otvet
}
