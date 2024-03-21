package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/mini_func"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/crud_generator/pkg/dbmeta"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TableColumn struct {
	TableName         string `json:"table_name"   gorm:"column:table_name;default:''"`
	ColumnName        string `json:"column_name"   gorm:"column:column_name;default:''"`
	ColumnType        string `json:"type_name"   gorm:"column:type_name;default:''"`
	ColumnIsIdentity  string `json:"is_identity"   gorm:"column:is_identity;default:''"`
	ColumnIsNullable  string `json:"is_nullable"   gorm:"column:is_nullable;default:''"`
	ColumnDescription string `json:"description"   gorm:"column:description;default:''"`
	ColumnTableKey    string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnColumnKey   string `json:"column_key"   gorm:"column:column_key;default:''"`
	TableComment      string `json:"table_comment"   gorm:"column:table_comment;default:''"`
}

type TableRowsStruct struct {
	IDMinimum sql.NullString `json:"id_min"   gorm:"column:id_min;default:0"`
	RowsCount sql.NullInt64  `json:"rows_count"   gorm:"column:rows_count;default:0"`
}

// FillMapTable - возвращает массив MassTable данными из БД
func FillMapTable() (map[string]*types.Table, error) {
	var err error
	//MassTable := make([]types.Table, 0)
	MapTable := make(map[string]*types.Table, 0)

	TextSQL := `
drop table if exists temp_keys; 
CREATE TEMPORARY TABLE temp_keys (table_from text,  column_from text, table_to text, column_to text);

------------------------------------------- Все внешние ключи ------------------------------
insert into temp_keys
SELECT 
       (select r.relname from pg_class r where r.oid = c.conrelid) as table_from,
       UNNEST((select array_agg(attname) from pg_attribute where attrelid = c.conrelid and array[attnum] <@ c.conkey)) as column_from,
       (select  r.relname from pg_class r where r.oid = c.confrelid) as table_to,
       a.attname as column_to

FROM 
	pg_constraint c 
	
join 
	pg_attribute a 
on 
	c.confrelid=a.attrelid and a.attnum = ANY(confkey)
	
WHERE 1=1
	--and c.confrelid = (select oid from pg_class where relname = 'lawsuit_invoices')
	--AND c.confrelid!=c.conrelid
;

------------------------------------------- Все таблицы и колонки ------------------------------

SELECT 
	c.table_name, 
	c.column_name,
	c.udt_name as type_name,
	c.is_identity as is_identity,
	c.is_nullable as is_nullable, 
	COALESCE(pgd.description, '') as description,
	COALESCE(keys.table_to, '') as table_key,
	COALESCE(keys.column_to, '') as column_key, 
    (SELECT obj_description(oid) FROM pg_class as r WHERE relkind = 'r' and r.oid = st.relid) as table_comment

FROM 
	pg_catalog.pg_statio_all_tables as st
	
right join 
	information_schema.columns c 
on 
	
	c.table_schema = st.schemaname
	and c.table_name   = st.relname

left join 
	pg_catalog.pg_description pgd 
on 
	pgd.objoid = st.relid
	and pgd.objsubid   = c.ordinal_position
	

LEFT JOIN --внешние ключи
	temp_keys as keys
ON
	keys.table_from = c.table_name
	and keys.column_from = c.column_name

	
LEFT JOIN --вьюхи
	INFORMATION_SCHEMA.views as v
ON
	v.table_schema = 'public'
	and v.table_name = c.table_name


where 1=1
	and c.table_schema='public'
	and v.table_name is null
	--INCLUDE_TABLES
	--EXCLUDE_TABLES
	--and c.table_name = 'lawsuit_invoices'

order by 
	table_name, 
	is_identity desc,
	column_name
	

	
`

	SCHEMA := strings.Trim(postgres_gorm.Settings.DB_SCHEMA, " ")
	if SCHEMA != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "public", SCHEMA)
	}

	if config.Settings.INCLUDE_TABLES != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "--INCLUDE_TABLES", "and c.table_name ~* '"+config.Settings.INCLUDE_TABLES+"'")
	}

	if config.Settings.EXCLUDE_TABLES != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "--EXCLUDE_TABLES", "and c.table_name !~* '"+config.Settings.EXCLUDE_TABLES+"'")
	}

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	//запрос
	//запустим все запросы отдельно
	var tx *gorm.DB
	sqlSlice := strings.Split(TextSQL, ";")
	len1 := len(sqlSlice)
	for i, TextSQL1 := range sqlSlice {
		//batch.Queue(TextSQL1)
		if i == len1-1 {
			tx = db.Raw(TextSQL1)
		} else {
			tx = db.Exec(TextSQL1)
			//rows.Close()
		}
		err = tx.Error
		if err != nil {
			log.Panic("DB.Raw() error:", err)
		}
	}

	//tx := db.Raw(TextSQL)
	//err = tx.Error
	//if err != nil {
	//	sError := fmt.Sprint("db.Raw() error: ", err)
	//	log.Panicln(sError)
	//	return MassTable, err
	//}

	//ответ в структуру
	MassTableColumn := make([]TableColumn, 0)
	tx = tx.Scan(&MassTableColumn)
	err = tx.Error
	if err != nil {
		sError := fmt.Sprint("Get_error()  error: ", err)
		log.Panicln(sError)
		return MapTable, err
	}

	//проверка 0 строк
	if tx.RowsAffected == 0 {
		sError := fmt.Sprint("db.Raw() RowsAffected =0 ")
		log.Warn(sError)
		err = errors.New(sError)
		//log.Panicln(sError)
		return MapTable, err
	}

	//заполним MapTable
	MapColumns := make(map[string]*types.Column, 0)
	OrderNumberColumn := 0
	OrderNumberTable := 0
	TableName0 := ""
	Table1 := CreateTable()
	for _, v := range MassTableColumn {
		if v.TableName != TableName0 {
			OrderNumberColumn = 0
			Table1.MapColumns = MapColumns
			MapColumns = make(map[string]*types.Column, 0)
			if TableName0 != "" {
				//MassTable = append(MassTable, Table1)
				MapTable[TableName0] = Table1
				OrderNumberTable++
			}

			//найдём имя модели golang
			TableName := v.TableName
			ModelName := create_files.FindSingularName(TableName)
			ModelName = create_files.FormatName(ModelName)

			//
			Table1 = CreateTable()
			Table1.Name = TableName
			Table1.OrderNumber = OrderNumberTable
			Table1.Comment = v.TableComment
			Table1.NameGo = ModelName

		}

		Column1 := types.Column{}
		Column1.Name = v.ColumnName
		Column1.Type = v.ColumnType
		FillNameGo(&Column1)

		//Type_go
		FillTypeGo(&Column1)
		//SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
		//if ok == false {
		//	log.Panic("GetMappings() ", Column1.Type, " error: not found")
		//}
		//Type_go := SQLMapping1.GoType
		//Column1.TypeGo = Type_go

		//
		if v.ColumnIsIdentity == "YES" {
			Column1.IsIdentity = true
		}
		if v.ColumnIsNullable == "YES" {
			Column1.IsNullable = true
		}
		Column1.Description = v.ColumnDescription
		Column1.OrderNumber = OrderNumberColumn
		Column1.TableKey = v.ColumnTableKey
		Column1.ColumnKey = v.ColumnColumnKey

		MapColumns[v.ColumnName] = &Column1
		//Table1.Columns = append(Table1.Columns, Column1)

		OrderNumberColumn++
		TableName0 = v.TableName
	}

	//
	if Table1.Name != "" {
		Table1.MapColumns = MapColumns
		MapTable[TableName0] = Table1
	}

	//FillTypeGo(MapTable)

	FillIDMinimum(MapTable)
	FillRowsCount(MapTable)

	return MapTable, err
}

