package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_ReadAll - создаёт 1 файл в папке grpc_client
func CreateFiles_ReadAll(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_readall.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCClientFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ProtoURL)

		GRPCClient_func_URL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClient_func_URL)

		GRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCConstantsURL)

		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		NRPC_Client_URL := create_files.Find_NRPC_Client_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, NRPC_Client_URL)

	}

	//создание функций
	TextGRPCClientFunc := CreateFiles_ReadAll_Table(Table1, TextTemplatedFunction)
	if TextGRPCClientFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextGRPCClientFunc

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//uuid
	TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	//alias
	TextGRPCClient = create_files.CheckAndAdd_ImportAlias(TextGRPCClient)

	//time
	TextGRPCClient = create_files.CheckAndAdd_ImportTime_FromText(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateFiles_ReadAll_Table - создаёт текст всех функций
func CreateFiles_ReadAll_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	_, ok := types.MapReadAll[Table1]
	if ok == false {
		return Otvet
	}

	Otvet1 := CreateFiles_ReadAll_Table1(Table1, TextTemplateFunction)
	Otvet = Otvet + Otvet1

	return Otvet
}

// CreateFiles_ReadAll_Table1 - создаёт текст всех функций
func CreateFiles_ReadAll_Table1(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	//
	return Otvet
}

// CreateFiles_ReadAll_Test - создаёт 1 файл в папке grpc_client
func CreateFiles_ReadAll_Test(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReady := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_readall_test.go"

	//создадим каталог
	ok, err := micro.FileExists(DirReadyTable)
	if ok == false {
		err = os.MkdirAll(DirReadyTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirReadyTable, " error: ", err)
		}
	}

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCClientFunction := DirTemplatesGRPCClient + config.Settings.TEMPLATES_GRPC_CLIENT_TABLES_READALL_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCClientFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClientFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//ModelTableURL := create_files.Find_ModelTableURL(TableName)
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableURL)

		//GRPCClient_func_URL := create_files.Find_GRPCClient_func_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClient_func_URL)

		GRPClientURL := create_files.Find_GRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

		CrudFuncURL := create_files.Find_CrudFuncURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, CrudFuncURL)
	}

	//создание функций
	TextGRPCClientFunc := CreateFiles_ReadAll_Test_Table(Table1, TextTemplatedFunction)
	if TextGRPCClientFunc == "" {
		return err
	}
	TextGRPCClient = TextGRPCClient + TextGRPCClientFunc

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//uuid
	TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	//alias
	TextGRPCClient = create_files.CheckAndAdd_ImportAlias(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateFiles_ReadAll_Test_Table - создаёт текст 1 функции
func CreateFiles_ReadAll_Test_Table(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := ""

	_, ok := types.MapReadAll[Table1]
	if ok == false {
		return Otvet
	}

	Otvet1 := CreateFiles_ReadAll_Test_Table1(Table1, TextTemplateFunction)
	Otvet = Otvet + Otvet1

	return Otvet
}

// CreateFiles_ReadAll_Test_Table1 - создаёт текст 1 функции
func CreateFiles_ReadAll_Test_Table1(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	return Otvet
}
