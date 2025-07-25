package tables_tables

import (
	"errors"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке model
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	MassAll := micro.MassFrom_Map(MapAll)
	for _, Table1 := range MassAll {
		//проверка имени таблицы "DELETED_"
		err1 := create_files.IsGood_TableName(Table1)
		if err1 != nil {
			log.Warn("CreateFiles() table: ", Table1.Name, " warning: ", err)
			continue
		}

		//
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

// CreateFiles - создаёт 1 файл в папке db
func CreateFiles(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesTable := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_TABLES + micro.SeparatorFile()
	DirReadyTable := DirReady + config.Settings.TEMPLATE_FOLDERNAME_TABLES + micro.SeparatorFile() + config.Settings.PREFIX_TABLE + TableName + micro.SeparatorFile()

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	// создание файла struct
	if config.Settings.NEED_CREATE_DB_TABLES == true {
		err = CreateFiles_Table_struct(Table1, DirTemplatesTable, DirReadyTable)
		if err != nil {
			log.Error("CreateFiles_Table_struct() table: ", Table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

// CreateFiles_Table_struct - создаёт 1 файл со структурой в папке model
func CreateFiles_Table_struct(Table1 *types.Table, DirTemplatesTable, DirReadyTable string) error {
	var err error
	var ModelName string

	TableName := strings.ToLower(Table1.Name)
	FilenameTemplateModel := DirTemplatesTable + "table.go_"
	FilenameReadyModel := DirReadyTable + config.Settings.PREFIX_TABLE + TableName + ".go"

	//чтение файла шаблона
	bytes, err := micro.ReadFile_Linux_Windows(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//заменим имя пакета на новое
	TextModel = create_files.Replace_PackageName(TextModel, DirReadyTable)

	//удалим все свои импорты
	TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)

	//создание текста
	TextModel, TextModelStruct, ModelName, err := FindText_ModelStruct(TextModel, Table1)
	TextModel = Replace_ModelStruct(TextModel, TextModelStruct)

	//
	TextModel = create_files.Replace_TemplateModel_to_Model(TextModel, Table1.NameGo)
	TextModel = create_files.Replace_TemplateTableName_to_TableName(TextModel, Table1.Name)

	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextModel = config.Settings.TEXT_MODULE_GENERATED + TextModel

	if config.Settings.HAS_IS_DELETED == true {
		TextModel = DeleteFunc_Delete(TextModel, ModelName, Table1)
		TextModel = DeleteFunc_Restore(TextModel, ModelName, Table1)
	}
	TextModel = DeleteFunc_Find_byExtID(TextModel, ModelName, Table1)

	//import time
	TextModel = create_files.CheckAndAdd_ImportTime_FromText(TextModel)

	//import uuid
	TextModel = create_files.CheckAndAdd_ImportUUID_FromText(TextModel)

	//
	TextModel = create_files.DeleteImportModel(TextModel)

	//замена импортов на новые URL
	TextModel = create_files.Replace_RepositoryImportsURL(TextModel)

	//удаление пустого импорта
	TextModel = create_files.Delete_EmptyImport(TextModel)

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), config.Settings.FILE_PERMISSIONS)

	return err
}

// FindText_ModelStruct - возвращает текст структуры и тегов gorm
func FindText_ModelStruct(TextModel string, Table1 *types.Table) (string, string, string, error) {
	var Otvet string
	var ModelName string
	var err error

	TableName := Table1.Name
	ModelName = Table1.NameGo
	//if config.Settings.SINGULAR_TABLE_NAMES == true {
	//	ModelName = create_files.Find_SingularName(TableName)
	//}
	//ModelName = create_files.FormatName(ModelName)
	//Table1.NameGo = ModelName

	//удалим старые импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextModel = create_files.Delete_TemplateRepositoryImports(TextModel)
	}

	//	Otvet = `// ` + ModelName + ` - ` + COMMENT_MODEL_STRUCT + TableName + `: ` + Table1.Comment + `
	//type ` + ModelName + ` struct {
	//`

	Prefix := micro.StringFromUpperCase(config.Settings.PREFIX_TABLE)
	ModelNameWithPrefix := Prefix + ModelName
	Otvet = create_files.Find_ModelNameComment(ModelNameWithPrefix, Table1)
	Otvet = Otvet + `
type ` + ModelNameWithPrefix + ` struct {
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

	// если у id есть alias то колонка id будет отдельно
	ColumnIDName, _ := create_files.Find_PrimaryKeyNameType(Table1)
	_, ok := types.MapConvertID[TableName+"."+ColumnIDName]
	if ok == true {
		has_Columns_CommonStruct = false
	}

	//
	ImportModelsName := micro.LastWord(config.Settings.TEMPLATE_FOLDERNAME_TABLES)

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

		if create_files.Is_Need_Сolumn(Column1) == false {
			continue
		}

		var TextColumn string
		TextModel, TextColumn = FindText_Column(TextModel, Table1, Column1)
		Otvet = Otvet + TextColumn + "\n"
		Table1.MapColumns[key1] = Column1
	}

	//добавим новый импорт
	if has_Columns_CommonStruct == true || has_Columns_NameStruct == true || has_Columns_Groups == true || has_Columns_ExtLinks == true {
		TablesURL := create_files.FindURL_Tables()
		TextModel = create_files.AddImport(TextModel, TablesURL)
	}

	Otvet = Otvet + "\n}"
	return TextModel, Otvet, ModelName, err
}

// FindText_Column - возвращает текст gorm
func FindText_Column(TextModel string, Table1 *types.Table, Column1 *types.Column) (string, string) {
	Otvet := ""
	//	Code string `json:"code" gorm:"column:code;default:0"`

	ColumnName := Column1.Name
	ColumnNameLowerCase := strings.ToLower(ColumnName)
	ColumnModelName := create_files.FormatName(Column1.Name)

	Type_go := Column1.TypeGo
	TextModel, Type_go = Find_ColumnTypeGoImport(TextModel, Table1, Column1)
	//Column1.TypeGo = Type_go
	TextDefaultValue := ""
	if Column1.IsPrimaryKey == false {
		TextDefaultValue = create_files.FindText_DefaultGORMValue(Column1)
	}
	TextPrimaryKey := FindText_PrimaryKey(Column1.IsPrimaryKey)
	Description := Column1.Description
	Description = create_files.PrintableString(Description) //экранирование символов

	TextAutoCreateTime := ""
	TextAutoUpdateTime := ""
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		if ColumnNameLowerCase == "created_at" {
			TextAutoCreateTime = ";autoCreateTime;<-:create;"
			TextDefaultValue = ""
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

// FindText_PrimaryKey - возвращает строку gorm для primaryKey
func FindText_PrimaryKey(Is_identity bool) string {
	Otvet := ""

	if Is_identity == true {
		Otvet = ";primaryKey;autoIncrement:true"
	}

	return Otvet
}

// Replace_ModelStruct - заменяет структуру модели на новую
func Replace_ModelStruct(TextTemplateModel, TextModelStruct string) string {
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
		log.Panic("Replace_ModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	s2 := TextTemplateModel[pos1:]
	TextFind1 = "}\n"
	posEnd := strings.Index(s2, TextFind1)
	if posEnd < 0 {
		log.Panic("Replace_ModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	//
	Otvet = TextTemplateModel[:pos1] + TextModelStruct + TextTemplateModel[pos1+posEnd+1:]

	return Otvet
}

// DeleteFunc_Delete - удаляет функцию Delete()
func DeleteFunc_Delete(TextModel, ModelName string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	//FirstSymbol := strings.ToLower(ModelName)[:1]
	TextFind := "Delete(*" + ModelName + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	TextFind = "\n// Delete "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFunc_Restore - удаляет функцию Restore()
func DeleteFunc_Restore(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	//FirstSymbol := strings.ToLower(Modelname)[:1]
	TextFind := "Restore(*" + Modelname + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	TextFind = "\n// Restore "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFunc_Find_byExtID - удаляет функцию Find_ByExtID()
func DeleteFunc_Find_byExtID(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	//
	_, ok1 := Table1.MapColumns["ext_id"]

	//
	_, ok2 := Table1.MapColumns["connection_id"]
	if ok1 == true && ok2 == true {
		return Otvet
	}

	//FirstSymbol := strings.ToLower(Modelname)[:1]
	TextFind := "Find_ByExtID(*" + Modelname + ") error"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	TextFind = "\n// Find_ByExtID "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// Find_ColumnTypeGoImport - заменяет ID на Alias
func Find_ColumnTypeGoImport(TextModel string, Table1 *types.Table, Column1 *types.Column) (string, string) {
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
