package create_files

import (
	"errors"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/mini_func"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"strconv"
	"strings"
)

// FindSingularName - возвращает наименование в единственном числе
func FindSingularName(s string) string {
	var Otvet string

	if s == "" {
		return Otvet
	}

	Otvet = inflection.Singular(s)

	return Otvet
}

// FormatName - возвращает наименование в формате CamelCase
func FormatName(Name string) string {
	Otvet := Name

	Otvet1, ok := types.MapReplaceName[Name]
	if ok == true {
		return Otvet1
	}

	switch strings.ToLower(Name) {
	case "id":
		Otvet = "ID"
		return Otvet
	}

	Otvet = strcase.ToCamel(Otvet)

	//_id в конце заменяем на ID
	lenName := len(Name)
	lenOtvet := len(Otvet)
	if lenName >= 3 {
		last3 := strings.ToLower(Name[lenName-3:])
		if last3 == "_id" {
			Otvet = Otvet[:lenOtvet-2] + "ID"
		}
	}

	return Otvet
}

// DeleteFuncFromFuncName - удаляет функцию из текста начиная с объявления функции
func DeleteFuncFromFuncName(Text, FuncName string) string {
	Otvet := Text

	TextFind := "\nfunc " + FuncName + "("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncFromComment - удаляет функцию из текста начиная с комментария
func DeleteFuncFromComment(Text, Comment string) string {
	Otvet := Text

	TextFind := Comment
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet

}

// DeleteCommentFromFuncName - удаляет комментарий с названием функции
func DeleteCommentFromFuncName(Text, FuncName string) string {
	Otvet := Text

	TextFind := "//" + FuncName + "()"
	Otvet = strings.ReplaceAll(Otvet, TextFind, "")

	return Otvet
}

// DeleteLineWithComment - удаляет текст от комментария до конца строки
func DeleteLineWithComment(Text, FuncName string) string {
	Otvet := Text

	TextFind := "//" + FuncName + "()"
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1:]

	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+1:]

	return Otvet
}

// FindPrimaryKeyNameTypeGo - возвращает наименование и тип golang колонки PrimaryKey
func FindPrimaryKeyNameTypeGo(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

	for _, Column1 := range Table1.MapColumns {
		if Column1.IsPrimaryKey == true {
			return Column1.NameGo, Column1.TypeGo
		}
	}

	return Otvet, Type
}

// FindPrimaryKeyNameType - возвращает наименование и тип БД колонки PrimaryKey
func FindPrimaryKeyNameType(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

	for _, Column1 := range Table1.MapColumns {
		if Column1.IsPrimaryKey == true {
			return Column1.Name, Column1.Type
		}
	}

	return Otvet, Type
}

// FindPrimaryKeyColumn - возвращает Column для колонки PrimaryKey
func FindPrimaryKeyColumn(Table1 *types.Table) (Column1 *types.Column) {
	var Otvet *types.Column

	for _, Column1 := range Table1.MapColumns {
		if Column1.IsPrimaryKey == true {
			return Column1
		}
	}

	return Otvet
}

//// FindPrimaryKeyNameTypeGo - возвращает наименование и тип golang колонки PrimaryKey
//func FindPrimaryKeyNameTypeGo(Table1 *types.Table) (string, string) {
//	Otvet := ""
//	Type := ""
//
//	for _, Column1 := range Table1.MapColumns {
//		if Column1.IsIdentity == true {
//			return Column1.NameGo, Column1.TypeGo
//		}
//	}
//
//	return Otvet, Type
//}

// ReplacePrimaryKeyOtvetID - заменяет "Otvet.ID" на название колонки PrimaryKey
func ReplacePrimaryKeyOtvetID(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnNamePK, ColumnTypeGoPK := FindPrimaryKeyNameTypeGo(Table1)

	//заменим ID-Alias на ID
	TableName := Table1.Name
	IDName, _ := FindPrimaryKeyNameType(Table1)
	Alias, ok := types.MapConvertID[TableName+"."+IDName]
	OtvetColumnName := "Otvet." + ColumnNamePK
	if ok == true {
		OtvetColumnName = Alias + "(" + OtvetColumnName + ")"
	}

	//заменим int64(Otvet.ID) на Otvet.ID
	if mini_func.IsNumberType(ColumnTypeGoPK) == true {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "string" {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) == 0", OtvetColumnName+" == \"\"")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) != 0", OtvetColumnName+" != \"\"")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "uuid.UUID" || ColumnTypeGoPK == "uuid.NullUUID" {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) == 0", OtvetColumnName+" == uuid.Nil")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) != 0", OtvetColumnName+" != uuid.Nil")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "time.Time" {
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) == 0", OtvetColumnName+".IsZero() == true")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID) != 0", OtvetColumnName+".IsZero() == false")
		Otvet = strings.ReplaceAll(Otvet, "int64(Otvet.ID)", OtvetColumnName)
	}
	//Otvet = strings.ReplaceAll(Otvet, "Otvet.ID = ", OtvetColumnName+" = ")
	//Otvet = strings.ReplaceAll(Otvet, "Otvet.ID != ", OtvetColumnName+" != ")
	//Otvet = strings.ReplaceAll(Otvet, " Otvet.ID)", " "+OtvetColumnName+")")
	Otvet = strings.ReplaceAll(Otvet, " Otvet.ID)", " Otvet."+ColumnNamePK+")")

	//Alias преобразуем в int64, и наоборот
	if Alias != "" {
		Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(Otvet.ID)", ColumnTypeGoPK+"(Otvet."+ColumnNamePK+")")
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(Otvet.ID)", OtvetColumnName)
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(ID)", Alias+"("+ColumnNamePK+")")
	} else {
		Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(Otvet.ID)", "Otvet."+ColumnNamePK+"")
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(Otvet.ID)", OtvetColumnName)
		Otvet = strings.ReplaceAll(Otvet, "AliasFromInt(ID)", "ID")
	}

	return Otvet
}

