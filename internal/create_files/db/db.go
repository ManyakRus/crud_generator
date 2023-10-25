package db

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/constants"
	"github.com/ManyakRus/crud_generator/internal/create_files/model"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strings"
)

// CreateDBFiles - создаёт все файлы в папке db
func CreateDBFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, table1 := range MapAll {
		//файлы db
		err = CreateDBFiles1(table1)
		if err != nil {
			log.Error("CreateDBFiles1() table: ", table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы db
		err = CreateDBTestFiles1(table1)
		if err != nil {
			log.Error("CreateDBTestFiles1() table: ", table1.Name, " error: ", err)
			return err
		}
	}

	return err
}

// CreateDBFiles1 - создаёт 1 файл в папке db
func CreateDBFiles1(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()
	DirReadyDB := DirReady + "pkg" + micro.SeparatorFile() + "db" + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + "db.go_"
	TableName := strings.ToLower(Table1.Name)
	FilenameReadyDB := DirReadyDB + TableName + micro.SeparatorFile() + TableName + ".go"

	//создадим каталог
	DirTable := DirReadyDB + TableName
	ok, err := micro.FileExists(DirTable)
	if ok == false {
		err = os.Mkdir(DirTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
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
	TextDB = model.DeleteFuncFind_byExtID(TextDB, ModelName, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// CreateDBTestFiles1 - создаёт 1 файл в папке db
func CreateDBTestFiles1(Table1 *types.Table) error {
	var err error

	//чтение файлов
	DirBin := micro.ProgramDir_bin()
	DirTemplates := DirBin + constants.FolderTemplates + micro.SeparatorFile()
	DirReady := DirBin + constants.FolderReady + micro.SeparatorFile()
	DirTemplatesDB := DirTemplates + config.Settings.TEMPLATE_FOLDERNAME_DB + micro.SeparatorFile()
	DirReadyDB := DirReady + "pkg" + micro.SeparatorFile() + "db" + micro.SeparatorFile()

	FilenameTemplateDB := DirTemplatesDB + "db_test.go_"
	TableName := strings.ToLower(Table1.Name)
	FilenameReadyDB := DirReadyDB + TableName + micro.SeparatorFile() + TableName + "_test.go"

	//создадим каталог
	DirTable := DirReadyDB + TableName
	ok, err := micro.FileExists(DirTable)
	if ok == false {
		err = os.Mkdir(DirTable, 0777)
		if err != nil {
			log.Panic("Mkdir() ", DirTable, " error: ", err)
		}
	}

	bytes, err := os.ReadFile(FilenameTemplateDB)
	if err != nil {
		log.Panic("ReadFile() ", FilenameTemplateDB, " error: ", err)
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
	TextDB = model.DeleteFuncFind_byExtID(TextDB, ModelName, Table1)

	//запись файла
	err = os.WriteFile(FilenameReadyDB, []byte(TextDB), constants.FILE_PERMISSIONS)

	return err
}

// DeleteFuncDelete - удаляет функцию Delete()
func DeleteFuncDelete(TextModel, ModelName string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	TextFind := "\n// Delete "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncRestore - удаляет функцию Restore()
func DeleteFuncRestore(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	TextFind := "\n// Restore "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncDeleteCtx - удаляет функцию Delete_ctx()
func DeleteFuncDeleteCtx(TextModel, ModelName string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	TextFind := "\n// Delete_ctx "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncRestoreCtx - удаляет функцию Restore_ctx()
func DeleteFuncRestoreCtx(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	TextFind := "\n// Restore_ctx "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncFind_byExtID(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

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

	TextFind := "\n// Find_ByExtID "
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncTestDelete - удаляет функцию Delete()
func DeleteFuncTestDelete(TextModel, ModelName string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	TextFind := "\nfunc TestDelete("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncTestRestore - удаляет функцию Restore()
func DeleteFuncTestRestore(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

	_, ok := Table1.MapColumns["is_deleted"]
	if ok == true {
		return Otvet
	}

	TextFind := "\nfunc TestRestore("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}

// DeleteFuncFind_byExtID - удаляет функцию Find_ByExtID()
func DeleteFuncTestFind_byExtID(TextModel, Modelname string, Table1 *types.Table) string {
	Otvet := TextModel

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

	TextFind := "func TestFind_ByExtID("
	pos1 := strings.Index(Otvet, TextFind)
	if pos1 < 0 {
		return Otvet
	}
	s2 := Otvet[pos1+1:]

	posEnd := strings.Index(s2, "\n}")
	if posEnd < 0 {
		return Otvet
	}

	Otvet = Otvet[:pos1-1] + Otvet[pos1+posEnd+3:]

	return Otvet
}
