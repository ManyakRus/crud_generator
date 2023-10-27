package create_files

import "testing"

func TestFindSingularName(t *testing.T) {
	Otvet := FindSingularName("lawsuit_status_types")
	if Otvet == "" {
		t.Error("TestFindSingularName() error: Otvet =''")
	}
}

func TestFormatName(t *testing.T) {
	Name := "contract_id"
	Otvet := FormatName(Name)
	if Otvet == "" {
		t.Error("TestFormatName() error")
	}
}