// ReplacePrimaryKeyM_ID - заменяет "m.ID" на название колонки PrimaryKey
func ReplacePrimaryKeyM_ID(Text string, Table1 *types.Table) string {
	Otvet := Text

	ColumnNamePK, ColumnTypeGoPK := FindPrimaryKeyNameTypeGo(Table1)

	//заменим ID-Alias на ID
	TableName := Table1.Name
	IDName, _ := FindPrimaryKeyNameType(Table1)
	Alias, ok := types.MapConvertID[TableName+"."+IDName]
	OtvetColumnName := "m." + ColumnNamePK
	if ok == true {
		OtvetColumnName = Alias + "(" + OtvetColumnName + ")"
	}

	//заменим int64(m.ID) на m.ID
	if mini_func.IsNumberType(ColumnTypeGoPK) == true {
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "string" {
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", OtvetColumnName+" == \"\"")
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", OtvetColumnName+" != \"\"")
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", OtvetColumnName)
	} else if ColumnTypeGoPK == "uuid.UUID" || ColumnTypeGoPK == "uuid.NullUUID" {
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", OtvetColumnName+" == uuid.Nil")
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", OtvetColumnName+" != uuid.Nil")
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", OtvetColumnName+".String()")
	} else if ColumnTypeGoPK == "time.Time" {
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", OtvetColumnName+".IsZero() == true")
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", OtvetColumnName+".IsZero() == false")
		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", OtvetColumnName)
	}
	//Otvet = strings.ReplaceAll(Otvet, "m.ID = ", OtvetColumnName+" = ")
	//Otvet = strings.ReplaceAll(Otvet, " = m.ID", " = "+OtvetColumnName)
	Otvet = strings.ReplaceAll(Otvet, ", m.ID,", ", "+OtvetColumnName+",")
	//Otvet = strings.ReplaceAll(Otvet, "(m.ID)", "("+OtvetColumnName+")")

	//Alias преобразуем в int64, и наоборот
	if Alias != "" {
		Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(m.ID)", ColumnTypeGoPK+"(m."+ColumnNamePK+")")
		Otvet = strings.ReplaceAll(Otvet, "IntToAlias(m.ID)", OtvetColumnName)
		Otvet = strings.ReplaceAll(Otvet, "IntToAlias(ID)", Alias+"("+ColumnNamePK+")")
	} else {
		Otvet = strings.ReplaceAll(Otvet, "IntFromAlias(m.ID)", "m."+ColumnNamePK+"")
		Otvet = strings.ReplaceAll(Otvet, "IntToAlias(m.ID)", OtvetColumnName)
		Otvet = strings.ReplaceAll(Otvet, "IntToAlias(ID)", "ID")
	}

	return Otvet
}

// AddSkipNow - добавляет строку t.SkipNow()
func AddSkipNow(Text string, Table1 *types.Table) string {
	Otvet := Text

	if Table1.IDMinimum == "" || Table1.IDMinimum == "0" {
		TextFind := "(t *testing.T) {"
		Otvet = strings.ReplaceAll(Otvet, TextFind, TextFind+"\n\tt.SkipNow() //now rows in DB\n")
	}

	return Otvet
}

// IsGoodTable - возвращает ошибку если таблица неправильная
func IsGoodTable(Table1 *types.Table) error {
	var err error

	//TableName := Table1.Name
	//ColumnName, _ := FindPrimaryKeyNameTypeGo(Table1)
	//if ColumnName == "" {
	//	TextError := fmt.Sprint("Wrong table: ", Table1.Name, " error: not found Primary key")
	//	err = errors.New(TextError)
	//}

	err = IsGoodTableNamePrefix(Table1)
	if err != nil {
		return err
	}

	err = IsGoodPrimaryKeyColumnsCount(Table1)
	if err != nil {
		return err
	}

	return err
}

// IsGoodPrimaryKeyColumnsCount - возвращает ошибку если количество колонок PrimaryKey неправильное
func IsGoodPrimaryKeyColumnsCount(Table1 *types.Table) error {
	var err error

	if Table1.PrimaryKeyColumnsCount <= 0 || Table1.PrimaryKeyColumnsCount >= 2 {
		TextError := fmt.Sprint("Wrong table: ", Table1.Name, " error: can not use many Primary key columns count: ", Table1.PrimaryKeyColumnsCount)
		err = errors.New(TextError)
	}

	return err
}

// IsGoodTableNamePrefix - возвращает ошибку если префикс таблицы = "DELETED_"
func IsGoodTableNamePrefix(Table1 *types.Table) error {
	var err error

	TableName := Table1.Name
	if strings.HasPrefix(TableName, "DELETED_") == true {
		TextError := fmt.Sprint("Wrong table: ", Table1.Name, " error: name = DELETED_")
		err = errors.New(TextError)
	}

	return err
}

// PrintableString - возвращает строку без запрещённых символов
func PrintableString(s string) string {
	Otvet := s
	Otvet = strconv.Quote(Otvet) //экранирование символов
	len1 := len(Otvet)
	if len1 > 0 {
		Otvet = Otvet[1 : len1-1]
	}

	return Otvet
}

// Find_Template_DB_Foldername - возвращает путь к папке
func Find_Template_DB_Foldername() string {
	Otvet := ""

	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD + micro.SeparatorFile()

	Otvet = DirTemplatesDB
	return Otvet
}

// DeleteImportModel - удаляет лишний импорт модели
func DeleteImportModel(s string) string {
	Otvet := s

	ModelURL := FindModelURL()
	ImportName := micro.LastWord(ModelURL)

	pos1 := strings.Index(Otvet, ImportName+".")
	if pos1 < 0 {
		Otvet = strings.ReplaceAll(Otvet, `"`+ModelURL+`"\n`, "")
		Otvet = strings.ReplaceAll(Otvet, `"`+ModelURL+`"`, "")
	}

	return Otvet
}

// FindModelURL - возвращает URL репозитория с пакетом "model"
func FindModelURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_MODEL

	return Otvet
}

// FindGRPCServerURL - возвращает URL репозитория с пакетом "server_grpc"
func FindGRPCServerURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER

	return Otvet
}

// FindGRPCClientURL - возвращает URL репозитория с пакетом "client_grpc"
func FindGRPClientURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT

	return Otvet
}

// FindNRPClientURL - возвращает URL репозитория с пакетом "client_nrpc"
func FindNRPClientURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT

	return Otvet
}

// FindGRPCConstantsURL - возвращает URL репозитория с пакетом "client_grpc"
func FindGRPCConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC + "/" + "grpc_constants"

	return Otvet
}

// FindNRPCServerlURL - возвращает URL репозитория с пакетом "server_nrpc"
func FindNRPCServerlURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER

	return Otvet
}

// FindProtobufURL - возвращает URL репозитория с файлом .proto
func FindProtobufURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO

	return Otvet
}

// FindTablesURL - возвращает URL репозитория с пакетом "tables"
func FindTablesURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_TABLES

	return Otvet
}

// FindTableNameURL - возвращает URL репозитория с пакетом "tables" + TableName
func FindTableNameURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_TABLES + "/" + config.Settings.PREFIX_TABLE + TableName

	return Otvet
}

// FindProtoURL - возвращает URL репозитория с пакетом "grpc_proto"
func FindProtoURL() string {
	Otvet := ""

	TEMPLATE_FOLDERNAME_GRPC_PROTO := strings.ToLower(config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO)
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + TEMPLATE_FOLDERNAME_GRPC_PROTO + "/" + "grpc_proto"

	return Otvet
}

// FindModelTableURL - возвращает URL репозитория model для таблицы TableName
func FindModelTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_MODEL + "/" + TableName

	return Otvet
}

// FindCrudStarterTableURL - возвращает URL репозитория crud_starter для таблицы TableName
func FindCrudStarterTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + "/" + constants.STARTER_TABLES_PREFIX + TableName

	return Otvet
}

// FindCrudTableURL - возвращает URL репозитория crud для таблицы TableName
func FindCrudTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD + "/" + constants.CRUD_TABLES_FREFIX + TableName

	return Otvet
}

