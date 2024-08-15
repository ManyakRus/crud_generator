package entities_tables

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

// CreateFilesFindMassBy - создаёт 1 файл в папке crud
func CreateFilesFindMassBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindMassBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()

	FilenameTemplateModel := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_FINDMASSBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyModel + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_MODEL + TableName + "_findmassby.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateModelFunction := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_FINDMASSBY_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateModelFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModelFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextModel = create_files.ReplacePackageName(TextModel, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextModel = create_files.DeleteTemplateRepositoryImports(TextModel)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextModel = create_files.AddImport(TextModel, ModelTableURL)

		ConstantsURL := create_files.FindDBConstantsURL()
		TextModel = create_files.AddImport(TextModel, ConstantsURL)

	}

	//создание функций
	TextModelFunc := CreateFilesFindMassByTable(Table1, TextTemplatedFunction)
	if TextModelFunc == "" {
		return err
	}
	TextModel = TextModel + TextModelFunc

	//создание текста
	TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextModel = config.Settings.TEXT_MODULE_GENERATED + TextModel

	//замена импортов на новые URL
	TextModel = create_files.ReplaceServiceURLImports(TextModel)

	//uuid
	TextModel = create_files.CheckAndAddImportUUID_FromText(TextModel)

	//alias
	TextModel = create_files.CheckAndAddImportAlias(TextModel)

	//удаление пустого импорта
	TextModel = create_files.DeleteEmptyImport(TextModel)

	//удаление пустых строк
	TextModel = create_files.DeleteEmptyLines(TextModel)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesFindMassByTable - создаёт текст всех функций
func CreateFilesFindMassByTable(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindMassBy {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFilesFindMassByTable1(Table1, TextTemplateFunction, TableColumns1.MassColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFilesFindMassByTable1 - создаёт текст всех функций
func CreateFilesFindMassByTable1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	ColumnNamesWithComma := ""

	//
	TextFind := "\t" + `tx = tx.Where("ColumnName = ?", m.FieldName)` + "\n"
	TextWhere := ""
	Underline := ""
	Plus := ""
	Comma := ""
	for _, ColumnName1 := range MassColumns1 {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		TextWhere = TextWhere + "\t" + `tx = tx.Where("` + ColumnName1 + ` = ?", m.` + Column1.NameGo + `)` + "\n"
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
		FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo
		ColumnNamesWithComma = ColumnNamesWithComma + Comma + Column1.Name

		Underline = "_"
		Plus = "+"
		Comma = ", "
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextWhere)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNamesWithComma", ColumnNamesWithComma)

	return Otvet
}

// CreateFilesFindMassByTest - создаёт 1 файл в папке crud
func CreateFilesFindMassByTest(Table1 *types.Table) error {
	var err error

	if len(types.MassFindMassBy) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDMASSBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_findmassby_test.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_FINDMASSBY_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateCrudFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrudFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextCrud = create_files.ReplacePackageName(TextCrud, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.DeleteTemplateRepositoryImports(TextCrud)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

	}

	//создание функций
	TextCrudFunc := CreateFilesFindMassByTestTable(Table1, TextTemplatedFunction)
	if TextCrudFunc == "" {
		return err
	}
	TextCrud = TextCrud + TextCrudFunc

	//создание текста
	TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//замена импортов на новые URL
	TextCrud = create_files.ReplaceServiceURLImports(TextCrud)

	//uuid
	TextCrud = create_files.CheckAndAddImportUUID_FromText(TextCrud)

	//alias
	TextCrud = create_files.CheckAndAddImportAlias(TextCrud)

	//удаление пустого импорта
	TextCrud = create_files.DeleteEmptyImport(TextCrud)

	//удаление пустых строк
	TextCrud = create_files.DeleteEmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextCrud), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesFindMassByTestTable - создаёт текст всех функций
func CreateFilesFindMassByTestTable(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindMassBy {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFilesFindMassByTestTable1(Table1, TextTemplateFunction, TableColumns1.MassColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFilesFindMassByTestTable1 - создаёт текст всех функций
func CreateFilesFindMassByTestTable1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
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
		DefaultValue := create_files.FindTextDefaultValue(Column1.TypeGo)
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

// AddInterfacesFindMassBy - добавляет функцию внутрь интерфейса
func AddInterfacesFindMassBy(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	if len(types.MassFindMassBy) == 0 {
		return Otvet
	}

	TextFunc := ""
	for _, TableColumns1 := range types.MassFindMassBy {
		if TableColumns1.TableName != Table1.Name {
			continue
		}

		FieldNamesWithUnderline := ""
		Underline := ""
		for _, ColumnName1 := range TableColumns1.MassColumnNames {
			Column1, ok := Table1.MapColumns[ColumnName1]
			if ok == false {
				log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
			}
			FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
			Underline = "_"
		}
		TextFunc1 := "\n\tFindMassBy_" + FieldNamesWithUnderline + "(*" + Table1.NameGo + ") ([]" + Table1.NameGo + ", error)"
		TextFunc = TextFunc + TextFunc1

	}

	Otvet = create_files.AddInterfaceFunction(Otvet, TextFunc)

	return Otvet
}
