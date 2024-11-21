package create_files

import (
	"github.com/gobeam/stringy"
	"testing"
)

func TestFindSingularName(t *testing.T) {
	Otvet := Find_SingularName("lawsuit_status_types")
	if Otvet == "" {
		t.Error("TestFindSingularName() error: Otvet =''")
	}
}

func TestFormatName(t *testing.T) {
	Name := "first_1min_candle_date"
	Otvet := FormatName(Name)
	if Otvet == "" {
		t.Error("TestFormatName() error")
	}
}

func TestDeleteLineWithComment(t *testing.T) {
	s := `import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_types"
	"testing"
	//TestFind_ByExtID() "gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
)
`
	Otvet := DeleteLineWithComment(s, "TestFind_ByExtID")
	if Otvet == "" {
		t.Error("TestFormatName() error")
	}
}

func TestPrintableString(t *testing.T) {
	s := `123 // \\ 456
789`
	Otvet := PrintableString(s)
	if Otvet == "" {
		t.Error("TestPrintableString() error: Otvet = ''")
	}
}

func TestPascalCase(t *testing.T) {
	Name := "first_1min_candle_date"
	str := stringy.New(Name)
	Otvet := str.PascalCase("1m", "1M").Get()
	if Otvet == "" {
		t.Error("TestFormatName() error")
	}
}
