package crud_starter

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"io/fs"
	"os"
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке crud_starter
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	//crud_starter.go
	err = CreateFile_CrudStarter(MapAll)
	if err != nil {
		log.Error("CreateFile_CrudStarter() error: ", err)
		return err
	}

	//crud_starter_manual.go
	if config.Settings.NEED_CREATE_MANUAL_FILES == true {
		err = CreateFileCrudStarter_manual(MapAll)
		if err != nil {
			log.Error("CreateFileCrudStarter_manual() error: ", err)
			return err
		}
	}

	return err
}

// CreateFile_CrudStarter - создаёт 1 файл в папке crud_starter
func CreateFile_CrudStarter(MapAll map[string]*types.Table) error {
	var err error

	if config.Settings.NEED_CREATE_DB == false {
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

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyCrudStarter)

	//
	TextCrudStarter := config.Settings.TEXT_MODULE_GENERATED + `package crud_starter`

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
	if config.Settings.NEED_CREATE_GRPC == true {
		TextGRPC := FindTextGRPC(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextGRPC
	}

	//NRPC
	if config.Settings.NEED_CREATE_NRPC == true {
		TextNRPC := FindTextNRPC(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextNRPC
	}

	//замена импортов на новые URL
	TextCrudStarter = create_files.Replace_RepositoryImportsURL(TextCrudStarter)

	//удаление пустого импорта
	TextCrudStarter = create_files.Delete_EmptyImport(TextCrudStarter)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudStarter, []byte(TextCrudStarter), fs.FileMode(config.Settings.FILE_PERMISSIONS))

	return err
}

// FindTextImport - возвращает текст всех функций .proto для таблицы
func FindTextImport(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
import (`
	//сортировка по названию таблиц
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//все таблицы
	TextStarter := ""
	//TextModel := ""
	TextDB := ""
	TextGRPC := ""
	//TextNRPC := ""
	for _, key1 := range keys {
		Table1, ok := MapAll[key1]
		if ok == false {
			log.Panic("MapAll[key1] not found, key: ", key1)
		}

		//проверка что таблица нормальная
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		TextStarter = TextStarter + FindTextImportStarter1(Table1)
		//TextModel = TextModel + FindTextImportModel1(Table1)
		TextDB = TextDB + FindTextImportDB1(Table1)
		TextGRPC = TextGRPC + FindTextImportGRPC1(Table1)
		//TextNRPC = TextNRPC + FindTextImportNRPC1(Table1)
	}

	Otvet = Otvet + TextStarter + "\n" + TextDB + "\n" + TextGRPC

	Otvet = Otvet + "\n)"

	return Otvet
}

// FindTextImportStarter1 - возвращает текст импорта crud_starter_ для 1 таблицы
func FindTextImportStarter1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	DB_URL := config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + "/" + config.Settings.STARTER_TABLES_PREFIX + TableName
	Otvet := "\n\t\"" + DB_URL + `"`

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
	DB_URL := config.Settings.SERVICE_REPOSITORY_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_CRUD
	TableNameWithPrefix := config.Settings.PREFIX_CRUD + TableName
	Otvet := "\n\t" + "\"" + DB_URL + "/" + TableNameWithPrefix + `"`

	return Otvet
}

// FindTextImportGRPC1 - возвращает текст импорта GRPC для 1 таблицы
func FindTextImportGRPC1(Table1 *types.Table) string {
	GRPC_URL := config.Settings.SERVICE_REPOSITORY_URL
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

// FindText_InitCrudTransport_DB - возвращает текст для всех таблиц
func FindText_InitCrudTransport_DB(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_DB - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_DB() {`
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = Otvet + ` 
	initCrudTransport_manual_DB()
`
	}
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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextDB1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextDB1 - возвращает текст для 1 таблицы
func FindTextDB1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudInterface(" + config.Settings.PREFIX_CRUD + TableName + ".Crud_DB{})"

	return Otvet
}

// FindTextGRPC - возвращает текст для всех таблиц
func FindTextGRPC(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_GRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_GRPC() {`

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = Otvet + ` 
	initCrudTransport_manual_GRPC()
`
	}

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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextGRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextNRPC1 - возвращает текст для 1 таблицы
func FindTextGRPC1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudInterface(grpc_" + TableName + ".Crud_GRPC{})"

	return Otvet
}

// FindTextNRPC - возвращает текст для всех таблиц
func FindTextNRPC(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_NRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_NRPC() {`

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = Otvet + ` 
	initCrudTransport_manual_NRPC()
`
	}
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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextGRPC1(Table1) //GRPC будет делать функции NRPC
		//Otvet = Otvet + FindTextNRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextNRPC1 - возвращает текст 1 таблицы
func FindTextNRPC1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudInterface(nrpc_" + TableName + ".Crud_NRPC{})"

	return Otvet
}

// CreateFileCrudStarter_manual - создаёт 1 файл в папке crud_starter
func CreateFileCrudStarter_manual(MapAll map[string]*types.Table) error {
	var err error

	if config.Settings.NEED_CREATE_DB == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	//DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	//DirTemplatesCrudStarter := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	//FilenameTemplateCrudStarter := DirTemplatesCrudStarter + "crud_starter.go_"
	DirReadyCrudStarter := DirReady + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER + micro.SeparatorFile()
	FilenameReadyCrudStarter := DirReadyCrudStarter + "crud_starter_manual.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyCrudStarter)

	//
	TextCrudStarter := config.Settings.TEXT_MODULE_GENERATED + `package crud_starter`

	//найдём новый текст для каждой таблицы
	RepositoryURL := config.Settings.SERVICE_REPOSITORY_URL
	ModelURL := RepositoryURL + config.Settings.TEMPLATE_FOLDERNAME_MODEL

	//импорт
	TextImport := FindTextImport(MapAll, ModelURL)
	TextCrudStarter = TextCrudStarter + "\n" + TextImport

	//DB
	TextDB := FindText_InitCrudTransport_DB_manual(MapAll, ModelURL)
	TextCrudStarter = TextCrudStarter + "\n" + TextDB

	//GRPC
	if config.Settings.NEED_CREATE_GRPC == true {
		TextGRPC := FindTextGRPC_manual(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextGRPC
	}

	//NRPC
	if config.Settings.NEED_CREATE_NRPC == true {
		TextNRPC := FindTextNRPC_manual(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextNRPC
	}

	//замена импортов на новые URL
	TextCrudStarter = create_files.Replace_RepositoryImportsURL(TextCrudStarter)

	//удаление пустого импорта
	TextCrudStarter = create_files.Delete_EmptyImport(TextCrudStarter)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudStarter, []byte(TextCrudStarter), fs.FileMode(config.Settings.FILE_PERMISSIONS))

	return err
}

// FindText_InitCrudTransport_DB_manual - возвращает текст для каждой таблицы
func FindText_InitCrudTransport_DB_manual(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// initCrudTransport_manual_DB - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_DB() {`
	//	if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//		Otvet = Otvet + `
	//`
	//	}
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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextDB_manual1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextDB_manual1 - возвращает текст для 1 таблицы
func FindTextDB_manual1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudManualInterface(" + config.Settings.PREFIX_CRUD + TableName + ".Crud_DB{})"

	return Otvet
}

// FindTextGRPC_manual - возвращает текст для всех таблиц
func FindTextGRPC_manual(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// initCrudTransport_manual_GRPC - заполняет объекты crud для работы с БД через протокол GRPC
func initCrudTransport_manual_GRPC() {`

	//	if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//		Otvet = Otvet + `
	//`
	//	}

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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextGRPC_manual1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextGRPC_manual1 - возвращает текст для 1 таблицы
func FindTextGRPC_manual1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudManualInterface(grpc_" + TableName + ".Crud_GRPC{})"

	return Otvet
}

// FindTextNRPC - возвращает текст для всех таблиц
func FindTextNRPC_manual(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// initCrudTransport_manual_NRPC - заполняет объекты crud для работы с БД через протокол NRPC
func initCrudTransport_manual_NRPC() {`

	//	if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//		Otvet = Otvet + `
	//`
	//	}
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
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		Otvet = Otvet + FindTextGRPC_manual1(Table1) //GRPC будет делать функции NRPC
		//Otvet = Otvet + FindTextNRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}
