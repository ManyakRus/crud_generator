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

// CreateFiles_Create - создаёт 1 файл
func CreateFiles_Create(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.TEMPLATES_CRUD_CREATE_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_CRUD + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_create.go"

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

	TextDB = CreateFiles_Create1(TextDB, Table1)

	//создание текста
	TextDB = create_files.Replace_TemplateModel_to_Model(TextDB, Table1.NameGo)
	TextDB = create_files.Replace_TemplateTableName_to_TableName(TextDB, Table1.Name)
	TextDB = create_files.AddText_ModuleGenerated(TextDB)

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

	//импорт postgres_func
	TextDB = create_files.CheckAndAdd_ImportPostgresFunc(TextDB)

	//импорт micro
	TextDB = create_files.CheckAndAdd_ImportMicro(TextDB)

	//удаление пустых строк
	TextDB = create_files.Delete_EmptyLines(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_Create1 - заполняет Text
func CreateFiles_Create1(Text string, Table1 *types.Table) string {
	Otvet := Text

	MassAllColumns := micro.MassFrom_Map(Table1.MapColumns)

	//Primary key
	ReplacePKNotEqual0 := ""
	ReplaceMassValuesIDAppend := ""
	ReplaceTextSQLCreateWithoutID_id := ""
	ReplaceTextSQLCreateWithoutID_values_id := ""
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	ReplacePKFieldsWithComma := ""
	ReplacePKFieldNamesFormat := ""
	ReplaceIDNot0 := ""
	Comma := ""
	NewLine := ""
	TextAnd := ""
	TextOR := ""
	//CommaNewLine := ""
	NumberID := len(MassAllColumns) - len(ColumnsPK)
	for _, Column1 := range ColumnsPK {
		NumberID = NumberID + 1
		sNumberID := strconv.Itoa(NumberID)
		ReplacePKFieldsWithComma = ReplacePKFieldsWithComma + Comma + "m." + Column1.NameGo
		ReplacePKFieldNamesFormat = ReplacePKFieldNamesFormat + Comma + Column1.NameGo + ": %v"
		TextEmpty := create_files.FindText_NotEqualEmpty(Column1, "m."+Column1.NameGo)
		ReplaceIDNot0 = ReplaceIDNot0 + TextAnd + TextEmpty
		ReplaceTextSQLCreateWithoutID_id = ReplaceTextSQLCreateWithoutID_id + ",\n\t" + Column1.Name
		ReplaceTextSQLCreateWithoutID_values_id = ReplaceTextSQLCreateWithoutID_values_id + ", $" + sNumberID
		TextNotEqual0 := create_files.FindText_NotEqualEmpty(Column1, "m."+Column1.NameGo)
		ReplacePKNotEqual0 = ReplacePKNotEqual0 + TextOR + TextNotEqual0

		TextValue := "m." + Column1.NameGo
		if Column1.IsNullable == true {
			TextValue = create_files.FindText_NullValue(Column1.TypeGo, TextValue)
		}
		ReplaceMassValuesIDAppend = ReplaceMassValuesIDAppend + NewLine + "MassValues = append(MassValues, " + TextValue + ")"

		Comma = ", "
		TextAnd = " && "
		NewLine = "\n\t"
		TextOR = " || "
		//CommaNewLine = ",\n"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldsWithComma", ReplacePKFieldsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldNamesFormat", ReplacePKFieldNamesFormat)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceIDNot0", ReplaceIDNot0)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceMassValuesIDAppend", ReplaceMassValuesIDAppend)
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKNotEqual0", ReplacePKNotEqual0)

	//все колонки
	ReplaceMassValuesAppend := ""
	TextSQLCreateWithoutID_values := ""
	TextSQLCreateWithoutID := ""
	ReplaceAllFieldsWithComma := ""
	ReplaceAllColumnNamesWithComma := ""
	ReplaceDollarsWithComma := ""
	Comma = ""
	CommaNewline := ""
	CommaNewline2 := ""
	NewLine = ""
	CommaNewLine := ""
	Number := 0
	NumberNotID := 0
	for _, Column1 := range MassAllColumns {
		//кроме ненужных колонок
		if create_files.Is_Need_Сolumn(Column1) == false {
			continue
		}
		//if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
		//	continue
		//}

		Number = Number + 1
		sNumber := strconv.Itoa(Number)

		ReplaceAllColumnNamesWithComma = ReplaceAllColumnNamesWithComma + CommaNewline + Column1.Name
		ReplaceDollarsWithComma = ReplaceDollarsWithComma + Comma + "$" + sNumber

		TextValue := "m." + Column1.NameGo
		if Column1.IsNullable == true {
			TextValue = create_files.FindText_NullValue(Column1.TypeGo, TextValue)
		}
		if Column1.IsPrimaryKey == true {
			//TextValue := "m." + Column1.NameGo
			//TextValue = create_files.FindText_NilValue(Column1.TypeGo, TextValue)
			//ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + TextValue
			//
		} else if Column1.Name == "created_at" {
			ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + "time.Now()"
		} else {
			ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + TextValue
		}

		if Column1.IsPrimaryKey == false {
			NumberNotID = NumberNotID + 1
			sNumberNotID := strconv.Itoa(NumberNotID)
			TextSQLCreateWithoutID = TextSQLCreateWithoutID + CommaNewLine + Column1.Name
			TextSQLCreateWithoutID_values = TextSQLCreateWithoutID_values + Comma + "$" + sNumberNotID
			ReplaceMassValuesAppend = ReplaceMassValuesAppend + NewLine + "MassValues = append(MassValues, " + TextValue + ")"
		}

		Comma = ", "
		CommaNewline = ",\n\t"
		CommaNewline2 = ",\n\t\t"
		NewLine = "\n\t"
		CommaNewLine = ",\n\t"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceAllFieldsWithComma", ReplaceAllFieldsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceAllColumnNamesWithComma", ReplaceAllColumnNamesWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceDollarsWithComma", ReplaceDollarsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTableName", Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceMassValuesAppend", ReplaceMassValuesAppend)

	ReplaceTextSQLCreateWithoutID := "const TextSQL_Create_WithoutID = " +
		"`" + "\nINSERT INTO " + Table1.Name + "(" + TextSQLCreateWithoutID + ")\n" +
		"VALUES (" + TextSQLCreateWithoutID_values + ")\n" +
		"`"
	ReplaceTextSQLCreate := "const TextSQL_Create = " +
		"`" + "\nINSERT INTO " + Table1.Name + "(" + TextSQLCreateWithoutID + ReplaceTextSQLCreateWithoutID_id + ")\n" +
		"VALUES (" + TextSQLCreateWithoutID_values + ReplaceTextSQLCreateWithoutID_values_id + ")\n" +
		"`"
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTextSQLCreateWithoutID", ReplaceTextSQLCreateWithoutID)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTextSQLCreateWithID", ReplaceTextSQLCreate)

	return Otvet
}
