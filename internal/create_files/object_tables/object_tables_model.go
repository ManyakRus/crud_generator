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

// CreateFiles - создаёт 1 файл в папке db
func CreateFiles(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesObjects := DirTemplates + config.Settings.TEMPLATES_READOBJECT_FOLDERNAME + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATES_READOBJECT_FOLDERNAME + micro.SeparatorFile()

	FilenameTemplateObject := DirTemplatesObjects + config.Settings.TEMPLATES_MODEL_READOBJECT_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_READOBJECT + TableName
	FilenameReadyObjects := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_READOBJECT + TableName + ".go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)
	//ok, err := micro.FileExists(DirReadyTable)
	//if ok == false {
	//	err = os.MkdirAll(DirReadyTable, 0777)
	//	if err != nil {
	//		log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
	//	}
	//}

	//загрузим шаблон файла
	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateObject)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateObject, " error: ", err)
	}
	TextGo := string(bytes)

	//заменим имя пакета на новое
	TextGo = create_files.Replace_PackageName(TextGo, DirReadyTable)

	//заменим имя модели
	TextGo = create_files.Replace_ObjectTemplateModel_to_Model(TextGo, Table1.NameGo)
	TextGo = create_files.Replace_ObjectTemplateTableName_to_TableName(TextGo, Table1.Name)
	TextGo = create_files.Replace_ModelAndTableName(TextGo, Table1)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGo = create_files.Delete_TemplateRepositoryImports(TextGo)

		TextGo = CreateFiles1(MapAll, Table1, TextGo)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGo = create_files.AddImport(TextGo, ModelTableURL)

	}

	////замена импортов на новые URL
	//TextGo = create_files.Replace_RepositoryImportsURL(TextGo)

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

// CreateFiles1 - возвращает текст заполненный
func CreateFiles1(MapAll map[string]*types.Table, Table1 *types.Table, TextGo string) string {
	Otvet := TextGo
	var err error

	//
	TextField := ""
	for _, Column1 := range Table1.MapColumns {
		//проверка имени колонки "DELETED_"
		err = create_files.IsGood_Column(Column1)
		if err != nil {
			log.Warn("CreateFiles1() table: ", Table1.Name, " column: ", Column1.Name, " warning: ", err)
			continue
		}

		//проверка есть внешний ключ
		TableNameFK := Column1.TableKey
		if TableNameFK == "" {
			continue
		}

		//ColumnNameFK := Column1.ColumnKey
		//if ColumnNameFK == "" {
		//	continue
		//}

		//
		TableFK, ok := MapAll[TableNameFK]
		if ok == false {
			continue
		}

		//ColumnFK := TableFK.MapColumns[ColumnNameFK]
		//if ColumnFK == nil {
		//	continue
		//}

		//добавим поле
		ModelNameColumn := create_files.Find_ObjectColumnModelName(Table1, Column1.NameGo)
		JsonName := create_files.SnakeCase_lower(ModelNameColumn)
		TextField1 := "\t" + ModelNameColumn + " " + TableFK.Name + "." + TableFK.NameGo + "\t`" + `json:"` + JsonName + `"	gorm:"-:all"` + "`" + "\n"
		TextField = TextField + TextField1

		//добавим import
		ModelTableURL := create_files.Find_ModelTableURL(TableNameFK)
		Otvet = create_files.AddImport(Otvet, ModelTableURL)

	}

	//новые поля
	Otvet = strings.ReplaceAll(Otvet, "//TextNewFields\n", TextField)

	//описание таблицы
	Otvet = strings.ReplaceAll(Otvet, "//TextDescription", Table1.Comment)

	return Otvet
}
