package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"sort"
	"strings"
)

// FindTextProtoTable1_UpdateEveryColumn - возвращает текст всех функций .proto для таблицы, обновления каждого поля таблицы
func FindTextProtoTable1_UpdateEveryColumn(TextProto string, Table1 *types.Table) string {
	Otvet := "" //"\n\t//\n"

	//ModelName := Table1.NameGo

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой колонки
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		Otvet1 := FindTextUpdateEveryColumn(TextProto, Table1, Column1)
		Otvet = Otvet + Otvet1
	}

	return Otvet
}

// FindTextUpdateEveryColumn - возвращает текст .proto функции Update_ColumnName()
func FindTextUpdateEveryColumn(TextProto string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := ""
	Otvet2 := TextUpdateEveryColumn(Table1, Column1)

	//добавим текст message RequestId_Float64 {
	//Otvet = Otvet + AddTextMessageRequestID_ColumnType(TextProto, Table1, Column1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// TextUpdateEveryColumn - возвращает текст .proto функции Update_ColumnName()
func TextUpdateEveryColumn(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := ""

	ModelName := Table1.NameGo

	TextRequest := ""
	//TypeGo := Column1.TypeGo

	TextRequest = create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)
	//TextRequest, _, _, _ = create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "")
	ColumnName := Column1.NameGo

	Otvet = "rpc " + ModelName + "_Update_" + ColumnName + "(" + TextRequest + ") returns (ResponseEmpty) {}"
	//Otvet = Otvet + "\n"

	return Otvet
}

// FindTextProtoTable1_UpdateManyFields - возвращает текст функции UpdateManyFields() .proto для таблицы
func FindTextProtoTable1_UpdateManyFields(TextProto string, Table1 *types.Table) string {
	Otvet := "" //"\n\t//\n"

	ModelName := Table1.NameGo
	Otvet = Otvet + FindTextUpdateManyFields(TextProto, ModelName)

	return Otvet
}

// FindTextUpdateManyFields - возвращает текст .proto
func FindTextUpdateManyFields(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextUpdateManyFields(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// TextUpdateManyFields - возвращает текст .proto
func TextUpdateManyFields(ModelName string) string {
	Otvet := "rpc " + ModelName + "_UpdateManyFields(Request_Model_MassString) returns (ResponseEmpty) {}"

	return Otvet
}
