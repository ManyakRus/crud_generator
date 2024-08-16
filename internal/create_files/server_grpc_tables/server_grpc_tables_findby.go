package server_grpc_tables

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

// CreateFilesFindBy - создаёт 1 файл в папке server_grpc
func CreateFilesFindBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_findby.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCServerFunction := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		CrudTableURL := create_files.FindCrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

	}

	//создание функций
	TextServerGRPCFunc := CreateFilesFindByTable(Table1, TextTemplatedFunction)
	if TextServerGRPCFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextServerGRPCFunc

	//создание текста
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена импортов на новые URL
	TextGRPCServer = create_files.ReplaceServiceURLImports(TextGRPCServer)

	//uuid
	TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

	//alias
	TextGRPCServer = create_files.CheckAndAddImportAlias(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.DeleteEmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

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
	TextFind := "\t" + `Model.FieldName = Request.RequestFieldName` + "\n"
	Underline := ""
	Plus := ""
	RequestName := "Request_"
	for _, ColumnName1 := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		//RequestFieldName := create_files.FindRequestFieldName_FromMass(Column1, MassColumns)
		TextRequest, TextRequestCode := create_files.ConvertProtobufVariableToGolangVariable_with_MassColumns(Column1, MassColumns, "Request.")
		if TextRequestCode != "" {
			TextAssign = TextAssign + TextRequestCode + "\n"
		} else {
			TextAssign = TextAssign + "\t" + "Model." + Column1.NameGo + " = " + TextRequest + "\n"
		}
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

// CreateFilesFindByTest - создаёт 1 файл в папке server_grpc
func CreateFilesFindByTest(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_findby_test.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCServerFunction := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDBY_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		CrudStarterURL := create_files.FindCrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

	}

	//создание функций
	TextGRPCServerFunc := CreateFilesFindByTestTable(Table1, TextTemplatedFunction)
	if TextGRPCServerFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextGRPCServerFunc

	//создание текста
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена импортов на новые URL
	TextGRPCServer = create_files.ReplaceServiceURLImports(TextGRPCServer)

	//uuid
	TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

	//alias
	TextGRPCServer = create_files.CheckAndAddImportAlias(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.DeleteEmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

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
	TextAssignFind := "\t" + `Request.RequestFieldName = 0` + "\n"
	TextAssign := ""
	TextFieldName_TEST := ""

	MassColumns := create_files.FindMassColumns_from_MassColumnsString(Table1, MassColumnsString)

	Underline := ""
	Comma := ""
	for _, ColumnName1 := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		DefaultValue := create_files.FindTextDefaultValue(Column1.TypeGo)
		RequestFieldName := create_files.FindRequestFieldName_FromMass(Column1, MassColumns)
		TextAssign = TextAssign + "\t" + `Request.` + RequestFieldName + ` = ` + DefaultValue + "\n"
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
