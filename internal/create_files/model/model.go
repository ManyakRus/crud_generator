package model

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/micro"
	"log"
	"os"
	"strings"
)

func CreateModelFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, table1 := range MapAll {
		err = CreateModelFiles1(table1)
		if err != nil {
			return err
		}
	}

	return err
}

func CreateModelFiles1(Table1 *types.Table) error {
	var err error

	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()
	DirTemplatesModel := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_MODEL + micro.SeparatorFile()
	DirReadyModel := DirReady + "pkg" + micro.SeparatorFile() + "model" + micro.SeparatorFile()

	FilenameTemplateModel := DirTemplatesModel + "model.go_"
	FilenameReadyModel := DirReadyModel + strings.ToLower(Table1.Name) + ".go"

	bytes, err := os.ReadFile(FilenameTemplateModel)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateModel, " error: ", err)
	}
	TextTemplateModel := string(bytes)

	TextModelStruct, ModelName, err := FindTextModelStruct(Table1)
	TextTemplateModel = ReplaceModelStruct(TextTemplateModel, TextModelStruct)

	return err
}

func FindTextModelStruct(Table1 *types.Table) (string, string, error) {
	var Otvet string
	var ModelName string
	var err error

	TableName := Table1.Name
	ModelName = create_files.FindSingularName(TableName)
	ModelName = create_files.FormatName(ModelName)

	Otvet = `// ` + ModelName + ` - model from table ` + TableName + `
type ` + ModelName + ` struct {
`
	for _, Column1 := range Table1.MapColumns {
		TextColumn := FindTextColumn(Column1)
		Otvet = Otvet + TextColumn + "\n"
	}

	Otvet = Otvet + "\n}"
	return Otvet, ModelName, err
}

func FindTextColumn(Column1 types.Column) string {
	Otvet := ""
	//	Code string `json:"code" gorm:"column:code;default:0"`

	ColumnName := Column1.Name
	ColumnModelName := create_files.FormatName(Column1.Name)
	SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
	if ok == false {
		log.Panic("GetMappings() ", Column1.Type, " error: not found")
	}
	Type_go := SQLMapping1.GoType
	TextDefaultValue := FindTextDefaultValue(Type_go)
	TextPrimaryKey := FindTextPrimaryKey(Column1.Is_identity)

	Otvet = Otvet + "\t" + ColumnModelName + " " + Type_go + " `json:\"" + ColumnName + "\"" + "gorm:\"column:" + ColumnName + TextPrimaryKey + TextDefaultValue
	Otvet = Otvet + "`"

	return Otvet
}

func FindTextDefaultValue(Type_go string) string {
	var Otvet string

	sValue := ""
	switch Type_go {
	case "string":
		sValue = "\"\""
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		sValue = "0"
	case "time.Time":
		sValue = "null"
	}

	if sValue != "" {
		Otvet = ";default:" + sValue
	}

	return Otvet
}

func FindTextPrimaryKey(Is_identity bool) string {
	Otvet := ""

	if Is_identity == true {
		Otvet = ";primaryKey;autoIncrement:true"
	}

	return Otvet
}

func ReplaceModelStruct(TextTemplateModel, TextModelStruct string) string {
	Otvet := ""

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL

	TextFind1 := "// " + ModelName
	pos1 := strings.Index(TextTemplateModel, TextFind1)
	if pos1 == 0 {
		TextFind1 := "type " + ModelName + " struct {"
		pos1 = strings.Index(TextTemplateModel, TextFind1)
	}

	if pos1 == 0 {
		log.Panic("ReplaceModelStruct() error: in model.go_ not found text: ", TextFind1)
	}

	return Otvet
}
