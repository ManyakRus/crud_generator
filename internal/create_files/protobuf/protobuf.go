package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strconv"
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
		err1 := create_files.IsGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		TextProtoNew = TextProtoNew + FindTextProtoTable1(TextProto, Table1)
		TextProtoNew = TextProtoNew + FindTextProtoTable1_UpdateManyFields(TextProto, Table1)
		TextProtoNew = TextProtoNew + FindTextProtoTable1_UpdateEveryColumn(TextProto, Table1)

		if config.Settings.NEED_CREATE_CACHE_API == true {
			TextProtoNew = TextProtoNew + FindTextProtoTable1_Cache(TextProto, Table1)
		}
		TextProto = AddTextMessageRequestID(TextProto, Table1)
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
		//TextFind := "rpc " + Table1.NameGo + "_"
		//pos1 := FindLastGoodPos(TextProto, TextFind)
		//if pos1 > 0 {
		//	Otvet = Otvet[:pos1] + Otvet1 + Otvet[pos1:]
		//} else {
		//	Otvet = Otvet + Otvet1
		//}
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

	TextRequest = create_files.FindTextProtobufRequest_Column_ManyPK(Table1, Column1)
	//TextRequest, _, _, _ = create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "")
	ColumnName := Column1.NameGo

	Otvet = "rpc " + ModelName + "_Update_" + ColumnName + "(" + TextRequest + ") returns (ResponseEmpty) {}"
	//Otvet = Otvet + "\n"

	return Otvet
}

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

	PrimaryKeyColumn := create_files.FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequestId, _ := create_files.FindTextProtobufRequest(Table1)
	ModelName := Table1.NameGo
	Otvet = "rpc " + ModelName + "_ReadFromCache(" + TextRequestId + ") returns (Response) {}"

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

// AddTextMessageRequestID1 - в текст .proto добавляет message с RequestID
func AddTextMessageRequestID1(Text string, Table1 *types.Table) string {
	Otvet := Text

	//найдём текст RequestID
	PrimaryKeyColumn := create_files.FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, TextFieldName := create_files.FindTextProtobufRequest(Table1)

	//найдём уже есть message
	TextFind := "message " + TextRequest + " {"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 >= 0 {
		return Otvet
	}

	//найдём ProtobufTypePK
	MappingPK, ok := dbmeta.GetMappings()[PrimaryKeyColumn.Type]
	if ok == false {
		log.Error("Неизвестный тип столбца " + PrimaryKeyColumn.Type)
		return Otvet
	}
	ProtobufTypePK := MappingPK.ProtobufType

	//добавим message
	TextMessage := `
// ` + TextRequest + ` - параметры запроса на сервер
message ` + TextRequest + ` {
    uint32 VersionModel= 1; //версия структуры модели
    ` + ProtobufTypePK + ` ` + TextFieldName + `   = 2; // id записи в БД
}
`

	Otvet = Otvet + TextMessage

	return Otvet
}

// AddTextMessageRequestID_ManyPK - в текст .proto добавляет message с RequestID
func AddTextMessageRequestID_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	//найдём текст RequestID
	PrimaryKeyColumns := create_files.FindPrimaryKeyColumns(Table1)
	if len(PrimaryKeyColumns) == 0 {
		return Otvet
	}

	Otvet = AddTextMessageRequestID_ColumnType_ManyPK(Otvet, Table1, PrimaryKeyColumns[0])

	return Otvet
}

// AddTextMessageRequestID_ColumnType - в текст .proto добавляет message с RequestID_Int64
func AddTextMessageRequestID_ColumnType(Text string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := Text

	if Table1.PrimaryKeyColumnsCount == 1 {
		Otvet = AddTextMessageRequestID_ColumnType1(Otvet, Table1, Column1)
	} else {
		Otvet = AddTextMessageRequestID_ColumnType_ManyPK(Otvet, Table1, Column1)
	}

	return Otvet
}

