package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindText_ReadAll - добавляет текст ReadAll
func FindText_ReadAll(TextProto string, Table1 *types.Table) (TextProtoNew string, TextAdd string) {
	TextProtoNew = TextProto

	TextAdd = ""
	for Table1, _ := range types.MapReadAll {

		Text1 := FindText_ReadAll1(Table1)

		//проверим такой текст уже есть
		pos1 := strings.Index(TextProto, Text1)
		if pos1 >= 0 {
			continue
		}

		//
		TextAdd = TextAdd + Text1
	}

	return TextProtoNew, TextAdd
}

// FindText_ReadAll1 - находит текст FindBy
func FindText_ReadAll1(Table1 *types.Table) string {
	Otvet := "\n\trpc "

	FuncName := "ReadAll"
	TextRequest := "Request_Empty"

	//
	Otvet = Otvet + Table1.NameGo + "_" + FuncName + "(" + TextRequest + ") returns (ResponseMass) {}\n"

	return Otvet
}
