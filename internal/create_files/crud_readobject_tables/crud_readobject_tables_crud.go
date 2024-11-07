package crud_readobject_tables

import (
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_ReadObject - создаёт 1 файл в папке crud
func CreateFiles_ReadObject(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATES_CRUD_READOBJECT_FOLDERNAME + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATES_CRUD_READOBJECT_FOLDERNAME + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READOBJECT_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD_READOBJECT + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD_READOBJECT + TableName + ".go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READOBJECT_FUNCTION_FILENAME
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

		////
		//ModelTableURL := create_files.Find_ModelTableURL(TableName)
		//TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		//crud foreign
		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, CrudTableURL)

		//
		ObjectTableURL := create_files.Find_ObjectTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ObjectTableURL)

		//
		ConstantsURL := create_files.Find_DBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, ConstantsURL)

	}

	//
	//FieldNamesWithPercent
	FieldNamesWithPercent := create_files.Find_FieldNamesWithPercent_from_Table(Table1)
	TextCrud = strings.ReplaceAll(TextCrud, "FieldNamesWithPercent", FieldNamesWithPercent)

	//FieldNamesWithCommaM
	FieldNamesWithCommaM := create_files.Find_FieldNamesWithComma_from_Table_VariableName(Table1, "m")
	TextCrud = strings.ReplaceAll(TextCrud, "FieldNamesWithCommaM", FieldNamesWithCommaM)

	//создание текста
	TextCrud = create_files.Replace_TemplateModel_to_Model(TextCrud, Table1.NameGo)
	TextCrud = create_files.Replace_TemplateTableName_to_TableName(TextCrud, Table1.Name)
	TextCrud = create_files.AddText_ModuleGenerated(TextCrud)

	//создание функций
	TextCrud, TextCrudFunc := CreateFiles_ReadObjectTable(MapAll, Table1, TextCrud, TextTemplatedFunction)
	//if TextCrudFunc == "" {
	//	return err
	//}
	//TextCrud = TextCrud + TextCrudFunc
	TextCrud = strings.ReplaceAll(TextCrud, "\t//TextFillManyFields", TextCrudFunc)

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

// CreateFiles_ReadObjectTable - создаёт текст всех функций
func CreateFiles_ReadObjectTable(MapAll map[string]*types.Table, Table1 *types.Table, TextCrud0, TextTemplateFunction0 string) (TextCrud, TextTemplateFunction string) {
	TextCrud = TextCrud0
	TextTemplateFunction = ""

	//
	MassImports := make([]string, 0)

	//
	for _, Column1 := range Table1.MapColumns {

		IsGoodColumn := create_files.IsGood_Column(Column1)
		if IsGoodColumn != nil {
			continue
		}

		IsForeignColumn := create_files.IsForeignColumn(MapAll, Column1)
		if IsForeignColumn == false {
			continue
		}

		TableF, _ := create_files.Find_TableF_ColumnF(MapAll, Column1)
		if TableF == nil {
			err := fmt.Errorf("Find_TableF_ColumnF() error: Foreign table for column %s not found", Column1.Name)
			log.Error(err)
			break
		}
		TableNameF := TableF.Name

		//шаблон функции
		TextTemplateFunction1 := CreateFiles_ReadObject_Table1(MapAll, Table1, Column1, TextTemplateFunction0)
		TextTemplateFunction = TextTemplateFunction + TextTemplateFunction1

		//добавим импорты
		//Model
		ModelFTableURL := create_files.Find_ModelTableURL(TableNameF)
		TextCrud = create_files.AddImport(TextCrud, ModelFTableURL)

		//crud foreign
		MassImports = append(MassImports, TableNameF)
	}

	//заполним импорты отсортированно
	for _, Import1 := range MassImports {
		CrudFTableURL := create_files.Find_CrudTableURL(Import1)
		TextCrud = create_files.AddImport(TextCrud, CrudFTableURL)

	}

	return TextCrud, TextTemplateFunction
}

