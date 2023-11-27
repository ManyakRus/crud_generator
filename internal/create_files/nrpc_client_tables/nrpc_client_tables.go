package nrpc_client_tables

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

// CreateAllFiles - создаёт все файлы в папке nrpc_client
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы nrpc_client
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы nrpc_client
		if config.Settings.NEED_CREATE_NRPC_CLIENT_TEST == true {
			err = CreateTestFiles(Table1)
			if err != nil {
				log.Error("CreateTestFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}
	}
	return err
}

// CreateFiles - создаёт 1 файл в папке nrpc_client
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesNRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	DirReadyNRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateNRPCClient := DirTemplatesNRPCClient + constants.NRPC_CLIENT_TABLE_FILENAME + "_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyNRPCClient + "nrpc_" + TableName + micro.SeparatorFile()
	FilenameReadyNRPCClient := DirReadyTable + "nrpc_" + TableName + ".go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateNRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateNRPCClient, " error: ", err)
	}
	TextNRPCClient := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextNRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextNRPCClient = create_files.DeleteTemplateRepositoryImports(TextNRPCClient)

		GRPCProtoURL := create_files.FindProtoURL()
		TextNRPCClient = create_files.AddImport(TextNRPCClient, GRPCProtoURL)

		NRPCClientURL := create_files.FindNRPCClientURL()
		TextNRPCClient = create_files.AddImport(TextNRPCClient, NRPCClientURL)

		GRPCConstantsURL := create_files.FindGRPCConstantsURL()
		TextNRPCClient = create_files.AddImport(TextNRPCClient, GRPCConstantsURL)

		TableURL := create_files.FindModelTableURL(TableName)
		TextNRPCClient = create_files.AddImport(TextNRPCClient, TableURL)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextNRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextNRPCClient

	if config.Settings.HAS_IS_DELETED == true {
		TextNRPCClient = DeleteFuncDelete(TextNRPCClient, ModelName, Table1)
		//TextNRPCClient = DeleteFuncDeleteCtx(TextNRPCClient, ModelName, Table1)
		TextNRPCClient = DeleteFuncRestore(TextNRPCClient, ModelName, Table1)
		//TextNRPCClient = DeleteFuncRestoreCtx(TextNRPCClient, ModelName, Table1)
	}
	TextNRPCClient = DeleteFuncFind_byExtID(TextNRPCClient, ModelName, Table1)

	//замена импортов на новые URL
	TextNRPCClient = create_files.ReplaceServiceURLImports(TextNRPCClient)

	//удаление пустого импорта
	TextNRPCClient = create_files.DeleteEmptyImport(TextNRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyNRPCClient, []byte(TextNRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateTestFiles - создаёт 1 файл в папке nrpc_client
func CreateTestFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesNRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()
	DirReadyNRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateNRPCClient := DirTemplatesNRPCClient + constants.NRPC_CLIENT_TABLE_TEST_FILENAME + "_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyNRPCClient + "nrpc_" + TableName + micro.SeparatorFile()
	FilenameReadyNRPCClient := DirReadyTable + "nrpc_" + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateNRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateNRPCClient, " error: ", err)
	}
	TextNRPCClient := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextNRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextNRPCClient = create_files.DeleteTemplateRepositoryImports(TextNRPCClient)

		NRPCClientURL := create_files.FindNRPCClientURL()
		TextNRPCClient = create_files.AddImport(TextNRPCClient, NRPCClientURL)

		TableURL := create_files.FindModelTableURL(TableName)
		TextNRPCClient = create_files.AddImport(TextNRPCClient, TableURL)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextNRPCClient = strings.ReplaceAll(TextNRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextNRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextNRPCClient

	if config.Settings.HAS_IS_DELETED == true {
		TextNRPCClient = DeleteFuncTestDelete(TextNRPCClient, ModelName, Table1)
		TextNRPCClient = DeleteFuncTestRestore(TextNRPCClient, ModelName, Table1)
	}
	TextNRPCClient = DeleteFuncTestFind_byExtID(TextNRPCClient, ModelName, Table1)

	//Postgres_ID_Test = ID Minimum
	if Table1.IDMinimum != "" {
		TextFind := "const Postgres_ID_Test = "
		TextNRPCClient = strings.ReplaceAll(TextNRPCClient, TextFind+"1", TextFind+Table1.IDMinimum)
	}

	// замена ID на PrimaryKey
	TextNRPCClient = create_files.ReplacePrimaryKeyID(TextNRPCClient, Table1)

	//SkipNow()
	TextNRPCClient = create_files.AddSkipNow(TextNRPCClient, Table1)

	//замена импортов на новые URL
	TextNRPCClient = create_files.ReplaceServiceURLImports(TextNRPCClient)

	//удаление пустого импорта
	TextNRPCClient = create_files.DeleteEmptyImport(TextNRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyNRPCClient, []byte(TextNRPCClient), constants.FILE_PERMISSIONS)

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
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
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
	ok := create_files.Has_Column_ExtID_ConnectionID(Table1)
	if ok == true {
		return Otvet
	}

	//

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "TestFindByExtID")

	return Otvet
}
