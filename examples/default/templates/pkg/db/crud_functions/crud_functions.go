package crud_functions

import "gorm.io/gorm"

// MapOmit_from_MassOmit - создает MapOmit из MassOmit
func MapOmit_from_MassOmit(MassOmit []string) map[string]interface{} {
	MapOmit := make(map[string]interface{}, 0)
	for _, v := range MassOmit {
		MapOmit[v] = gorm.Expr("NULL")
	}

	return MapOmit
}
