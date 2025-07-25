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

// CreateFiles_Soft_Restore - создаёт 1 файл
func CreateFiles_Soft_Restore(MapAll map[string]*types.Table, Table1 *types.Table) error {
	var err error

	//
	if create_files.Has_Column_IsDeleted_Bool(Table1) == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()
	DirReadyDB := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + config.Settings.TEMPLATES_CRUD_SOFT_RESTORE_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyDB + config.Settings.PREFIX_CRUD + TableName
	FilenameReadyDB := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_CRUD + TableName + "_soft_restore.go"

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

	TextDB = CreateFiles_Soft_Restore1(TextDB, Table1)

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

	//удаление пустых строк
	TextDB = create_files.Delete_EmptyLines(TextDB)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_Soft_Restore1 - заполняет Text
func CreateFiles_Soft_Restore1(Text string, Table1 *types.Table) string {
	Otvet := Text

	//Primary key
	ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)
	//TableAlias := create_files.Find_TableAlias(Table1)
	ReplacePKFieldsWithComma := ""
	ReplacePKFieldNamesFormat := ""
	ReplaceID0 := ""
	ReplaceWhereID := ""
	Comma := ""
	TextAnd := ""
	for i, Column1 := range ColumnsPK {
		sNumber := strconv.Itoa(i + 1)
		ReplacePKFieldsWithComma = ReplacePKFieldsWithComma + Comma + "m." + Column1.NameGo
		ReplacePKFieldNamesFormat = ReplacePKFieldNamesFormat + Comma + Column1.NameGo + ": %v"
		TextEmpty := create_files.FindText_EqualEmpty(Column1, "m."+Column1.NameGo)
		ReplaceID0 = ReplaceID0 + TextAnd + TextEmpty
		ReplaceWhereID = ReplaceWhereID + "\tand " + Column1.Name + " = $" + sNumber + "\n"

		Comma = ", "
		TextAnd = " && "
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldsWithComma", ReplacePKFieldsWithComma)
	Otvet = strings.ReplaceAll(Otvet, "ReplacePKFieldNamesFormat", ReplacePKFieldNamesFormat)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceID0", ReplaceID0)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceWhereID", ReplaceWhereID)

	ReplaceDeletedAtEqual := ""
	ReplaceFieldDeletedAt := ""
	if create_files.Has_Column_DeletedAt_Time(Table1) == true {
		ReplaceDeletedAtEqual = ", deleted_at = null"
		ReplaceFieldDeletedAt = "m.DeletedAt = time.Time{}"
	}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceDeletedAtEqual", ReplaceDeletedAtEqual)
	Otvet = strings.ReplaceAll(Otvet, "ReplaceFieldDeletedAt", ReplaceFieldDeletedAt)

	//все колонки
	//ReplaceAllColumnNamesWithComma := ""
	//ReplaceDollarsWithComma := ""
	//Comma = ""
	//CommaNewline := "\t"
	//CommaNewline2 := ""
	//MassColumns := micro.MassFrom_Map(Table1.MapColumns)
	//for i, Column1 := range MassColumns {
	//	sNumber := strconv.Itoa(i + 1)
	//
	//	if Column1.IsPrimaryKey == true {
	//		ReplaceWhereID = ReplaceWhereID + "\tand " + Column1.Name + " = $" + sNumber + "\n"
	//	}
	//
	//	//Comma = ", "
	//	CommaNewline = ",\n\t"
	//	CommaNewline2 = ",\n\t\t"
	//}
	Otvet = strings.ReplaceAll(Otvet, "ReplaceTableName", Table1.Name)

	return Otvet
}