// FindNRPC_Client_URL - возвращает URL репозитория с пакетом "nrpc_client"
func FindNRPC_Client_URL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT

	return Otvet
}

// FindDBConstantsURL - возвращает URL репозитория с пакетом db "db_constants"
func FindDBConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_DB + "/" + "db_constants"

	return Otvet
}

// FindConstantsURL - возвращает URL репозитория с пакетом pkg "constants"
func FindConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CONSTANTS

	return Otvet
}

// FindNRPCConstantsURL - возвращает URL репозитория с пакетом db "constants"
func FindNRPCConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC + "/" + "nrpc_constants"

	return Otvet
}

// FindCrudStarterURL - возвращает URL репозитория с пакетом "crud_starter"
func FindCrudStarterURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER

	return Otvet
}

// Find_GRPC_NRPC_URL - возвращает URL репозитория с пакетом "crud_starter"
func Find_GRPC_NRPC_URL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_NRPC

	return Otvet
}

// FindCalcStructVersionURL - возвращает URL репозитория с пакетом "calc_struct_version"
func FindCalcStructVersionURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_DB + "/" + "calc_struct_version"

	return Otvet
}

// FindCrudFunctionsURL - возвращает URL репозитория с пакетом crud_functions
func FindCrudFunctionsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDER_CRUD_FUNCTIONS

	return Otvet
}

// FindTextDefaultGORMValue - возвращает значение по умолчанию для столбца Column1 для тегов в GORM
func FindTextDefaultGORMValue(Column1 *types.Column) string {
	var Otvet string

	ColumnName := Column1.Name

	//ищем в файле настроек nullable.json
	is_nullable_config, _ := types.MapNullableFileds[ColumnName]

	//
	sValue := ""
	Type_go := Column1.TypeGo
	//	if Column1.TableKey != "" && Column1.IsNullable == true && (is_nullable_config == true || has_is_nullable_config == false) {
	if Column1.IsNullable == true || (is_nullable_config == true) {
		sValue = "null"
	} else {

		switch Type_go {
		case "string":
			sValue = "\\\"\\\""
		case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
			sValue = "0"
		case "time.Time":
			sValue = "null"
		case "uuid.UUID", "uuid.NullUUID":
			sValue = "null"
		}
	}

	if sValue != "" {
		Otvet = ";default:" + sValue
	}

	return Otvet
}

// FindTextDefaultValue - возвращает golang значение по умолчанию для типа
func FindTextDefaultValue(Type_go string) string {
	var Otvet string

	switch Type_go {
	case "string":
		Otvet = `""`
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		Otvet = "0"
	case "time.Time":
		Otvet = "time.Time{}"
	case "bool":
		Otvet = "false"
	case "uuid.UUID", "uuid.NullUUID":
		Otvet = "uuid.Nil"
	}

	return Otvet
}

// FindTextDefaultValueSQL - возвращает значение по умолчанию для типа
func FindTextDefaultValueSQL(Type_go string) string {
	var Otvet string

	switch Type_go {
	case "string":
		Otvet = `''`
	case "int", "int32", "int64", "float32", "float64", "uint", "uint32", "uint64":
		Otvet = "0"
	case "time.Time":
		Otvet = "null"
	case "bool":
		Otvet = "false"
	case "uuid.UUID", "uuid.NullUUID":
		Otvet = "null"
	}

	return Otvet
}

// FindURL_Alias - возвращает URL репозитория с пакетом "alias"
func FindURL_Alias() string {
	Otvet := ""
	if config.Settings.TEMPLATE_FOLDERNAME_ALIAS == "" {
		return Otvet
	}
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/"
	Otvet = Otvet + config.Settings.TEMPLATE_FOLDERNAME_ALIAS

	return Otvet
}

// FindURL_Tables - возвращает URL репозитория с пакетом "tables"
func FindURL_Tables() string {
	Otvet := ""
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/"
	Otvet = Otvet + config.Settings.TEMPLATE_FOLDERNAME_TABLES

	return Otvet
}

// AddImport - добавляет RepositoryURL в секцию Import, если его там нет
func AddImport(Text, RepositoryURL string) string {
	Otvet := Text

	//если уже есть импорт
	pos1 := strings.Index(Otvet, `"`+RepositoryURL+`"`)
	if pos1 >= 0 {
		return Otvet
	}

	//
	TextFind := "import ("
	LenFind := len(TextFind)
	pos1 = strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		log.Error("not found word: import (")
		return Otvet
	}

	Otvet = Otvet[:pos1+LenFind] + "\n\t" + `"` + RepositoryURL + `"` + Otvet[pos1+LenFind:]

	return Otvet
}

// CheckAndAddImport - добавляет URL в секцию Import, если его там нет, если он нужен
func CheckAndAddImport(Text, URL string) string {
	Otvet := Text

	//проверим используется или нет
	ModuleName := micro.LastWord(URL)
	pos1 := strings.Index(Otvet, ModuleName+".")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImport(Text, URL)

	return Otvet
}

// AddImportTime - добавляет пакет в секцию Import, если его там нет
func AddImportTime(Text string) string {
	Otvet := Text

	//если уже есть импорт
	RepositoryURL := `time`
	Otvet = AddImport(Text, RepositoryURL)
	//pos1 := strings.Index(Otvet, RepositoryURL)
	//if pos1 >= 0 {
	//	return Otvet
	//}
	//
	////
	//TextImport := "import ("
	//pos1 = strings.Index(Otvet, TextImport)
	//if pos1 < 0 {
	//	log.Error("not found word: ", TextImport)
	//	return Text
	//}
	//
	//Otvet = Otvet[:pos1+len(TextImport)] + "\n\t" + RepositoryURL + Otvet[pos1+len(TextImport):]

	return Otvet
}

// AddImportUUID - добавляет пакет в секцию Import, если его там нет
func AddImportUUID(Text string) string {
	Otvet := Text

	//если уже есть импорт
	RepositoryURL := `github.com/google/uuid`
	Otvet = AddImport(Text, RepositoryURL)
	//pos1 := strings.Index(Otvet, RepositoryURL)
	//if pos1 >= 0 {
	//	return Otvet
	//}
	//
	////
	//TextImport := "import ("
	//pos1 = strings.Index(Otvet, TextImport)
	//if pos1 < 0 {
	//	log.Error("not found word: ", TextImport)
	//	return TextModel
	//}
	//
	//Otvet = Otvet[:pos1+len(TextImport)] + "\n\t" + RepositoryURL + Otvet[pos1+len(TextImport):]

	return Otvet
}

