package entities_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"io/fs"
	"os"
	"strings"
)

// CreateFiles_FindMassBy - создаёт 1 файл в папке crud
func CreateFiles_FindMassBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindMassBy_String) == 0 {
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
	TextModel = create_files.Replace_PackageName(TextModel, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)

		//ModelTableURL := create_files.Find_ModelTableURL(TableName)
		//TextModel = create_files.AddImport(TextModel, ModelTableURL)

		ConstantsURL := create_files.Find_DBConstantsURL()
		TextModel = create_files.AddImport(TextModel, ConstantsURL)

	}

	//создание функций
	TextModelFunc := CreateFiles_FindMassBy_Table(Table1, TextTemplatedFunction)
	if TextModelFunc == "" {
		return err
	}
	TextModel = TextModel + TextModelFunc

	//создание текста
	TextModel = create_files.Replace_TemplateModel_to_Model(TextModel, Table1.NameGo)
	TextModel = create_files.Replace_TemplateTableName_to_TableName(TextModel, Table1.Name)
	TextModel = create_files.AddText_ModuleGenerated(TextModel)

	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextModel = config.Settings.TEXT_MODULE_GENERATED + TextModel

	//замена импортов на новые URL
	TextModel = create_files.Replace_RepositoryImportsURL(TextModel)

	//uuid
	TextModel = create_files.CheckAndAdd_ImportUUID_FromText(TextModel)

	//alias
	TextModel = create_files.CheckAndAdd_ImportAlias(TextModel)

	//удаление пустого импорта
	TextModel = create_files.Delete_EmptyImport(TextModel)

	//удаление пустых строк
	TextModel = create_files.Delete_EmptyLines(TextModel)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextModel), fs.FileMode(config.Settings.FILE_PERMISSIONS))

	return err
}

// CreateFiles_FindMassBy_Table - создаёт текст всех функций
func CreateFiles_FindMassBy_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindMassBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFiles_FindMassBy_Table1(Table1, TextTemplateFunction, TableColumns1.MassColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindMassBy_Table1 - создаёт текст всех функций
func CreateFiles_FindMassBy_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
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

	//функция ReadAll()
	if len(MassColumns1) == 0 {
		FuncName := config.Settings.TEXT_READALL
		Otvet = strings.ReplaceAll(Otvet, "FindMassBy_FieldNamesWithUnderline", FuncName)
		Otvet = strings.ReplaceAll(Otvet, "(m)", "()")
	}

	//
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextWhere)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNamesWithComma", ColumnNamesWithComma)

	return Otvet
}

// CreateFiles_FindMassBy_Test - создаёт 1 файл в папке crud
func CreateFiles_FindMassBy_Test(Table1 *types.Table) error {
	var err error

	if len(types.MassFindMassBy_String) == 0 {
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
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

	}

	//создание функций
	TextCrudFunc := CreateFiles_FindMassBy_Test_Table(Table1, TextTemplatedFunction)
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
	err = os.WriteFile(FilenameReady, []byte(TextCrud), fs.FileMode(config.Settings.FILE_PERMISSIONS))

	return err
}

// CreateFiles_FindMassBy_Test_Table - создаёт текст всех функций
func CreateFiles_FindMassBy_Test_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindMassBy_String {
		if TableColumns1.TableName != Table1.Name {
			continue
		}
		Otvet1 := CreateFiles_FindMassBy_Test_Table1(Table1, TextTemplateFunction, TableColumns1.MassColumnNames)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindMassBy_Test_Table1 - создаёт текст всех функций
func CreateFiles_FindMassBy_Test_Table1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
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

	//функция ReadAll()
	if len(MassColumns1) == 0 {
		FuncName := config.Settings.TEXT_READALL
		Otvet = strings.ReplaceAll(Otvet, "FindMassBy_FieldNamesWithUnderline", FuncName)
	}

	//
	Otvet = strings.ReplaceAll(Otvet, TextAssignFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldName_TEST)

	return Otvet
}

// AddInterfaces_FindMassBy - добавляет функцию внутрь интерфейса
func AddInterfaces_FindMassBy(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	if len(types.MassFindMassBy_String) == 0 {
		return Otvet
	}

	TextFunc := ""
	for _, TableColumns1 := range types.MassFindMassBy_String {
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

		TextFunc1 := ""
		FuncName := "FindMassBy_" + FieldNamesWithUnderline
		//функция ReadAll()
		if len(TableColumns1.MassColumnNames) == 0 {
			FuncName = config.Settings.TEXT_READALL
			TextFunc1 = "\n\t" + FuncName + "() ([]" + Table1.NameGo + ", error)"
		} else {
			TextFunc1 = "\n\t" + FuncName + "(*" + Table1.NameGo + ") ([]" + Table1.NameGo + ", error)"
		}

		//
		TextFunc = TextFunc + TextFunc1

	}

	Otvet = create_files.AddInterfaceFunction(Otvet, TextFunc)

	return Otvet
}
