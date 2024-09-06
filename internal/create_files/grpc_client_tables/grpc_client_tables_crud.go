package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/folders"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateFiles - создаёт 1 файл в папке grpc_client
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + "grpc_client_table.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + ".go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//создание текста
	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		//proto
		RepositoryGRPCProtoURL := create_files.Find_ProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)

		//model
		RepositoryModelURL := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryModelURL)

		////grpc client
		//RepositoryGRPCClientlURL := create_files.Find_GRPClientURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCClientlURL)

		//grpc client func
		GRPCClientFuncURL := create_files.Find_GRPCClient_func_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPCClientFuncURL)

		//nrpc client
		RepositoryNRPCClientlURL := create_files.Find_NRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)

		//constants GRPC
		RepositoryGRPCConstantsURL := create_files.Find_GRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		//DBConstantsURL := create_files.Find_DBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//grpc_nrpc
		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)

		//замена ID на PrimaryKey
		//TextGRPCClient = create_files.Replace_PrimaryKeyOtvetID(TextGRPCClient, Table1)

		//замена ID на PrimaryKey
		TextGRPCClient = create_files.Replace_PrimaryKeyM_ID(TextGRPCClient, Table1)

		//замена RequestId{}
		TextGRPCClient = create_files.ReplaceText_RequestID_PrimaryKey(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	}

	//удалим лишние функции
	TextGRPCClient = create_files.DeleteFunc_Delete(TextGRPCClient, Table1)
	TextGRPCClient = create_files.DeleteFunc_Restore(TextGRPCClient, Table1)
	TextGRPCClient = create_files.DeleteFunc_Find_byExtID(TextGRPCClient, Table1)

	//замена имени таблицы
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	//ModelName := Table1.NameGo
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClient, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesTest - создаёт 1 файл в папке grpc_client
func CreateFilesTest(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + "grpc_client_table_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile() + "tests" + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.Replace_PackageName(TextGRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.Delete_TemplateRepositoryImports(TextGRPCClient)

		GRPClientURL := create_files.Find_GRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		ModelTableName := create_files.Find_ModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableName)

		GRPClientTableURL := create_files.Find_GRPCClientTableURL(Table1.Name)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientTableURL)

		//GRPClientFuncURL := create_files.Find_GRPCClient_func_URL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientFuncURL)

		//Postgres_ID_Test = ID Minimum
		TextGRPCClient = create_files.Replace_Postgres_ID_Test(TextGRPCClient, Table1)

		//замена Otvet.ID = -1
		TextGRPCClient = create_files.Replace_OtvetIDEqual1(TextGRPCClient, Table1)

		//замена Otvet.ID = 0
		TextGRPCClient = create_files.Replace_OtvetIDEqual0(TextGRPCClient, Table1)

		//замена ID на PrimaryKey
		TextGRPCClient = create_files.Replace_PrimaryKeyOtvetID(TextGRPCClient, Table1)

		//добавим импорт uuid
		TextGRPCClient = create_files.CheckAndAdd_ImportUUID_FromText(TextGRPCClient)

	}

	//создание текста
	TextGRPCClient = create_files.Replace_TemplateModel_to_Model(TextGRPCClient, Table1.NameGo)
	TextGRPCClient = create_files.Replace_TemplateTableName_to_TableName(TextGRPCClient, Table1.Name)
	TextGRPCClient = create_files.AddText_ModuleGenerated(TextGRPCClient)

	ModelName := Table1.NameGo
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	//TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	if config.Settings.HAS_IS_DELETED == true {
		TextGRPCClient = DeleteFuncTestDelete(TextGRPCClient, ModelName, Table1)
		TextGRPCClient = DeleteFuncTestRestore(TextGRPCClient, ModelName, Table1)
	}
	TextGRPCClient = DeleteFuncTestFind_byExtID(TextGRPCClient, ModelName, Table1)

	//SkipNow()
	TextGRPCClient = create_files.AddSkipNow(TextGRPCClient, Table1)

	//замена импортов на новые URL
	TextGRPCClient = create_files.Replace_RepositoryImportsURL(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.Delete_EmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.Delete_EmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClient, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete ")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore ")

	return Otvet
}

// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
func DeleteFuncDeleteCtx(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Delete_ctx ")

	return Otvet
}

// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
func DeleteFuncRestoreCtx(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Restore_ctx ")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID_Int64(Table1)
	if ok == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// Find_ByExtID ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestDelete")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestRestore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID_Int64(Table1)
	if ok == true {
		return Otvet
	}

	//
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestFindByExtID")

	return Otvet
}
