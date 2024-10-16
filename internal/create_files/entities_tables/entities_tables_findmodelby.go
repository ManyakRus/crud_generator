package entities_tables

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
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()

	FilenameTemplateModel := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_FINDMODELBY_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyModel + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_MODEL + TableName + "_findmodelby.go"

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
	FilenameTemplateModelFunction := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_FINDMODELBY_FUNCTION_FILENAME
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
	TextModelFunc := CreateFiles_FindModelBy_Table(MapAll, Table1, &TextModel, TextTemplatedFunction)
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
	err = os.WriteFile(FilenameReady, []byte(TextModel), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindModelBy_Table - создаёт текст всех функций
func CreateFiles_FindModelBy_Table(MapAll map[string]*types.Table, Table1 *types.Table, TextModel *string, TextTemplateFunction string) string {
	Otvet := ""

	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}
		Otvet1 := CreateFiles_FindModelBy_Table1(MapAll, Table1, TextModel, TextTemplateFunction, TableColumns1.Column)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFiles_FindModelBy_Table1 - создаёт текст всех функций
func CreateFiles_FindModelBy_Table1(MapAll map[string]*types.Table, Table1 *types.Table, TextModel *string, TextTemplateFunction string, Column1 *types.Column) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""

	//
	Underline := ""
	Plus := ""
	FieldNamesWithUnderline = FieldNamesWithUnderline + Underline + Column1.NameGo
	FieldNamesWithComma = FieldNamesWithComma + Plus + Column1.NameGo
	Underline = "_"
	Plus = "+"
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
	ModelTableURL := create_files.Find_ModelTableURL(ForeignTableName)
	*TextModel = create_files.AddImport(*TextModel, ModelTableURL)

	return Otvet
}

// AddInterfaces_FindModelBy - добавляет функцию внутрь интерфейса
func AddInterfaces_FindModelBy(MapAll map[string]*types.Table, TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	if len(types.MassFindModelBy) == 0 {
		return Otvet
	}

	TextFunc := ""
	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}

		Column1 := TableColumns1.Column

		//
		ForeignTableName := Column1.TableKey
		ForeignTable, ok := MapAll[ForeignTableName]
		if ok == false {
			log.Panic("Table not found: ", ForeignTableName)
		}

		FieldNamesWithUnderline := ""
		FieldNamesWithUnderline = FieldNamesWithUnderline + Column1.NameGo
		TextFunc1 := "\n\tFind" + ForeignTable.NameGo + "By_" + FieldNamesWithUnderline + "(*" + Table1.NameGo + ") (" + ForeignTable.Name + "." + ForeignTable.NameGo + ",error)"
		TextFunc = TextFunc + TextFunc1

	}

	Otvet = create_files.AddInterfaceFunction(Otvet, TextFunc)

	return Otvet
}
