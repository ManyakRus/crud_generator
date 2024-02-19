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
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc_client
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы grpc_client
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы grpc_client
		if config.Settings.NEED_CREATE_GRPC_CLIENT_TEST == true {
			err = CreateTestFiles(Table1)
			if err != nil {
				log.Error("CreateTestFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//
		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы grpc_client update
			err = CreateFilesUpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFilesUpdateEveryColumn() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы grpc_client update
			err = CreateTestFilesUpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateTestFilesUpdateEveryColumn() table: ", Table1.Name, " error: ", err)
				return err
			}

		}
	}

	return err
}

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
	TextGRPCClient = create_files.ReplacePackageName(TextGRPCClient, DirReadyTable)

	////заменим импорты
	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	//	TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)
	//
	//	ConstantsURL := create_files.FindGRPCConstantsURL()
	//	TextGRPCClient = create_files.AddImport(TextGRPCClient, ConstantsURL)
	//
	//	GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
	//	TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)
	//
	//	DBConstantsURL := create_files.FindDBConstantsURL()
	//	TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)
	//
	//	//удалим лишние функции
	//	TextGRPCClient = create_files.DeleteFuncDelete(TextGRPCClient, Table1)
	//	TextGRPCClient = create_files.DeleteFuncRestore(TextGRPCClient, Table1)
	//	TextGRPCClient = create_files.DeleteFuncFind_byExtID(TextGRPCClient, Table1)
	//}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//TextGRPCClient = create_files.ReplaceServiceURLImports(TextGRPCClient)
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		//proto
		RepositoryGRPCProtoURL := create_files.FindProtoURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCProtoURL)

		//model
		RepositoryModelURL := create_files.FindModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryModelURL)

		//grpc client
		RepositoryGRPCClientlURL := create_files.FindGRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCClientlURL)

		//nrpc client
		RepositoryNRPCClientlURL := create_files.FindNRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryNRPCClientlURL)

		//constants GRPC
		RepositoryGRPCConstantsURL := create_files.FindGRPCConstantsURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, RepositoryGRPCConstantsURL)

		//DBConstantsURL := create_files.FindDBConstantsURL()
		//TextGRPCClient = create_files.AddImport(TextGRPCClient, DBConstantsURL)

		//grpc_nrpc
		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPC_NRPC_URL)
	}

	//удалим лишние функции
	TextGRPCClient = create_files.DeleteFuncDelete(TextGRPCClient, Table1)
	TextGRPCClient = create_files.DeleteFuncRestore(TextGRPCClient, Table1)
	TextGRPCClient = create_files.DeleteFuncFind_byExtID(TextGRPCClient, Table1)

	//удаление пустого импорта
	TextGRPCClient = create_files.DeleteEmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.DeleteEmptyLines(TextGRPCClient)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCClient, []byte(TextGRPCClient), constants.FILE_PERMISSIONS)

	return err
}

