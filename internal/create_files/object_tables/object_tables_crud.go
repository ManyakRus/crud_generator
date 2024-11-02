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
	DirTemplatesObjects := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_OBJECTS + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_OBJECTS + micro.SeparatorFile()

	FilenameTemplateObject := DirTemplatesObjects + config.Settings.TEMPLATES_OBJECTS_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_OBJECT + TableName
	FilenameReadyObjects := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_OBJECT + TableName + ".go"

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

		TextGo = CreateFiles1(MapAll, Table1, TextGo)

		//ModelTableURL := create_files.Find_ModelTableURL(TableName)
		//TextGo = create_files.AddImport(TextGo, ModelTableURL)

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
		err1 := create_files.IsGood_Column(Column1)
		if err1 != nil {
			log.Warn("CreateFiles1() table: ", Table1.Name, " column: ", Column1.Name, " warning: ", err)
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

		//добавим import
		ModelTableURL := create_files.Find_ModelTableURL(TableNameFK)
		TextGo = create_files.AddImport(TextGo, ModelTableURL)

		//добавим поле
		TextField1 := "\t" + Column1.NameGo + " " + TableFK.Name + "." + TableFK.NameGo + "\n"
		TextField = TextField + TextField1

	}

	Otvet = strings.ReplaceAll(Otvet, "//TextNewFields\n", TextField)

	return Otvet
}
