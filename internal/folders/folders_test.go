package folders

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCreateAllFolders(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()
	CreateAllFolders()
}
