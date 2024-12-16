package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_FindModelBy - создаёт 1 файл в папке crud
func CreateFiles_FindModelBy(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	if len(types.MassFindModelBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDMODELBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_findmodelby.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDMODELBY_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrudFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		ConstantsURL := create_files.Find_DBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, ConstantsURL)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextCrud = create_files.Replace_Connect_WithApplicationName(TextCrud)

	}

	//создание функций
	TextCrudFunc := CreateFiles_FindModelBy_Table(MapAll, Table1, &TextCrud, TextTemplatedFunction)
	if TextCrudFunc == "" {
		return err
	}
	TextCrud = TextCrud + TextCrudFunc

	//создание текста
	TextCrud = create_files.Replace_TemplateModel_to_Model(TextCrud, Table1.NameGo)
	TextCrud = create_files.Replace_TemplateTableName_to_TableName(TextCrud, Table1.Name)
	TextCrud = create_files.AddText_ModuleGenerated(TextCrud)

	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//замена импортов на новые URL
	TextCrud = create_files.Replace_RepositoryImportsURL(TextCrud)

	//uuid
	TextCrud = create_files.CheckAndAdd_ImportUUID_FromText(TextCrud)

	//alias
	TextCrud = create_files.CheckAndAdd_ImportAlias(TextCrud)

	//fmt
	TextCrud = create_files.CheckAndAdd_ImportFmt(TextCrud)

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.Delete_EmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextCrud), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindModelBy_Table - создаёт текст всех функций
func CreateFiles_FindModelBy_Table(MapAll map[string]*types.Table, Table1 *types.Table, TextCrud *string, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}
		Otvet1 := CreateFiles_FindModelBy_Table1(MapAll, Table1, TextCrud, TextTemplateFunction, TableColumns1.Column)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindModelBy_Table1 - создаёт текст всех функций
func CreateFiles_FindModelBy_Table1(MapAll map[string]*types.Table, Table1 *types.Table, TextCrud *string, TextTemplateFunction string, Column1 *types.Column) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	TextWhere := ""
	//ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)

	//
	TextFind := "\t" + `//tx = tx.Where("ColumnName = ?", m.PrimaryKeyName)` + "\n"
	Underline := ""
	Plus := ""
	//TextWhere = TextWhere + "\t" + `tx = tx.Where("` + ColumnPK.Name + ` = ?", m.` + ColumnPK.NameGo + `)` + "\n"
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
	FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo
	Underline = "_"
	Plus = "+"

	//кроме помеченных на удаление
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		TextWhere = TextWhere + "\t" + `tx = tx.Where("is_deleted = ?", false)` + "\n"
	}

	//
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextWhere)
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
	Foreign_ColumnPK := create_files.Find_PrimaryKeyColumn(ForeignTable)
	TextForeignPrimaryKeyName := Foreign_ColumnPK.NameGo
	Otvet = strings.ReplaceAll(Otvet, "ForeignPrimaryKeyName", TextForeignPrimaryKeyName)

	//
	TextFieldName := Column1.NameGo
	Otvet = strings.ReplaceAll(Otvet, "FieldName", TextFieldName)

	//
	ModelTableURL := create_files.Find_ModelTableURL(ForeignTableName)
	*TextCrud = create_files.AddImport(*TextCrud, ModelTableURL)

	//ForeignPrimaryKeysWithQuotes
	ForeignColumnsPK := create_files.Find_PrimaryKeyColumns(ForeignTable)
	ForeignPrimaryKeysWithQuotes := create_files.ColumnNamesGo_WithQuotes(ForeignColumnsPK)
	Otvet = strings.ReplaceAll(Otvet, "ForeignPrimaryKeysWithQuotes", ForeignPrimaryKeysWithQuotes)

	//PrimaryKeysWithQuotes
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	PrimaryKeysWithQuotes := create_files.ColumnNamesGo_WithQuotes(ColumnsPK)
	Otvet = strings.ReplaceAll(Otvet, "PrimaryKeysWithQuotes", PrimaryKeysWithQuotes)

	return Otvet
}

// CreateFiles_FindModelBy_Test - создаёт 1 файл в папке crud
func CreateFiles_FindModelBy_Test(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	if len(types.MassFindModelBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDMODELBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_findmodelby_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDMODELBY_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrudFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextCrud = create_files.AddImport(TextCrud, CrudFuncURL)

		ConstantsURL := create_files.Find_ConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, ConstantsURL)

	}

	//создание функций
	TextCrudFunc := CreateFiles_FindModelBy_Test_Table(MapAll, Table1, TextTemplatedFunction)
	if TextCrudFunc == "" {
		return err
	}
	TextCrud = TextCrud + TextCrudFunc

	//создание текста
	TextCrud = create_files.Replace_TemplateModel_to_Model(TextCrud, Table1.NameGo)
	TextCrud = create_files.Replace_TemplateTableName_to_TableName(TextCrud, Table1.Name)
	TextCrud = create_files.AddText_ModuleGenerated(TextCrud)

	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//замена импортов на новые URL
	TextCrud = create_files.Replace_RepositoryImportsURL(TextCrud)

	//uuid
	TextCrud = create_files.CheckAndAdd_ImportUUID_FromText(TextCrud)

	//alias
	TextCrud = create_files.CheckAndAdd_ImportAlias(TextCrud)

	//удаление пустого импорта
	TextCrud = create_files.Delete_EmptyImport(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.Delete_EmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextCrud), config.Settings.FILE_PERMISSIONS)

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

	Underline := ""
	Comma := ""
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
	FieldNamesWithComma = FieldNamesWithComma + Comma + Column1.NameGo

	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)

	for _, ColumnPK1 := range ColumnsPK {
		//DefaultValue := create_files.FindText_DefaultValue(ColumnPK1.TypeGo)
		Value := create_files.FindText_ColumnNameTest(ColumnPK1)
		TextAssign = TextAssign + "\t" + `Otvet.` + ColumnPK1.NameGo + ` = ` + Value + "\n"
		TextFieldName_TEST = TextFieldName_TEST + Comma + Value
		Comma = ", "
	}

	Underline = "_"
	Comma = ", "
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