func CreateTable() *types.Table {
	Otvet := &types.Table{}
	Otvet.MapColumns = make(map[string]*types.Column, 0)

	return Otvet
}

// FillIDMinimum - находим минимальный ID, для тестов с этим ID
func FillIDMinimum(MapTable map[string]*types.Table) {
	var err error

	//соединение
	db := postgres_gorm.GetConnection()
	ctxMain := contextmain.GetContext()

	for TableName, Table1 := range MapTable {
		//текст запроса
		NameID, TypeGo := FindNameTypeID(Table1)
		if NameID == "" {
			continue
		}
		DefaultValueSQL := create_files.FindTextDefaultValueSQL(TypeGo)
		TextSQL := "SELECT Min(" + NameID + ") as id_minimum FROM \"" + postgres_gorm.Settings.DB_SCHEMA + "\".\"" + TableName + "\" WHERE " + NameID + " <> " + DefaultValueSQL
		ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
		defer ctxCancelFunc()
		db.WithContext(ctx)

		//запрос
		tx := db.Raw(TextSQL)
		err = tx.Error
		if err != nil {
			log.Panic("Wrong SQL query: ", TextSQL, " error: ", err)
		}

		var IDMinimum sql.NullString
		//TableRows := TableRowsStruct{}
		tx = tx.Scan(&IDMinimum)
		err = tx.Error
		if err != nil {
			log.Panic("Wrong SQL Scan(): ", TextSQL, " error: ", err)
		}

		//
		if TypeGo == "string" {
			Table1.IDMinimum = `"` + IDMinimum.String + `"`
		} else if mini_func.IsNumberType(TypeGo) == true {
			Table1.IDMinimum = IDMinimum.String
		}
	}

}