// CreateFiles_ReadObject_Table1 - создаёт текст всех функций
func CreateFiles_ReadObject_Table1(MapAll map[string]*types.Table, Table1 *types.Table, Column1 *types.Column, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	TableFK, ColumnFK := create_files.Find_TableF_ColumnF(MapAll, Column1)
	ColumnPKF := create_files.Find_PrimaryKeyColumn(TableFK)

	//IntFromAlias
	IntFromAlias := create_files.Convert_GolangVariableToProtobufVariableType(Table1, Column1, "ModelNameColumnID", ColumnFK.TypeGo)
	Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(ModelNameColumnID)", IntFromAlias)

	//PrimaryKeyNameFK
	PrimaryKeyNameFK := ColumnPKF.NameGo
	Otvet = strings.ReplaceAll(Otvet, "PrimaryKeyNameFK", PrimaryKeyNameFK)

	//FieldNameForeign
	FieldNameForeign := ColumnFK.NameGo
	Otvet = strings.ReplaceAll(Otvet, "FieldNameForeign", FieldNameForeign)

	//TableNameForeign
	TableNameForeign := TableFK.Name
	Otvet = strings.ReplaceAll(Otvet, "TableNameForeign", TableNameForeign)

	//FieldNamesWithPercent
	FieldNamesWithPercent := create_files.Find_FieldNamesWithPercent_from_Table(Table1)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithPercent", FieldNamesWithPercent)

	//FieldNamesWithCommaM
	FieldNamesWithCommaM := create_files.Find_FieldNamesWithComma_from_Table_VariableName(Table1, "m")
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithCommaM", FieldNamesWithCommaM)

	//FieldNamesWithComma
	FieldNamesWithComma := create_files.Find_FieldNamesWithComma_from_Table(Table1)
	Otvet = strings.ReplaceAll(Otvet, "FieldNamesWithComma", FieldNamesWithComma)

	//ModelNameForeign
	ModelNameForeign := TableFK.NameGo
	Otvet = strings.ReplaceAll(Otvet, "ModelNameForeign", ModelNameForeign)

	//FieldNameTable
	FieldNameTable := Column1.NameGo
	Otvet = strings.ReplaceAll(Otvet, "FieldNameTable", FieldNameTable)

	//ModelNameColumn
	ModelNameColumn := create_files.Find_ObjectColumnModelName(Column1.NameGo)
	Otvet = strings.ReplaceAll(Otvet, "ModelNameColumn", ModelNameColumn)

	//crud_
	Otvet = strings.ReplaceAll(Otvet, " crud_", " "+config.Settings.PREFIX_CRUD)

	return Otvet
}

// CreateFiles_ReadObject_Test - создаёт 1 файл в папке crud
func CreateFiles_ReadObject_Test(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATES_CRUD_READOBJECT_FOLDERNAME + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATES_CRUD_READOBJECT_FOLDERNAME + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READOBJECT_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyCrud + config.Settings.PREFIX_CRUD_READOBJECT + TableName + micro.SeparatorFile() + config.Settings.TESTS_FOLDERNAME
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD_READOBJECT + TableName + "_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrud := string(bytes)

	////загрузим шаблон файла функции
	//FilenameTemplateCrudFunction := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_READOBJECT_FUNCTION_TEST_FILENAME
	//bytes, err = os.ReadFile(FilenameTemplateCrudFunction)
	//if err != nil {
	//	log.Panic("ReadFile() ", FilenameTemplateCrudFunction, " error: ", err)
	//}
	//TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextCrud = create_files.Replace_PackageName(TextCrud, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextCrud = create_files.Delete_TemplateRepositoryImports(TextCrud)

		////
		//ModelTableURL := create_files.Find_ModelTableURL(TableName)
		//TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		//
		ObjectTableURL := create_files.Find_ObjectTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ObjectTableURL)

		//
		CrudObjectTableURL := create_files.Find_CrudObjectTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, CrudObjectTableURL)

		//
		ConstantsURL := create_files.Find_ConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, ConstantsURL)

		//
		CrudStarterURL := create_files.Find_CrudStarterURL()
		TextCrud = create_files.AddImport(TextCrud, CrudStarterURL)

		//
		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextCrud = create_files.AddImport(TextCrud, CrudFuncURL)

	}

	//создание функций
	TextCrud = CreateFiles_ReadObject_TableTest(MapAll, Table1, TextCrud)

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

// CreateFiles_ReadObject_TableTest - создаёт текст функции
func CreateFiles_ReadObject_TableTest(MapAll map[string]*types.Table, Table1 *types.Table, TextGo string) string {
	Otvet := TextGo

	return Otvet
}
