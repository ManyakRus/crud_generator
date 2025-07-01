package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strconv"
	"strings"
)

// CreateFiles_FindBy - создаёт 1 файл в папке crud
func CreateFiles_FindBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_findby.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_FILENAME
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
	TextCrudFunc := CreateFiles_FindBy_Table(Table1, TextTemplatedFunction)
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
func CreateFiles_FindBy_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
	Otvet := TextTemplateFunction

	TableAlias := create_files.Find_TableAlias(Table1)

	//MassColumns1
	FieldNamesWithUnderline := ""
	ReplaceWhereID := ""
	Underline := ""
	Number := 0
	for _, sColumn1 := range MassColumns1 {
		Column1, ok := Table1.MapColumns[sColumn1]
		if ok == false {
			log.Panic("Column not found: ", sColumn1)
		}
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
		Number = Number + 1
		sNumber := strconv.Itoa(Number)
		ReplaceWhereID = ReplaceWhereID + "\tand " + TableAlias + "." + Column1.Name + " = $" + sNumber + "\n"

		Underline = "_"
	}
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceWhereID", ReplaceWhereID)

	//все колонки
	ReplaceAllFieldsWithComma := ""
	//CommaNewline := ""
	CommaNewline2 := ""
	MassColumns := micro.MassFrom_Map(Table1.MapColumns)
	for _, Column1 := range MassColumns {
		//кроме ненужных колонок
		if create_files.Is_Need_Сolumn(Column1) == false {
			continue
		}
		//if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
		//	continue
		//}

		//Number = Number + 1
		//sNumber := strconv.Itoa(Number)
		//if Column1.IsNullable == true {
		//	DefaultValueSQL := create_files.FindText_DefaultValueSQL_NotNull(Column1.TypeGo)
		//	ReplaceTextSQL = ReplaceTextSQL + CommaNewline + "COALESCE(" + TableAlias + "." + Column1.Name + ", " + DefaultValueSQL + ") as " + Column1.Name
		//} else {
		//	ReplaceTextSQL = ReplaceTextSQL + CommaNewline + TableAlias + "." + Column1.Name
		//}
		ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + "&m." + Column1.NameGo

		//if slice.Contains(MassColumns1, Column1.Name) == true {
		//	ReplaceWhereID = ReplaceWhereID + "\tand " + TableAlias + "." + Column1.Name + " = $" + sNumber + "\n"
		//}

		//CommaNewline = ",\n\t"
		CommaNewline2 = ",\n\t\t"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceAllFieldsWithComma", ReplaceAllFieldsWithComma)
	//	ReplaceTextSQL = ReplaceTextSQL + `
	//FROM
	//	` + Table1.Name + ` as ` + TableAlias + "\n"
	//	Otvet = strings.ReplaceAll(Otvet, "ReplaceTextSQL", ReplaceTextSQL)

	//
	FieldNamesWithComma := ""

	//
	ReplaceFieldNamesFormat := ""
	ReplaceFieldsWithComma := ""
	Plus := ""
	Comma := ""
	for _, ColumnName1 := range MassColumns1 {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo
		ReplaceFieldNamesFormat = ReplaceFieldNamesFormat + Comma + Column1.NameGo + ": %v"
		ReplaceFieldsWithComma = ReplaceFieldsWithComma + Comma + "m." + Column1.NameGo

		Comma = ", "
		Underline = "_"
		Plus = "+"
	}

	//кроме помеченных на удаление
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		ReplaceWhereID = ReplaceWhereID + "\tand " + TableAlias + ".is_deleted = false\n"
	}

	//
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceFieldNamesFormat", ReplaceFieldNamesFormat)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceFieldsWithComma", ReplaceFieldsWithComma)

	return Otvet
}

// CreateFiles_FindBy_Test - создаёт 1 файл в папке crud
func CreateFiles_FindBy_Test(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDBY_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_findby_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDBY_FUNCTION_TEST_FILENAME
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
	TextCrudFunc := CreateFiles_FindBy_Test_Table(Table1, TextTemplatedFunction)
	if TextCrudFunc == "" {
		return err
	}
	TextCrud = TextCrud + TextCrudFunc

	//замена "postgres_gorm.Connect_WithApplicationName("
	TextCrud = create_files.Replace_Connect_WithApplicationName(TextCrud)

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
func CreateFiles_FindBy_Test_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
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
	for _, ColumnName1 := range MassColumns1 {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)
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
