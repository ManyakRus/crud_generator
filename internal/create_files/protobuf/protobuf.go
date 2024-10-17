package protobuf

import (
	"github.com/ManyakRus/crud_generator/internal/config"
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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		TextProtoNew = TextProtoNew + FindText_ProtoTable1(TextProto, Table1)
		TextProtoNew = TextProtoNew + FindText_ProtoTable1_UpdateManyFields(TextProto, Table1)
		TextProtoNew = TextProtoNew + FindText_ProtoTable1_UpdateEveryColumn(TextProto, Table1)

		//добавим текст FindBy
		TextProtoNew1 := ""
		TextProto, TextProtoNew1 = FindText_FindBy(TextProto, Table1)
		TextProtoNew = TextProtoNew + TextProtoNew1

		//добавим текст FindMassBy
		TextProto, TextProtoNew1 = FindText_FindMassBy(TextProto, Table1)
		TextProtoNew = TextProtoNew + TextProtoNew1

		//добавим текст ReadAll
		TextProto, TextProtoNew1 = FindText_ReadAll(TextProto, Table1)
		TextProtoNew = TextProtoNew + TextProtoNew1

		//добавим текст FindModelBy
		TextProto, TextProtoNew1 = FindText_FindModelBy(MapAll, TextProto, Table1)
		TextProtoNew = TextProtoNew + TextProtoNew1

		//
		if config.Settings.NEED_CREATE_CACHE_API == true {
			TextProtoNew = TextProtoNew + FindText_ProtoTable1_Cache(TextProto, Table1)
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
	TextProto = create_files.Delete_EmptyLines(TextProto)

	//запись файла
	err = os.WriteFile(FilenameReadyProto, []byte(TextProto), config.Settings.FILE_PERMISSIONS)

	return err
}

//func FillTableProto1(TextProto string, Table1 *types.Table) string {
//	Otvet := TextProto
//
//	//ModelName := Table1.NameGo
//
//
//	//создание текста
//	TextProtoTable := FindText_ProtoTable1(Table1)
//
//	return Otvet
//}

// AddTextMessageRequestID1 - в текст .proto добавляет message с RequestID
func AddTextMessageRequestID1(Text string, Table1 *types.Table) string {
	Otvet := Text

	//найдём текст RequestID
	PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	TextRequest, TextFieldName := create_files.FindText_ProtobufRequest(Table1)

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
    uint32 VersionModel = 1; //версия структуры модели
    ` + ProtobufTypePK + ` ` + TextFieldName + ` = 2; // id записи в БД
}
`

	Otvet = Otvet + TextMessage

	return Otvet
}

// AddTextMessageRequestID_ManyPK - в текст .proto добавляет message с RequestID
func AddTextMessageRequestID_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	//найдём текст RequestID
	PrimaryKeyColumns := create_files.Find_PrimaryKeyColumns(Table1)
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
	PrimaryKeyColumn := create_files.Find_PrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}
	//
	_, FieldNamePK := create_files.FindText_ProtobufRequest(Table1)

	TextRequest, FieldName, _, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "_")

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
    uint32 VersionModel = 1; //версия структуры модели
    ` + ProtobufTypePK + ` ` + FieldNamePK + ` = 2; // id записи в БД
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
	PrimaryKeyColumns := create_files.Find_PrimaryKeyColumns(Table1)
	if len(PrimaryKeyColumns) == 0 {
		return Otvet
	}
	//

	TextRequest := create_files.FindText_ProtobufRequest_Column_ManyPK(Table1, Column1)

	//найдём уже есть message
	TextFind := "message " + TextRequest + " {"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 >= 0 {
		return Otvet
	}

	TextMessage := `
// ` + TextRequest + ` - параметры запроса на сервер
message ` + TextRequest + ` {
    uint32 VersionModel = 1; //версия структуры модели`

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
		_, FieldNamePK, _, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, ColumnPK1, "")

		//добавим message
		TextMessage = TextMessage + `
    ` + ProtobufTypePK + ` ` + FieldNamePK + ` = ` + sNumber + `; //id записи в БД`
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
		_, FieldName, _, _ := create_files.FindText_ProtobufRequest_ID_Type(Table1, Column1, "")
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

// AddTextMessageRequestID_Columns - в текст .proto добавляет message из присланных колонок
func AddTextMessageRequestID_Columns(Text string, Columns []*types.Column) string {
	Otvet := Text

	TextRequest := "Request_" + create_files.Find_RequestFieldNames_FromMass(Columns)

	//найдём уже есть message
	TextFind := "message " + TextRequest + " {"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 >= 0 {
		return Otvet
	}

	TextMessage := `
// ` + TextRequest + ` - параметры запроса на сервер
message ` + TextRequest + ` {
    uint32 VersionModel = 1; //версия структуры модели`

	//
	for i, Column1 := range Columns {
		ProtoType := create_files.Convert_GolangTypeNameToProtobufTypeName(Column1.TypeGo)
		ProtoName := create_files.Find_RequestFieldName_FromMass(Column1, Columns)
		TextMessage = TextMessage + `
    ` + ProtoType + ` ` + ProtoName + ` = ` + strconv.Itoa(i+2) + `; //значение поиска`
	}

	TextMessage = TextMessage + `
}`
	Otvet = TextMessage

	return Otvet
}

// AddTextMessageRequestModel_Column - в текст .proto добавляет message из присланных колонок
func AddTextMessageRequestModel_Column(Text string, Column1 *types.Column) string {
	Otvet := Text

	ProtoName := create_files.Convert_GolangTypeNameToProtobufFieldName(Column1.TypeGo)
	TextRequest := "Request_Model_" + ProtoName

	//найдём уже есть message
	TextFind := "message " + TextRequest + " {"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 >= 0 {
		return Otvet
	}

	TextMessage := `
// ` + TextRequest + ` - параметры запроса на сервер
message ` + TextRequest + ` {
    uint32 VersionModel = 1; //версия структуры модели
    string ModelString = 2; //объект-модель в формате json
`

	//
	//for i, Column1 := range Columns {
	ProtoType := create_files.Convert_GolangTypeNameToProtobufTypeName(Column1.TypeGo)
	//ProtoName := create_files.Find_RequestFieldName(Table1, Column1)
	TextMessage = TextMessage + "\t" + ProtoType + ` ` + ProtoName + ` = 3; //значение поиска`
	//}

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
			log.Panic("FindText_ProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}

		Otvet = AddTextMessageRequestID_ColumnType_ManyPK(Otvet, Table1, Column1)
	}

	return Otvet
}
