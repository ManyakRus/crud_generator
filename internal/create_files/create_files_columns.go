package create_files

import (
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/log"
	"sort"
	"strings"
)

// Has_Column_ExtID_ConnectionID - возвращает true если есть поля ExtId и ConnectionID, если они int64
func Has_Column_ExtID_ConnectionID(Table1 *types.Table) bool {
	Otvet := false

	//
	_, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}

	//
	_, ok = Table1.MapColumns["connection_id"]
	if ok == false {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ExtID_ConnectionID_Int64 - возвращает true если есть поля ExtId и ConnectionID, если они int64
func Has_Column_ExtID_ConnectionID_Int64(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	//
	Column1, ok = Table1.MapColumns["connection_id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_IsDeleted_Bool - возвращает true если есть поле is_deleted
func Has_Column_IsDeleted_Bool(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["is_deleted"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "bool" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_ColumnType_Time - возвращает true если есть колонка с типом время
func Has_ColumnType_Time(Table1 *types.Table) bool {
	Otvet := false

	//сортировка по названию колонок
	keys := make([]string, 0, len(Table1.MapColumns))
	for k := range Table1.MapColumns {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key1 := range keys {
		Column1, ok := Table1.MapColumns[key1]
		if ok == false {
			log.Panic("Table1.MapColumns[key1] = false")
		}
		SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
		if ok == false {
			log.Panic("GetMappings() ", Column1.Type, " error: not found")
		}
		if SQLMapping1.GoType == "time.Time" {
			Otvet = true
			break
		}
	}

	return Otvet
}

// Has_Column_ID_Int64 - возвращает true если есть поле id типа int64
func Has_Column_ID_Int64(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ExtID_Int64 - возвращает true если есть поле ext_id типа int64
func Has_Column_ExtID_Int64(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_CreatedAt_Time - возвращает true если есть поле created_at
func Has_Column_CreatedAt_Time(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["created_at"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "time.Time" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ModifiedAt_Time - возвращает true если есть поле modified_at
func Has_Column_ModifiedAt_Time(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["modified_at"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "time.Time" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_DeletedAt_Time - возвращает true если есть поле deleted_at
func Has_Column_DeletedAt_Time(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["deleted_at"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "time.Time" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_TableNameID_Int64 - возвращает true если есть поле table_name_id
func Has_Column_TableNameID_Int64(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["table_name_id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_TableRowID_Int64 - возвращает true если есть поле table_row_id
func Has_Column_TableRowID_Int64(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["table_row_id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_IsGroup_Bool - возвращает true если есть поле is_group
func Has_Column_IsGroup_Bool(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["is_group"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "bool" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_ParentID_Int64 - возвращает true если есть поле parent_id
func Has_Column_ParentID_Int64(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["parent_id"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "int64" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_Name_String - возвращает true если есть поле name
func Has_Column_Name_String(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["name"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "string" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// Has_Column_Description_String - возвращает true если есть поле description
func Has_Column_Description_String(Table1 *types.Table) bool {
	Otvet := false

	//
	Column1, ok := Table1.MapColumns["description"]
	if ok == false {
		return Otvet
	}

	//
	if Column1.TypeGo != "string" {
		return Otvet
	}

	Otvet = true
	return Otvet
}

// ----

// Has_Columns_CommonStruct - возвращает true если есть все общие структуры
func Has_Columns_CommonStruct(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_ExtID_Int64(Table1) && Has_Column_CreatedAt_Time(Table1) && Has_Column_ModifiedAt_Time(Table1) && Has_Column_DeletedAt_Time(Table1) && Has_Column_IsDeleted_Bool(Table1) && Has_Column_ID_Int64(Table1)

	return Otvet
}

// Has_Columns_NameStruct - возвращает true если есть колонки name + description
func Has_Columns_NameStruct(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_Name_String(Table1) && Has_Column_Description_String(Table1)

	return Otvet
}

// Has_Columns_Groups - возвращает true если есть колонки is_group + parent_id
func Has_Columns_Groups(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_IsGroup_Bool(Table1) && Has_Column_ParentID_Int64(Table1)

	return Otvet
}

// Has_Columns_ExtLink - возвращает true если есть колонки table_name_id + table_row_id
func Has_Columns_ExtLink(Table1 *types.Table) bool {
	Otvet := false

	Otvet = Has_Column_TableNameID_Int64(Table1) && Has_Column_TableRowID_Int64(Table1)

	return Otvet
}

// ----

// Is_Column_CommonStruct - возвращает true если это колонка ext_id, created_at, modified_at, deleted_at, id
func Is_Column_CommonStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "ext_id", "created_at", "modified_at", "deleted_at", "is_deleted", "id":
		Otvet = true
	}

	return Otvet
}

// Is_Column_NameStruct - возвращает true если это колонка name или description
func Is_Column_NameStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "name", "description":
		Otvet = true
	}

	return Otvet
}

// Is_Column_GroupsStruct - возвращает true если это колонка is_group, parent_id
func Is_Column_GroupsStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "is_group", "parent_id":
		Otvet = true
	}

	return Otvet
}

// Is_Column_ExtLinksStruct - возвращает true если это колонка table_name_id, table_row_id
func Is_Column_ExtLinksStruct(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name

	switch ColumnName {
	case "table_name_id", "table_row_id":
		Otvet = true
	}

	return Otvet
}

// Is_Common_Сolumn - возвращает true если это общая колонка: table_name_id, table_row_id, is_group, parent_id, name или description, ext_id, created_at, modified_at, deleted_at, id
func Is_Common_Сolumn(Column1 *types.Column) bool {
	Otvet := false

	Otvet = Is_Column_CommonStruct(Column1) || Is_Column_NameStruct(Column1) || Is_Column_GroupsStruct(Column1) || Is_Column_ExtLinksStruct(Column1)

	return Otvet
}

// Is_NotNeedUpdate_Сolumn - возвращает true если не нужна функция UpdateColumnNAme(), например если это общая колонка: table_name_id, table_row_id, is_group, parent_id, ext_id, created_at, modified_at, deleted_at, id
func Is_NotNeedUpdate_Сolumn(Column1 *types.Column) bool {
	Otvet := false

	Otvet = Is_Column_CommonStruct(Column1) || Is_Column_GroupsStruct(Column1) //|| Is_Column_ExtLinksStruct(Column1)

	if Is_Need_Сolumn(Column1) == false {
		Otvet = true
	}

	return Otvet
}

// Is_NotNeedSave_Сolumn_SQL - возвращает true если нужно записывать эту колонку
func Is_NeedSave_Сolumn_SQL(Column1 *types.Column) bool {
	Otvet := false

	//
	if Is_Need_Сolumn(Column1) == false {
		Otvet = true
	}

	//
	if Column1.IsGenerated == true {
		Otvet = false
	}

	return Otvet
}

// Is_NotNeedUpdate_Сolumn_SQL - возвращает true если не нужна функция UpdateColumnNAme(), например если это общая колонка: table_name_id, table_row_id, is_group, parent_id, ext_id, created_at, modified_at, deleted_at, id
func Is_NotNeedUpdate_Сolumn_SQL(Column1 *types.Column) bool {
	Otvet := false

	ColumnName := Column1.Name
	switch ColumnName {
	case "created_at":
		Otvet = true
	}

	if Is_Need_Сolumn(Column1) == false {
		Otvet = true
	}

	//
	if Column1.IsGenerated == true {
		Otvet = true
	}

	return Otvet
}

// Is_Need_Сolumn - возвращает true если эта колонка нужна
func Is_Need_Сolumn(Column1 *types.Column) bool {
	Otvet := true

	//
	if strings.HasPrefix(Column1.Name, "DELETED_") == true {
		Otvet = false
	}

	return Otvet
}

// FindColumn_ExtID - возвращает колонку ExtID
func FindColumn_ExtID(Table1 *types.Table) *types.Column {
	Column1, ok := Table1.MapColumns["ext_id"]
	if ok == false {
		return nil
	}
	return Column1
}

// FindColumn_ConnectionID - возвращает колонку ConnectionID
func FindColumn_ConnectionID(Table1 *types.Table) *types.Column {
	Column1, ok := Table1.MapColumns["connection_id"]
	if ok == false {
		return nil
	}
	return Column1
}
