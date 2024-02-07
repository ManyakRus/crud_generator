package grpc_client_vars

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
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	err = CreateFileGRPCClientVars(MapAll)
	if err != nil {
		log.Error("CreateFileGRPCClientVars() error: ", err)
		return err
	}

	err = CreateFileNRPCClientVars(MapAll)
	if err != nil {
		log.Error("CreateFileNRPCClientVars() error: ", err)
		return err
	}

	return err
}

// CreateFileGRPCClientVars - создаёт 1 файл в папке grpc
func CreateFileGRPCClientVars(MapAll map[string]*types.Table) error {
	var err error

	if config.Settings.NEED_CREATE_GRPC == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyGRPCClientVars := DirReadyGRPCClient + constants.GRPC_CLIENT_VARS_FILENAME

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyGRPCClient)

	FilenameTemplateVars := DirTemplatesGRPCClient + constants.GRPC_CLIENT_VARS_FILENAME + "_"
	bytes, err := os.ReadFile(FilenameTemplateVars)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateVars, " error: ", err)
	}
	TextGRPCClientVars := string(bytes)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClientVars = create_files.DeleteTemplateRepositoryImports(TextGRPCClientVars)

		GRPCProtoURL := create_files.FindGRPCProtoURL()
		TextGRPCClientVars = create_files.AddImport(TextGRPCClientVars, GRPCProtoURL)
	}

	//
	TextVars := FindTextVars(MapAll, false)
	TextGRPCClientVars = TextGRPCClientVars + TextVars

	//
	TextFuncFillClients := FindTextFuncFillClients(MapAll)
	TextGRPCClientVars = TextGRPCClientVars + `
// FillClients - создание клиентов GRPC` + TextFuncFillClients

	//
	TextGRPCClientVars = TextGRPCClientVars + "\n}\n"

	//
	TextGRPCClientVars = create_files.DeleteEmptyLines(TextGRPCClientVars)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClientVars, []byte(TextGRPCClientVars), constants.FILE_PERMISSIONS)

	return err
}

// CreateFileNRPCClientVars - создаёт 1 файл в папке grpc
func CreateFileNRPCClientVars(MapAll map[string]*types.Table) error {
	var err error

	if config.Settings.NEED_CREATE_NRPC == false {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesNRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	DirReadyNRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	FilenameReadyGRPCClientVars := DirReadyNRPCClient + constants.NRPC_CLIENT_VARS_FILENAME

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyNRPCClient)

	FilenameTemplateVars := DirTemplatesNRPCClient + constants.NRPC_CLIENT_VARS_FILENAME + "_"
	bytes, err := os.ReadFile(FilenameTemplateVars)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateVars, " error: ", err)
	}
	TextNRPCClientVars := string(bytes)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextNRPCClientVars = create_files.DeleteTemplateRepositoryImports(TextNRPCClientVars)

		NRPCProtoURL := create_files.FindNRPCProtoURL()
		TextNRPCClientVars = create_files.AddImport(TextNRPCClientVars, NRPCProtoURL)
	}

	//
	TextVars := FindTextVars(MapAll, true)
	TextNRPCClientVars = TextNRPCClientVars + TextVars

	//
	TextFuncFillClients := FindTextFuncFillClients(MapAll)
	TextNRPCClientVars = TextNRPCClientVars + `
// FillClients - создание клиентов NRPC` + TextFuncFillClients

	//
	TextNRPCClientVars = TextNRPCClientVars + "\n}\n"

	//
	TextNRPCClientVars = create_files.DeleteEmptyLines(TextNRPCClientVars)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClientVars, []byte(TextNRPCClientVars), constants.FILE_PERMISSIONS)

	return err
}

// FindTextVars - возвращает текст создания переменных
func FindTextVars(MapAll map[string]*types.Table, IsNRPC bool) string {
	Otvet := ""

	//сортировка
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextStar := ""
	if IsNRPC == true {
		TextStar = "*"
	}

	//создание переменных
	for _, key1 := range keys {
		Table1 := MapAll[key1]
		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}
		Comment := create_files.FindModelComment(Table1)
		Text1 := Comment + "\n" + "var " + Table1.NameGo + "Client " + TextStar + "grpc_proto." + Table1.NameGo + "Client"
		Otvet = Otvet + "\n" + Text1 + "\n"
	}
	return Otvet
}

// FindTextFuncFillClients - возвращает текст создания функции FillClients
func FindTextFuncFillClients(MapAll map[string]*types.Table) string {
	Otvet := ""

	//сортировка
	keys := make([]string, 0, len(MapAll))
	for k := range MapAll {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//создание функции
	Otvet = Otvet + `
func FillClients() {`

	for _, key1 := range keys {
		Table1 := MapAll[key1]
		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}
		Text1 := "\t" + Table1.NameGo + "Client = " + "grpc_proto.New" + Table1.NameGo + "Client(Conn)"
		Otvet = Otvet + "\n" + Text1

	}

	return Otvet
}