// AddImportGorm - добавляет пакет в секцию Import, если его там нет
func AddImportGorm(Text string) string {
	Otvet := Text

	//если уже есть импорт
	RepositoryURL := `gorm.io/gorm`
	Otvet = AddImport(Text, RepositoryURL)
	//pos1 := strings.Index(Otvet, RepositoryURL)
	//if pos1 >= 0 {
	//	return Otvet
	//}
	//
	////
	//TextImport := "import ("
	//pos1 = strings.Index(Otvet, TextImport)
	//if pos1 < 0 {
	//	log.Error("not found word: ", TextImport)
	//	return TextModel
	//}
	//
	//Otvet = Otvet[:pos1+len(TextImport)] + "\n\t" + RepositoryURL + Otvet[pos1+len(TextImport):]

	return Otvet
}

// AddImportTimestamp - добавляет покет в секцию Import, если его там нет
func AddImportTimestamp(Text string) string {
	Otvet := Text

	RepositoryURL := `google.golang.org/protobuf/types/known/timestamppb`
	Otvet = AddImport(Text, RepositoryURL)

	////если уже есть импорт
	//pos1 := strings.Index(Otvet, `"google.golang.org/protobuf/types/known/timestamppb"`)
	//if pos1 >= 0 {
	//	return Otvet
	//}
	//
	////
	//pos1 = strings.Index(Otvet, "import (")
	//if pos1 < 0 {
	//	log.Error("not found word: import (")
	//	return TextModel
	//}
	//
	//Otvet = Otvet[:pos1+8] + "\n\t" + `"google.golang.org/protobuf/types/known/timestamppb"` + Otvet[pos1+8:]

	return Otvet
}

// CheckAndAddImportAlias - добавляет покет в секцию Alias, если его там нет
func CheckAndAddImportAlias(TextModel string) string {
	Otvet := TextModel

	//если уже есть импорт
	pos1 := strings.Index(Otvet, `/alias`)
	if pos1 >= 0 {
		return Otvet
	}

	//если нету alias
	pos1 = strings.Index(Otvet, `alias.`)
	if pos1 < 0 {
		return Otvet
	}

	//
	pos1 = strings.Index(Otvet, "import (")
	if pos1 < 0 {
		log.Error("not found word: import (")
		return TextModel
	}

	AliasURL := FindURL_Alias()
	Otvet = Otvet[:pos1+8] + "\n\t" + `"` + AliasURL + `"` + Otvet[pos1+8:]

	return Otvet
}

// CheckAndAddImportTime_FromTable - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAddImportTime_FromTable(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	HasTimeColumn := Has_ColumnType_Time(Table1)
	if HasTimeColumn == false {
		return Otvet
	}

	Otvet = AddImportTime(Otvet)

	return Otvet
}

// CheckAndAddImportTime_FromText - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAddImportTime_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, " time.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImportTime(Otvet)

	return Otvet
}

// CheckAndAddImportUUID_FromText - добавляет пакет "uuid" в секцию Import, если его там нет
func CheckAndAddImportUUID_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, "uuid.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImportUUID(Otvet)

	return Otvet
}

// CheckAndAddImportGorm_FromText - добавляет пакет "gorm.io/gorm" в секцию Import, если его там нет
func CheckAndAddImportGorm_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, `"gorm.io/gorm"`)
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImportGorm(Otvet)

	return Otvet
}

// CheckAndAddImportTimestamp_FromText - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAddImportTimestamp_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, " timestamppb.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImportTimestamp(Otvet)

	return Otvet
}

// DeleteTemplateRepositoryImports - удаляет импорты репозитория шаблона
func DeleteTemplateRepositoryImports(Text string) string {
	Otvet := Text

	if config.Settings.TEMPLATE_REPOSITORY_URL == "" {
		return Otvet
	}

	//
	TextFind := "import ("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		log.Error("not found word: import (")
		return Otvet
	}

	TextFind = `"` + config.Settings.TEMPLATE_REPOSITORY_URL
	pos1 = strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}

	s2 := Otvet[pos1:]
	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+1:]

	Otvet = DeleteTemplateRepositoryImports(Otvet)

	return Otvet
}

// ReplaceServiceURLImports - заменяет URL репозитория шаблона на URL репозитория сервиса
func ReplaceServiceURLImports(Text string) string {
	Otvet := Text

	if config.Settings.SERVICE_REPOSITORY_URL == "" {
		return Otvet
	}

	if config.Settings.TEMPLATE_REPOSITORY_URL == "" {
		return Otvet
	}

	if config.Settings.USE_DEFAULT_TEMPLATE == false {
		return Otvet
	}

	//
	TextFind := "import ("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}

	TEMPLATE_REPOSITORY_URL := config.Settings.TEMPLATE_REPOSITORY_URL
	SERVICE_REPOSITORY_URL := config.Settings.SERVICE_REPOSITORY_URL
	Otvet = strings.ReplaceAll(Otvet, TEMPLATE_REPOSITORY_URL, SERVICE_REPOSITORY_URL)

	return Otvet
}

// ReplaceModelAndTableName - заменяет имя модели и имя таблицы в шаблоне на новые
func ReplaceModelAndTableName(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, Table1.NameGo)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)

	return Otvet
}

// FindModelComment - возвращает комментарий для модели
func FindModelComment(Table1 *types.Table) string {
	Otvet := ""

	TableName := Table1.Name
	ModelName := Table1.NameGo
	COMMENT_MODEL_STRUCT := config.Settings.COMMENT_MODEL_STRUCT

	Otvet = `// ` + ModelName + ` - ` + COMMENT_MODEL_STRUCT + TableName + `: ` + Table1.Comment

	return Otvet
}

// FindModelNameComment - возвращает комментарий для названия модели
func FindModelNameComment(ModelName string, Table1 *types.Table) string {
	Otvet := ""

	TableName := Table1.Name
	//ModelName := Table1.NameGo
	COMMENT_MODEL_STRUCT := config.Settings.COMMENT_MODEL_STRUCT

	Otvet = `// ` + ModelName + ` - ` + COMMENT_MODEL_STRUCT + TableName + `: ` + Table1.Comment

	return Otvet
}

// ReplacePackageName - заменяет имя пакета в шаблоне на новое
func ReplacePackageName(Text, PackageName string) string {
	Otvet := Text

	//найдём имя каталога, это будет имя пакета
	PackageName = micro.DeleteEndSlash(PackageName)
	PackageName = micro.LastWord(PackageName)

	//
	TextFind := "package "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		log.Error("not found word: package ")
		return Otvet
	}

	s2 := Otvet[pos1:]
	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		log.Error("not found word: \n")
		return Otvet
	}

	Otvet = Otvet[:pos1+len(TextFind)] + PackageName + Otvet[pos1+posEnd:]

	return Otvet
}

