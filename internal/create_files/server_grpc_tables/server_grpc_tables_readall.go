package server_grpc_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles_ReadAll - создаёт 1 файл в папке server_grpc
func CreateFiles_ReadAll(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_READALL_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_readall.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCServerFunction := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_READALL_FUNCTION_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		CrudTableURL := create_files.Find_CrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

	}

	//создание функций
	TextGRPCServerFunc := CreateFiles_ReadAll_Table(Table1, TextTemplatedFunction)
	if TextGRPCServerFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextGRPCServerFunc

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена импортов на новые URL
	TextGRPCServer = create_files.Replace_RepositoryImportsURL(TextGRPCServer)

	//uuid
	TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	//alias
	TextGRPCServer = create_files.CheckAndAdd_ImportAlias(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.Delete_EmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

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

	return Otvet
}

// CreateFiles_ReadAll_Test - создаёт 1 файл в папке server_grpc
func CreateFiles_ReadAll_Test(Table1 *types.Table) error {
	var err error

	if len(types.MapReadAll) == 0 {
		return err
	}

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_READALL_TEST_FILENAME
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReady := DirReadyTable + micro.SeparatorFile() + config.Settings.PREFIX_SERVER_GRPC + TableName + "_readall_test.go"

	//создадим каталог
	create_files.CreateDirectory(DirReadyTable)

	//загрузим шаблон файла
	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//загрузим шаблон файла функции
	FilenameTemplateGRPCServerFunction := DirTemplatesGRPCServer + config.Settings.TEMPLATES_GRPC_SERVER_READALL_FUNCTION_TEST_FILENAME
	bytes, err = os.ReadFile(FilenameTemplateGRPCServerFunction)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServerFunction, " error: ", err)
	}
	TextTemplatedFunction := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.Replace_PackageName(TextGRPCServer, DirReadyTable)

	//ModelName := Table1.NameGo
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.Delete_TemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		CrudStarterURL := create_files.Find_CrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ProtoURL := create_files.Find_ProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		//
		CrudFuncURL := create_files.Find_CrudFuncURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudFuncURL)

	}

	//создание функций
	TextGRPCServerFunc := CreateFiles_ReadAll_Test_Table(Table1, TextTemplatedFunction)
	if TextGRPCServerFunc == "" {
		return err
	}
	TextGRPCServer = TextGRPCServer + TextGRPCServerFunc

	//создание текста
	TextGRPCServer = create_files.Replace_TemplateModel_to_Model(TextGRPCServer, Table1.NameGo)
	TextGRPCServer = create_files.Replace_TemplateTableName_to_TableName(TextGRPCServer, Table1.Name)
	TextGRPCServer = create_files.AddText_ModuleGenerated(TextGRPCServer)

	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//замена импортов на новые URL
	TextGRPCServer = create_files.Replace_RepositoryImportsURL(TextGRPCServer)

	//uuid
	TextGRPCServer = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCServer)

	//alias
	TextGRPCServer = create_files.CheckAndAdd_ImportAlias(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.Delete_EmptyImport(TextGRPCServer)

	//удаление пустых строк
	TextGRPCServer = create_files.Delete_EmptyLines(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReady, []byte(TextGRPCServer), config.Settings.FILE_PERMISSIONS)

	return err
}

// CreateFiles_ReadAll_Test_Table - создаёт текст всех функций
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

// CreateFiles_ReadAll_Test_Table1 - создаёт текст всех функций
func CreateFiles_ReadAll_Test_Table1(Table1 *types.Table, TextTemplateFunction string) string {
	Otvet := TextTemplateFunction

	return Otvet
}
