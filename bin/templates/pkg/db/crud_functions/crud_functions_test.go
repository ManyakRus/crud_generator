package crud_functions

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestMapOmit_from_MassOmit(t *testing.T) {
	t.Run("Empty input slice", func(t *testing.T) {
		MassOmit := []string{}
		result := MapOmit_from_MassOmit(MassOmit)
		if len(result) != 0 {
			t.Errorf("Expected an empty map, but got %w", result)
		}
	})

	t.Run("Non-empty input slice", func(t *testing.T) {
		MassOmit := []string{"key1", "key2"}
		result := MapOmit_from_MassOmit(MassOmit)
		expected := map[string]interface{}{
			"key1": gorm.Expr("NULL"),
			"key2": gorm.Expr("NULL"),
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestMassNeedFields_from_MassNeedUpdateFields(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"Empty input", []string{}, []string{"ModifiedAt"}},
		//{"Contains IsDeleted and ID", []string{"IsDeleted", "ID"}, []string{"IsDeleted", "ID", "ModifiedAt", "DeletedAt", "CreatedAt"}},
		{"Contains only IsDeleted", []string{"IsDeleted"}, []string{"IsDeleted", "ModifiedAt", "DeletedAt"}},
		//{"Contains only ID", []string{"ID"}, []string{"ID", "ModifiedAt", "CreatedAt"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MassNeedFields_from_MassNeedUpdateFields(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
