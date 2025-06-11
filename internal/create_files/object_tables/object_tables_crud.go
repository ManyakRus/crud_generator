package object_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_crud - создаёт 1 файл в папке db
func CreateFiles_crud(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesObjects := DirTemplates + config.Settings.TEMPLATES_READOBJECT_FOLDERNAME + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATES_READOBJECT_FOLDERNAME + micro.SeparatorFile()

	FilenameTemplateObject := DirTemplatesObjects + config.Settings.TEMPLATES_MODEL_CRUD_READOBJECT_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_READOBJECT + TableName
	FilenameReadyObjects := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_READOBJECT + TableName + config.Settings.SUFFIX_CRUD + ".go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateObject)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateObject, " error: ", err)
	}
	TextGo := string(bytes)

	//заменим имя пакета на новое
	TextGo = create_files.Replace_PackageName(TextGo, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGo = create_files.Delete_TemplateRepositoryImports(TextGo)

		//TextGo = CreateFiles1(MapAll, Table1, TextGo)

		//ModelTableURL := create_files.Find_ModelTableURL(TableName)
		//TextGo = create_files.AddImport(TextGo, ModelTableURL)

		//Find_DBConstants
		ConstantsURL := create_files.Find_DBConstantsURL()
		TextGo = create_files.AddImport(TextGo, ConstantsURL)

		//calc_struct_version
		CalcStructVersionURL := create_files.Find_CalcStructVersionURL()
		TextGo = create_files.AddImport(TextGo, CalcStructVersionURL)

	}

	////замена импортов на новые URL
	//TextGo = create_files.Replace_RepositoryImportsURL(TextGo)

	//заменим имя модели
	TextGo = create_files.Replace_ObjectTemplateModel_to_Model(TextGo, Table1.NameGo)
	TextGo = create_files.Replace_ObjectTemplateTableName_to_TableName(TextGo, Table1.Name)
	TextGo = create_files.Replace_ModelAndTableName(TextGo, Table1)

	//uuid
	TextGo = create_files.CheckAndAdd_ImportUUID_FromText(TextGo)

	//alias
	TextGo = create_files.CheckAndAdd_ImportAlias(TextGo)

	//удаление пустого импорта
	TextGo = create_files.Delete_EmptyImport(TextGo)

	//импорт "fmt"
	TextGo = create_files.CheckAndAdd_ImportFmt(TextGo)

	//удаление пустых строк
	TextGo = create_files.Delete_EmptyLines(TextGo)

	//запись файла
	err = os.WriteFile(FilenameReadyObjects, []byte(TextGo), config.Settings.FILE_PERMISSIONS)

	return err
}
