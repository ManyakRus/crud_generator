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
	TextDB = create_files.Replace_PackageName(TextDB, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.Delete_TemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

		ConstantsURL := create_files.Find_DBConstantsURL()
		TextDB = create_files.AddImport(TextDB, ConstantsURL)

		CrudFunctionsURL := create_files.Find_CrudFunctionsURL()
		TextDB = create_files.AddImport(TextDB, CrudFunctionsURL)

		//удалим лишние функции
		TextDB = create_files.DeleteFunc_Delete(TextDB, Table1)
		TextDB = create_files.DeleteFunc_Restore(TextDB, Table1)
		TextDB = create_files.DeleteFunc_Find_byExtID(TextDB, Table1)

		//удалим лишние функции ctx
		TextDB = DeleteFunc_DeleteCtx(TextDB, Table1)
		TextDB = DeleteFunc_RestoreCtx(TextDB, Table1)
		TextDB = DeleteFunc_Find_byExtIDCtx(TextDB, Table1)

		//кэш
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//исправление Save()
			TextDB = create_files.CommentLineInText(TextDB, constants.TEXT_CACHE_REMOVE)
			//TextDB = strings.ReplaceAll(TextDB, `//`+constants.TEXT_CACHE_REMOVE, constants.TEXT_CACHE_REMOVE)
		}

		//
		TextDB = Replace_ExtID_equal0_string(TextDB, Table1)
	}

	//TextDB = create_files.DeleteFunc_Find_byExtID(TextDB, Table1)
	//TextDB = create_files.DeleteFunc_Find_byExtIDCtx(TextDB, Table1)
	TextDB = AddTextOmit(TextDB, Table1)
	TextDB = ReplaceText_modified_at(TextDB, Table1)
	TextDB = ReplaceText_created_at(TextDB, Table1)
	TextDB = ReplaceText_is_deleted_deleted_at(TextDB, Table1)
	TextDB = create_files.DeleteImportModel(TextDB)

	TextDB = ReplaceCacheRemove(TextDB, Table1)

	TextDB = ReplacePrimaryKeyM_ID(TextDB, Table1)

	//id := m.ID
	TextDB = ReplaceColumnNamePK(TextDB, Table1)

	//"Replace_ColumnNameM(m.ID)"
	//TextDB = create_files.Replace_ColumnNameM(TextDB, Table1)

	//создание текста
	TextDB = create_files.Replace_TemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_TemplateTableName_to_TableName(TextDB, Table1.Name)
	TextDB = create_files.AddText_ModuleGenerated(TextDB)

	//TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextDB = config.Settings.TEXT_MODULE_GENERATED + TextDB

	//замена импортов на новые URL
	TextDB = create_files.Replace_RepositoryImportsURL(TextDB)

	//uuid
	TextDB = create_files.CheckAndAdd_ImportUUID_FromText(TextDB)

	//alias
	TextDB = create_files.CheckAndAdd_ImportAlias(TextDB)

	//удаление пустого импорта
	TextDB = create_files.Delete_EmptyImport(TextDB)

	//переименование функций
	TextDB = RenameFunctions(TextDB, Table1)

	//импорт "fmt"
	TextDB = create_files.CheckAndAdd_ImportFmt(TextDB)

	//удаление пустых строк
	TextDB = create_files.Delete_EmptyLines(TextDB)

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
	TextDB = create_files.Replace_PackageName(TextDB, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextDB = create_files.Delete_TemplateRepositoryImports(TextDB)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextDB = create_files.AddImport(TextDB, ModelTableURL)

		ConstantsURL := create_files.Find_ConstantsURL()
		TextDB = create_files.AddImport(TextDB, ConstantsURL)

		//удалим лишние функции
		TextDB = DeleteFunc_TestDelete(TextDB, Table1)
		TextDB = DeleteFunc_TestRestore(TextDB, Table1)
		TextDB = DeleteFunc_TestFind_byExtID(TextDB, Table1)

		//Postgres_ID_Test = ID Minimum
		TextDB = create_files.Replace_Postgres_ID_Test(TextDB, Table1)

		//замена ID на PrimaryKey
		//TextDB = Replace_PrimaryKeyOtvetID_Test(TextDB, Table1)

		//добавим импорт uuid
		TextDB = create_files.CheckAndAdd_ImportUUID_FromText(TextDB)

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextDB = create_files.Replace_Connect_WithApplicationName(TextDB)

	}

	//создание текста
	TextDB = create_files.Replace_TemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_TemplateTableName_to_TableName(TextDB, Table1.Name)
	TextDB = create_files.AddText_ModuleGenerated(TextDB)

	if config.Settings.HAS_IS_DELETED == true {
		TextDB = DeleteFunc_TestDelete(TextDB, Table1)
		TextDB = DeleteFunc_TestRestore(TextDB, Table1)
	}
	TextDB = DeleteFunc_TestFind_byExtID(TextDB, Table1)

	//SkipNow() если нет строк в БД
	TextDB = create_files.AddSkipNow(TextDB, Table1)

	//замена импортов на новые URL
	TextDB = create_files.Replace_RepositoryImportsURL(TextDB)

	//удаление пустого импорта
	TextDB = create_files.Delete_EmptyImport(TextDB)

	//удаление пустых строк
	TextDB = create_files.Delete_EmptyLines(TextDB)

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

