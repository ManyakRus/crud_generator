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

// CreateFiles_FindModelBy - создаёт 1 файл в папке server_grpc
func CreateFiles_FindModelBy(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	if len(types.MassFindModelBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDMODELBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_findmodelby.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCServerFunction := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDMODELBY_FUNCTION_FILENAME
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
	TextServerGRPCFunc := CreateFiles_FindModelBy_Table(MapAll, Table1, TextTemplatedFunction)
	if TextServerGRPCFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextServerGRPCFunc

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

// CreateFiles_FindModelBy_Table - создаёт текст всех функций
func CreateFiles_FindModelBy_Table(MapAll map[string]*types.Table, Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}
		Otvet1 := CreateFiles_FindModelBy_Table1(MapAll, Table1, TextTemplateFunction, TableColumns1.Column)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindModelBy_Table1 - создаёт текст всех функций
func CreateFiles_FindModelBy_Table1(MapAll map[string]*types.Table, Table1 *types.Table, TextTemplateFunction string, Column1 *types.Column) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	TextAssign := ""

	//MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

	//
	TextFind := "\t" + `Model.FieldName = Request.RequestFieldName` + "\n"
	Underline := ""
	Plus := ""
	//RequestName := "Request_"
	//RequestFieldName := create_files.Find_RequestFieldName_FromMass(Column1, MassColumns)
	TextRequest, TextRequestCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
	if TextRequestCode != "" {
		TextAssign = TextAssign + TextRequestCode + "\n"
	} else {
		TextAssign = TextAssign + "\t" + "Model." + Column1.NameGo + " = " + TextRequest + "\n"
	}
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo_translit
	FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo

	//ProtoTypeName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
	//RequestName = RequestName + Underline + ProtoTypeName
	RequestName := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)

	Underline = "_"
	Plus = "+"
	//}

	Otvet = strings.ReplaceAll(Otvet, "RequestName", RequestName)
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)

	//найдём внешнюю таблицу
	ForeignTableName := Column1.TableKey
	ForeignTable, ok := MapAll[ForeignTableName]
	if ok == false {
		log.Panic("Table not found: ", ForeignTableName)
	}

	//
	TextFindModelBy := "Find" + ForeignTable.NameGo_translit + "By"
	Otvet = strings.ReplaceAll(Otvet, "FindModelBy", TextFindModelBy)

	//
	TextForeignModel := ForeignTable.NameGo
	Otvet = strings.ReplaceAll(Otvet, "ForeignModel", TextForeignModel)

	//
	TextIDEqual := ""
	TextModelFieldName := ""
	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	MassPK_and_Column := create_files.AppendColumn(MassPK, Column1)
	for _, ColumnPK1 := range MassPK_and_Column {
		RequestColumnName, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, ColumnPK1, "Request.")
		if GolangCode == "" {
			TextIDEqual = TextIDEqual + "\t" + ColumnPK1.NameGo + " := " + RequestColumnName + "\n"
		} else {
			TextIDEqual = TextIDEqual + GolangCode
		}
		TextModelFieldName = TextModelFieldName + "\t" + "Model." + ColumnPK1.NameGo + " = " + ColumnPK1.NameGo + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, "\tID := Request.RequestFieldName", TextIDEqual)
	Otvet = strings.ReplaceAll(Otvet, "\tModel.FieldName = ID", TextModelFieldName)

	return Otvet
}

// CreateFiles_FindModelBy_Test - создаёт 1 файл в папке server_grpc
func CreateFiles_FindModelBy_Test(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	if len(types.MassFindModelBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDMODELBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_findmodelby_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCServerFunction := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_FINDMODELBY_FUNCTION_TEST_FILENAME
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

		CrudFuncURL := create_files.Find_CrudFuncURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudFuncURL)

	}

	//создание функций
	TextGRPCServerFunc := CreateFiles_FindModelBy_Test_Table(MapAll, Table1, TextTemplatedFunction)
	if TextGRPCServerFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextGRPCServerFunc

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
	TextAssign := ""
	TextFieldName_TEST := ""

	//MassColumns := create_files.FindMass_Columns_from_MassColumnsString(Table1, MassColumnsString)

	Underline := ""
	Comma := ""
	//RequestName := "Request_Model_"
	//for _, ColumnName1 := range MassColumnsString {
	DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo_translit
	FieldNamesWithComma = FieldNamesWithComma + Comma + Column1.NameGo
	TextFieldName_TEST = TextFieldName_TEST + Comma + DefaultValue

	//ProtoTypeName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
	//RequestName = RequestName + Underline + ProtoTypeName
	RequestName := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)

	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	MassPK_and_Column := create_files.AppendColumn(MassPK, Column1)
	for _, ColumnPK1 := range MassPK_and_Column {
		RequestFieldName, _ := create_files.Convert_ProtobufVariableToGolangVariable(Table1, ColumnPK1, "Request.")
		TextAssign = TextAssign + "\t" + `` + RequestFieldName + ` = ` + DefaultValue + "\n"
	}
	//RequestFieldName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
	//TextAssign = TextAssign + "\t" + `` + RequestFieldName + ` = ` + DefaultValue + "\n"

	Underline = "_"
	Comma = ", "
	//}
	Otvet = strings.ReplaceAll(Otvet, "RequestName", RequestName)
	Otvet = strings.ReplaceAll(Otvet, "\t"+`Request.RequestFieldName = 0`+"\n", TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldName_TEST)

	//найдём внешнюю таблицу
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
