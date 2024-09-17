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

// CreateFiles_ReadAll - создаёт 1 файл в папке crud
func CreateFiles_ReadAll(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()

	FilenameTemplateModel := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_READALL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyModel + TableName
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_MODEL + TableName + "_readall.go"

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
	FilenameTemplateModelFunction := DirTemplatesModel + config.Settings.TEMPLATES_MODEL_READALL_FUNCTION_FILENAME
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
	TextModelFunc := CreateFiles_ReadAll_Table(Table1, TextTemplatedFunction)
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
	err = os.WriteFile(FilenameReady, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// CreateFiles_ReadAll_Table - создаёт текст всех функций
func CreateFiles_ReadAll_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	_, ok := types.MapReadAll[Table1]
	if ok == false {
		return Otvet
	}

	Otvet = CreateFiles_ReadAll_Table1(Table1, TextTemplateFunction)

	return Otvet
}

// CreateFiles_ReadAll_Table1 - создаёт текст всех функций
func CreateFiles_ReadAll_Table1(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	//

	return Otvet
}

// AddInterfaces_ReadAll - добавляет функцию внутрь интерфейса
func AddInterfaces_ReadAll(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := types.MapReadAll[Table1]
	if ok == false {
		return Otvet
	}

	TextFunc := "\n\tReadAll() ([]" + Table1.NameGo + ", error)"

	Otvet = create_files.AddInterfaceFunction(Otvet, TextFunc)

	return Otvet
}