// CreateTestFiles - создаёт 1 файл в папке grpc_client
func CreateTestFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCClient := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPCClient := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()

	FilenameTemplateGRPCClient := DirTemplatesGRPCClient + "grpc_client_table_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCClient + "grpc_" + TableName + micro.SeparatorFile()
	FilenameReadyGRPCClient := DirReadyTable + "grpc_" + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCClient)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCClient, " error: ", err)
	}
	TextGRPCClient := string(bytes)

	//заменим имя пакета на новое
	TextGRPCClient = create_files.ReplacePackageName(TextGRPCClient, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCClient = create_files.DeleteTemplateRepositoryImports(TextGRPCClient)

		GRPClientURL := create_files.FindGRPClientURL()
		TextGRPCClient = create_files.AddImport(TextGRPCClient, GRPClientURL)

		ModelTableName := create_files.FindModelTableURL(TableName)
		TextGRPCClient = create_files.AddImport(TextGRPCClient, ModelTableName)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCClient = strings.ReplaceAll(TextGRPCClient, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCClient = config.Settings.TEXT_MODULE_GENERATED + TextGRPCClient

	if config.Settings.HAS_IS_DELETED == true {
		TextGRPCClient = DeleteFuncTestDelete(TextGRPCClient, ModelName, Table1)
		TextGRPCClient = DeleteFuncTestRestore(TextGRPCClient, ModelName, Table1)
	}
	TextGRPCClient = DeleteFuncTestFind_byExtID(TextGRPCClient, ModelName, Table1)

	//Postgres_ID_Test = ID Minimum
	if Table1.IDMinimum != "" {
		TextFind := "const Postgres_ID_Test = "
		TextGRPCClient = strings.ReplaceAll(TextGRPCClient, TextFind+"0", TextFind+Table1.IDMinimum)
	}

	// замена ID на PrimaryKey
	TextGRPCClient = create_files.ReplacePrimaryKeyID(TextGRPCClient, Table1)

	//SkipNow()
	TextGRPCClient = create_files.AddSkipNow(TextGRPCClient, Table1)

	//замена импортов на новые URL
	TextGRPCClient = create_files.ReplaceServiceURLImports(TextGRPCClient)

	//удаление пустого импорта
	TextGRPCClient = create_files.DeleteEmptyImport(TextGRPCClient)

	//удаление пустых строк
	TextGRPCClient = create_files.DeleteEmptyLines(TextGRPCClient)

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

// CreateFilesUpdateEveryColumn - создаёт 1 файл в папке grpc_client
func CreateFilesUpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPC_Client := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPC_Client := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()

	FilenameTemplateGRPC_Client := DirTemplatesGRPC_Client + constants.GRPC_CLIENT_TABLE_UPDATE_FUNC_FILENAME
	DirReadyTable := DirReadyGRPC_Client
	FilenameReadyGRPC_ClientUpdate := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_update.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPC_Client)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPC_Client, " error: ", err)
	}
	TextGRPC_ClientUpdateFunc := string(bytes)

	TextGRPC_Client := "package " + config.Settings.PREFIX_CLIENT_GRPC + TableName + "\n\n"
	TextGRPC_Client = TextGRPC_Client + `import (
	"context"
	"time"
	"github.com/ManyakRus/starter/log"
)

`

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		ProtoURL := create_files.FindProtoURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ProtoURL)

		GRPCClientURL := create_files.FindGRPClientURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCClientURL)

		GRPCConstantsURL := create_files.FindGRPCConstantsURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCConstantsURL)

		GRPC_NRPC_URL := create_files.Find_GRPC_NRPC_URL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPC_NRPC_URL)

		NRPCClientURL := create_files.FindNRPClientURL()
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, NRPCClientURL)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ModelTableURL)

		//TextGRPC_Client = create_files.ConvertIdToAlias(TextGRPC_Client, Table1)
	}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumn(TextGRPC_ClientUpdateFunc, Table1)
	// пустой файл не нужен
	if TextUpdateEveryColumn == "" {
		return err
	}

	//ModelName := Table1.NameGo
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPC_Client = TextGRPC_Client + TextUpdateEveryColumn

	TextGRPC_Client = config.Settings.TEXT_MODULE_GENERATED + TextGRPC_Client

	//удаление пустого импорта
	TextGRPC_Client = create_files.DeleteEmptyImport(TextGRPC_Client)
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPC_Client = create_files.CheckAndAddImportTime_FromText(TextGRPC_Client)
		TextGRPC_Client = create_files.CheckAndAddImportTimestamp_FromText(TextGRPC_Client)
		TextGRPC_Client = create_files.CheckAndAddImportAlias(TextGRPC_Client)
	}

	//удаление пустых строк
	TextGRPC_Client = create_files.DeleteEmptyLines(TextGRPC_Client)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPC_ClientUpdate, []byte(TextGRPC_Client), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumn - возвращает текст для всех таблиц
func FindTextUpdateEveryColumn(TextGRPC_ClientUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumn1(TextGRPC_ClientUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumn1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumn1(TextGRPC_ClientUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPC_ClientUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	TextRequest, TextRequestFieldName, _ := create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "Request.")

	ColumnNameGolang := create_files.FindTextConvertGolangTypeToProtobufType(Table1, Column1, "m.")

	_, IDTypeGo := create_files.FindPrimaryKeyNameTypeGo(Table1)

	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL+"_Update", ModelName+"_"+FuncName)
	Otvet = strings.ReplaceAll(Otvet, " Update ", " "+FuncName+" ")
	Otvet = strings.ReplaceAll(Otvet, " Update(", " "+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "m.ColumnName", ColumnNameGolang)
	Otvet = strings.ReplaceAll(Otvet, " m.ID", " "+IDTypeGo+"(m.ID)")
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.FieldName", "Request."+TextRequestFieldName)

	return Otvet
}

