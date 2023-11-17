package grpc_server

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

// CreateAllFiles - создаёт все файлы в папке grpc_server
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.CheckGoodTable(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы grpc_server
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы grpc_server
		if config.Settings.NEED_CREATE_GRPC_SERVER_TEST == true {
			err = CreateTestFiles(Table1)
			if err != nil {
				log.Error("CreateTestFiles() table: ", Table1.Name, " error: ", err)
				return err
			}
		}
	}
	return err
}

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

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		ProtoURL := create_files.FindProtoURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		if config.Settings.HAS_IS_DELETED == true {
			TextGRPCServer = DeleteFuncDelete(TextGRPCServer, ModelName, Table1)
			TextGRPCServer = DeleteFuncDeleteCtx(TextGRPCServer, ModelName, Table1)
			TextGRPCServer = DeleteFuncRestore(TextGRPCServer, ModelName, Table1)
			TextGRPCServer = DeleteFuncRestoreCtx(TextGRPCServer, ModelName, Table1)
		}
		TextGRPCServer = DeleteFuncFind_byExtID(TextGRPCServer, ModelName, Table1)
		TextGRPCServer = ConvertID(TextGRPCServer, Table1)

		//замена импортов на новые URL
		//TextGRPCServer = create_files.ReplaceServiceURLImports(TextGRPCServer)
		//TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		////proto
		//RepositoryGRPCProtoURL := create_files.FindProtoURL()
		//TextGRPCServer = create_files.AddImport(TextGRPCServer, RepositoryGRPCProtoURL)
		//
		////model
		//RepositoryModelURL := create_files.FindModelTableURL(TableName)
		//TextGRPCServer = create_files.AddImport(TextGRPCServer, RepositoryModelURL)
	}

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// CreateTestFiles - создаёт 1 файл в папке grpc_server
func CreateTestFiles(Table1 *types.Table) error {
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

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextGRPCServer := string(bytes)

	//заменим имя пакета на новое
	create_files.ReplacePackageName(TextGRPCServer, DirReadyTable)

	//заменим импорты
	if config.Settings.USE_DEFAULT_TEMPLATE == true {
		TextGRPCServer = create_files.DeleteTemplateRepositoryImports(TextGRPCServer)

		ModelTableURL := create_files.FindModelTableURL(TableName)
		TextGRPCServer = create_files.AddImport(TextGRPCServer, ModelTableURL)

		//ProtoURL := create_files.FindProtobufURL()
		//TextGRPCServer = create_files.AddImport(TextGRPCServer, ProtoURL)

		CrudStarterURL := create_files.FindCrudStarterURL()
		TextGRPCServer = create_files.AddImport(TextGRPCServer, CrudStarterURL)
	}

	//создание текста
	ModelName := Table1.NameGo
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextGRPCServer = strings.ReplaceAll(TextGRPCServer, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextGRPCServer = config.Settings.TEXT_MODULE_GENERATED + TextGRPCServer

	if config.Settings.HAS_IS_DELETED == true {
		TextGRPCServer = DeleteFuncTestDelete(TextGRPCServer, ModelName, Table1)
		TextGRPCServer = DeleteFuncTestRestore(TextGRPCServer, ModelName, Table1)
	}
	TextGRPCServer = DeleteFuncTestFind_byExtID(TextGRPCServer, ModelName, Table1)

	//Postgres_ID_Test = ID Minimum
	if Table1.IDMinimum != "" {
		TextFind := "const " + ModelName + "_ID_Test = "
		TextGRPCServer = strings.ReplaceAll(TextGRPCServer, TextFind+"1", TextFind+Table1.IDMinimum)
	}

	// замена ID на PrimaryKey
	TextGRPCServer = create_files.ReplacePrimaryKeyID(TextGRPCServer, Table1)

	//SkipNow()
	TextGRPCServer = create_files.AddSkipNow(TextGRPCServer, Table1)

	//замена импортов на новые URL
	//TextGRPCServer = create_files.ReplaceServiceURLImports(TextGRPCServer)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextGRPCServer), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Otvet, "\n// "+ModelName+"_Delete ")

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore ")

	return Otvet
}

// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
func DeleteFuncDeleteCtx(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Delete_ctx ")

	return Otvet
}

// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
func DeleteFuncRestoreCtx(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Restore_ctx ")

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
	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_FindByExtID ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_Delete")

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"Restore")

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
	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"_FindByExtID")

	return Otvet
}

// ConvertID - заменяет ID на Alias
func ConvertID(Text string, Table1 *types.Table) string {
	Otvet := Text

	TableName := Table1.Name
	IDName, _ := create_files.FindPrimaryKeyNameType(Table1)
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == false {
		return Otvet
	}

	Otvet = strings.ReplaceAll(Otvet, "Request.Id", TextConvert+"(Request.Id)")
	if TextConvert[:6] != "alias." {
		return Otvet
	}

	URL := create_files.FindURL_Alias()
	if URL == "" {
		return Otvet
	}

	Otvet = create_files.AddImport(Otvet, URL)

	return Otvet
}
