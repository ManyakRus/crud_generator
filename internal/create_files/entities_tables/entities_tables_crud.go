package entities_tables

import (
	"errors"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFiles - создаёт 1 файл в папке model
func CreateFiles(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile() + TableName + micro.SeparatorFile()

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyModel)

	// создание файла struct
	if config.Settings.NEED_CREATE_MODEL_STRUCT == true {
		err = CreateFilesModel_struct(Table1, DirTemplatesModel, DirReadyModel)
		if err != nil {
			log.Error("CreateFilesModel_struct() table: ", Table1.Name, " error: ", err)
			return err
		}
	}

	// создание файла crud
	if config.Settings.NEED_CREATE_MODEL_CRUD == true {
		err = CreateFilesModel_crud(Table1, DirTemplatesModel, DirReadyModel)
		if err != nil {
			log.Error("CreateFilesModel_crud() table: ", Table1.Name, " error: ", err)
			return err
		}
	}

	// создание файла manual
	if config.Settings.NEED_CREATE_MANUAL_FILES == true {
		err = CreateFilesModel_manual(Table1, DirTemplatesModel, DirReadyModel)
		if err != nil {
			log.Error("CreateFilesModel_manual() table: ", Table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

// CreateFilesModel_struct - создаёт 1 файл со структурой в папке model
func CreateFilesModel_struct(Table1 *types.Table, DirTemplatesModel, DirReadyModel string) error {
	var err error
	//var ModelName string

	TableName := strings.ToLower(Table1.Name)
	ModelName := Table1.NameGo
	FilenameTemplateModel := DirTemplatesModel + "model.go_"
	FilenameReadyModel := DirReadyModel + config.Settings.PREFIX_MODEL + TableName + ".go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyModel)

	//чтение файла шаблона
	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//заменим имя пакета на новое
	TextModel = create_files.Replace_PackageName(TextModel, DirReadyModel)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)

		TableURL := create_files.Find_TableNameURL(TableName)
		TextModel = create_files.AddImport(TextModel, TableURL)
	}

	TextModel = create_files.CheckAndAdd_ImportTime_FromText(TextModel)
	TextModel = create_files.Replace_ModelAndTableName(TextModel, Table1)

	//замена импортов на новые URL
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Comment := create_files.Find_ModelComment(Table1)
		TextTemplate := "// " + ModelName
		TextModel = strings.ReplaceAll(TextModel, TextTemplate, Comment)

		TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)

		TableNameURL := create_files.Find_TableNameURL(TableName)
		TextModel = create_files.AddImport(TextModel, TableNameURL)
	}

	//удаление пустого импорта
	TextModel = create_files.Delete_EmptyImport(TextModel)

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesModel_crud - создаёт 1 файл с crud операциями
func CreateFilesModel_crud(Table1 *types.Table, DirTemplatesModel, DirReadyModel string) error {
	var err error

	//ModelName := Table1.NameGo

	TableName := strings.ToLower(Table1.Name)
	FilenameTemplateModel := DirTemplatesModel + "model_crud.go_"
	FilenameReadyModel := DirReadyModel + TableName + "_crud.go"

	//чтение файла шаблона
	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//заменим имя пакета на новое
	TextModel = create_files.Replace_PackageName(TextModel, DirReadyModel)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)

		TableURL := create_files.Find_CalcStructVersionURL()
		TextModel = create_files.AddImport(TextModel, TableURL)

		DBConstantsURL := create_files.Find_DBConstantsURL()
		TextModel = create_files.AddImport(TextModel, DBConstantsURL)

		//удалим лишние функции
		TextModel = create_files.DeleteFunc_Delete(TextModel, Table1)
		TextModel = create_files.DeleteFunc_Restore(TextModel, Table1)
		TextModel = create_files.DeleteFunc_Find_byExtID(TextModel, Table1)

		//удалим лишние функции из интерфейса
		TextModel = DeleteFromInterfaceDelete(TextModel, Table1)
		TextModel = DeleteFromInterfaceRestore(TextModel, Table1)
		TextModel = DeleteFromInterfaceFind_ByExtID(TextModel, Table1)
		TextModel = DeleteFromInterfaceUpdateManyFields(TextModel, Table1)
		//кэш
		if config.Settings.NEED_CREATE_CACHE_API == false {
			TextModel = DeleteFromInterfaceReadFromCache(TextModel, Table1)
			TextModel = DeleteFunc_ReadFromCache(TextModel, Table1)
		}
		TextModel = Replace_IDtoID(TextModel, Table1)

		//
		TextModel = AddFunctionStringIdentifier(TextModel, Table1)

		//добавим импорт uuid
		TextModel = create_files.CheckAndAdd_ImportUUID_FromText(TextModel)

		//добавим импорт strconv
		TextModel = create_files.CheckAndAdd_ImportStrconv(TextModel)

		//добавим импорт fmt
		TextModel = create_files.CheckAndAdd_ImportFmt(TextModel)

		//
		TextModel = AddInterfacesFindBy(TextModel, Table1)

		//
		TextModel = AddInterfacesFindMassBy(TextModel, Table1)
	}

	//
	if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
		TextModel = AddInterfaceUpdateEveryColumn(TextModel, Table1)
	}

	//создание текста
	TextModel = create_files.Replace_TemplateModel_to_Model(TextModel, Table1.NameGo)
	TextModel = create_files.Replace_TemplateTableName_to_TableName(TextModel, Table1.Name)
	TextModel = create_files.AddText_ModuleGenerated(TextModel)

	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextModel = config.Settings.TEXT_MODULE_GENERATED + TextModel

	TextModel = create_files.CheckAndAdd_ImportTime_FromText(TextModel)
	TextModel = create_files.DeleteImportModel(TextModel)

	//замена импортов на новые URL
	TextModel = create_files.Replace_RepositoryImportsURL(TextModel)

	//удаление пустого импорта
	TextModel = create_files.Delete_EmptyImport(TextModel)

	//удаление функций
	TextModel = DeleteFunctions(TextModel, TableName, types.MapModelCrudDeleteFunctions)

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFunctions - удаляет функции из текста, по карте MapModelCrudDeleteFunctions
func DeleteFunctions(Text, TableName string, MapModelCrudDeleteFunctions map[string]string) string {
	Otvet := Text

	TextDelete, ok := MapModelCrudDeleteFunctions[TableName]
	if ok == false {
		return Otvet
	}

	MassDelete := strings.Split(TextDelete, ",")
	for _, FunctionName1 := range MassDelete {
		Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// "+FunctionName1)
		Otvet = create_files.DeleteFuncFromFuncName(Otvet, FunctionName1)
	}

	return Otvet
}

