package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindText_FindBy - возвращает TextProto и текст FindBy
func FindText_FindBy(TextProto string, Table1 *types.Table) (string, string) {
	Otvet := TextProto

	Otvet2 := ""
	for _, TableColumns1 := range types.MassFindBy {
		if TableColumns1.Table.Name != Table1.Name {
			continue
		}

		Text1 := FindText_FindBy1(TableColumns1)

		//проверим такой текст функции уже есть
		pos1 := strings.Index(Otvet, Text1)
		if pos1 >= 0 {
			continue
		}

		//добавим message
		TextMess := AddTextMessageRequestID_Columns(Otvet, TableColumns1.Columns)

		//проверим такой текст message уже есть
		pos1 = strings.Index(Otvet, TextMess)
		if pos1 < 0 {
			Otvet = Otvet + "\n" + TextMess
		}

		//
		Otvet2 = Otvet2 + Text1
	}

	return Otvet, Otvet2
}

// FindText_FindBy1 - находит текст FindBy
func FindText_FindBy1(TableColumns1 types.TableColumns) string {
	Otvet := "\n\trpc "

	TextFields := ""
	TextRequest := "Request_"
	Underline := ""
	for _, Column1 := range TableColumns1.Columns {
		TextFields = TextFields + Underline + Column1.NameGo_translit
		TextRequest1 := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
		TextRequest = TextRequest + Underline + TextRequest1
		Underline = "_"
	}

	Otvet = Otvet + TableColumns1.Table.NameGo_translit + "_FindBy_" + TextFields + "(" + TextRequest + ") returns (Response) {}\n"

	return Otvet
}
