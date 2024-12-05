package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindText_ProtoTable1 - возвращает текст всех функций .proto для таблицы
func FindText_ProtoTable1(TextProto string, Table1 *types.Table) string {
	Otvet := "" //"\n\t//\n"

	ModelName := Table1.NameGo_translit
	//Otvet = Otvet + AddTextMessageRequestID(TextProto, Table1)

	Otvet = Otvet + FindText_Read(TextProto, Table1)
	Otvet = Otvet + FindText_Create(TextProto, ModelName)
	Otvet = Otvet + FindText_Update(TextProto, ModelName)
	Otvet = Otvet + FindText_Save(TextProto, ModelName)
	if create_files.Has_Column_ExtID_ConnectionID(Table1) == true {
		Otvet = Otvet + FindText_FindByExtId(TextProto, ModelName)
	}

	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		Otvet = Otvet + FindText_Delete(TextProto, Table1)

		if config.Settings.HAS_IS_DELETED == true {
			Otvet = Otvet + FindText_Restore(TextProto, Table1)
		}
	}

	return Otvet
}

// FindText_Read - возвращает текст .proto
func FindText_Read(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := Text_Read(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindText_Create - возвращает текст .proto
func FindText_Create(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := Text_Create(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindText_Update - возвращает текст .proto
func FindText_Update(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := Text_Update(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindText_Save - возвращает текст .proto
func FindText_Save(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := Text_Save(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindText_Delete - возвращает текст .proto
func FindText_Delete(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := Text_Delete(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindText_Restore - возвращает текст .proto
func FindText_Restore(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := Text_Restore(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindText_FindByExtId - возвращает текст .proto
func FindText_FindByExtId(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := Text_FindByExtId(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// Text_Read - возвращает текст .proto
func Text_Read(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo_translit
	PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, _ := create_files.FindText_ProtobufRequest(Table1)
	Otvet = "rpc " + ModelName + "_Read(" + TextRequest + ") returns (Response) {}"

	return Otvet
}

// Text_Create - возвращает текст .proto
func Text_Create(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Create(RequestModel) returns (Response) {}"

	return Otvet
}

// Text_Update - возвращает текст .proto
func Text_Update(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Update(RequestModel) returns (Response) {}"

	return Otvet
}

// Text_Save - возвращает текст .proto
func Text_Save(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Save(RequestModel) returns (Response) {}"

	return Otvet
}

// Text_Delete - возвращает текст .proto
func Text_Delete(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo_translit
	PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, _ := create_files.FindText_ProtobufRequest(Table1)
	Otvet = "rpc " + ModelName + "_Delete(" + TextRequest + ") returns (Response) {}"

	return Otvet
}

// Text_Restore - возвращает текст .proto
func Text_Restore(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo_translit
	PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, _ := create_files.FindText_ProtobufRequest(Table1)
	Otvet = "rpc " + ModelName + "_Restore(" + TextRequest + ") returns (Response) {}"

	return Otvet
}

// Text_FindByExtId - возвращает текст .proto
func Text_FindByExtId(ModelName string) string {
	Otvet := "rpc " + ModelName + "_FindByExtID(RequestExtID) returns (Response) {}"

	return Otvet
}