// CreateTestFilesUpdateEveryColumn - создаёт 1 файл в папке grpc_client
func CreateTestFilesUpdateEveryColumn(Table1 *types.Table) error {
	var err error

	TableName := strings.ToLower(Table1.Name)

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPC_Client := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile()
	DirReadyGRPC_Client := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT + micro.SeparatorFile() + config.Settings.PREFIX_CLIENT_GRPC + TableName + micro.SeparatorFile()

	FilenameTemplateGRPC_Client := DirTemplatesGRPC_Client + constants.GRPC_CLIENT_TABLE_UPDATE_FUNC_TEST_FILENAME
	DirReadyTable := DirReadyGRPC_Client
	FilenameReadyGRPC_ClientUpdate := DirReadyTable + config.Settings.PREFIX_CLIENT_GRPC + TableName + "_update_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPC_Client)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPC_Client, " error: ", err)
	}
	TextGRPC_ClientUpdateFunc := string(bytes)

	TextGRPC_Client := "package " + config.Settings.PREFIX_CLIENT_GRPC + TableName + "\n\n"
	TextGRPC_Client = TextGRPC_Client + `import (
	"testing"
	"github.com/ManyakRus/starter/config_main"
)

`

	//заменим импорты
	//if config.Settings.USE_DEFAULT_TEMPLATE == true {
	GRPCClientURL := create_files.FindGRPClientURL()
	TextGRPC_Client = create_files.AddImport(TextGRPC_Client, GRPCClientURL)

	ModelTableURL := create_files.FindModelTableURL(TableName)
	TextGRPC_Client = create_files.AddImport(TextGRPC_Client, ModelTableURL)

	//TextGRPC_Client = create_files.ConvertIdToAlias(TextGRPC_Client, Table1)
	//}

	//создание текста
	TextUpdateEveryColumn := FindTextUpdateEveryColumnTest(TextGRPC_ClientUpdateFunc, Table1)
	// пустой файл не нужен
	if TextUpdateEveryColumn == "" {
		return err
	}
	//ModelName := Table1.NameGo
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	//TextGRPC_Client = strings.ReplaceAll(TextGRPC_Client, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPC_Client = TextGRPC_Client + TextUpdateEveryColumn

	TextGRPC_Client = config.Settings.TEXT_MODULE_GENERATED + TextGRPC_Client

	//SkipNow() если нет строк в БД
	TextGRPC_Client = create_files.AddSkipNow(TextGRPC_Client, Table1)

	//удаление пустого импорта
	TextGRPC_Client = create_files.DeleteEmptyImport(TextGRPC_Client)

	//удаление пустых строк
	TextGRPC_Client = create_files.DeleteEmptyLines(TextGRPC_Client)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPC_ClientUpdate, []byte(TextGRPC_Client), constants.FILE_PERMISSIONS)

	return err
}

// FindTextUpdateEveryColumnTest - возвращает текст для всех таблиц
func FindTextUpdateEveryColumnTest(TextGRPC_ClientUpdateFunc string, Table1 *types.Table) string {
	Otvet := ""

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//найдём новый текст для каждой таблицы
	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("FindTextProtoTable1_UpdateEveryColumn() Table1.MapColumns[key1] = false")
		}
		if create_files.Is_NotNeedUpdate_Сolumn(Column1) == true {
			continue
		}

		TextColumn1 := FindTextUpdateEveryColumnTest1(TextGRPC_ClientUpdateFunc, Table1, Column1)
		Otvet = Otvet + TextColumn1 + "\n\n"

	}

	return Otvet
}

// FindTextUpdateEveryColumnTest1 - возвращает текст для одной таблицы
func FindTextUpdateEveryColumnTest1(TextGRPC_ClientUpdateFunc string, Table1 *types.Table, Column1 *types.Column) string {
	Otvet := TextGRPC_ClientUpdateFunc

	ModelName := Table1.NameGo
	ColumnName := Column1.NameGo
	FuncName := "Update_" + ColumnName
	TextRequest, TextRequestFieldName, _ := create_files.FindTextProtobufRequest_ID_Type(Table1, Column1, "Request.")
	DefaultValue := create_files.FindTextDefaultValue(Column1.TypeGo)

	Otvet = strings.ReplaceAll(Otvet, "TestCrud_GRPC_Update(", "TestCrud_GRPC_"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	Otvet = strings.ReplaceAll(Otvet, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestId", "grpc_proto."+TextRequest)
	Otvet = strings.ReplaceAll(Otvet, "ColumnName", ColumnName)
	Otvet = strings.ReplaceAll(Otvet, "Request.ID", "Request."+TextRequestFieldName)
	Otvet = strings.ReplaceAll(Otvet, "TestRead(", "Test"+FuncName+"(")
	Otvet = strings.ReplaceAll(Otvet, "error: ID =0", "error: "+ColumnName+" ="+DefaultValue)
	Otvet = strings.ReplaceAll(Otvet, "_Update(", "_"+FuncName+"(")

	return Otvet
}
