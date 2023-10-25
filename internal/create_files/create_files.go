package create_files

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

// FindSingularName - возвращает наименование в единственном числе
func FindSingularName(s string) string {
	var Otvet string

	if s == "" {
		return Otvet
	}

	Otvet = inflection.Singular(s)

	return Otvet
}

// FormatName - возвращает наименование в формате CamelCase
func FormatName(Name string) string {
	Otvet := Name

	Otvet = strcase.ToCamel(Otvet)

	return Otvet
}
