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

	for _, table1 := range MapAll {
		//файлы grpc_server
		err = CreateFiles(table1)
		if err != nil {
			log.Error("CreateFiles() table: ", table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы grpc_server
		err = CreateTestFiles(table1)
		if err != nil {
			log.Error("CreateTestFiles() table: ", table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

// CreateFiles - создаёт 1 файл в папке grpc_server
func CreateFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + "internal" + micro.SeparatorFile() + "grpc_server" + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "grpc_server.go_"
	TableName := strings.ToLower(Table1.Name)
	DirTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirTable + TableName + ".go"

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextDB := string(bytes)

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = constants.TEXT_GENERATED + TextDB

	if config.Settings.HAS_IS_DELETED == true {
		TextDB = DeleteFuncDelete(TextDB, ModelName, Table1)
		TextDB = DeleteFuncDeleteCtx(TextDB, ModelName, Table1)
		TextDB = DeleteFuncRestore(TextDB, ModelName, Table1)
		TextDB = DeleteFuncRestoreCtx(TextDB, ModelName, Table1)
	}
	TextDB = DeleteFuncFind_byExtID(TextDB, ModelName, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// CreateTestFiles - создаёт 1 файл в папке grpc_server
func CreateTestFiles(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()
	DirTemplatesGRPCServer := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER + micro.SeparatorFile()
	DirReadyGRPCServer := DirReady + "internal" + micro.SeparatorFile() + "grpc_server" + micro.SeparatorFile()

	FilenameTemplateGRPCServer := DirTemplatesGRPCServer + "grpc_server_test.go_"
	TableName := strings.ToLower(Table1.Name)
	DirTable := DirReadyGRPCServer
	FilenameReadyGRPCServer := DirTable + TableName + "_test.go"

	bytes, err := os.ReadFile(FilenameTemplateGRPCServer)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateGRPCServer, " error: ", err)
	}
	TextDB := string(bytes)

	//создание текста
	ModelName := Table1.NameGo
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_MODEL, ModelName)
	TextDB = strings.ReplaceAll(TextDB, config.Settings.TEXT_TEMPLATE_TABLENAME, Table1.Name)
	TextDB = constants.TEXT_GENERATED + TextDB

	if config.Settings.HAS_IS_DELETED == true {
		TextDB = DeleteFuncDelete(TextDB, ModelName, Table1)
		TextDB = DeleteFuncDeleteCtx(TextDB, ModelName, Table1)
		TextDB = DeleteFuncRestore(TextDB, ModelName, Table1)
		TextDB = DeleteFuncRestoreCtx(TextDB, ModelName, Table1)
	}
	TextDB = DeleteFuncFind_byExtID(TextDB, ModelName, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyGRPCServer, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Delete ")

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

	//
	_, ok := Table1.MapColumns["ext_id"]
	if ok == true {
		return Otvet
	}

	//
	_, ok = Table1.MapColumns["connection_id"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromComment(Text, "\n// "+ModelName+"_Find_ByExtID ")

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(Text, ModelName string, Table1 *types.Table) string {
	Otvet := Text

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"Delete")

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

	//
	_, ok := Table1.MapColumns["ext_id"]
	if ok == true {
		return Otvet
	}

	//
	_, ok = Table1.MapColumns["connection_id"]
	if ok == true {
		return Otvet
	}

	Otvet = create_files.DeleteFuncFromFuncName(Otvet, "Test_server_"+ModelName+"Find_ByExtID")

	return Otvet
}
