package server_grpc_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc_server
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	MassAll := micro.MassFrom_Map(MapAll)
	for _, Table1 := range MassAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGood_Table(Table1)
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
			err = CreateFiles_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//UPDATE_EVERY_COLUMN
		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы grpc_server update
			err = CreateFiles_UpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы grpc_server update
			if config.Settings.NEED_CREATE_GRPC_SERVER_TEST == true {
				err = CreateFiles_UpdateEveryColumn_Test(Table1)
				if err != nil {
					log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

		}

		//NEED_CREATE_CACHE_API
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_FILES == true {
				err = CreateFiles_Cache(Table1)
				if err != nil {
					log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

			//тестовые файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_TEST_FILES == true {
				err = CreateFiles_Cache_Test(Table1)
				if err != nil {
					log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
					return err
				}
			}
		}

		//FindBy
		err = CreateFiles_FindBy(Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_FindBy_Test(Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//FindMassBy
		err = CreateFiles_FindMassBy(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_FindMassBy_Test(Table1)
		if err != nil {
			log.Error("CreateFiles_FindMassBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//ReadAll
		err = CreateFiles_ReadAll(Table1)
		if err != nil {
			log.Error("CreateFiles_ReadAll() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_ReadAll_Test(Table1)
		if err != nil {
			log.Error("CreateFiles_ReadAll_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//FindModelBy
		err = CreateFiles_FindModelBy(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFiles_FindModelBy_Test(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_FindBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

	}
	return err
}

// Replace_Model_ID_Test - заменяет текст "const LawsuitStatusType_ID_Test = 0" на нужный ИД
func Replace_Model_ID_Test(Text string, Table1 *types.Table) string {
	Otvet := Text

	//if Table1.PrimaryKeyColumnsCount == 1 {
	//	PrimaryKeyColumn := Find_PrimaryKeyColumn(Table1)
	//	if PrimaryKeyColumn == nil {
	//		return Otvet
	//	}
	//
	//	Otvet = Replace_Model_ID_Test1(Otvet, Table1, PrimaryKeyColumn)
	//} else {
	Otvet = Replace_Model_ID_Test_ManyPK(Otvet, Table1)
	//}

	return Otvet
}

// Replace_Model_ID_Test_ManyPK - заменяет текст "const Postgres_ID_Test = 0" на нужные ИД, для много колонок PrimaryKey
func Replace_Model_ID_Test_ManyPK(Text string, Table1 *types.Table) string {
	Otvet := Text

	MassPK := create_files.Find_PrimaryKeyColumns(Table1)
	if len(MassPK) == 0 {
		return Otvet
	}

	//заменим const Postgres_ID_Test = 0
	TextFind := "const LawsuitStatusType_ID_Test = 0\n"
	TextNew := ""
	for _, Column1 := range MassPK {
		TextNew = TextNew + create_files.Replace_Model_ID_Test1(TextFind, Table1, Column1)
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим Request.ID = LawsuitStatusType_ID_Test
	TextFind = "\tRequest.ID = LawsuitStatusType_ID_Test\n"
	TextNew = ""
	for _, Column1 := range MassPK {
		Name := strings.ToUpper(Column1.NameGo)
		VariableName := Table1.NameGo + "_" + Name + "_Test"
		//Text1 := Convert_GolangVariableToProtobufVariable(Table1, Column1, VariableName)
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		TextNew = TextNew + "\tRequest." + RequestColumnName + " = " + VariableName + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	//заменим Request.ID = LawsuitStatusType_ID_Test
	TextFind = "\tRequest2.ID = LawsuitStatusType_ID_Test\n"
	TextNew = ""
	for _, Column1 := range MassPK {
		Name := strings.ToUpper(Column1.NameGo)
		VariableName := Table1.NameGo + "_" + Name + "_Test"
		//Text1 := Convert_GolangVariableToProtobufVariable(Table1, Column1, VariableName)
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		TextNew = TextNew + "\tRequest2." + RequestColumnName + " = " + VariableName + "\n"
	}
	Otvet = strings.ReplaceAll(Otvet, TextFind, TextNew)

	return Otvet
}

// Convert_RequestIdToAlias - заменяет ID на Alias
func Convert_RequestIdToAlias(Text string, Table1 *types.Table) string {
	Otvet := Text

	TableName := Table1.Name
	IDName, _ := create_files.Find_PrimaryKeyNameType(Table1)
	TextConvert, ok := types.MapConvertID[TableName+"."+IDName]
	if ok == false {
		return Otvet
	}

	Otvet = strings.ReplaceAll(Otvet, "Request.ID", TextConvert+"(Request.ID)")
	if TextConvert[:6] != "alias." {
		return Otvet
	}

	Otvet = create_files.CheckAndAdd_ImportAlias(Otvet)

	return Otvet
}