// FindTextModelStruct - возвращает текст структуры и тегов gorm
func FindTextModelStruct(TextModel string, Table1 *types.Table) (string, string, string, error) {
	var Otvet string
	var ModelName string
	var err error

	TableName := Table1.Name
	//ModelName = create_files.Find_SingularName(TableName)
	//ModelName = create_files.FormatName(ModelName)
	//Table1.NameGo = ModelName
	COMMENT_MODEL_STRUCT := config.Settings.COMMENT_MODEL_STRUCT

	Otvet = `// ` + ModelName + ` - ` + COMMENT_MODEL_STRUCT + TableName + `: ` + Table1.Comment + `
type ` + ModelName + ` struct {
`

	//сортировка
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	has_Columns_CommonStruct := create_files.Has_Columns_CommonStruct(Table1)
	has_Columns_NameStruct := create_files.Has_Columns_NameStruct(Table1)
	has_Columns_Groups := create_files.Has_Columns_Groups(Table1)
	has_Columns_ExtLinks := create_files.Has_Columns_ExtLink(Table1)

	ImportModelsName := micro.LastWord(config.Settings.TEMPLATE_FOLDERNAME_MODEL)

	if has_Columns_CommonStruct == true {
		Otvet = Otvet + "\t" + ImportModelsName + ".CommonStruct\n"
	}

	if has_Columns_NameStruct == true {
		Otvet = Otvet + "\t" + ImportModelsName + ".NameStruct\n"
	}

	if has_Columns_Groups == true {
		Otvet = Otvet + "\t" + ImportModelsName + ".GroupStruct\n"
	}

	if has_Columns_ExtLinks == true {
		Otvet = Otvet + "\t" + ImportModelsName + ".ExtLinkStruct\n"
	}

	//цикл по всем колонкам
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]

		//пропускаем колонки если они уже есть в CommonStruct
		if has_Columns_CommonStruct == true && create_files.Is_Column_CommonStruct(Column1) == true {
			continue
		} else if has_Columns_NameStruct == true && create_files.Is_Column_NameStruct(Column1) == true {
			continue
		} else if has_Columns_Groups == true && create_files.Is_Column_GroupsStruct(Column1) == true {
			continue
		} else if has_Columns_ExtLinks == true && create_files.Is_Column_ExtLinksStruct(Column1) == true {
			continue
		}

		var TextColumn string
		TextModel, TextColumn = FindTextColumn(TextModel, Table1, Column1)
		Otvet = Otvet + TextColumn + "\n"
		Table1.MapColumns[key1] = Column1
	}

	Otvet = Otvet + "\n}"
	return TextModel, Otvet, ModelName, err
}

