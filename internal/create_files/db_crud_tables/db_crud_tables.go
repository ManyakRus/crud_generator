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
	"sort"
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

		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы db update
			err = CreateFilesUpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы db update
			err = CreateTestFilesUpdateEveryColumn(Table1)
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

// CreateFilesUpdateEveryColumn - создаёт 1 файл в папке grpc_client
func CreateFilesUpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_UPDATE_FUNC_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCrudUpdate := DirReadyTable + config.Settings.PREFIX_CRUD + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrudUpdateFunc := string(bytes)

	TextCrud := "package " + config.Settings.PREFIX_CRUD + TableName + "\n\n"
	TextCrud = TextCrud + `import (
	"errors"
	"context"
	"fmt"
	"time"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	)

`

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		DBConstantsURL := create_files.FindDBConstantsURL()
		TextCrud = create_files.AddImport(TextCrud, DBConstantsURL)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

		//TextCrud = create_files.ConvertIdToAlias(TextCrud, Table1)
	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextCrudUpdateFunc, Table1)
	// пустой файл не нужен
	if TextUpdateEveryColumn == "" {
		return err
	}

	//ModelName := Table1.NameGo
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextCrud = TextCrud + TextUpdateEveryColumn

	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//удаление пустого импорта
	TextCrud = create_files.DeleteEmptyImport(TextCrud)
	TextCrud = create_files.DeleteEmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudUpdate, []byte(TextCrud), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextCrudUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextUpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumn1(TextCrudUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextCrudUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextCrudUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	TextRequest, TextRequestFieldName := create_files.FindTextProtobufRequest(Column1.TypeGo)

	//заменяем Read_ctx()
	Otvet = strings.ReplaceAll(Otvet, " Read_ctx ", " "+FuncName+"_ctx ")
	Otvet = strings.ReplaceAll(Otvet, " Read_ctx(", " "+FuncName+"_ctx(")
	Otvet = strings.ReplaceAll(Otvet, ".Read_ctx(", "."+FuncName+"_ctx(")

	//заменяем Read()
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Read", ModelName+"_"+FuncName)
	Otvet = strings.ReplaceAll(Otvet, " Read ", " "+FuncName+" ")
	Otvet = strings.ReplaceAll(Otvet, " Read(", " "+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, `"Read()`, `"`+FuncName+"()")
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Model.ID", "Model."+ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	//Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	//Otvet = strings.ReplaceAll(Otvet, "m.ID", "m."+ColumnName)

	return Otvet
}

// CreateTestFilesUpdateEveryColumn - создаёт 1 файл в папке grpc_client
func CreateTestFilesUpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesCrud := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyCrud := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + micro.SeparatorFile()

	FilenameTemplateCrud := DirTemplatesCrud + config.Settings.TEMPLATES_CRUD_TABLE_UPDATE_FUNC_TEST_FILENAME
	DirReadyTable := DirReadyCrud
	FilenameReadyCrudUpdate := DirReadyTable + config.Settings.PREFIX_CRUD + TableName + "_update_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateCrud)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateCrud, " error: ", err)
	}
	TextCrudUpdateFunc := string(bytes)

	TextCrud := "package " + config.Settings.PREFIX_CRUD + TableName + "\n\n"
	TextCrud = TextCrud + `import (
	"testing"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	)

`

	//заменим импорты
	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	ModelTableURL := create_files.FindModelTableURL(TableName)
	TextCrud = create_files.AddImport(TextCrud, ModelTableURL)

	TextCrud = create_files.ConvertIdToAlias(TextCrud, Table1)
	//}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumnTest(TextCrudUpdateFunc, Table1)
	// пустой файл не нужен
	if TextUpdateEveryColumn == "" {
		return err
	}
	//ModelName := Table1.NameGo
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextCrud = strings.ReplaceAll(TextCrud, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextCrud = TextCrud + TextUpdateEveryColumn

	TextCrud = config.Settings.TEXT_MODULE_GENERATED + TextCrud

	//удаление пустого импорта
	TextCrud = create_files.DeleteEmptyImport(TextCrud)
	TextCrud = create_files.DeleteEmptyLines(TextCrud)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudUpdate, []byte(TextCrud), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumnTest - возвращает текст для всех таблиц
func FindTextUpdateEveryColumnTest(TextCrudUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextUpdateEveryColumnTest() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumnTest1(TextCrudUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumnTest1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumnTest1(TextCrudUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextCrudUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	TextRequest, TextRequestFieldName := create_files.FindTextProtobufRequest(Column1.TypeGo)
	DefaultValue := create_files.FindTextDefaultValue(Column1.TypeGo)

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	Otvet = strings.ReplaceAll(Otvet, "Otvet.Name", "Otvet."+ColumnName)
	//Otvet = strings.ReplaceAll(Otvet, "Postgres_ID_Test", DefaultValue)
	Otvet = strings.ReplaceAll(Otvet, "TestUpdate(", "Test"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, ".Update(", "."+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, " DefaultValue", " "+DefaultValue)

	return Otvet
}
