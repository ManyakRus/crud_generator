package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/mini_func"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFiles - создаёт 1 файл в папке db
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.TEMPLATES_CRUD_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_CRUD + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + ".go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//заменим имя пакета на новое
	TextDB = create_files.ReplacePackageName(TextDB, DirReadyTable)

	ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.DeleteTemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

		ConstantsURL := create_files.FindDBConstantsURL()
		TextDB = create_files.AddImport(TextDB, ConstantsURL)

		CrudFunctionsURL := create_files.FindCrudFunctionsURL()
		TextDB = create_files.AddImport(TextDB, CrudFunctionsURL)

		//удалим лишние функции
		TextDB = create_files.DeleteFuncDelete(TextDB, Table1)
		TextDB = create_files.DeleteFuncRestore(TextDB, Table1)
		TextDB = create_files.DeleteFuncFind_byExtID(TextDB, Table1)

		//удалим лишние функции ctx
		TextDB = create_files.DeleteFuncDeleteCtx(TextDB, Table1)
		TextDB = create_files.DeleteFuncRestoreCtx(TextDB, Table1)
		TextDB = create_files.DeleteFuncFind_byExtIDCtx(TextDB, Table1)

		//кэш
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//исправление Save()
			TextDB = strings.ReplaceAll(TextDB, `//`+constants.TEXT_CACHE_REMOVE, constants.TEXT_CACHE_REMOVE)
		}

		//
		TextDB = Replace_ExtID_equal0_string(TextDB, Table1)
	}

	//TextDB = create_files.DeleteFuncFind_byExtID(TextDB, Table1)
	//TextDB = create_files.DeleteFuncFind_byExtIDCtx(TextDB, Table1)
	TextDB = AddTextOmit(TextDB, Table1)
	TextDB = ReplaceText_modified_at(TextDB, Table1)
	TextDB = ReplaceText_created_at(TextDB, Table1)
	TextDB = ReplaceText_is_deleted_deleted_at(TextDB, Table1)
	TextDB = create_files.DeleteImportModel(TextDB)

	TextDB = create_files.ReplaceCacheRemove(TextDB, Table1)

	TextDB = create_files.ReplacePrimaryKeyM_ID(TextDB, Table1)

	//id := m.ID
	TextDB = create_files.ReplaceColumnNamePK(TextDB, Table1)

	//"ReplaceColumnNameM(m.ID)"
	//TextDB = create_files.ReplaceColumnNameM(TextDB, Table1)

	//создание текста
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

	//uuid
	TextDB = create_files.CheckAndAddImportUUID_FromText(TextDB)

	//alias
	TextDB = create_files.CheckAndAddImportAlias(TextDB)

	//удаление пустого импорта
	TextDB = create_files.DeleteEmptyImport(TextDB)

	//переименование функций
	TextDB = RenameFunctions(TextDB, Table1)

	//импорт "fmt"
	TextDB = create_files.CheckAndAddImportFmt(TextDB)

	//удаление пустых строк
	TextDB = create_files.DeleteEmptyLines(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesTest - создаёт 1 файл в папке db
func CreateFilesTest(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.TEMPLATES_CRUD_TEST_FILENAME
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_CRUD + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//заменим имя пакета на новое
	TextDB = create_files.ReplacePackageName(TextDB, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.DeleteTemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

		ConstantsURL := create_files.FindConstantsURL()
		TextDB = create_files.AddImport(TextDB, ConstantsURL)

		//удалим лишние функции
		TextDB = create_files.DeleteFuncTestDelete(TextDB, Table1)
		TextDB = create_files.DeleteFuncTestRestore(TextDB, Table1)
		TextDB = create_files.DeleteFuncTestFind_byExtID(TextDB, Table1)

		//Postgres_ID_Test = ID Minimum
		TextDB = create_files.Replace_Postgres_ID_Test(TextDB, Table1)

		//замена ID на PrimaryKey
		TextDB = create_files.ReplacePrimaryKeyOtvetID(TextDB, Table1)

		//добавим импорт uuid
		TextDB = create_files.CheckAndAddImportUUID_FromText(TextDB)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextDB = create_files.ReplaceConnect_WithApplicationName(TextDB)

	}

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	if config.Settings.HAS_IS_DELETED == true {
		TextDB = create_files.DeleteFuncTestDelete(TextDB, Table1)
		TextDB = create_files.DeleteFuncTestRestore(TextDB, Table1)
	}
	TextDB = create_files.DeleteFuncTestFind_byExtID(TextDB, Table1)

	//SkipNow() если нет строк в БД
	TextDB = create_files.AddSkipNow(TextDB, Table1)

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

	//удаление пустого импорта
	TextDB = create_files.DeleteEmptyImport(TextDB)

	//удаление пустых строк
	TextDB = create_files.DeleteEmptyLines(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// AddTextOmit - добавляет код для записи null в колонки Nullable
func AddTextOmit(TextDB string, Table1 *types.Table) string {
	Otvet := TextDB

	TextFind := "\t//игнор пустых колонок"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}

	TextOmit := ""
	NullableCount := 0

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("Table1.MapColumns[key1] = false")
		}
		ColumnName := Column1.Name
		ColumnNameGo := Column1.NameGo
		TypeGo := Column1.TypeGo

		if Column1.IsNullable == false {
			continue
		}

		//ищем в файле настроек nullable.json
		is_nullable_config, has_is_nullable_config := types.MapNullableFileds[ColumnName]
		if has_is_nullable_config == true && is_nullable_config == false {
			continue
		}

		//
		if TypeGo == "time.Time" {
			NullableCount = NullableCount + 1
			TextFind := `if m.` + ColumnNameGo + `.IsZero() == true {`
			pos1 := strings.Index(TextDB, TextFind)
			if pos1 >= 0 {
				continue
			}

			TextOmit = TextOmit + "\t" + `ColumnName = "` + ColumnNameGo + `"
	if m.` + ColumnNameGo + `.IsZero() == true {
		MassOmit = append(MassOmit, ColumnName)
	}

`
		} else if mini_func.IsNumberType(TypeGo) == true && (Column1.TableKey != "" || is_nullable_config == true) {
			NullableCount = NullableCount + 1
			TextFind := `if m.` + ColumnNameGo + ` == 0 {`
			pos1 := strings.Index(TextDB, TextFind)
			if pos1 >= 0 {
				continue
			}

			TextOmit = TextOmit + "\t" + `ColumnName = "` + ColumnNameGo + `"
	if m.` + ColumnNameGo + ` == 0 {
		MassOmit = append(MassOmit, ColumnName)
	}

`
		}

	}

	Otvet = Otvet[:pos1] + TextOmit + Otvet[pos1:]

	if NullableCount == 0 && config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = strings.ReplaceAll(Otvet, "\n\tvar ColumnName string", "")
	}

	return Otvet
}

