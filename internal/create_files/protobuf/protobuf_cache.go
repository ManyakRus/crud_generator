package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindTextProtoTable1_Cache - возвращает текст функции ReadFromCache() .proto для таблицы
func FindTextProtoTable1_Cache(TextProto string, Table1 *types.Table) string {
	Otvet := "" //"\n\t//\n"

	Otvet = Otvet + FindTextReadFromCache(TextProto, Table1)

	return Otvet
}

// FindTextReadFromCache - возвращает текст .proto
func FindTextReadFromCache(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := TextReadFromCache(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// TextReadFromCache - возвращает текст .proto
func TextReadFromCache(Table1 *types.Table) string {
	Otvet := ""

	PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequestId, _ := create_files.FindText_ProtobufRequest(Table1)
	ModelName := Table1.NameGo
	Otvet = "rpc " + ModelName + "_ReadFromCache(" + TextRequestId + ") returns (Response) {}"

	return Otvet
}
