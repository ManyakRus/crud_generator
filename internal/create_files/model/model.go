package model

import (
	"errors"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
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

	for _, table1 := range MapAll {
		err = CreateFiles(table1)
		if err != nil {
			log.Error("CreateFiles() table: ", table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

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

	//создадим каталог
	ok, err := micro.FileExists(DirReadyModel)
	if ok == false {
		err = os.MkdirAll(DirReadyModel, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyModel, " error: ", err)
		}
	}

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
			log.Error("CreateFilesModel_struct() table: ", Table1.Name, " error: ", err)
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

	//чтение файла шаблона
	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextModel := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextModel, DirReadyModel)

	//создание текста
	//TextModel, TextModelStruct, ModelName, err := FindTextModelStruct(TextModel, Table1)
	//TextModel = ReplaceModelStruct(TextModel, TextModelStruct)
	//
	////
	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	////TextModel = config.Settings.TEXT_MODULE_GENERATED + TextModel
	//
	//if config.Settings.HAS_IS_DELETED == true {
	//	TextModel = DeleteFuncDelete(TextModel, ModelName, Table1)
	//	TextModel = DeleteFuncRestore(TextModel, ModelName, Table1)
	//}
	//TextModel = DeleteFuncFind_byExtID(TextModel, ModelName, Table1)

	TextModel = create_files.CheckAndAddImportTime_FromText(TextModel)
	//TextModel = create_files.DeleteImportModel(TextModel)

	//замена импортов на новые URL
	//TextModel = create_files.ReplaceServiceURLImports(TextModel)

	TextModel = create_files.ReplaceModelAndTableName(TextModel, Table1)

	//замена импортов на новые URL
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Comment := create_files.FindModelComment(Table1)
		TextTemplate := "// " + ModelName
		TextModel = strings.ReplaceAll(TextModel, TextTemplate, Comment)

		TextModel = create_files.DeleteTemplateRepositoryImports(TextModel)

		TableNameURL := create_files.FindTableNameURL(TableName)
		TextModel = create_files.AddImport(TextModel, TableNameURL)
	}

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesModel_crud - создаёт 1 файл с crud операциями
func CreateFilesModel_crud(Table1 *types.Table, DirTemplatesModel, DirReadyModel string) error {
	var err error

	ModelName := Table1.NameGo

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
	create_files.ReplacePackageName(TextModel, DirReadyModel)

	//создание текста
	TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextModel = strings.ReplaceAll(TextModel, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextModel = config.Settings.TEXT_MODULE_GENERATED + TextModel

	if config.Settings.HAS_IS_DELETED == true {
		TextModel = DeleteFuncDelete(TextModel, ModelName, Table1)
		TextModel = DeleteFuncRestore(TextModel, ModelName, Table1)
	}
	TextModel = DeleteFuncFind_byExtID(TextModel, ModelName, Table1)

	TextModel = create_files.CheckAndAddImportTime_FromText(TextModel)
	TextModel = create_files.DeleteImportModel(TextModel)

	//замена импортов на новые URL
	TextModel = create_files.ReplaceServiceURLImports(TextModel)

	//запись файла
	err = os.WriteFile(FilenameReadyModel, []byte(TextModel), constants.FILE_PERMISSIONS)

	return err
}

// FindTextModelStruct - возвращает текст структуры и тегов gorm
func FindTextModelStruct(TextModel string, Table1 *types.Table) (string, string, string, error) {
	var Otvet string
	var ModelName string
	var err error

	TableName := Table1.Name
	//ModelName = create_files.FindSingularName(TableName)
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
	if Column1.IsIdentity == false {
		TextDefaultValue = create_files.FindTextDefaultValue(Column1)
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

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(TextModel, ModelName string, Table1 *types.Table) string {
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

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(TextModel, Modelname string, Table1 *types.Table) string {
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

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(TextModel, Modelname string, Table1 *types.Table) string {
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
