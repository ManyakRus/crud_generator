package entities_tables

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"sort"
)

// AddFunctionStringIdentifier - добавляет функцию StringIdentifier(), для таблиц где много PrimaryKey
func AddFunctionStringIdentifier(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	if Table1.PrimaryKeyColumnsCount == 1 {
		return Otvet
	}

	_, TextNamesTypes, _ := create_files.FindText_IDMany(Table1)

	Text := `
// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(` + TextNamesTypes + `) string {
	Otvet := ""
`

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey == false {
			continue
		}
		TextConvert := create_files.FindText_ConvertToString(Column1, Column1.NameGo)
		Text = Text + "\tOtvet = Otvet + " + TextConvert + ` + "_"` + "\n"
	}

	Text = Text + `
	return Otvet
}`

	Otvet = Otvet + Text

	return Otvet
}
