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

// CreateFilesFindBy - создаёт 1 файл в папке crud
func CreateFilesFindBy(Table1 *types.Table) error {
	var err error

	if len(types.MassFindBy_String) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()

	FilenameTemplateModel := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_FINDBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyModel + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_MODEL + TableName + "_findby.go"

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
	FilenameTemplateModelFunction := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_FINDBY_FUNCTION_FILENAME
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
	TextModelFunc := CreateFilesFindByTable(Table1, TextTemplatedFunction)
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
func CreateFilesFindByTable1(Table1 *types.Table, TextTemplateFunction string, MassColumns1 []string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	TextWhere := ""

	//
	TextFind := "\t" + `tx = tx.Where("ColumnName = ?", m.FieldName)` + "\n"
	Underline := ""
	Plus := ""
	for _, ColumnName1 := range MassColumns1 {
		Column1, ok := Table1.MapColumns[ColumnName1]
		if ok == false {
			log.Panic(Table1.Name + " .MapColumns[" + ColumnName1 + "] = false")
		}
		TextWhere = TextWhere + "\t" + `tx = tx.Where("` + ColumnName1 + ` = ?", m.` + Column1.NameGo + `)` + "\n"
		FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
		FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo
		Underline = "_"
		Plus = "+"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextWhere)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)

	return Otvet
}

// AddInterfacesFindBy - добавляет функцию внутрь интерфейса
func AddInterfacesFindBy(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	if len(types.MassFindBy_String) == 0 {
		return Otvet
	}

	TextFunc := ""
	for _, TableColumns1 := range types.MassFindBy_String {
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
		TextFunc1 := "\n\tFindBy_" + FieldNamesWithUnderline + "(*" + Table1.NameGo + ") error"
		TextFunc = TextFunc + TextFunc1

	}

	Otvet = create_files.AddInterfaceFunction(Otvet, TextFunc)

	return Otvet
}