// AddTextMessageRequestID_ColumnType1 - в текст .proto добавляет message с RequestID_Int64, 1PK
func AddTextMessageRequestID_ColumnType1(Text string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := Text

	//найдём текст RequestID
	PrimaryKeyColumn := create_files.FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}
	//
	_, FieldNamePK := create_files.FindTextProtobufRequest(Table1)

	TextRequest, FieldName, _, _ := create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "_")

	//найдём уже есть message
	TextFind := "message " + TextRequest + " {"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 >= 0 {
		return Otvet
	}

	//найдём ProtobufTypePK
	MappingPK, ok := dbmeta.GetMappings()[PrimaryKeyColumn.Type]
	if ok == false {
		log.Error("Неизвестный тип столбца " + PrimaryKeyColumn.Type)
		return Otvet
	}
	ProtobufTypePK := MappingPK.ProtobufType

	//найдём ProtobufTypeColumn
	Mapping1, ok := dbmeta.GetMappings()[Column1.Type]
	if ok == false {
		log.Error("Неизвестный тип столбца " + Column1.Type)
		return Otvet
	}
	ProtobufTypeColumn := Mapping1.ProtobufType

	//добавим message
	TextMessage := `
// ` + TextRequest + ` - параметры запроса на сервер
message ` + TextRequest + ` {
    uint32 VersionModel= 1; //версия структуры модели
    ` + ProtobufTypePK + ` ` + FieldNamePK + `   = 2; // id записи в БД
    ` + ProtobufTypeColumn + ` ` + FieldName + ` = 3; // значение поиска
}
`

	Otvet = Otvet + TextMessage

	return Otvet
}

// AddTextMessageRequestID_ColumnType_ManyPK - в текст .proto добавляет message с RequestID_Int64, много PK
func AddTextMessageRequestID_ColumnType_ManyPK(Text string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := Text

	//найдём текст RequestID
	PrimaryKeyColumns := create_files.FindPrimaryKeyColumns(Table1)
	if len(PrimaryKeyColumns) == 0 {
		return Otvet
	}
	//

	TextRequest := create_files.FindTextProtobufRequest_Column_ManyPK(Table1, Column1)

	//найдём уже есть message
	TextFind := "message " + TextRequest + " {"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 >= 0 {
		return Otvet
	}

	TextMessage := `
// ` + TextRequest + ` - параметры запроса на сервер
message ` + TextRequest + ` {
    uint32 VersionModel= 1; //версия структуры модели`

	//заполним строки про PrimaryKey
	isPrimaryKey := false
	Number := 1
	for _, ColumnPK1 := range PrimaryKeyColumns {
		Number = Number + 1
		sNumber := strconv.Itoa(Number)

		if Column1 == ColumnPK1 {
			isPrimaryKey = true
		}

		//найдём ProtobufTypePK
		MappingPK, ok := dbmeta.GetMappings()[ColumnPK1.Type]
		if ok == false {
			log.Error("Неизвестный тип столбца " + ColumnPK1.Type)
			return Otvet
		}
		ProtobufTypePK := MappingPK.ProtobufType
		_, FieldNamePK, _, _ := create_files.FindTextProtobufRequest_ID_Type(Table1, ColumnPK1, "")

		//добавим message
		TextMessage = TextMessage + `
    ` + ProtobufTypePK + ` ` + FieldNamePK + `   = ` + sNumber + `; //id записи в БД`
	}

	//заполним строку про Column1
	if isPrimaryKey == false {

		//найдём ProtobufTypeColumn
		Mapping1, ok := dbmeta.GetMappings()[Column1.Type]
		if ok == false {
			log.Error("Неизвестный тип столбца " + Column1.Type)
			return Otvet
		}
		ProtobufTypeColumn := Mapping1.ProtobufType
		_, FieldName, _, _ := create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "")
		Number = Number + 1
		sNumber := strconv.Itoa(Number)

		TextMessage = TextMessage + `
    ` + ProtobufTypeColumn + ` ` + FieldName + ` = ` + sNumber + `; //значение поиска`
	}

	TextMessage = TextMessage + `
}`
	Otvet = Otvet + "\n" + TextMessage

	return Otvet
}

// AddTextMessageRequestID - возвращает текст в .proto для таблицы
func AddTextMessageRequestID(TextProto string, Table1 *types.Table) string {
	Otvet := TextProto //"\n\t//\n"

	Otvet = AddTextMessageRequestID_ManyPK(Otvet, Table1)

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//для каждой колонки добавим добавим message RequestId_Int64
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}

		Otvet = AddTextMessageRequestID_ColumnType_ManyPK(Otvet, Table1, Column1)
	}

	return Otvet
}
