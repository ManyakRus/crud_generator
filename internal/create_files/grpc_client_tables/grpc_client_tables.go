package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"sort"
	"strings"
)

// CreateAllFiles - создаёт все файлы в папке grpc_client
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

		//файлы grpc_client
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы grpc_client
		if config.Settings.NEED_CREATE_GRPC_CLIENT_TEST == true {
			err = CreateFiles_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//
		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы grpc_client update
			err = CreateFiles_UpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFiles_UpdateEveryColumn() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы grpc_client update
			err = CreateFiles_UpdateEveryColumn_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_UpdateEveryColumn_Test() table: ", Table1.Name, " error: ", err)
				return err
			}

		}
		//NEED_CREATE_CACHE_API
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_FILES == true {
				err = CreateFiles_GRPC_Client_Cache(Table1)
				if err != nil {
					log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

			//тестовые файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_TEST_FILES == true {
				err = CreateFiles_GRPC_Client_Cache_Test(Table1)
				if err != nil {
					log.Error("CreateFiles_Test() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

		}

		//
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
			log.Error("CreateFiles_FindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		err = CreateFiles_FindModelBy_Test(MapAll, Table1)
		if err != nil {
			log.Error("CreateFiles_FindModelBy_Test() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		if config.Settings.NEED_CREATE_READOBJECT == true {
			err = CreateFiles_ReadObject(Table1)
			if err != nil {
				log.Error("CreateFiles_ReadObject() table: ", Table1.Name, " error: ", err)
				return err
			}

			err = CreateFiles_ReadObject_Test(Table1)
			if err != nil {
				log.Error("CreateFiles_ReadObject_Test() table: ", Table1.Name, " error: ", err)
				return err
			}

		}

	}

	return err
}

// Replace_PrimaryKey_ID - заменяет "m.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyM_ID(Text string, Table1 *types.Table) string {
	Otvet := Text

	VariableName := "m"

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"

		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	//Otvet = strings.ReplaceAll(Otvet, "\tRequest.ID = int64(ID)", TextRequestIDInt64ID)
	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)

	//заменим ID := Request.ID
	Otvet = strings.ReplaceAll(Otvet, "\tID := Request.ID\n", TextIDRequestID)

	return Otvet
}

// Replace_PrimaryKeyOtvetID - заменяет "Otvet.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyOtvetID(Text string, Table1 *types.Table) string {
	Otvet := Text

	VariableName := "Otvet"

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDVariableID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDVariableID = TextOtvetIDVariableID + "\t" + VariableName + "." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + VariableName + "." + Column1.NameGo + " = m." + Column1.NameGo + "\n"

		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + VariableName + "." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = AliasFromInt(ID)", TextOtvetIDAliasID)
	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = ProtoFromInt(m.ID)", TextRequestIDmID)
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = int64(ID)", TextRequestIDInt64ID)
	//Otvet = strings.ReplaceAll(Otvet, "\tOtvet.ID = "+VariableName+".ID\n", TextOtvetIDVariableID)
	Otvet = strings.ReplaceAll(Otvet, " IntFromAlias("+VariableName+".ID) == 0", TextMID0)
	//Otvet = strings.ReplaceAll(Otvet, "\tm2.ID = int64(m.ID)", TextM2ID)
	//Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) == 0", TextIfMId)
	//Otvet = strings.ReplaceAll(Otvet, "int64(m.ID) != 0", TextIfMIdNot0)
	Otvet = strings.ReplaceAll(Otvet, " IntFromAlias("+VariableName+".ID) != 0", " "+TextIfMIdNot0)

	//VariableName = "Request"
	//Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = ProtoFromInt(m.ID)", TextRequestIDmID)
	//Otvet = strings.ReplaceAll(Otvet, " IntFromAlias("+VariableName+".ID) == 0", TextMID0)

	//m
	Otvet = strings.ReplaceAll(Otvet, "\tOtvet.ID = m.ID\n", TextOtvetIDmID)

	return Otvet
}

