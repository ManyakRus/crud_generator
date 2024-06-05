package crud_functions

import (
	"gorm.io/gorm"
	"slices"
)

// MapOmit_from_MassOmit - создает MapOmit из MassOmit
func MapOmit_from_MassOmit(MassOmit []string) map[string]interface{} {
	MapOmit := make(map[string]interface{}, 0)
	for _, v := range MassOmit {
		if v == "CreatedAt" {
			continue
		}
		MapOmit[v] = gorm.Expr("NULL")
	}

	return MapOmit
}

// MassNeedFields_from_MassNeedUpdateFields - создает MassFields из MassNeedUpdateFields
// копирует все элементты и добавляет "ModifiedAt" и др.
func MassNeedFields_from_MassNeedUpdateFields(MassNeedUpdateFields []string) []string {
	//колонки для обновления
	MassFields := make([]string, len(MassNeedUpdateFields))
	copy(MassFields, MassNeedUpdateFields)
	MassFields = append(MassFields, "ModifiedAt")

	// DeletedAt вместе с IsDeleted
	if slices.Contains(MassNeedUpdateFields, "IsDeleted") == true {
		MassFields = append(MassFields, "DeletedAt")
	}

	//// CreatedAt вместе с ID=0
	//if slices.Contains(MassNeedUpdateFields, "ID") == true {
	//	MassFields = append(MassFields, "CreatedAt")
	//}

	return MassFields
}
