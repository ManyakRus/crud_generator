package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindText_FindModelBy - возвращает TextProto и текст FindModelBy
func FindText_FindModelBy(TextProto string, Table1 *types.Table) (string, string) {
	Otvet := TextProto

	Otvet2 := ""
	for _, TableColumns1 := range types.MassFindModelBy {
		if TableColumns1.Table != Table1 {
			continue
		}

		Text1 := FindText_FindModelBy1(TableColumns1.Table, TableColumns1.Column)

		//проверим такой текст функции уже есть
		pos1 := strings.Index(TextProto, Text1)
		if pos1 >= 0 {
			continue
		}

		//добавим message
		TextMess := AddTextMessageRequestModel_Column(TextProto, TableColumns1.Column)

		//проверим такой текст message уже есть
		pos1 = strings.Index(TextProto, TextMess)
		if pos1 < 0 {
			Otvet = TextMess
		}

		//
		Otvet2 = Otvet2 + Text1
	}

	return Otvet, Otvet2
}

// FindText_FindModelBy1 - находит текст FindModelBy
func FindText_FindModelBy1(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := "\n\trpc "

	TextFields := ""
	//TextRequest := "Request_"
	Underline := ""
	//for _, Column1 := range TableColumns1.Columns {
	TextFields = TextFields + Underline + Column1.NameGo_translit
	//TextRequest1 := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
	//TextRequest = TextRequest + Underline + TextRequest1
	Underline = "_"
	//}
	TextRequest := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)

	Otvet = Otvet + Table1.NameGo_translit + "_FindModelBy_" + TextFields + "(" + TextRequest + ") returns (Response) {}\n"

	return Otvet
}