// DeleteEmptyImport - удаляет пустой импорт
func DeleteEmptyImport(Text string) string {
	Otvet := Text

	sFind := `import (
)
`
	Otvet = strings.ReplaceAll(Otvet, sFind, "")

	return Otvet
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Delete ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Delete")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Restore ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Restore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверка есть ли колонки ExtID и ConnectionID
	if Has_Column_ExtID_ConnectionID(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Find_ByExtID ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Find_ByExtID")

	return Otvet
}

// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
func DeleteFuncDeleteCtx(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Delete_ctx ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Delete_ctx")

	return Otvet
}

// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
func DeleteFuncRestoreCtx(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Restore_ctx ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Restore_ctx")

	return Otvet
}

// DeleteFuncFind_byExtIDCtx - удаляет функцию Find_ByExtID_ctx()
func DeleteFuncFind_byExtIDCtx(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверка есть ли колонки ExtID и ConnectionID
	if Has_Column_ExtID_ConnectionID(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// Find_ByExtID_ctx ")
	Otvet = DeleteFuncFromFuncName(Otvet, "Find_ByExtID_ctx")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// TestDelete ")
	Otvet = DeleteFuncFromFuncName(Otvet, "TestDelete")

	return Otvet
}

// DeleteFuncReadFromCache - удаляет функцию ReadFromCache()
func DeleteFuncReadFromCache(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// ReadFromCache ")
	Otvet = DeleteFuncFromFuncName(Otvet, "ReadFromCache")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if Has_Column_IsDeleted(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	Otvet = DeleteFuncFromComment(Otvet, "\n// TestRestore ")
	Otvet = DeleteFuncFromFuncName(Otvet, "TestRestore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверка есть ли колонки ExtID и ConnectionID
	if Has_Column_ExtID_ConnectionID(Table1) == true {
		return Otvet
	}

	//
	Otvet = DeleteFuncFromComment(Otvet, "\n// TestFind_ByExtID ")
	Otvet = DeleteFuncFromFuncName(Otvet, "TestFind_ByExtID")

	return Otvet
}

// DeleteEmptyLines - удаляет пустые строки
func DeleteEmptyLines(Text string) string {
	Otvet := Text
	Otvet = strings.ReplaceAll(Otvet, "\n\n\n", "\n\n")
	Otvet = strings.ReplaceAll(Otvet, "\n//\n\n", "\n\n")
	Otvet = strings.ReplaceAll(Otvet, "\n\t//\n\n", "\n\n")
	//Otvet = strings.ReplaceAll(Otvet, "\r\r", "\r")
	//Otvet = strings.ReplaceAll(Otvet, "\r\n", "\n")
	//Otvet = strings.ReplaceAll(Otvet, "}\n\n", "}\n")
	pos1 := strings.Index(Otvet, "\n\n\n")
	if pos1 >= 0 {
		Otvet = DeleteEmptyLines(Otvet)
	}

	//удалим последние 2 пустые строки
	HasSuffix := strings.HasSuffix(Otvet, "\n\n")
	if HasSuffix == true {
		Otvet = Otvet[:len(Otvet)-1]
	}

	return Otvet
}

// DeleteLastUnderline - удаляет последний символ подчёркивания
func DeleteLastUnderline(s string) string {
	Otvet := s
	if s == "" {
		return Otvet
	}

	Otvet = strings.TrimSuffix(Otvet, "_")

	return Otvet
}

// FindLastGoodPos - возвращает позицию последнего нахождения, с новой строки
func FindLastGoodPos(s, TextFind string) int {
	Otvet := -1
	pos1 := strings.LastIndex(s, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	pos2 := strings.Index(s[pos1:], "\n")
	if pos2 < 0 {
		return Otvet
	}
	Otvet = pos1 + pos2 + 1

	return Otvet
}

// AddInterfaceFunction - добавляет функцию в интерфейс
func AddInterfaceFunction(s, TextAdd string) string {
	Otvet := s

	//Проверим такая функция уже есть
	pos1 := strings.Index(Otvet, TextAdd)
	if pos1 >= 0 {
		return Otvet
	}

	//найдём начало интефейса
	sFind := " interface {"
	pos1 = FindLastGoodPos(Otvet, sFind)
	if pos1 < 0 {
		log.Error("FindLastGoodPos() error: not found: ", sFind)
		return Otvet
	}

	s2 := Otvet[pos1:]
	pos2 := strings.Index(s2, "\n}")
	if pos2 < 0 {
		log.Error("FindLastGoodPos() error: not found: \\n")
		return Otvet
	}
	PosStart := pos1 + pos2

	Otvet = Otvet[:PosStart] + TextAdd + Otvet[PosStart:]

	return Otvet
}

// FindTextProtobufRequest - возвращает "RequestID" и "ID" - имя message из .proto, в зависимости от типа, а также название поля
func FindTextProtobufRequest(Table1 *types.Table) (string, string) {
	Otvet := ""
	TextRequestFieldName := ""

	TextRequest := "Request_"

	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return "", ""
	}

	PrimaryKeyTypeGo := PrimaryKeyColumn.TypeGo
	//switch PrimaryKeyTypeGo {
	//case "string", "uuid.UUID":
	//	TextRequest = "Request"
	//}

	TextRequestFieldName = PrimaryKeyColumn.TypeGo

	TextRequestFieldName = micro.StringFromUpperCase(TextRequestFieldName)
	Otvet = TextRequest + TextRequestFieldName

	switch PrimaryKeyTypeGo {
	case "string", "uuid.UUID":
		TextRequestFieldName = "String_1"
		Otvet = TextRequest + "String"
	case "time.Time":
		TextRequestFieldName = "Date"
		Otvet = TextRequest + "Date"
	}

	//switch TypeGo {
	//case "int", "int64":
	//	{
	//		if PrimaryKeyTypeGo == "int64" {
	//			Otvet = TextRequest + "Id"
	//			TextRequestFieldName = "ID"
	//		} else {
	//			Otvet = TextRequest + "Int64"
	//			TextRequestFieldName = "Int64"
	//		}
	//	}
	//
	//case "int32":
	//	{
	//		Otvet = TextRequest + "Int32"
	//		TextRequestFieldName = "Int32"
	//	}
	//case "string":
	//	{
	//		Otvet = TextRequest + "String"
	//		TextRequestFieldName = "String_1"
	//	}
	//case "uuid.UUID":
	//	{
	//		Otvet = TextRequest + "String"
	//		TextRequestFieldName = "String_1"
	//	}
	//case "time.Time":
	//	{
	//		Otvet = TextRequest + "Date"
	//		TextRequestFieldName = "Date"
	//	}
	//case "float32":
	//	{
	//		Otvet = TextRequest + "Float32"
	//		TextRequestFieldName = "Float32"
	//	}
	//case "float64":
	//	{
	//		Otvet = TextRequest + "Float64"
	//		TextRequestFieldName = "Float64"
	//	}
	//case "bool":
	//	{
	//		Otvet = TextRequest + "Bool"
	//		TextRequestFieldName = "Bool"
	//	}
	//}

	return Otvet, TextRequestFieldName
}

//// FindTextProtobufRequestPrimaryKey - возвращает "RequestID" и "ID" - имя message из .proto, в зависимости от типа, а также название поля
//func FindTextProtobufRequestPrimaryKey(Table1 *types.Table, TypeGo string) (string, string) {
//	Otvet := "RequestId"
//	TextRequestFieldName := "ID"
//
//	TextRequest := "Request"
//
//	switch TypeGo {
//	case "int", "int64":
//		{
//			Otvet = TextRequest + "Id"
//			TextRequestFieldName = "ID"
//		}
//
//	case "int32":
//		{
//			Otvet = TextRequest + "Int32"
//			TextRequestFieldName = "Int32"
//		}
//	case "string":
//		{
//			Otvet = TextRequest + "String"
//			TextRequestFieldName = "String_1"
//		}
//	case "uuid.UUID":
//		{
//			Otvet = TextRequest + "String"
//			TextRequestFieldName = "String_1"
//		}
//	case "time.Time":
//		{
//			Otvet = TextRequest + "Date"
//			TextRequestFieldName = "Date"
//		}
//	case "float32":
//		{
//			Otvet = TextRequest + "Float32"
//			TextRequestFieldName = "Float32"
//		}
//	case "float64":
//		{
//			Otvet = TextRequest + "Float64"
//			TextRequestFieldName = "Float64"
//		}
//	case "bool":
//		{
//			Otvet = TextRequest + "Bool"
//			TextRequestFieldName = "Bool"
//		}
//	}
//
//	return Otvet, TextRequestFieldName
//}

// FindTextProtobufRequest_ID_Type - возвращает имя message из .proto для двух параметров ID + Type,в зависимости от типа, а также название поля
// возвращает:
// Otvet - имя message из .proto
// TextRequestFieldName - название поля в Request
// TextRequestFieldGolang - название поля в Request с преобразованием в тип гоу
// TextGolangLine - замена всей строки в го
func FindTextProtobufRequest_ID_Type(Table1 *types.Table, Column1 *types.Column, VariableName string) (string, string, string, string) {
	Otvet := "RequestId"
	TextRequestFieldName := "ID"
	TextRequestFieldGolang := "ID"
	TextGolangLine := ""

	TypeGo := Column1.TypeGo
	TableName := Table1.Name
	ColumnName := Column1.Name

	//найдём тип колонки PrimaryKey
	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return "", "", "", ""
	}

	PrimaryKey_TypeGo := PrimaryKeyColumn.TypeGo
	//Text_Request_ID := "Request_ID"
	Otvet, _ = FindTextProtobufRequest(Table1)
	//Text_Request_ID = "Request_" + TextID

	TextRequestProtoName := ""

	//найдём строку по типу колонки
	switch TypeGo {
	case "int", "int64":
		{
			TextRequestProtoName = "Int64"
			TextRequestFieldName = "Int64"
			if PrimaryKey_TypeGo == "int64" {
				TextRequestFieldName = "Int64_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}

	case "int32":
		{
			TextRequestProtoName = "Int32"
			TextRequestFieldName = "Int32"
			if PrimaryKey_TypeGo == "int32" {
				TextRequestFieldName = "Int32_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}
	case "string":
		{
			TextRequestProtoName = "String"
			TextRequestFieldName = "String_1"
			if PrimaryKey_TypeGo == "string" || PrimaryKey_TypeGo == "uuid.UUID" {
				TextRequestFieldName = "String_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}
	case "uuid.UUID":
		{
			TextRequestProtoName = "String"
			TextRequestFieldName = "String_1"
			TextGolangLine = "value, err := uuid.Parse(" + VariableName + "" + TextRequestFieldName + ")" + `
	if Request.` + TextRequestFieldName + ` == "" {
		value = uuid.Nil
		err = nil
	}
	if err != nil {
		return &Otvet, err
	}
`
			if PrimaryKey_TypeGo == "string" || PrimaryKey_TypeGo == "uuid.UUID" {
				TextRequestFieldName = "String_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}
	case "time.Time":
		{
			TextRequestProtoName = "Date"
			TextRequestFieldName = "Date"
			if PrimaryKey_TypeGo == "time.Time" {
				TextRequestFieldName = "Date_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName + ".AsTime()"
		}
	case "float32":
		{
			TextRequestProtoName = "Float32"
			TextRequestFieldName = "Float32"
			if PrimaryKey_TypeGo == "float32" {
				TextRequestFieldName = "Float32_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}
	case "float64":
		{
			TextRequestProtoName = "Float64"
			TextRequestFieldName = "Float64"
			if PrimaryKey_TypeGo == "float64" {
				TextRequestFieldName = "Float64_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}
	case "bool":
		{
			TextRequestProtoName = "Bool"
			TextRequestFieldName = "Bool"
			if PrimaryKey_TypeGo == "bool" {
				TextRequestFieldName = "Bool_2"
			}
			TextRequestFieldGolang = VariableName + TextRequestFieldName
		}
	}

	TextConvert, ok := types.MapConvertID[TableName+"."+ColumnName]
	if ok == true {
		TextRequestFieldGolang = TextConvert + "(" + VariableName + TextRequestFieldName + ")"
	}

	Otvet = Otvet + "_" + TextRequestProtoName

	return Otvet, TextRequestFieldName, TextRequestFieldGolang, TextGolangLine
}

// ConvertRequestIdToAlias - заменяет ID на Alias
func ConvertRequestIdToAlias(Text string, Table1 *types.Table) string {
	Otvet := Text

	TableName := Table1.Name
	IDName, _ := FindPrimaryKeyNameType(Table1)
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == false {
		return Otvet
	}

	Otvet = strings.ReplaceAll(Otvet, "Request.ID", TextConvert+"(Request.ID)")
	if TextConvert[:6] != "alias." {
		return Otvet
	}

	URL := FindURL_Alias()
	if URL == "" {
		return Otvet
	}

	Otvet = AddImport(Otvet, URL)

	return Otvet
}

// ConvertIDToAlias_OtvetID - заменяет "Otvet.ID = ID" на "Otvet.ID = alias.Name(ID)"
func ConvertIDToAlias_OtvetID(Text string, Table1 *types.Table) string {
	Otvet := Text

	TableName := Table1.Name
	IDName, _ := FindPrimaryKeyNameType(Table1)
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == false {
		return Otvet
	}

	if TextConvert[:6] != "alias." {
		return Otvet
	}

	TextFrom := constants.TEXT_OTVET_ID_ALIAS
	TextTo := TextFrom
	TextTo = strings.ReplaceAll(TextFrom, " ID", " "+TextConvert+"(ID)")

	Otvet = strings.ReplaceAll(Otvet, TextFrom, TextTo)
	URL := FindURL_Alias()
	if URL == "" {
		return Otvet
	}

	Otvet = AddImport(Otvet, URL)

	return Otvet
}

// DeleteCommentFromString - удаляет комментарий из строки //
func DeleteCommentFromString(TextFrom string) string {
	Otvet := TextFrom

	pos1 := strings.Index(Otvet, "//")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1]

	return Otvet
}

// FindTextConvertGolangTypeToProtobufType - возвращает имя переменной +  имя колонки, преобразованное в тип protobuf
func FindTextConvertGolangTypeToProtobufType(Table1 *types.Table, Column1 *types.Column, VariableName string) string {
	Otvet := ""

	if Column1 == nil {
		return Otvet
	}

	Otvet = VariableName + Column1.NameGo

	//TableName := Table1.Name
	//IDName := Column1.Name

	////alias в Int64
	//TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	//if ok == true {
	//	Otvet = TextConvert + "(" + VariableName + Column1.NameGo + ")"
	//}

	//time.Time в timestamppb
	//if Column1.TypeGo == "time.Time" {
	//	Otvet = "timestamppb.New(" + VariableName + Column1.NameGo + ")"
	//}

	//преобразуем alias в обычный тип, и дату в timestamp
	switch Column1.TypeGo {
	case "time.Time":
		Otvet = "timestamppb.New(" + VariableName + Column1.NameGo + ")"
	case "string":
		Otvet = "string(" + VariableName + Column1.NameGo + ")"
	case "int64":
		Otvet = "int64(" + VariableName + Column1.NameGo + ")"
	case "int32":
		Otvet = "int32(" + VariableName + Column1.NameGo + ")"
	case "bool":
		Otvet = "bool(" + VariableName + Column1.NameGo + ")"
	case "float32":
		Otvet = "float32(" + VariableName + Column1.NameGo + ")"
	case "float64":
		Otvet = "float64(" + VariableName + Column1.NameGo + ")"
	case "uuid.UUID":
		Otvet = VariableName + Column1.NameGo + ".String()"
	}

	return Otvet
}

// FindTextConvertProtobufTypeToGolangType - возвращает имя переменной +  имя колонки, преобразованное в тип golang из protobuf
func FindTextConvertProtobufTypeToGolangType(Table1 *types.Table, Column1 *types.Column, VariableName string) (string, string) {
	Otvet := VariableName + Column1.NameGo
	GolangCode := ""

	TableName := Table1.Name
	IDName := Column1.Name

	RequestColumnName := Column1.NameGo
	RequestColumnName = FindRequestColumnName(Table1, Column1)

	//alias в Int64
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == true {
		Otvet = TextConvert + "(" + VariableName + Column1.NameGo + ")"
		return Otvet, GolangCode
	}

	//time.Time в timestamppb
	switch Column1.TypeGo {
	case "time.Time":
		{
			Otvet = VariableName + Column1.NameGo + ".AsTime()"
			return Otvet, GolangCode
		}
	case "uuid.UUID":
		{
			Otvet = "uuid.FromBytes([]byte(" + VariableName + RequestColumnName + "))"
			GolangCode = `ID, err := uuid.FromBytes([]byte(Request.` + RequestColumnName + `))
	if err != nil {
		return &Otvet, err
	}
`
			return Otvet, GolangCode
		}
	}
	//if Column1.TypeGo == "time.Time" {
	//	Otvet = VariableName + Column1.NameGo + ".AsTime()"
	//	return Otvet
	//}

	////преобразуем alias в обычный тип, и дату в timestamp
	//switch Column1.TypeGo {
	//case "time.Time":
	//	Otvet = "timestamppb.New(" + VariableName + Column1.NameGo + ")"
	//case "string":
	//	Otvet = "string(" + VariableName + Column1.NameGo + ")"
	//case "int64":
	//	Otvet = "int64(" + VariableName + Column1.NameGo + ")"
	//case "int32":
	//	Otvet = "int32(" + VariableName + Column1.NameGo + ")"
	//case "bool":
	//	Otvet = "bool(" + VariableName + Column1.NameGo + ")"
	//case "float32":
	//	Otvet = "float32(" + VariableName + Column1.NameGo + ")"
	//case "float64":
	//	Otvet = "float64(" + VariableName + Column1.NameGo + ")"
	//}

	return Otvet, GolangCode
}

// FindTextEqualEmpty - находит текст сравнение с пустым значением
func FindTextEqualEmpty(Column1 *types.Column, VariableName string) string {
	Otvet := ""

	DefaultValue := FindTextDefaultValue(Column1.TypeGo)
	Otvet = VariableName + " == " + DefaultValue

	if DefaultValue == "time.Time{}" {
		Otvet = VariableName + ".IsZero() == true"
	}

	return Otvet
}

//// AddSkipNowEveryFunc - добавляет функцию SkipNow() для каждой тестовой функции
//func AddSkipNowEveryFunc(Text string) string {
//	Otvet := ""
//
//	sFind := "\nfunc "
//	Mass := make([]string, 0)
//	Mass = strings.Split(Text, sFind)
//	for _, v := range Mass {
//		pos1 := strings.Index(v, sFind)
//		if pos1 < 0 {
//			continue
//		}
//
//		s2 := Text[pos1:]
//		pos2 := strings.Index(s2, "\n")
//		if pos2 < 0 {
//			continue
//		}
//		v = v[:pos1+pos2] + "\n\tt.SkipNow() //нет строк в БД \n" + v[pos1+pos2:]
//		Otvet = Otvet + v
//	}
//
//	return Otvet
//}

// FilenameWithoutLastUnderline - удаляет последний символ, если он = "_"
func FilenameWithoutLastUnderline(Filename string) string {
	Otvet := Filename

	if strings.HasSuffix(Filename, "_") == true {
		Otvet = Filename[:len(Filename)-1]
	}

	return Otvet
}

// FillVariable - заменяет переменную в тексте
func FillVariable(Text, VariableName, Value string) string {
	Otvet := Text

	sFind := VariableName + " = "
	pos1 := strings.Index(Otvet, sFind)
	if pos1 < 0 {
		return Otvet
	}

	s2 := Text[pos1:]
	posEnd := strings.Index(s2, "\n")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1] + VariableName + " = " + Value + Otvet[pos1+posEnd:]

	return Otvet
}

// Is_UUID_Type - проверяет является ли тип UUID
func Is_UUID_Type(TypeGo string) bool {
	Otvet := TypeGo == "uuid.UUID" || TypeGo == "uuid.NullUUID"
	return Otvet
}

//// ConvertID_toTypeID - заменяет int64(m.ID) на m.ID
//func ConvertID_toTypeID(Text string, Table1 *types.Table) string {
//	Otvet := Text
//
//	//заменим ID-Alias на ID
//	TableName := Table1.Name
//	IDName, _ := FindPrimaryKeyNameType(Table1)
//	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
//	if ok == true {
//		Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", ""+TextConvert+"(m.ID)")
//		return Otvet
//	}
//
//	//заменим int64(m.ID) на m.ID
//	Otvet = strings.ReplaceAll(Otvet, "int64(m.ID)", "m.ID")
//
//	return Otvet
//}

// Replace_Postgres_ID_Test - заменяет текст "const Postgres_ID_Test = 0" на нужный ИД
func Replace_Postgres_ID_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextFind := "const Postgres_ID_Test = 0"
	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	IDMinimum := Table1.IDMinimum
	if IDMinimum == "" {
		IDMinimum = FindTextDefaultValue(PrimaryKeyColumn.TypeGo)
	}

	switch PrimaryKeyColumn.TypeGo {
	case "uuid.UUID":
		{
			if Table1.IDMinimum == "" {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var Postgres_ID_Test = `+IDMinimum+``)
			} else {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var Postgres_ID_Test, _ = uuid.Parse("`+IDMinimum+`")`)
			}
		}
	case "string":
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const Postgres_ID_Test = "`+IDMinimum+`"`)
		}
	default:
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, "const Postgres_ID_Test = "+IDMinimum)
		}
	}

	return Otvet
}

// Replace_Model_ID_Test - заменяет текст "const LawsuitStatusType_ID_Test = 0" на нужный ИД
func Replace_Model_ID_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	TEXT_TEMPLATE_MODEL := config.Settings.TEXT_TEMPLATE_MODEL
	ModelName := Table1.NameGo
	TextFind := "const " + TEXT_TEMPLATE_MODEL + "_ID_Test = 0"
	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	IDMinimum := Table1.IDMinimum
	if IDMinimum == "" {
		IDMinimum = FindTextDefaultValue(PrimaryKeyColumn.TypeGo)
	}
	DefaultModelName := config.Settings.TEXT_TEMPLATE_MODEL

	switch PrimaryKeyColumn.TypeGo {
	case "uuid.UUID":
		{
			if Table1.IDMinimum == "" {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+ModelName+`_ID_Test = `+IDMinimum+``)
			} else {
				Otvet = strings.ReplaceAll(Otvet, TextFind, `var `+ModelName+`_ID_Test, _ = uuid.Parse("`+IDMinimum+`")`)
			}
			Otvet = strings.ReplaceAll(Otvet, ``+DefaultModelName+`_ID_Test`, ``+ModelName+`_ID_Test.String()`)
		}
	case "string":
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const `+ModelName+`_ID_Test = "`+IDMinimum+`"`)
		}
	default:
		{
			Otvet = strings.ReplaceAll(Otvet, TextFind, `const `+ModelName+`_ID_Test = `+IDMinimum)
		}
	}

	return Otvet
}

// ReplaceTextRequestID_and_Column - заменяет RequestId{} на RequestString{}
func ReplaceTextRequestID_and_Column(Text string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := Text

	//TypeGo := Column1.TypeGo

	TextRequestID, _, _, _ := FindTextProtobufRequest_ID_Type(Table1, Column1, "Request")
	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
	//Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextID)

	return Otvet
}

// ReplaceTextRequestID_PrimaryKey - заменяет RequestId{} на RequestString{}
func ReplaceTextRequestID_PrimaryKey(Text string, Table1 *types.Table) string {
	Otvet := Text

	TextRequest := "Request"
	Otvet = ReplaceTextRequestID_PrimaryKey1(Otvet, Table1, TextRequest)

	TextRequest = "Request2"
	Otvet = ReplaceTextRequestID_PrimaryKey1(Otvet, Table1, TextRequest)

	return Otvet
}

// ReplaceTextRequestID_PrimaryKey1 - заменяет RequestId{} на RequestString{}
func ReplaceTextRequestID_PrimaryKey1(Text string, Table1 *types.Table, TextRequest string) string {
	Otvet := Text

	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	//TypeGo := PrimaryKeyColumn.TypeGo

	TextRequestID, TextID := FindTextProtobufRequest(Table1)

	_, GolangCode := FindTextConvertProtobufTypeToGolangType(Table1, PrimaryKeyColumn, "Request.")
	if GolangCode != "" {
		Otvet = strings.ReplaceAll(Otvet, "ID := "+TextRequest+".ID", GolangCode)
		Otvet = strings.ReplaceAll(Otvet, TextRequest+".ID = ", TextRequest+"."+TextID+" = ")
	}

	Otvet = strings.ReplaceAll(Otvet, "RequestId{}", TextRequestID+"{}")
	Otvet = strings.ReplaceAll(Otvet, "*grpc_proto.RequestId", "*grpc_proto."+TextRequestID)
	//Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextID)
	Otvet = strings.ReplaceAll(Otvet, TextRequest+".ID", TextRequest+"."+TextID)

	return Otvet
}

// ReplaceIDtoID - заменяет int64(ID) на ID
func ReplaceIDtoID(Text string, Table1 *types.Table) string {
	Otvet := Text

	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	OtvetColumnName := FindTextConvertGolangTypeToProtobufType(Table1, PrimaryKeyColumn, "")
	if OtvetColumnName == "" {
		return Otvet
	}

	Otvet = strings.ReplaceAll(Otvet, "int64(ID)", OtvetColumnName)
	Otvet = strings.ReplaceAll(Otvet, "(ID int64", "("+PrimaryKeyColumn.NameGo+" "+PrimaryKeyColumn.TypeGo)
	Otvet = strings.ReplaceAll(Otvet, "(ID)", "("+PrimaryKeyColumn.NameGo+")")

	return Otvet
}

// ReplaceOtvetIDEqual1 - заменяет Otvet.ID = -1
func ReplaceOtvetIDEqual1(Text string, Table1 *types.Table) string {
	Otvet := Text

	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	Value := FindNegativeValue(PrimaryKeyColumn.TypeGo)

	Otvet = strings.ReplaceAll(Otvet, "Otvet.ID = -1", "Otvet.ID = "+Value)

	return Otvet
}

// ReplaceModelIDEqual1 - заменяет Otvet.ID = -1
func ReplaceModelIDEqual1(Text string, Table1 *types.Table) string {
	Otvet := Text

	PrimaryKeyColumn := FindPrimaryKeyColumn(Table1)
	if PrimaryKeyColumn == nil {
		return Otvet
	}

	Value := FindNegativeValue(PrimaryKeyColumn.TypeGo)

	Otvet = strings.ReplaceAll(Otvet, "m.ID = -1", "m.ID = "+Value)

	return Otvet
}

// FindNegativeValue - возвращает -1 для числовых типов
func FindNegativeValue(TypeGo string) string {
	Otvet := ""

	Otvet = FindTextDefaultValue(TypeGo)
	if mini_func.IsNumberType(TypeGo) == true {
		Otvet = "-1"
	}

	return Otvet
}

// FindRequestColumnName - возвращает название колонки в Request
func FindRequestColumnName(Table1 *types.Table, Column1 *types.Column) string {
	Otvet := ""

	_, Otvet = FindTextProtobufRequest(Table1)

	return Otvet
}

// ReplaceConnect_WithApplicationName - заменяет Connect_WithApplicationName() на Connect_WithApplicationName_SingularTableName()
func ReplaceConnect_WithApplicationName(Text string) string {
	Otvet := Text

	if config.Settings.SINGULAR_TABLE_NAMES == false {
		return Otvet
	}

	Otvet = strings.ReplaceAll(Otvet, "postgres_gorm.Connect_WithApplicationName(", "postgres_gorm.Connect_WithApplicationName_SingularTableName(")
	Otvet = strings.ReplaceAll(Otvet, "postgres_gorm.Start(", "postgres_gorm.Start_SingularTableName(")

	return Otvet
}
