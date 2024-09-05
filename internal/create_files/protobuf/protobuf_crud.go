package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"strings"
)

// FindTextProtoTable1 - возвращает текст всех функций .proto для таблицы
func FindTextProtoTable1(TextProto string, Table1 *types.Table) string {
	Otvet := "\n" //"\n\t//\n"

	ModelName := Table1.NameGo
	//Otvet = Otvet + AddTextMessageRequestID(TextProto, Table1)

	Otvet = Otvet + FindTextRead(TextProto, Table1)
	Otvet = Otvet + FindTextCreate(TextProto, ModelName)
	Otvet = Otvet + FindTextUpdate(TextProto, ModelName)
	Otvet = Otvet + FindTextSave(TextProto, ModelName)
	if create_files.Has_Column_ExtID_ConnectionID_Int64(Table1) == true {
		Otvet = Otvet + FindTextFindByExtId(TextProto, ModelName)
	}

	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		Otvet = Otvet + FindTextDelete(TextProto, Table1)

		if config.Settings.HAS_IS_DELETED == true {
			Otvet = Otvet + FindTextRestore(TextProto, Table1)
		}
	}

	return Otvet
}

// FindTextRead - возвращает текст .proto
func FindTextRead(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := TextRead(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextCreate - возвращает текст .proto
func FindTextCreate(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextCreate(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextUpdate - возвращает текст .proto
func FindTextUpdate(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextUpdate(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextSave - возвращает текст .proto
func FindTextSave(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextSave(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextDelete - возвращает текст .proto
func FindTextDelete(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := TextDelete(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextRestore - возвращает текст .proto
func FindTextRestore(TextProto string, Table1 *types.Table) string {
	Otvet := ""
	Otvet2 := TextRestore(Table1)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextFindByExtId - возвращает текст .proto
func FindTextFindByExtId(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextFindByExtId(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// TextRead - возвращает текст .proto
func TextRead(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo
	PrimaryKeyColumn := create_files.FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, _ := create_files.FindTextProtobufRequest(Table1)
	Otvet = "rpc " + ModelName + "_Read(" + TextRequest + ") returns (Response) {}"

	return Otvet
}

// TextCreate - возвращает текст .proto
func TextCreate(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Create(RequestModel) returns (Response) {}"

	return Otvet
}

// TextUpdate - возвращает текст .proto
func TextUpdate(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Update(RequestModel) returns (Response) {}"

	return Otvet
}

// TextSave - возвращает текст .proto
func TextSave(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Save(RequestModel) returns (Response) {}"

	return Otvet
}

// TextDelete - возвращает текст .proto
func TextDelete(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo
	PrimaryKeyColumn := create_files.FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, _ := create_files.FindTextProtobufRequest(Table1)
	Otvet = "rpc " + ModelName + "_Delete(" + TextRequest + ") returns (Response) {}"

	return Otvet
}

// TextRestore - возвращает текст .proto
func TextRestore(Table1 *types.Table) string {
	Otvet := ""

	ModelName := Table1.NameGo
	PrimaryKeyColumn := create_files.FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, _ := create_files.FindTextProtobufRequest(Table1)
	Otvet = "rpc " + ModelName + "_Restore(" + TextRequest + ") returns (Response) {}"

	return Otvet
}

// TextFindByExtId - возвращает текст .proto
func TextFindByExtId(ModelName string) string {
	Otvet := "rpc " + ModelName + "_FindByExtID(RequestExtID) returns (Response) {}"

	return Otvet
}
