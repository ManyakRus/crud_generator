package postgres

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestFillMassTable(t *testing.T) {
	//ConfigMain.LoadEnv()
	config.LoadSettingsTxt()
	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	Otvet, err := FillMapTable()
	if err != nil {
		t.Error("TestFillMassTable() error: ", err)
	}

	if len(Otvet) == 0 {
		t.Error("TestFillMassTable() error: len =0")
	}
}
