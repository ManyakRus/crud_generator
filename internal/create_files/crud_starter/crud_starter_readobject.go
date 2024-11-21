package crud_starter

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"sort"
	"strings"
)

// CreateFile_CrudStarter_ReadObject - создаёт 1 файл в папке crud_starter
func CreateFile_CrudStarter_ReadObject(MapAll map[string]*types.Table) error {
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
	FilenameReadyCrudStarter := DirReadyCrudStarter + "crud_starter_readobject.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyCrudStarter)

	//
	TextCrudStarter := config.Settings.TEXT_MODULE_GENERATED + `package crud_starter`

	//найдём новый текст для каждой таблицы
	RepositoryURL := config.Settings.SERVICE_REPOSITORY_URL
	ModelURL := RepositoryURL + config.Settings.TEMPLATE_FOLDERNAME_MODEL

	//импорт
	TextImport := FindTextReadObjectImport(MapAll, ModelURL)
	TextCrudStarter = TextCrudStarter + "\n" + TextImport

	//ReadObject
	TextDB := FindText_InitCrudTransport_ReadObject(MapAll, ModelURL)
	TextCrudStarter = TextCrudStarter + "\n" + TextDB

	//GRPC
	if config.Settings.NEED_CREATE_GRPC == true {
		TextGRPC := FindTextReadObjectGRPC(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextGRPC
	}

	//NRPC
	if config.Settings.NEED_CREATE_NRPC == true {
		TextNRPC := FindTextReadObjectNRPC(MapAll, ModelURL)
		TextCrudStarter = TextCrudStarter + "\n" + TextNRPC
	}

	//заменим grpc_proto на новое
	TextProto := create_files.TextProto()
	TextCrudStarter = strings.ReplaceAll(TextCrudStarter, "grpc_proto.", TextProto+".")

	////GRPC
	//if config.Settings.NEED_CREATE_GRPC == true {
	//	TextGRPC := FindTextGRPC(MapAll, ModelURL)
	//	TextCrudStarter = TextCrudStarter + "\n" + TextGRPC
	//}
	//
	////NRPC
	//if config.Settings.NEED_CREATE_NRPC == true {
	//	TextNRPC := FindTextNRPC(MapAll, ModelURL)
	//	TextCrudStarter = TextCrudStarter + "\n" + TextNRPC
	//}

	//замена импортов на новые URL
	TextCrudStarter = create_files.Replace_RepositoryImportsURL(TextCrudStarter)

	//удаление пустого импорта
	TextCrudStarter = create_files.Delete_EmptyImport(TextCrudStarter)

	//запись файла
	err = os.WriteFile(FilenameReadyCrudStarter, []byte(TextCrudStarter), config.Settings.FILE_PERMISSIONS)

	return err
}

// FindTextReadObjectImport - возвращает текст всех функций .proto для таблицы
func FindTextReadObjectImport(MapAll map[string]*types.Table, ModelURL string) string {
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
		TextDB = TextDB + FindTextReadObjectImportDB1(Table1)
		TextGRPC = TextGRPC + FindTextReadObjectImportGRPC1(Table1)
	}

	Otvet = Otvet + TextStarter + "\n" + TextDB + "\n" + TextGRPC

	Otvet = Otvet + "\n)"

	return Otvet
}

// FindTextReadObjectImportDB1 - возвращает текст импорта DB для 1 таблицы
func FindTextReadObjectImportDB1(Table1 *types.Table) string {
	CrudObjectURL := create_files.Find_CrudObjectTableURL(Table1.Name)
	Otvet := "\n\t" + `"` + CrudObjectURL + `"`

	return Otvet
}

// FindTextReadObjectImportGRPC1 - возвращает текст импорта GRPC для 1 таблицы
func FindTextReadObjectImportGRPC1(Table1 *types.Table) string {
	GRPC_URL := config.Settings.SERVICE_REPOSITORY_URL
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t\"" + GRPC_URL + "/" + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + "/grpc_" + TableName + `"`

	return Otvet
}

// FindText_InitCrudTransport_ReadObject - возвращает текст для всех таблиц
func FindText_InitCrudTransport_ReadObject(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_ReadObject_DB - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_ReadObject_DB() {`
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = Otvet + ` 
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

		Otvet = Otvet + FindTextReadObject1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextReadObject1 - возвращает текст для 1 таблицы
func FindTextReadObject1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudReadObjectInterface(" + config.Settings.PREFIX_CRUD + "object_" + TableName + ".Crud_DB{})"

	return Otvet
}

// FindTextReadObjectGRPC - возвращает текст для всех таблиц
func FindTextReadObjectGRPC(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_ReadObject_GRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_ReadObject_GRPC() {`

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = Otvet + ` 
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

		Otvet = Otvet + FindTextReadObjectGRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

// FindTextReadObjectGRPC1 - возвращает текст для 1 таблицы
func FindTextReadObjectGRPC1(Table1 *types.Table) string {
	TableName := strings.ToLower(Table1.Name)
	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudReadObjectInterface(grpc_" + TableName + ".Crud_GRPC{})"

	return Otvet
}

// FindTextReadObjectNRPC - возвращает текст для всех таблиц
func FindTextReadObjectNRPC(MapAll map[string]*types.Table, ModelURL string) string {
	Otvet := `
// InitCrudTransport_ReadObject_NRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_ReadObject_NRPC() {`

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		Otvet = Otvet + ` 
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

		Otvet = Otvet + FindTextReadObjectGRPC1(Table1) //GRPC будет делать функции NRPC
		//Otvet = Otvet + FindTextNRPC1(Table1)
	}

	Otvet = Otvet + "\n}"

	return Otvet
}

//// FindTextReadObjectNRPC1 - возвращает текст 1 таблицы
//func FindTextReadObjectNRPC1(Table1 *types.Table) string {
//	TableName := strings.ToLower(Table1.Name)
//	Otvet := "\n\t" + config.Settings.STARTER_TABLES_PREFIX + TableName + ".SetCrudInterface(nrpc_" + TableName + ".Crud_NRPC{})"
//
//	return Otvet
//}
