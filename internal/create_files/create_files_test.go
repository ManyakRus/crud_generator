package create_files

import "testing"

func TestFindSingularName(t *testing.T) {
	Otvet := FindSingularName("lawsuit_status_types")
	if Otvet == "" {
		t.Error("TestFindSingularName() error: Otvet =''")
	}
}

func TestFormatName(t *testing.T) {
	Name := "document_invoice_id"
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
