package server_grpc_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_FindBy - создаёт 1 файл в папке server_grpc
func CreateFiles_FindBy(Table1 *types.Table) error {
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
	create_files.CreateDirectory(DirReadyTable)

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
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

	}

	//создание функций
	TextServerGRPCFunc := CreateFiles_FindBy_Table(Table1, TextTemplatedFunction)
	if TextServerGRPCFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextServerGRPCFunc

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "grpc_proto.", TextProto+".")

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена импортов на новые URL
	TextGRPCServer = create_files.Replace_RepositoryImportsURL(TextGRPCServer)

	//uuid
	TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	//alias
	TextGRPCServer = create_files.CheckAndAdd_ImportAlias(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.Delete_EmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindBy_Table - создаёт текст всех функций
func CreateFiles_FindBy_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFiles_FindBy_Table1(Table1, TextTemplateFunction, TableColumns1.ColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindBy_Table1 - создаёт текст всех функций
func CreateFiles_FindBy_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumnsString []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	TextAssign := ""

	MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

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
		//RequestFieldName := create_files.Find_RequestFieldName_FromMass(Column1, MassColumns)
		TextRequest, TextRequestCode := create_files.Convert_ProtobufVariableToGolangVariable_with_MassColumns(Column1, MassColumns, "Request.")
		TextRequest = create_files.Convert_ColumnToAlias(Table1, Column1, TextRequest)
		if TextRequestCode != "" {
			TextAssign = TextAssign + TextRequestCode + "\n"
		} else {
			TextAssign = TextAssign + "\t" + "Model." + Column1.NameGo + " = " + TextRequest + "\n"
		}
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo_translit
		FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo

		ProtoTypeName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
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

// CreateFiles_FindBy_Test - создаёт 1 файл в папке server_grpc
func CreateFiles_FindBy_Test(Table1 *types.Table) error {
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

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_findby_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

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
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		CrudStarterURL := create_files.Find_CrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudFuncURL)

		ConstantsURL := create_files.Find_ConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

	}

	//создание функций
	TextGRPCServerFunc := CreateFiles_FindBy_Test_Table(Table1, TextTemplatedFunction)
	if TextGRPCServerFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextGRPCServerFunc

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, "grpc_proto.", TextProto+".")

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена импортов на новые URL
	TextGRPCServer = create_files.Replace_RepositoryImportsURL(TextGRPCServer)

	//uuid
	TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	//alias
	TextGRPCServer = create_files.CheckAndAdd_ImportAlias(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.Delete_EmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindBy_Test_Table - создаёт текст всех функций
func CreateFiles_FindBy_Test_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFiles_FindBy_Test_Table1(Table1, TextTemplateFunction, TableColumns1.ColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindBy_Test_Table1 - создаёт текст всех функций
func CreateFiles_FindBy_Test_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumnsString []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""

	//
	TextAssignFind := "\t" + `Request.RequestFieldName = 0` + "\n"
	TextAssign := ""
	TextFieldName_TEST := ""

	MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

	Underline := ""
	Comma := ""
	RequestName := "Request_"
	for _, ColumnName1 := range MassColumnsString {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)
		RequestFieldName := create_files.Find_RequestFieldName_FromMass(Column1, MassColumns)
		TextAssign = TextAssign + "\t" + `Request.` + RequestFieldName + ` = ` + DefaultValue + "\n"
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo_translit
		FieldNamesWithComma = FieldNamesWithComma + Comma + Column1.NameGo
		TextFieldName_TEST = TextFieldName_TEST + Comma + DefaultValue

		ProtoTypeName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
		RequestName = RequestName + Underline + ProtoTypeName

		Underline = "_"
		Comma = ", "
	}
	Otvet = strings.ReplaceAll(Otvet, "RequestName", RequestName)
	Otvet = strings.ReplaceAll(Otvet, TextAssignFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldName_TEST)

	return Otvet
}
