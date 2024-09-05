package server_grpc_tables

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

// CreateFiles - создаёт 1 файл в папке grpc_server
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "server_grpc.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + ".go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudTableURL := create_files.FindCrudTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudTableURL)

		if Table1.PrimaryKeyColumnsCount == 1 {
			TextGRPCServer = ReplaceIDRequestID_1PK(TextGRPCServer, Table1)
		}

		//замена "m.ID = AliasFromInt(ID)"
		TextGRPCServer = create_files.ReplacePrimaryKeyM_ID(TextGRPCServer, Table1)

		//замена "ID := Request.ID"
		TextGRPCServer = create_files.ReplacePrimaryKeyOtvetID(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

		//замена int64(ID) на ID
		TextGRPCServer = create_files.ReplaceIDtoID(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

		//удалим лишние функции
		TextGRPCServer = DeleteFuncDelete(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncRestore(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncFind_byExtID(TextGRPCServer, Table1)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.ConvertRequestIdToAlias(TextGRPCServer, Table1)
	}

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// CreateFilesTest - создаёт 1 файл в папке grpc_server
func CreateFilesTest(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + config.Settings.TEMPLATE_FOLDERNAME + micro.SeparatorFile()
	DirReady := DirBin + config.Settings.READY_FOLDERNAME + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "server_grpc_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirReadyTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirReadyTable + config.Settings.PREFIX_SERVER_GRPC + TableName + "_test.go"

	//создадим папку готовых файлов
	folders.CreateFolder(DirReadyTable)

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	TextGRPCServer = create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		//удалим лишние функции
		TextGRPCServer = DeleteFuncTestDelete(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncTestRestore(TextGRPCServer, Table1)
		TextGRPCServer = DeleteFuncTestFind_byExtID(TextGRPCServer, Table1)

		//добавим импорты
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudStarterURL := create_files.FindCrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)

		ConstantsURL := create_files.FindConstantsURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ConstantsURL)

		GRPClientTableURL := create_files.FindGRPCClientTableURL(Table1.Name)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, GRPClientTableURL)

		TextGRPCServer = create_files.CheckAndAddImport(TextGRPCServer, "encoding/json")

		//замена "postgres_gorm.Connect_WithApplicationName("
		TextGRPCServer = create_files.ReplaceConnect_WithApplicationName(TextGRPCServer)

		if Table1.PrimaryKeyColumnsCount > 1 {
		}

		//Postgres_ID_Test = ID Minimum
		TextGRPCServer = create_files.Replace_Model_ID_Test(TextGRPCServer, Table1)

		//замена RequestId{}
		TextGRPCServer = create_files.ReplaceTextRequestID_PrimaryKey(TextGRPCServer, Table1)

		//замена Otvet.ID = -1
		TextGRPCServer = create_files.ReplaceModelIDEqual1(TextGRPCServer, Table1)

		//добавим импорт uuid
		TextGRPCServer = create_files.CheckAndAddImportUUID_FromText(TextGRPCServer)

		//
		TextGRPCServer = create_files.ReplaceOtvetIDEqual0(TextGRPCServer, Table1)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	//Postgres_ID_Test = ID Minimum
	//if Table1.IDMinimum != "" {
	//	TextFind := "const " + ModelName + "_ID_Test = "
	//	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, TextFind+"0", TextFind+Table1.IDMinimum)
	//}

	// замена ID на PrimaryKey
	TextGRPCServer = create_files.ReplacePrimaryKeyOtvetID(TextGRPCServer, Table1)
	TextGRPCServer = create_files.ReplacePrimaryKeyM_ID(TextGRPCServer, Table1)

	//SkipNow()
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//замена импортов на новые URL
	//TextGRPCServer = create_files.ReplaceServiceURLImports(TextGRPCServer)

	//удаление пустого импорта
	TextGRPCServer = create_files.DeleteEmptyImport(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// "+ModelName+"_Delete ")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore ")

	return Otvet
}

//// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
//func DeleteFuncDeleteCtx(Text, ModelName string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Delete_ctx ")
//
//	return Otvet
//}
//
//// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
//func DeleteFuncRestoreCtx(Text, ModelName string, Table1 *types.Table) string {
//	Otvet := Text
//
//	_, ok := Table1.MapColumns["is_deleted"]
//	if ok == true {
//		return Otvet
//	}
//
//	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore_ctx ")
//
//	return Otvet
//}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID_Int64(Table1)
	if ok == true {
		return Otvet
	}

	//
	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_FindByExtID ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_Delete")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text string, Table1 *types.Table) string {
	Otvet := Text

	//проверим есть ли колонка IsDeleted
	if create_files.Has_Column_IsDeleted_Bool(Table1) == true && config.Settings.HAS_IS_DELETED == true {
		return Otvet
	}

	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"Restore")

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(Text string, Table1 *types.Table) string {
	Otvet := Text

	//если есть обе колонки - ничего не делаем
	ok := create_files.Has_Column_ExtID_ConnectionID_Int64(Table1)
	if ok == true {
		return Otvet
	}

	//
	ModelName := config.Settings.TEXT_TEMPLATE_MODEL
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_FindByExtID")

	return Otvet
}
