package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindText_FindMassBy - добавляет текст FindBy
func FindText_FindMassBy(TextProto string, Table1 *types.Table) (string, string) {
	Otvet := TextProto

	Otvet2 := ""
	for _, TableColumns1 := range types.MassFindMassBy {
		if TableColumns1.Table.Name != Table1.Name {
			continue
		}

		Text1 := FindText_FindMassBy1(TableColumns1)

		//проверим такой текст уже есть
		pos1 := strings.Index(TextProto, Text1)
		if pos1 >= 0 {
			continue
		}

		//добавим message
		TextMess := AddTextMessageRequestID_Columns(TextProto, TableColumns1.Columns)
		//проверим такой текст message уже есть
		pos1 = strings.Index(TextProto, TextMess)
		if pos1 < 0 {
			Otvet = Otvet + "\n" + TextMess
		}

		//
		Otvet2 = Otvet2 + Text1
	}

	return Otvet, Otvet2
}

// FindText_FindMassBy1 - находит текст FindBy
func FindText_FindMassBy1(TableColumns1 types.TableColumns) string {
	Otvet := "\n\trpc "

	TextFields := ""
	TextRequest := "Request_"
	Underline := ""
	for _, Column1 := range TableColumns1.Columns {
		TextFields = TextFields + Underline + Column1.NameGo
		TextRequest1 := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
		TextRequest = TextRequest + Underline + TextRequest1
		Underline = "_"
	}

	FuncName := "FindMassBy_" + TextFields
	//функция ReadAll()
	if len(TableColumns1.Columns) == 0 {
		FuncName = constants.TEXT_READALL
		TextRequest = "Request_Empty"
	}

	//
	Otvet = Otvet + TableColumns1.Table.NameGo + "_" + FuncName + "(" + TextRequest + ") returns (ResponseMass) {}\n"

	return Otvet
}
