package create_files

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"strings"
)

// DeleteImportModel - удаляет лишний импорт модели
func DeleteImportModel(s string) string {
	Otvet := s

	ModelURL := Find_ModelURL()
	ImportName := micro.LastWord(ModelURL)

	pos1 := strings.Index(Otvet, ImportName+".")
	if pos1 < 0 {
		Otvet = strings.ReplaceAll(Otvet, `"`+ModelURL+`"\n`, "")
		Otvet = strings.ReplaceAll(Otvet, `"`+ModelURL+`"`, "")
	}

	return Otvet
}

// Find_ModelURL - возвращает URL репозитория с пакетом "model"
func Find_ModelURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_MODEL
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_GRPCServerURL - возвращает URL репозитория с пакетом "server_grpc"
func Find_GRPCServerURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// FindGRPCClientURL - возвращает URL репозитория с пакетом "client_grpc"
func Find_GRPClientURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_NRPClientURL - возвращает URL репозитория с пакетом "client_nrpc"
func Find_NRPClientURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_GRPCConstantsURL - возвращает URL репозитория с пакетом "client_grpc"
func Find_GRPCConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC + "/" + "grpc_constants"
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_NRPCServerlURL - возвращает URL репозитория с пакетом "server_nrpc"
func Find_NRPCServerlURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_ProtobufURL - возвращает URL репозитория с файлом .proto
func Find_ProtobufURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.FOLDERNAME_GRPC_PROTO
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_TablesURL - возвращает URL репозитория с пакетом "tables"
func Find_TablesURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_TABLES
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_TableNameURL - возвращает URL репозитория с пакетом "tables" + TableName
func Find_TableNameURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_TABLES + "/" + config.Settings.PREFIX_TABLE + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_ProtoURL - возвращает URL репозитория с пакетом "grpc_proto"
func Find_ProtoURL() string {
	Otvet := ""

	FOLDERNAME_GRPC_PROTO := strings.ToLower(config.Settings.FOLDERNAME_GRPC_PROTO)
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + FOLDERNAME_GRPC_PROTO
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_ModelTableURL - возвращает URL репозитория model для таблицы TableName
func Find_ModelTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_MODEL + "/" + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_ObjectTableURL - возвращает URL репозитория object для таблицы TableName
func Find_ObjectTableURL(TableName string) string {
	Otvet := ""

	prefix_object := config.Settings.PREFIX_READOBJECT
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATES_READOBJECT_FOLDERNAME + "/" + prefix_object + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CrudObjectTableURL - возвращает URL репозитория crud_object для таблицы TableName
func Find_CrudObjectTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATES_CRUD_READOBJECT_FOLDERNAME + "/" + config.Settings.PREFIX_CRUD_READOBJECT + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CrudFuncURL - возвращает URL репозитория crud_func
func Find_CrudFuncURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATES_CRUD_FUNC_FOLDERNAME
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_GRPCClientTableURL - возвращает URL репозитория grpc_client для таблицы TableName
func Find_GRPCClientTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + "/" + config.Settings.PREFIX_CLIENT_GRPC + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CrudStarterTableURL - возвращает URL репозитория crud_starter для таблицы TableName
func Find_CrudStarterTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + "/" + config.Settings.STARTER_TABLES_PREFIX + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CrudTableURL - возвращает URL репозитория crud для таблицы TableName
func Find_CrudTableURL(TableName string) string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD + "/" + config.Settings.CRUD_TABLES_FREFIX + TableName
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_NRPC_Client_URL - возвращает URL репозитория с пакетом "nrpc_client"
func Find_NRPC_Client_URL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_DBConstantsURL - возвращает URL репозитория с пакетом db "db_constants"
func Find_DBConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_DB + "/" + "db_constants"
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_ConstantsURL - возвращает URL репозитория с пакетом pkg "constants"
func Find_ConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CONSTANTS
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_NRPCConstantsURL - возвращает URL репозитория с пакетом db "constants"
func Find_NRPCConstantsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC + "/" + "nrpc_constants"
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CrudStarterURL - возвращает URL репозитория с пакетом "crud_starter"
func Find_CrudStarterURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_GRPC_NRPC_URL - возвращает URL репозитория с пакетом "crud_starter"
func Find_GRPC_NRPC_URL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_NRPC
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_GRPCClient_func_URL - возвращает URL репозитория с пакетом "grpc_client_func"
func Find_GRPCClient_func_URL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT_FUNC
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CalcStructVersionURL - возвращает URL репозитория с пакетом "calc_struct_version"
func Find_CalcStructVersionURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_DB + "/" + "calc_struct_version"
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// Find_CrudFunctionsURL - возвращает URL репозитория с пакетом crud_functions
func Find_CrudFunctionsURL() string {
	Otvet := ""

	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDER_CRUD_FUNCTIONS
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

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
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// FindURL_Tables - возвращает URL репозитория с пакетом "tables"
func FindURL_Tables() string {
	Otvet := ""
	Otvet = config.Settings.SERVICE_REPOSITORY_URL + "/"
	Otvet = Otvet + config.Settings.TEMPLATE_FOLDERNAME_TABLES
	Otvet = strings.ReplaceAll(Otvet, `\`, `/`)

	return Otvet
}

// AddImport - добавляет RepositoryURL в секцию Import, если его там нет
func AddImport(Text, RepositoryURL string) string {
	Otvet := AddImport_WithAlias(Text, RepositoryURL, "")
	return Otvet
}

// AddImport - добавляет RepositoryURL в секцию Import, если его там нет
func AddImport_WithAlias(Text, RepositoryURL, AliasName string) string {
	Otvet := Text

	//если уже есть импорт
	pos1 := strings.Index(Otvet, `"`+RepositoryURL+`"`+"\n")
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

	if AliasName != "" {
		AliasName = AliasName + " "
	}

	Otvet = Otvet[:pos1+LenFind] + "\n\t" + AliasName + `"` + RepositoryURL + `"` + Otvet[pos1+LenFind:]

	return Otvet
}

// CheckAndAdd_Import - добавляет URL в секцию Import, если его там нет, если он нужен
func CheckAndAdd_Import(Text, URL string) string {
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

// AddImport_Time - добавляет пакет в секцию Import, если его там нет
func AddImport_Time(Text string) string {
	Otvet := Text

	RepositoryURL := `time`
	Otvet = AddImport(Text, RepositoryURL)

	return Otvet
}

// CheckAndAdd_ImportStrconv - добавляет пакет в секцию Import, если его там нет
func CheckAndAdd_ImportStrconv(Text string) string {
	Otvet := Text

	RepositoryURL := `strconv`
	Otvet = CheckAndAdd_Import(Text, RepositoryURL)

	return Otvet
}

// CheckAndAdd_ImportFmt - добавляет пакет fmt в секцию Import, если его там нет
func CheckAndAdd_ImportFmt(Text string) string {
	Otvet := Text

	RepositoryURL := `fmt`
	Otvet = CheckAndAdd_Import(Text, RepositoryURL)

	return Otvet
}

// CheckAndAdd_ImportPostgresFunc - добавляет пакет postgres_func в секцию Import, если его там нет
func CheckAndAdd_ImportPostgresFunc(Text string) string {
	Otvet := Text

	RepositoryURL := `github.com/ManyakRus/starter/postgres_func`
	Otvet = CheckAndAdd_Import(Text, RepositoryURL)

	return Otvet
}

// CheckAndAdd_ImportMicro - добавляет пакет micro в секцию Import, если его там нет
func CheckAndAdd_ImportMicro(Text string) string {
	Otvet := Text

	RepositoryURL := `github.com/ManyakRus/starter/micro`
	Otvet = CheckAndAdd_Import(Text, RepositoryURL)

	return Otvet
}

// AddImport_UUID - добавляет пакет в секцию Import, если его там нет
func AddImport_UUID(Text string) string {
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

// AddImport_Gorm - добавляет пакет в секцию Import, если его там нет
func AddImport_Gorm(Text string) string {
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

// AddImport_Timestamp - добавляет покет в секцию Import, если его там нет
func AddImport_Timestamp(Text string) string {
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

// CheckAndAdd_ImportAlias - добавляет покет в секцию Alias, если его там нет
func CheckAndAdd_ImportAlias(TextModel string) string {
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

// CheckAndAdd_ImportTime_FromTable - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAdd_ImportTime_FromTable(TextModel string, Table1 *types.Table) string {
	Otvet := TextModel

	HasTimeColumn := Has_ColumnType_Time(Table1)
	if HasTimeColumn == false {
		return Otvet
	}

	Otvet = AddImport_Time(Otvet)

	return Otvet
}

// CheckAndAdd_ImportTime_FromText - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAdd_ImportTime_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, " time.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImport_Time(Otvet)

	return Otvet
}

// CheckAndAdd_ImportUUID_FromText - добавляет пакет "uuid" в секцию Import, если его там нет
func CheckAndAdd_ImportUUID_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, "uuid.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImport_UUID(Otvet)

	return Otvet
}

// CheckAndAdd_ImportGorm_FromText - добавляет пакет "gorm.io/gorm" в секцию Import, если его там нет
func CheckAndAdd_ImportGorm_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, `"gorm.io/gorm"`)
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImport_Gorm(Otvet)

	return Otvet
}

// CheckAndAdd_ImportTimestamp_FromText - добавляет пакет "time" в секцию Import, если его там нет
func CheckAndAdd_ImportTimestamp_FromText(Text string) string {
	Otvet := Text

	pos1 := strings.Index(Text, " timestamppb.")
	if pos1 < 0 {
		return Otvet
	}

	Otvet = AddImport_Timestamp(Otvet)

	return Otvet
}

// Delete_TemplateRepositoryImports - удаляет импорты репозитория шаблона
func Delete_TemplateRepositoryImports(Text string) string {
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

	Otvet = Delete_TemplateRepositoryImports(Otvet)

	return Otvet
}

// Replace_RepositoryImportsURL - заменяет URL репозитория шаблона на URL репозитория сервиса
func Replace_RepositoryImportsURL(Text string) string {
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
