package create_files

import (
	"github.com/ManyakRus/crud_generator/internal/types"
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

	Otvet1, ok := types.MapReplaceName[Name]
	if ok == true {
		return Otvet1
	}

	switch strings.ToLower(Name) {
	case "id":
		Otvet = "ID"
	default:
		Otvet = strcase.ToCamel(Otvet)
	}

	//_id в конце заменяем на ID
	len1 := len(Name)
	if len1 >= 3 {
		last3 := strings.ToLower(Name[len1-3:])
		if last3 == "_id" {
			Otvet = Otvet[:len1-2-1] + "ID"
		}
	}

	return Otvet
}

// DeleteFuncFromFuncName - удаляет функцию из текста начиная с объявления функции
func DeleteFuncFromFuncName(Text, FuncName string) string {
	Otvet := Text

	TextFind := "\nfunc " + FuncName + "("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncFromComment - удаляет функцию из текста начиная с комментария
func DeleteFuncFromComment(Text, Comment string) string {
	Otvet := Text

	TextFind := Comment //"\n// Delete "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet

}

// Has_Column_ExtID_ConnectionID - возвращает true если есть поля ExtId и ConnectionID
func Has_Column_ExtID_ConnectionID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}

	//
	_, ok = Table1.MapColumns["connection_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_IsDeleted - возвращает true если есть поле is_deleted
func Has_Column_IsDeleted(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["is_deleted"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}
