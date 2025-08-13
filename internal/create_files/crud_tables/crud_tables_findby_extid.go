package crud_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_FindBy_ExtID - создаёт 1 файл
func CreateFiles_FindBy_ExtID(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.TEMPLATES_CRUD_FINDBY_EXTID_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_CRUD + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_findby_extid.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateDB)
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

	TextDB = CreateFiles_FindBy_ExtID1(TextDB, Table1)

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

	//импорт "postgres_func"
	TextDB = create_files.CheckAndAdd_ImportPostgresFunc(TextDB)

	//удаление пустых строк
	TextDB = create_files.Delete_EmptyLines(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_FindBy_ExtID1 - заполняет Text
func CreateFiles_FindBy_ExtID1(Text string, Table1 *types.Table) string {
	Otvet := Text

	//TableAlias := create_files.Find_TableAlias(Table1)

	////Primary key
	//ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	//ReplacePKFieldsWithComma := ""
	//ReplacePKFieldNamesFormat := ""
	//ReplaceWhereID := ""
	//Comma := ""
	//CommaNewline := ""
	//for i, Column1 := range ColumnsPK {
	//	ReplacePKFieldsWithComma = ReplacePKFieldsWithComma + Comma + "m." + Column1.NameGo
	//	ReplacePKFieldNamesFormat = ReplacePKFieldNamesFormat + Comma + Column1.NameGo + ": %v"
	//	sNumber := strconv.Itoa(i + 1)
	//	ReplaceWhereID = ReplaceWhereID + "\tand " + TableAlias + "." + Column1.Name + " = $" + sNumber + "\n"
	//	Comma = ", "
	//	CommaNewline = ",\n\t\t"
	//}
	//Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldsWithComma", ReplacePKFieldsWithComma)
	//Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldNamesFormat", ReplacePKFieldNamesFormat)
	//Otvet = strings.ReplaceAll(Otvet, "ReplaceWhereID", ReplaceWhereID)

	//все колонки
	ReplaceAllFieldsWithComma := ""
	CommaNewline2 := ""
	MassColumns := micro.MassFrom_Map(Table1.MapColumns)
	for _, Column1 := range MassColumns {
		if create_files.Is_Need_Сolumn(Column1) == false {
			continue
		}

		//
		ReplaceAllFieldsWithComma = ReplaceAllFieldsWithComma + CommaNewline2 + "&m." + Column1.NameGo

		CommaNewline2 = ",\n\t\t"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceAllFieldsWithComma", ReplaceAllFieldsWithComma)

	//0=null
	ReplaceFieldsWithComma := ""
	ColumnConnectionID := create_files.FindColumn_ConnectionID(Table1)
	if ColumnConnectionID.IsNullable == true {
		TextVariable := "m." + ColumnConnectionID.NameGo
		TextValue := create_files.FindText_NullValue(ColumnConnectionID.TypeGo, TextVariable)
		ReplaceFieldsWithComma = ReplaceFieldsWithComma + TextValue
	} else {
		ReplaceFieldsWithComma = ReplaceFieldsWithComma + "m.ConnectionID"
	}
	ReplaceFieldsWithComma = ReplaceFieldsWithComma + ", "
	ColumnExtID := create_files.FindColumn_ExtID(Table1)
	if ColumnExtID.IsNullable == true {
		TextVariable := "m." + ColumnExtID.NameGo
		TextValue := create_files.FindText_NullValue(ColumnExtID.TypeGo, TextVariable)
		ReplaceFieldsWithComma = ReplaceFieldsWithComma + TextValue
	} else {
		ReplaceFieldsWithComma = ReplaceFieldsWithComma + "m.ExtID"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceFieldsWithComma", ReplaceFieldsWithComma)

	return Otvet
}