// ReplacePrimaryKeyM_ID - заменяет "Otvet.ID" на название колонки PrimaryKey
func ReplacePrimaryKeyM_ID(Text string, Table1 *types.Table) string {
	Otvet := Text

	VariableName := "m"

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariable(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"

		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = ProtoFromInt(m.ID)", TextRequestIDmID)
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = int64(ID)", TextRequestIDInt64ID)
	//Otvet = strings.ReplaceAll(Otvet, "\tOtvet.ID = "+VariableName+".ID\n", TextOtvetIDmID)
	//Otvet = strings.ReplaceAll(Otvet, " IntFromAlias("+VariableName+".ID) == 0", TextMID0)
	Otvet = strings.ReplaceAll(Otvet, "\tm2.ID = int64(m.ID)", TextM2ID)
	Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", TextIfMId)
	Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", TextIfMIdNot0)

	//заменим ID := Request.ID
	//Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextIDRequestID)

	return Otvet
}

// DeleteFunc_TestRestore - удаляет функцию Restore()
func DeleteFunc_TestRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// TestRestore ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestRestore")

	return Otvet
}

// DeleteFunc_TestDelete - удаляет функцию Delete()
func DeleteFunc_TestDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// TestDelete ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestDelete")

	return Otvet
}

// DeleteFunc_Find_byExtID - удаляет функцию Find_ByExtID()
func DeleteFunc_TestFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверка есть ли колонки ExtID и ConnectionID
	if create_files.Has_Column_ExtID_ConnectionID_Int64(Table1) == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// TestFind_ByExtID ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestFind_ByExtID")

	return Otvet
}

// DeleteFunc_Find_byExtIDCtx - удаляет функцию Find_ByExtID_ctx()
func DeleteFunc_Find_byExtIDCtx(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверка есть ли колонки ExtID и ConnectionID
	if create_files.Has_Column_ExtID_ConnectionID_Int64(Table1) == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// Find_ByExtID_ctx ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Find_ByExtID_ctx")

	return Otvet
}

// DeleteFunc_RestoreCtx - удаляет функцию Restore_ctx()
func DeleteFunc_RestoreCtx(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// Restore_ctx ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Restore_ctx")

	return Otvet
}

// DeleteFunc_DeleteCtx - удаляет функцию Delete_ctx()
func DeleteFunc_DeleteCtx(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// Delete_ctx ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Delete_ctx")

	return Otvet
}