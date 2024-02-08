package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	err = CreateFileProto(MapAll)
	if err != nil {
		log.Error("CreateFileProto() error: ", err)
		return err
	}

	return err
}

// CreateFileProto - создаёт 1 файл в папке grpc
func CreateFileProto(MapAll map[string]*types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesProto := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile()
	DirReadyProto := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile()
	FilenameReadyProto := DirReadyProto + config.Settings.SERVICE_NAME + ".proto"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyProto)

	FilenameTemplateProto := DirTemplatesProto + "service.proto_"
	if config.Settings.TEMPLATE_EXTERNAL_PROTO_FILENAME != "" {
		FilenameTemplateProto = config.Settings.TEMPLATE_EXTERNAL_PROTO_FILENAME
	}
	bytes, err := os.ReadFile(FilenameTemplateProto)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateProto, " error: ", err)
	}
	TextProto := string(bytes)

	//заменим название сервиса
	ServiceName := config.Settings.SERVICE_NAME
	ServiceNameProto := micro.StringFromUpperCase(ServiceName)
	TEMPLATE_SERVICE_NAME := config.Settings.TEMPLATE_SERVICE_NAME
	TextProto = strings.ReplaceAll(TextProto, TEMPLATE_SERVICE_NAME, ServiceNameProto)
	//заменим ещё раз с большой буквы
	TEMPLATE_SERVICE_NAME = micro.StringFromUpperCase(TEMPLATE_SERVICE_NAME)
	TextProto = strings.ReplaceAll(TextProto, TEMPLATE_SERVICE_NAME, ServiceNameProto)

	//сортировка по названию таблиц
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	TextProtoNew := ""
	for _, key1 := range keys {
		Table1, ok := MapAll[key1]
		if ok == false {
			log.Panic("MapAll[key1] not found")
		}

		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		TextProtoNew = TextProtoNew + FindTextProtoTable1(TextProto, Table1)
	}

	//найдём куда вставить текст
	sFind := "\nservice "
	pos1 := strings.Index(TextProto, sFind)
	if pos1 < 0 {
		log.Panic("Not found text ", sFind)
	}

	s2 := TextProto[pos1+1:]
	sFind = "\n"
	posEnd := strings.Index(s2, sFind)
	if posEnd < 0 {
		log.Panic("Not found text ", sFind)
	}
	PosStart := pos1 + posEnd + 1

	//
	TextProto = TextProto[:PosStart] + TextProtoNew + TextProto[PosStart:]

	//
	TextProto = create_files.DeleteEmptyLines(TextProto)

	//запись файла
	err = os.WriteFile(FilenameReadyProto, []byte(TextProto), constants.FILE_PERMISSIONS)

	return err
}

//func FillTableProto1(TextProto string, Table1 *types.Table) string {
//	Otvet := TextProto
//
//	//ModelName := Table1.NameGo
//
//
//	//создание текста
//	TextProtoTable := FindTextProtoTable1(Table1)
//
//	return Otvet
//}

// FindTextProtoTable1 - возвращает текст всех функций .proto для таблицы
func FindTextProtoTable1(TextProto string, Table1 *types.Table) string {
	Otvet := "\n" //"\n\t//\n"

	ModelName := Table1.NameGo
	Otvet = Otvet + FindTextRead(TextProto, ModelName)
	Otvet = Otvet + FindTextCreate(TextProto, ModelName)
	Otvet = Otvet + FindTextUpdate(TextProto, ModelName)
	Otvet = Otvet + FindTextSave(TextProto, ModelName)
	if create_files.Has_Column_ExtID_ConnectionID(Table1) == true {
		Otvet = Otvet + FindTextFindByExtId(TextProto, ModelName)
	}

	if create_files.Has_Column_IsDeleted(Table1) == true {
		Otvet = Otvet + FindTextDelete(TextProto, ModelName)

		if config.Settings.HAS_IS_DELETED == true {
			Otvet = Otvet + FindTextRestore(TextProto, ModelName)
		}
	}

	return Otvet
}

// FindTextRead - возвращает текст .proto
func FindTextRead(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextRead(ModelName)

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
func FindTextDelete(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextDelete(ModelName)

	//проверка такой текст уже есть
	pos1 := strings.Index(TextProto, Otvet2)
	if pos1 >= 0 {
		return Otvet
	}

	Otvet = "\t" + Otvet2 + "\n"

	return Otvet
}

// FindTextRestore - возвращает текст .proto
func FindTextRestore(TextProto string, ModelName string) string {
	Otvet := ""
	Otvet2 := TextRestore(ModelName)

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
func TextRead(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Read(RequestId) returns (Response) {}"

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
func TextDelete(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Delete(RequestId) returns (Response) {}"

	return Otvet
}

// TextRestore - возвращает текст .proto
func TextRestore(ModelName string) string {
	Otvet := "rpc " + ModelName + "_Restore(RequestId) returns (Response) {}"

	return Otvet
}

// TextFindByExtId - возвращает текст .proto
func TextFindByExtId(ModelName string) string {
	Otvet := "rpc " + ModelName + "_FindByExtID(RequestExtID) returns (Response) {}"

	return Otvet
}
