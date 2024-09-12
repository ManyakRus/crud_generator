package crud_tables

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

// CreateFilesReadAll - создаёт 1 файл в папке crud
func CreateFilesReadAll(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READALL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_readall.go"

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
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READALL_FUNCTION_FILENAME
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

	}

	//создание функций
	TextCrudFunc := CreateFilesReadAllTable(Table1, TextTemplatedFunction)
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
	err = os.WriteFile(FilenameReady, []byte(TextCrud), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesReadAllTable - создаёт текст всех функций
func CreateFilesReadAllTable(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for TableReadAll1, _ := range types.MapReadAll {
		if TableReadAll1.Name != Table1.Name {
			continue
		}
		Otvet1 := CreateFilesReadAllTable1(Table1, TextTemplateFunction)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFilesReadAllTable1 - создаёт текст всех функций
func CreateFilesReadAllTable1(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""
	ColumnNamesWithComma := ""

	//
	TextFind := "\t" + `tx = tx.Where("ColumnName = ?", m.FieldName)` + "\n"
	TextWhere := ""
	//кроме помеченных на удаление
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		TextWhere = TextWhere + "\t" + `tx = tx.Where("is_deleted = ?", false)` + "\n"
	}

	//
	FuncName := constants.TEXT_READALL
	Otvet = strings.ReplaceAll(Otvet, "ReadAll_FieldNamesWithUnderline", FuncName)
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	ColumnNamesWithComma = create_files.Find_ColumnNamesWithComma(ColumnsPK)
	Otvet = strings.ReplaceAll(Otvet, ", m *lawsuit_status_types.LawsuitStatusType", "")
	Otvet = strings.ReplaceAll(Otvet, "m *lawsuit_status_types.LawsuitStatusType", "")
	Otvet = strings.ReplaceAll(Otvet, "(ctx, db, m)", "(ctx, db)")

	//
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextWhere)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPlus", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ColumnNamesWithComma", ColumnNamesWithComma)

	return Otvet
}

// CreateFilesReadAllTest - создаёт 1 файл в папке crud
func CreateFilesReadAllTest(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READALL_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_readall_test.go"

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
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READALL_FUNCTION_TEST_FILENAME
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

		CrudFuncURL := create_files.Find_CrudFuncURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, CrudFuncURL)

	}

	//создание функций
	TextCrudFunc := CreateFilesReadAllTestTable(Table1, TextTemplatedFunction)
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
	err = os.WriteFile(FilenameReady, []byte(TextCrud), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesReadAllTestTable - создаёт текст всех функций
func CreateFilesReadAllTestTable(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	for TableReadAll1, _ := range types.MapReadAll {
		if TableReadAll1.Name != Table1.Name {
			continue
		}
		Otvet1 := CreateFilesReadAllTestTable1(Table1, TextTemplateFunction)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// CreateFilesReadAllTestTable1 - создаёт текст всех функций
func CreateFilesReadAllTestTable1(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	//
	FieldNamesWithUnderline := ""
	FieldNamesWithComma := ""

	//
	TextAssignFind := "\t" + `Otvet.FieldName = 0` + "\n"
	TextAssign := ""
	TextFieldsDefaultValue := ""

	FuncName := constants.TEXT_READALL
	Otvet = strings.ReplaceAll(Otvet, "ReadAll_FieldNamesWithUnderline", FuncName)
	Otvet = strings.ReplaceAll(Otvet, "(&Otvet)", "()")

	//
	Otvet = strings.ReplaceAll(Otvet, TextAssignFind, TextAssign)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithUnderline", FieldNamesWithUnderline)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesDefault", TextFieldsDefaultValue)

	return Otvet
}
