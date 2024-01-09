package load_configs

import (
	"encoding/json"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/micro"
	"os"
	"testing"
)

func TestLoadCrudFunctionsRename(t *testing.T) {
	config.LoadSettingsTxt()
	config.FillSettings()
	config.FillFlags()

	LoadCrudFunctionsRename()
}

func TestSaveCrudFunctionsRename(t *testing.T) {
	t.SkipNow()

	var err error

	config.LoadSettingsTxt()
	config.FillSettings()
	config.FillFlags()

	dir := micro.ProgramDir_bin()
	FileName := dir + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile() + constants.CONFIG_FOLDER_NAME + micro.SeparatorFile() + "crud_functions_rename_test.json"

	//File, _ := os.OpenFile(FileName, 0, 666)
	//defer File.Close()

	var bytes []byte
	ReplaceStruct := types.ReplaceStruct{}
	ReplaceStruct.Old = "create_update_ctx"
	ReplaceStruct.New = "create_update_ctx_original"
	FunctionsReplace := make([]types.ReplaceStruct, 0)
	FunctionsReplace = append(FunctionsReplace, ReplaceStruct)
	types.MapRenameFunctions["lawsuits"] = FunctionsReplace
	//types.MapRenameFunctions["lawsuits2"] = FunctionsReplace
	bytes, err = json.MarshalIndent(types.MapRenameFunctions, "", "  ")
	if err != nil {

	}

	//_, err = File.Write([]byte("aaa"))
	os.WriteFile(FileName, bytes, 0666)
	if err != nil {

	}

}