// FindTextColumn - возвращает текст gorm
func FindTextColumn(TextModel string, Table1 *types.Table, Column1 *types.Column) (string, string) {
	Otvet := ""
	//	Code string `json:"code" gorm:"column:code;default:0"`

	ColumnName := Column1.Name
	ColumnNameLowerCase := strings.ToLower(ColumnName)
	ColumnModelName := create_files.FormatName(Column1.Name)
	//Column1.NameGo = ColumnModelName
	//SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
	//if ok == false {
	//	log.Panic("GetMappings() ", Column1.Type, " error: not found")
	//}
	//Type_go := SQLMapping1.GoType
	Type_go := Column1.TypeGo
	TextModel, Type_go = FindColumnTypeGoImport(TextModel, Table1, Column1)
	//Column1.TypeGo = Type_go
	TextDefaultValue := ""
	if Column1.IsPrimaryKey == false {
		TextDefaultValue = create_files.FindText_DefaultGORMValue(Column1)
	}
	TextPrimaryKey := FindTextPrimaryKey(Column1.IsIdentity)
	Description := Column1.Description
	Description = create_files.PrintableString(Description) //экранирование символов

	TextAutoCreateTime := ""
	TextAutoUpdateTime := ""
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		if ColumnNameLowerCase == "created_at" {
			TextAutoCreateTime = ";autoCreateTime"
		}

		if ColumnNameLowerCase == "modified_at" {
			TextAutoUpdateTime = ";autoUpdateTime"
		}
	}

	Otvet = Otvet + "\t" + ColumnModelName + " " + Type_go + "\t"
	Otvet = Otvet + "`json:\"" + ColumnName + "\""
	Otvet = Otvet + " gorm:\"column:" + ColumnName + TextPrimaryKey + TextDefaultValue + TextAutoCreateTime + TextAutoUpdateTime + "\""
	//Otvet = Otvet + " db:\"" + ColumnName + "\""
	Otvet = Otvet + "`"
	Otvet = Otvet + "\t//" + Description

	return TextModel, Otvet
}

// FindTextPrimaryKey - возвращает строку gorm для primaryKey
func FindTextPrimaryKey(Is_identity bool) string {
	Otvet := ""

	if Is_identity == true {
		Otvet = ";primaryKey;autoIncrement:true"
	}

	return Otvet
}

