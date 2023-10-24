package model

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
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

	TextModelStruct, err := FindTextModelStruct(Table1)

	return err
}

func FindTextModelStruct(Table1 *types.Table) (string, error) {
	var Otvet string
	var err error

	TableName := Table1.Name
	ModelName := create_files.FindSingularName(TableName)

	Otvet = `// ` + ModelName + ` - model from table ` + TableName + `
type ` + ModelName + ` struct {
`

	Otvet = Otvet + "}"
	return Otvet, err
}