// FillRowsCount - находим количество строк в таблице, для кэша
func FillRowsCount(MapTable map[string]*types.Table) {
	var err error

	//соединение
	db := postgres_gorm.GetConnection()
	ctxMain := contextmain.GetContext()

	for TableName, Table1 := range MapTable {
		//текст запроса
		//только Postgres SQL
		TextSQL := `
SELECT 
	reltuples::bigint AS rows_count
FROM
	pg_class
WHERE  
	oid = 'public."` + TableName + `"'::regclass
`
		ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
		defer ctxCancelFunc()
		db.WithContext(ctx)

		//запрос
		tx := db.Raw(TextSQL)
		err = tx.Error
		if err != nil {
			log.Panic("Wrong SQL query: ", TextSQL, " error: ", err)
		}

		var RowsCount sql.NullInt64
		//TableRows := TableRowsStruct{}
		tx = tx.Scan(&RowsCount)
		err = tx.Error
		if err != nil {
			log.Panic("Wrong SQL Scan(): ", TextSQL, " error: ", err)
		}

		Table1.RowsCount = RowsCount.Int64
	}

}

func FindNameTypeID(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

	for ColumnName, Column1 := range Table1.MapColumns {
		if Column1.IsIdentity == true {
			return ColumnName, Column1.TypeGo
		}
	}

	return Otvet, Type
}

//// FillTypeGo - заполняет тип golang из типа postgres
//func FillTypeGo(MapTable map[string]*types.Table) {
//
//	for _, Table1 := range MapTable {
//		for _, Column1 := range Table1.MapColumns {
//
//			SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
//			if ok == false {
//				log.Panic("GetMappings() ", Column1.Type, " error: not found")
//			}
//			Type_go := SQLMapping1.GoType
//			Column1.TypeGo = Type_go
//		}
//	}
//
//}

// FillNameGo - заполняет NameGo во все колонки
func FillNameGo(Column1 *types.Column) error {
	var err error

	ColumnName := Column1.Name
	ColumnNameGo := create_files.FormatName(ColumnName)
	Column1.NameGo = ColumnNameGo
	if ColumnNameGo == "" {
		err = errors.New("Column: " + ColumnName + " Type: " + Column1.Type + " NameGo= \"\"")
	}

	return err
}

// FillTypeGo - заполняет тип golang из типа postgres
func FillTypeGo(Column1 *types.Column) error {
	var err error

	SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
	if ok == false {
		log.Panic("GetMappings() ", Column1.Type, " error: not found")
	}

	ColumnName := Column1.Name
	Type_go := SQLMapping1.GoType
	Column1.TypeGo = Type_go
	if Type_go == "" {
		err = errors.New("Column: " + ColumnName + " Type: " + Column1.Type + " TypeGo= \"\"")
	}

	return err
}