// ReplaceModelStruct - заменяет структуру модели на новую
func ReplaceModelStruct(TextTemplateModel, TextModelStruct string) string {
	Otvet := ""

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL

	//найдём начало и конец
	TextFind1 := "// " + ModelName
	pos1 := strings.Index(TextTemplateModel, TextFind1)
	if pos1 < 0 {
		TextFind1 := "type " + ModelName + " struct {"
		pos1 = strings.Index(TextTemplateModel, TextFind1)
	}

	if pos1 < 0 {
		log.Panic("ReplaceModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	s2 := TextTemplateModel[pos1:]
	TextFind1 = "}\n"
	posEnd := strings.Index(s2, TextFind1)
	if posEnd < 0 {
		log.Panic("ReplaceModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	//
	Otvet = TextTemplateModel[:pos1] + TextModelStruct + TextTemplateModel[pos1+posEnd+1:]

	return Otvet
}

// FindColumnTypeGoImport - заменяет ID на Alias
func FindColumnTypeGoImport(TextModel string, Table1 *types.Table, Column1 *types.Column) (string, string) {
	Otvet := Column1.TypeGo

	//тип колонки из БД или из convert_id.json
	TableName := Table1.Name
	ColumnName := Column1.Name
	TextConvert, ok := types.MapConvertID[TableName+"."+ColumnName]
	if ok == false {
		return TextModel, Otvet
	}

	Otvet = TextConvert

	//добавим импорт
	URL := create_files.FindURL_Alias()
	if URL == "" {
		return TextModel, Otvet
	}

	TextModel = create_files.AddImport(TextModel, URL)

	return TextModel, Otvet
}

// FillColumnsNameGo - заполняет NameGo во все колонки
func FillColumnsNameGo(MapAll *map[string]*types.Table) error {
	var err error

	for _, Table1 := range *MapAll {
		for _, Column1 := range Table1.MapColumns {
			ColumnName := Column1.Name
			ColumnModelName := create_files.FormatName(ColumnName)
			Column1.NameGo = ColumnModelName
			if ColumnModelName == "" {
				err = errors.New("Table: " + Table1.Name + " Column: " + ColumnName + " = \"\"")
			}
		}

	}

	return err
}

// DeleteFromInterfaceDelete - удаляет функцию Delete() из интерфейса
func DeleteFromInterfaceDelete(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	TextFind := "\n\tDelete(*" + ModelName + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	return Otvet
}

// DeleteFromInterfaceRestore - удаляет функцию Restore() из интерфейса
func DeleteFromInterfaceRestore(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	TextFind := "\n\tRestore(*" + ModelName + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	return Otvet
}

// DeleteFromInterfaceFind_ByExtID - удаляет функцию Find_ByExtID() из интерфейса
func DeleteFromInterfaceFind_ByExtID(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонки ExtID и ConnectionID
	if create_files.Has_Column_ExtID_ConnectionID_Int64(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	TextFind := "\n\tFind_ByExtID(*" + ModelName + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	return Otvet
}

// CreateFilesModel_manual - создаёт 1 файл с _manual.go
func CreateFilesModel_manual(Table1 *types.Table, DirTemplatesModel, DirReadyModel string) error {
	var err error

	//
	//ModelName := Table1.NameGo

	TableName := strings.ToLower(Table1.Name)
	FilenameTemplateModel := DirTemplatesModel + constants.MODEL_TABLE_MANUAL_FILENAME
	FilenameReadyModel := DirReadyModel + TableName + "_manual.go"

	//чтение файла шаблона
	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//заменим имя пакета на новое
	TextModel = create_files.Replace_PackageName(TextModel, DirReadyModel)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)
	}

	//создание текста
	TextModel = create_files.Replace_TemplateModel_to_Model(TextModel, Table1.NameGo)
	TextModel = create_files.Replace_TemplateTableName_to_TableName(TextModel, Table1.Name)

	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	//замена импортов на новые URL
	TextModel = create_files.Replace_RepositoryImportsURL(TextModel)

	//удаление пустого импорта
	TextModel = create_files.Delete_EmptyImport(TextModel)

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFromInterfaceReadFromCache - удаляет функцию ReadFromCache() из интерфейса
func DeleteFromInterfaceReadFromCache(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	TextFind := "\n\tReadFromCache(*" + ModelName + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	return Otvet
}

// DeleteFromInterfaceUpdateManyFields - удаляет функцию UpdateManyFields() из интерфейса
func DeleteFromInterfaceUpdateManyFields(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	TextFind := "\n\tUpdateManyFields(*" + ModelName + ", []string) error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	return Otvet
}

// DeleteFunc_ReadFromCache - удаляет функцию ReadFromCache()
func DeleteFunc_ReadFromCache(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// ReadFromCache ")
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "ReadFromCache")

	return Otvet
}

// Replace_IDtoID_Many - заменяет int64(ID) на ID, и остальные PrimaryKey
func Replace_IDtoID(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextNames, TextNamesTypes, TextProtoNames := create_files.FindText_IDMany(Table1)

	Otvet = strings.ReplaceAll(Otvet, "ReplaceManyID(ID)", TextNames)
	Otvet = strings.ReplaceAll(Otvet, "int64(ID)", TextProtoNames)
	Otvet = strings.ReplaceAll(Otvet, "(ID int64", "("+TextNamesTypes)
	Otvet = strings.ReplaceAll(Otvet, "(ID)", "("+TextNames+")")
	Otvet = strings.ReplaceAll(Otvet, ", ID)", ", "+TextNames+")")
	Otvet = strings.ReplaceAll(Otvet, ", ID int64)", ", "+TextNamesTypes+")")

	return Otvet
}