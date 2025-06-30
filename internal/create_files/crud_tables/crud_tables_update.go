package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strconv"
	"strings"
)

// CreateFiles_Update - создаёт 1 файл
func CreateFiles_Update(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.TEMPLATES_CRUD_UPDATE_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_CRUD + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_update.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

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

		//CrudFunctionsURL := create_files.Find_CrudFunctionsURL()
		//TextDB = create_files.AddImport(TextDB, CrudFunctionsURL)

	}

	TextDB = CreateFiles_Update1(TextDB, Table1)

	//создание текста
	TextDB = create_files.Replace_TemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_TemplateTableName_to_TableName(TextDB, Table1.Name)
	TextDB = create_files.AddText_ModuleGenerated(TextDB)

	//замена импортов на новые URL
	TextDB = create_files.Replace_RepositoryImportsURL(TextDB)

	//postgres_func
	TextDB = create_files.CheckAndAdd_Import(TextDB, "github.com/ManyakRus/starter/postgres_func")

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
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_Update1 - заполняет Text
func CreateFiles_Update1(Text string, Table1 *types.Table) string {
	Otvet := Text

	//Primary key
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	//TableAlias := create_files.Find_TableAlias(Table1)
	ReplacePKFieldsWithComma := ""
	ReplacePKFieldNamesFormat := ""
	ReplaceID0 := ""
	//ReplaceWhereID := ""
	TextWhereID := ""
	//ReplaceTextMassFields := "MassFields := make([]any, 0)\n"
	ReplaceMassFieldsAppend := ""
	Comma := ""
	//TextNewLine := ""
	TextAnd := ""
	for i, Column1 := range ColumnsPK {
		sNumber := strconv.Itoa(i + 1)
		ReplacePKFieldsWithComma = ReplacePKFieldsWithComma + Comma + "m." + Column1.NameGo
		ReplacePKFieldNamesFormat = ReplacePKFieldNamesFormat + Comma + Column1.NameGo + ": %v"
		TextEmpty := create_files.FindText_EqualEmpty(Column1, "m."+Column1.NameGo)
		ReplaceID0 = ReplaceID0 + TextAnd + TextEmpty
		TextWhereID = TextWhereID + "\tand " + `"` + Column1.Name + `"` + " = $" + sNumber
		//ReplaceTextMassFields = ReplaceTextMassFields + "\tMassFields = append(MassFields, m." + Column1.NameGo + ")\n"
		ReplaceMassFieldsAppend = ReplaceMassFieldsAppend + "\tMassFields = append(MassFields, m." + Column1.NameGo + ")\n"

		Comma = ", "
		//TextNewLine = ",\n"
		TextAnd = " && "
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldsWithComma", ReplacePKFieldsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldNamesFormat", ReplacePKFieldNamesFormat)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceID0", ReplaceID0)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceMassFieldsAppend", ReplaceMassFieldsAppend)
	//Otvet = strings.ReplaceAll(Otvet, "ReplaceWhereID", ReplaceWhereID)

	//все колонки
	ReplaceAllFieldsWithComma := ""
	//ReplaceAllColumnNamesWithComma := ""
	//ReplaceDollarsWithComma := ""
	ReplaceColumnNameEqualDollarComma := ""
	ReplaceWhereID := ""
	ReplaceMassFieldsWithComma := ""
	//Comma = ""
	CommaNewline := "\t"
	CommaNewline2 := ""
	MassColumns := micro.MassFrom_Map(Table1.MapColumns)
	Number := 0
	for _, Column1 := range MassColumns {
		if create_files.Is_NotNeedUpdate_Сolumn_SQL(Column1) == true {
			continue
		}
		//if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
		//	continue
		//}

		Number = Number + 1
		sNumber := strconv.Itoa(Number)
		ReplaceColumnNameEqualDollarComma = ReplaceColumnNameEqualDollarComma + CommaNewline + `"` + Column1.Name + `"` + " = $" + sNumber
		if Column1.Name == "modified_at" {
			ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + "time.Now()"
		} else if Column1.IsNullable == true {
			TextVariable := "m." + Column1.NameGo
			TextValue := create_files.FindText_NullValue(Column1.TypeGo, TextVariable)
			ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + TextValue
		} else {
			ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + "m." + Column1.NameGo
		}

		if Column1.IsPrimaryKey == true {
			ReplaceWhereID = ReplaceWhereID + "\tand " + `"` + Column1.Name + `"` + " = $" + sNumber + "\n"
			ReplaceMassFieldsWithComma = ReplaceMassFieldsWithComma + Comma + "m." + Column1.NameGo
		}

		Comma = ", "
		CommaNewline = ",\n\t"
		CommaNewline2 = ",\n\t\t"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceAllFieldsWithComma", ReplaceAllFieldsWithComma)
	//Otvet = strings.ReplaceAll(Otvet, "ReplaceAllColumnNamesWithComma", ReplaceAllColumnNamesWithComma)
	//Otvet = strings.ReplaceAll(Otvet, "ReplaceDollarsWithComma", ReplaceDollarsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTableName", Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceColumnNameEqualDollarComma", ReplaceColumnNameEqualDollarComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceWhereID", ReplaceWhereID)

	//
	PK_count := len(ColumnsPK)
	sPK_count := strconv.Itoa(PK_count)
	//ReplaceTextSQLUpdateMass := `
	//MassFields := make([]any, 0)
	//Comma := ""
	//TextSQL := ` + "`" + `UPDATE "` + Table1.Name + `" SET` + "`" + `
	//for i, ColumnName1 := range MassNeedUpdateFields {
	//	ColumnNameDB, err := micro.Find_Tag_JSON(m, ColumnName1)
	//	if err != nil {
	//		return err
	//	}
	//	TextSQL = TextSQL + Comma + ColumnNameDB + " = $" + ` + `strconv.Itoa(` + strconv.Itoa(PK_count) + `+i+1)
	//	Value, err := micro.GetStructValue(m, ColumnName1)
	//	if err != nil {
	//		return err
	//	}
	//	MassFields = append(MassFields, Value)
	//	Comma = ",\n"
	//}
	//TextSQL = TextSQL + "\nWHERE 1=1 ` + TextWhereID + `"`
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTextSQLWhere", TextWhereID)
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKCount", sPK_count)

	Otvet = ReplaceCacheRemove(Otvet, Table1)

	return Otvet
}
