package db

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
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
		if config.Settings.NEED_CREATE_DB_TEST == true {
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
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + "db.go_"
	TableName := strings.ToLower(Table1.Name)
	DirTable := DirReadyDB + "db_" + TableName
	FilenameReadyDB := DirTable + micro.SeparatorFile() + "db_" + TableName + ".go"

	//создадим каталог
	ok, err := micro.FileExists(DirTable)
	if ok == false {
		err = os.MkdirAll(DirTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	if config.Settings.HAS_IS_DELETED == true {
		TextDB = DeleteFuncDelete(TextDB, Table1)
		TextDB = DeleteFuncDeleteCtx(TextDB, Table1)
		TextDB = DeleteFuncRestore(TextDB, Table1)
		TextDB = DeleteFuncRestoreCtx(TextDB, Table1)
	}
	TextDB = DeleteFuncFind_byExtID(TextDB, Table1)
	TextDB = DeleteFuncFind_byExtIDCtx(TextDB, Table1)
	TextDB = AddTextOmit(TextDB, Table1)
	TextDB = ReplaceText_modified_at(TextDB, Table1)
	TextDB = ReplaceText_is_deleted_deleted_at(TextDB, Table1)
	TextDB = create_files.DeleteImportModel(TextDB)

	//замена импортов на новые URL
	TextDB = create_files.ReplaceServiceURLImports(TextDB)

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
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + "db_test.go_"
	DirTable := DirReadyDB + "db_" + TableName
	FilenameReadyDB := DirTable + micro.SeparatorFile() + "db_" + TableName + "_test.go"

	//создадим каталог
	ok, err := micro.FileExists(DirTable)
	if ok == false {
		err = os.Mkdir(DirTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
	}
	TextDB := string(bytes)

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	if config.Settings.HAS_IS_DELETED == true {
		TextDB = DeleteFuncTestDelete(TextDB, Table1)
		TextDB = DeleteFuncTestRestore(TextDB, Table1)
	}
	TextDB = DeleteFuncTestFind_byExtID(TextDB, Table1)

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

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete ")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore ")

	return Otvet
}

// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
func DeleteFuncDeleteCtx(Text string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete_ctx ")

	return Otvet
}

// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
func DeleteFuncRestoreCtx(Text string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore_ctx ")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Find_ByExtID ")

	return Otvet
}

// DeleteFuncFind_byExtIDCtx - удаляет функцию Find_ByExtID_ctx()
func DeleteFuncFind_byExtIDCtx(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Find_ByExtID_ctx ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestDelete")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestRestore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	FuncName := "TestFind_ByExtID"

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		Otvet = create_files.DeleteCommentFromFuncName(Otvet, FuncName)
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, FuncName)
	//Otvet = create_files.DeleteLineWithComment(Otvet, FuncName)

	return Otvet
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
	for _, Column1 := range Table1.MapColumns {
		ColumnNameGo := Column1.NameGo
		TypeGo := Column1.TypeGo

		_, is_nullable_config := types.MapNullableFileds[ColumnNameGo]

		if Column1.IsNullable == false {
			continue
		}

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

	TextFind := "//Text_modified_at"
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

	TextFind := "//Text_is_deleted_deleted_at"
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}
