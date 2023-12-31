package db_crud_tables

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
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке db
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы db
		if config.Settings.NEED_CREATE_DB == true {
			err = CreateFiles(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}
		//тестовые файлы db
		if config.Settings.NEED_CREATE_DB_TEST == true {
			err = CreateTestFiles(Table1)
			if err != nil {
				log.Error("CreateTestFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}
	}

	return err
}

// CreateFiles - создаёт 1 файл в папке db
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + constants.TemplateFilenameCrudGo
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
	}

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//TextDB = create_files.DeleteFuncFind_byExtID(TextDB, Table1)
	//TextDB = create_files.DeleteFuncFind_byExtIDCtx(TextDB, Table1)
	TextDB = AddTextOmit(TextDB, Table1)
	TextDB = ReplaceText_modified_at(TextDB, Table1)
	TextDB = ReplaceText_created_at(TextDB, Table1)
	TextDB = ReplaceText_is_deleted_deleted_at(TextDB, Table1)
	TextDB = create_files.DeleteImportModel(TextDB)

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

	//удаление пустого импорта
	TextDB = create_files.DeleteEmptyImport(TextDB)

	//переименование функций
	TextDB = RenameFunctions(TextDB, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// CreateTestFiles - создаёт 1 файл в папке db
func CreateTestFiles(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + constants.TemplateFilenameCrudGoTest
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

		//удалим лишние функции
		TextDB = create_files.DeleteFuncTestDelete(TextDB, Table1)
		TextDB = create_files.DeleteFuncTestRestore(TextDB, Table1)
		TextDB = create_files.DeleteFuncTestFind_byExtID(TextDB, Table1)
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

	//Postgres_ID_Test = ID Minimum
	if Table1.IDMinimum != "" {
		TextFind := "const Postgres_ID_Test = "
		TextDB = strings.ReplaceAll(TextDB, TextFind+"1", TextFind+Table1.IDMinimum)
	}

	//SkipNow()
	TextDB = create_files.AddSkipNow(TextDB, Table1)

	// замена ID на PrimaryKey
	TextDB = create_files.ReplacePrimaryKeyID(TextDB, Table1)

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

	//удаление пустого импорта
	TextDB = create_files.DeleteEmptyImport(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

//// DeleteFuncDelete - удаляет функцию Delete()
//func DeleteFuncDelete(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete ")
//
//	return Otvet
//}
//
//// DeleteFuncRestore - удаляет функцию Restore()
//func DeleteFuncRestore(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore ")
//
//	return Otvet
//}
//
//// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
//func DeleteFuncDeleteCtx(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete_ctx ")
//
//	return Otvet
//}
//
//// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
//func DeleteFuncRestoreCtx(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore_ctx ")
//
//	return Otvet
//}
//
//// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
//func DeleteFuncFind_byExtID(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	//если есть обе колонки - ничего не делаем
//	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
//	if ok == true {
//		return Otvet
//	}
//
//	//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Find_ByExtID ")
//
//	return Otvet
//}
//
//// DeleteFuncFind_byExtIDCtx - удаляет функцию Find_ByExtID_ctx()
//func DeleteFuncFind_byExtIDCtx(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	//если есть обе колонки - ничего не делаем
//	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
//	if ok == true {
//		return Otvet
//	}
//
//	//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Find_ByExtID_ctx ")
//
//	return Otvet
//}

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
	for _, Column1 := range Table1.MapColumns {
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

func RenameFunctions(TextDB string, Table1 *types.Table) string {
	Otvet := TextDB

	TableName := strings.ToLower(Table1.Name)
	Rename1, ok := types.MapRenameFunctions[TableName]
	if ok == false {
		return Otvet
	}

	for _, v := range Rename1 {
		Otvet = strings.ReplaceAll(Otvet, " "+v.Old+"(", " "+v.New+"(")
	}

	return Otvet
}
