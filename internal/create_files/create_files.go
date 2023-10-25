package create_files

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"strings"
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

	switch strings.ToLower(Name) {
	case "id":
		Otvet = "ID"
	default:
		Otvet = strcase.ToCamel(Otvet)
	}

	return Otvet
}
