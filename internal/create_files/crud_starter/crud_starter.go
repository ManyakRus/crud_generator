package crud_starter

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	err = CreateFileCrudStarter(MapAll)
	if err != nil {
		log.Error("CreateFileCrudStarter() error: ", err)
		return err
	}

	return err
}

// CreateFileCrudStarter - создаёт 1 файл в папке grpc
func CreateFileCrudStarter(MapAll map[string]*types.Table) error {
	var err error

	if config.Settings.NEED_CRUD == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	//DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	//DirTemplatesCrudStarter := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	//FilenameTemplateCrudStarter := DirTemplatesCrudStarter + "crud_starter.go_"
	DirReadyCrudStarter := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	FilenameReadyCrudStarter := DirReadyCrudStarter + "crud_starter.go"

	//bytes, err := os.ReadFile(FilenameTemplateCrudStarter)
	//if err != nil {
	//	log.Panic("ReadFile() ", FilenameTemplateCrudStarter, " error: ", err)
	//}
	//TextCrudStarter := string(bytes)
	//
	////заменим название сервиса
	//TextCrudStarter = strings.ReplaceAll(TextCrudStarter, config.Settings.TEMPLATE_SERVICE_NAME, config.Settings.SERVICE_NAME)

	TextCrudStarter := constants.TEXT_GENERATED + `package crud_starter`

	//найдём новый текст для каждой таблицы
	RepositoryURL := config.Settings.SERVICE_REPOSITORY_URL
	ModelURL := RepositoryURL + config.Settings.TEMPLATE_FOLDERNAME_MODEL

	//импорт
	TextImport := FindTextImport(MapAll, ModelURL)
	TextCrudStarter = TextCrudStarter + "\n" + TextImport

	//DB
	TextDB := FindText_InitCrudTransport_DB(MapAll, ModelURL)
	TextCrudStarter = TextCrudStarter + "\n" + TextDB

	//GRPC
	if config.Settings.NEED_GRPC == true {
		TextGRPC := FindTextGRPC(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextGRPC
	}

	//NRPC
	if config.Settings.NEED_NRPC == true {
		TextNRPC := FindTextNRPC(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextNRPC
	}

	//запись файла
	err = os.WriteFile(FilenameReadyCrudStarter, []byte(TextCrudStarter), constants.FILE_PERMISSIONS)

	return err
}

// FindTextImport - возвращает текст всех функций .proto для таблицы
func FindTextImport(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
import (
	`
	//сортировка по названию таблиц
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//все таблицы
	TextModel := ""
	TextDB := ""
	TextGRPC := ""
	TextNRPC := ""
	for _, key1 := range keys {
		Table1, ok := MapAll[key1]
		if ok == false {
			log.Panic("MapAll[key1] not found, key: ", key1)
		}

		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		TextModel = TextModel + FindTextImportModel1(Table1)
		TextDB = TextDB + FindTextImportDB1(Table1)
		TextGRPC = TextGRPC + FindTextImportGRPC1(Table1)
		TextNRPC = TextNRPC + FindTextImportNRPC1(Table1)
	}

	Otvet = Otvet + TextModel + "\n" + TextDB + "\n" + TextGRPC + "\n" + TextNRPC

	Otvet = Otvet + "\n)"

	return Otvet
}

// FindTextImportModel1 - возвращает текст импорта Model для 1 таблицы
func FindTextImportModel1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	DB_URL := config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_MODEL + "/" + TableName
	Otvet := "\n\t\"" + DB_URL + `"`

	return Otvet
}

// FindTextImportDB1 - возвращает текст импорта DB для 1 таблицы
func FindTextImportDB1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	DB_URL := config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_DB
	Otvet := "\n\t\"" + DB_URL + "/db_" + TableName + `"`

	return Otvet
}

// FindTextImportGRPC1 - возвращает текст импорта GRPC для 1 таблицы
func FindTextImportGRPC1(Table1 *types.Table) string {
	GRPC_URL := config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t\"" + GRPC_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + "/grpc_" + TableName + `"`

	return Otvet
}

// FindTextImportNRPC1 - возвращает текст импорта NRPC для 1 таблицы
func FindTextImportNRPC1(Table1 *types.Table) string {
	//NRPC_URL := config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t\"" + config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + "/nrpc_" + TableName + `"`

	return Otvet
}

// FindText_InitCrudTransport_DB - возвращает текст всех функций .proto для таблицы
func FindText_InitCrudTransport_DB(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_DB - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_DB() {`
	//сортировка по названию таблиц
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Table1, ok := MapAll[key1]
		if ok == false {
			log.Panic("MapAll[key1] not found, key: ", key1)
		}

		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextDB1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextNRPC1 - возвращает текст всех функций .proto для таблицы
func FindTextDB1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	ModelName := Table1.NameGo
	Otvet := "\n\t" + TableName + "." + ModelName + "{}.SetCrudInterface(db_" + TableName + ".Crud_DB{})"

	return Otvet
}

// FindTextNRPC - возвращает текст всех функций .proto для таблицы
func FindTextGRPC(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_GRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_GRPC() {`
	//сортировка по названию таблиц
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Table1, ok := MapAll[key1]
		if ok == false {
			log.Panic("MapAll[key1] not found, key: ", key1)
		}

		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextGRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextNRPC1 - возвращает текст всех функций .proto для таблицы
func FindTextGRPC1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	ModelName := Table1.NameGo
	Otvet := "\n\t" + TableName + "." + ModelName + "{}.SetCrudInterface(grpc_" + TableName + ".Crud_GRPC{})"

	return Otvet
}

// FindTextNRPC - возвращает текст всех функций .proto для таблицы
func FindTextNRPC(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_NRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_NRPC() {`
	//сортировка по названию таблиц
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Table1, ok := MapAll[key1]
		if ok == false {
			log.Panic("MapAll[key1] not found, key: ", key1)
		}

		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextNRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextNRPC1 - возвращает текст всех функций .proto для таблицы
func FindTextNRPC1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	ModelName := Table1.NameGo
	Otvet := "\n\t" + TableName + "." + ModelName + "{}.SetCrudInterface(nrpc_" + TableName + ".Crud_NRPC{})"

	return Otvet
}