// Replace_PrimaryKeyRequest_ID - заменяет "m.ID" на название колонки PrimaryKey
func Replace_PrimaryKeyRequest_ID(Text string, Table1 *types.Table) string {
	Otvet := Text

	VariableName := "Request"

	//сортировка по названию таблиц
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	TextOtvetIDAliasID := ""
	TextIfMId := ""
	TextIfMIdNot0 := ""
	TextM2ID := ""
	TextIDRequestID := ""
	TextOtvetIDID := ""
	TextRequestIDmID := ""
	TextRequestIDInt64ID := ""
	TextOtvetIDmID := ""
	TextMID0 := ""
	TextOR := ""
	for _, key1 := range keys {
		Column1, _ := Table1.MapColumns[key1]
		if Column1.IsPrimaryKey != true {
			continue
		}
		TextOtvetIDID = TextOtvetIDID + "\t" + VariableName + "." + Column1.NameGo + " = " + Column1.NameGo + "\n"
		RequestColumnName := create_files.Find_RequestFieldName(Table1, Column1)
		Value, GolangCode := create_files.Convert_ProtobufVariableToGolangVariable(Table1, Column1, "Request.")
		if GolangCode == "" {
			TextIDRequestID = TextIDRequestID + "\t" + Column1.NameGo + " := " + Value + "\n"
		} else {
			TextIDRequestID = TextIDRequestID + "\t" + GolangCode + "\n"
		}
		TextM := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "m")
		TextRequestIDmID = TextRequestIDmID + "\t" + VariableName + "." + RequestColumnName + " = " + TextM + "\n"
		TextInt64ID := create_files.Convert_GolangVariableToProtobufVariableID(Table1, Column1, "")
		TextRequestIDInt64ID = TextRequestIDInt64ID + "\t" + VariableName + "." + RequestColumnName + " = " + TextInt64ID + "\n"
		TextOtvetIDmID = TextOtvetIDmID + "\t" + "Otvet." + Column1.NameGo + " = " + VariableName + "." + Column1.NameGo + "\n"

		DefaultValue := create_files.FindText_DefaultValue(Column1.TypeGo)

		TextM2ID = TextM2ID + "\t" + "m2." + Column1.NameGo + " = " + "m." + Column1.NameGo + "\n"
		TextIfMId = TextIfMId + TextOR + "m." + Column1.NameGo + " == " + DefaultValue
		TextIfMIdNot0 = TextIfMIdNot0 + TextOR + "m." + Column1.NameGo + " != " + DefaultValue

		TextMID0 = TextMID0 + TextOR + " (" + VariableName + "." + Column1.NameGo + " == " + DefaultValue + ")"
		TextAlias := create_files.Convert_IDToAlias(Table1, Column1, Column1.NameGo)
		TextOtvetIDAliasID = TextOtvetIDAliasID + "\t" + VariableName + "." + Column1.NameGo + " = " + TextAlias + "\n"
		TextOR = " || "
	}

	Otvet = strings.ReplaceAll(Otvet, "\t"+VariableName+".ID = ProtoFromInt(m.ID)", TextRequestIDmID)
	Otvet = strings.ReplaceAll(Otvet, "\tRequest.ID = int64(ID)", TextRequestIDInt64ID)

	return Otvet
}

// Replace_NRPC_CLIENT - включает NRPC, заменяет //Response, err = nrpc_client. на Response, err = nrpc_client.
func Replace_NRPC_CLIENT(Text string) string {
	Otvet := Text

	Otvet = strings.ReplaceAll(Otvet, "//Response, err = nrpc_client.", "Response, err = nrpc_client.")
	Otvet = strings.ReplaceAll(Otvet, "//_, err = nrpc_client.Client.", "_, err = nrpc_client.Client.")

	return Otvet
}

// Replace_RequestExtID - заменяет RequestExtID{} на Request_Int64_String{}
func Replace_RequestExtID(TextGRPCServer string, Table1 *types.Table) string {
	Otvet := TextGRPCServer

	//если нет таких колонок - ничего не делаем
	if create_files.Has_Column_ExtID_ConnectionID(Table1) == false {
		return Otvet
	}

	//если обе колонки Int64 - ничего не делаем
	if create_files.Has_Column_ExtID_ConnectionID_Int64(Table1) == true {
		return Otvet
	}

	//
	ColumnExtID := create_files.FindColumn_ExtID(Table1)
	if ColumnExtID == nil {
		return Otvet
	}

	//
	if ColumnExtID.TypeGo != "string" {
		return Otvet
	}

	//
	Otvet = strings.ReplaceAll(Otvet, "grpc_proto.RequestExtID", "grpc_proto.RequestExtIDString")
	Otvet = strings.ReplaceAll(Otvet, ".ExtID == 0", `.ExtID == ""`)

	return Otvet
}
