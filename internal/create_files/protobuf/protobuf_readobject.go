package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindText_ReadObject - добавляет текст FindText_ReadObject
func FindText_ReadObject(TextProto string, Table1 *types.Table) (TextProtoNew string, TextAdd string) {
	TextProtoNew = TextProto

	TextAdd = ""
	Text1 := FindText_ReadObject1(Table1)

	//проверим такой текст уже есть
	pos1 := strings.Index(TextProto, Text1)
	if pos1 >= 0 {
		return
	}

	//
	TextAdd = TextAdd + Text1

	return TextProtoNew, TextAdd
}

// FindText_ReadObject1 - находит текст FindBy
func FindText_ReadObject1(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo_translit
	//PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	//if PrimaryKeyColumn == nil {
	//	return Otvet
	//}

	TextRequest, _ := create_files.FindText_ProtobufRequest(Table1)
	Otvet = "\trpc " + ModelName + "_ReadObject(" + TextRequest + ") returns (Response) {}\n"

	return Otvet

	return Otvet
}