// ReplaceText_modified_at - заменяет текст "Text_modified_at" на текст из файла
func ReplaceText_modified_at(s string, Table1 *types.Table) string {
	Otvet := s

	TextNew := config.Settings.TEXT_DB_MODIFIED_AT
	_, ok := Table1.MapColumns["modified_at"]
	if ok == false {
		TextNew = ""
	}

	TextFind := "\t//Text_modified_at\n"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}

// ReplaceText_is_deleted_deleted_at - заменяет текст "Text_is_deleted_deleted_at" на текст из файла
func ReplaceText_is_deleted_deleted_at(s string, Table1 *types.Table) string {
	Otvet := s

	TextNew := config.Settings.TEXT_DB_IS_DELETED
	_, ok := Table1.MapColumns["is_deleted"]
	if ok == false {
		TextNew = ""
	}

	_, ok = Table1.MapColumns["deleted_at"]
	if ok == false {
		TextNew = ""
	}

	TextFind := "\t//Text_is_deleted_deleted_at\n"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}

// ReplaceText_created_at - заменяет текст "Text_created_at" на текст из файла
func ReplaceText_created_at(s string, Table1 *types.Table) string {
	Otvet := s

	TextNew := config.Settings.TEXT_DB_CREATED_AT
	_, ok := Table1.MapColumns["created_at"]
	if ok == false {
		TextNew = ""
	}

	TextFind := "\t//Text_created_at\n"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}

// Replace_ExtID_equal0_string - заменяет "ExtID == 0 " на "ExtID == "" "
func Replace_ExtID_equal0_string(TextDB string, Table1 *types.Table) string {
	Otvet := TextDB

	Column1, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}
	TypeGo := Column1.TypeGo
	if TypeGo == "string" {
		Otvet = strings.ReplaceAll(Otvet, "ExtID == 0 ", `ExtID == "" `)
	}

	return Otvet
}
