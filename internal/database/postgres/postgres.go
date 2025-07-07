package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
	"sort"
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
	IsPrimaryKey      bool   `json:"is_primary_key"   gorm:"column:is_primary_key;default:false"`
	TableIsView       bool   `json:"table_is_view"   gorm:"column:table_is_view;default:false"`
}

type TableRowsStruct struct {
	IDMinimum sql.NullString `json:"id_min"   gorm:"column:id_min;default:0"`
	RowsCount sql.NullInt64  `json:"rows_count"   gorm:"column:rows_count;default:0"`
}

// StartAll - заполняет MapTable данными из БД, и другие колонки
func StartAll(SettingsFill types.SettingsFillFromDatabase) (map[string]*types.Table, error) {
	var MapTable map[string]*types.Table
	var err error

	MapTable, err = FillMapTable1(SettingsFill)

	//заполним ID Minimum
	err = FillIDMinimum_ManyPK(MapTable)
	//err = FillIDMinimum(MapTable)
	if err != nil {
		log.Error("FillIDMinimum() error: ", err)
		return MapTable, err
	}

	//заполним Rows Count
	err = FillRowsCount(MapTable)
	if err != nil {
		log.Error("FillRowsCount() error: ", err)
		return MapTable, err
	}

	return MapTable, err
}

// FillMapTable1 - возвращает массив MassTable данными из БД
func FillMapTable1(SettingsFill types.SettingsFillFromDatabase) (map[string]*types.Table, error) {
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
	AND c.confrelid!=c.conrelid

GROUP BY	
       (select r.relname from pg_class r where r.oid = c.conrelid),
       UNNEST((select array_agg(attname) from pg_attribute where attrelid = c.conrelid and array[attnum] <@ c.conkey)),
       (select  r.relname from pg_class r where r.oid = c.confrelid),
       a.attname
;


------------------------------------------- Все primary keys ------------------------------
drop table if exists temp_primary_keys; 
CREATE TEMPORARY TABLE temp_primary_keys (table_name text,  column_name text);

insert into temp_primary_keys
select 
    ccu.table_name,
	(ccu.column_name) as column_name
       
from pg_constraint pgc
         join pg_namespace nsp on nsp.oid = pgc.connamespace
         join pg_class  cls on pgc.conrelid = cls.oid
         left join information_schema.constraint_column_usage ccu
                   on pgc.conname = ccu.constraint_name
                       and nsp.nspname = ccu.constraint_schema
WHERE 1=1
	and ccu.table_schema = 'public'
	and contype = 'p'
	
--GROUP BY
--	ccu.table_name
--HAVING sum(1)=1
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
    (SELECT obj_description(oid) FROM pg_class as r WHERE relkind = 'r' and r.oid = st.relid) as table_comment,
	CASE
		WHEN tpk.table_name is not null
		THEN true
		ELSE false END 
	    as is_primary_key,
	CASE
		WHEN v.table_name is not null
		THEN true
		ELSE false END 
	    as is_view
FROM 
	information_schema.columns c 
	
left join 
	pg_catalog.pg_statio_all_tables as st
on 
	st.schemaname = c.table_schema
	and st.relname = c.table_name

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


LEFT JOIN
	temp_primary_keys as tpk
ON 
	tpk.table_name = c.table_name
	and tpk.column_name = c.column_name

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

	if SettingsFill.INCLUDE_TABLES != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "--INCLUDE_TABLES", "and c.table_name ~* '"+SettingsFill.INCLUDE_TABLES+"'")
	}

	if SettingsFill.EXCLUDE_TABLES != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "--EXCLUDE_TABLES", "and c.table_name !~* '"+SettingsFill.EXCLUDE_TABLES+"'")
	}

	if SettingsFill.NEED_USE_DB_VIEWS == true {
		TextSQL = strings.ReplaceAll(TextSQL, "\tand v.table_name is null", "")
	}

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	//запрос
	//запустим все запросы отдельно
	tx := postgres_gorm.RawMultipleSQL(db, TextSQL)
	err = tx.Error
	if err != nil {
		log.Panic("RawMultipleSQL() error:", err)
	}
	//var tx *gorm.DB
	//sqlSlice := strings.Split(TextSQL, ";")
	//len1 := len(sqlSlice)
	//for i, TextSQL1 := range sqlSlice {
	//	//batch.Queue(TextSQL1)
	//	if i == len1-1 {
	//		tx = db.Raw(TextSQL1)
	//	} else {
	//		tx = db.Exec(TextSQL1)
	//		//rows.Close()
	//	}
	//	err = tx.Error
	//	if err != nil {
	//		log.Panic("DB.Raw() error:", err)
	//	}
	//}

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

	//trans := transliterator.NewTransliterator(nil)

	//заполним MapTable
	MapColumns := make(map[string]*types.Column, 0)
	OrderNumberColumn := 0
	OrderNumberTable := 0
	PrimaryKeyColumnsCount := 0
	TableName0 := ""
	Table1 := CreateTable()
	for _, v := range MassTableColumn {
		if v.TableName != TableName0 {
			OrderNumberColumn = 0
			Table1.MapColumns = MapColumns
			Table1.PrimaryKeyColumnsCount = PrimaryKeyColumnsCount
			MapColumns = make(map[string]*types.Column, 0)
			if TableName0 != "" {
				//MassTable = append(MassTable, Table1)
				MapTable[TableName0] = Table1
				OrderNumberTable++
			}

			//новая таблица
			//найдём имя модели golang
			TableName := v.TableName
			//ModelName := create_files.Find_SingularName(TableName)
			//ModelName = create_files.FormatName(ModelName)

			//NameGo_translit := trans.Transliterate(ModelName, "en")

			//
			TableComment := v.TableComment
			TableComment = strings.ReplaceAll(TableComment, "\n", "")
			TableComment = strings.ReplaceAll(TableComment, "\r", "")

			Table1 = CreateTable()
			Table1.Name = TableName
			Table1.OrderNumber = OrderNumberTable
			Table1.Comment = TableComment
			//Table1.NameGo = ModelName
			//Table1.NameGo_translit = NameGo_translit

			PrimaryKeyColumnsCount = 0
		}

		Column1 := types.Column{}
		Column1.Name = v.ColumnName
		Column1.Type = v.ColumnType
		//FillNameGo(&Column1)
		//
		////Type_go
		FillTypeGo(SettingsFill, &Column1)
		//SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
		//if ok == false {
		//	log.Panic("GetMappings() ", Column1.Type, " error: not found")
		//}
		//Type_go := SQLMapping1.GoType
		//Column1.TypeGo = Type_go

		//IsPrimaryKey_manual := database.Find_IsPrimaryKey_manual(SettingsFill, v.TableName, v.ColumnName)
		IsPrimaryKey := v.IsPrimaryKey //|| IsPrimaryKey_manual
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
		Column1.IsPrimaryKey = IsPrimaryKey
		if IsPrimaryKey == true {
			PrimaryKeyColumnsCount++
		}

		MapColumns[v.ColumnName] = &Column1
		//Table1.Column = append(Table1.Column, Column1)

		OrderNumberColumn++
		TableName0 = v.TableName
	}

	//последнюю таблицу заполним тут
	if Table1.Name != "" {
		Table1.MapColumns = MapColumns
		Table1.PrimaryKeyColumnsCount = PrimaryKeyColumnsCount
		MapTable[TableName0] = Table1
	}

	//FillTypeGo(MapTable)

	return MapTable, err
}

func CreateTable() *types.Table {
	Otvet := &types.Table{}
	Otvet.MapColumns = make(map[string]*types.Column, 0)

	return Otvet
}

//// FillIDMinimum - находим минимальный ID, для тестов с этим ID
//func FillIDMinimum(MapTable map[string]*types.Table) error {
//	var err error
//
//	//соединение
//	db := postgres_gorm.GetConnection()
//	ctxMain := contextmain.GetContext()
//
//	Schema := strings.Trim(postgres_gorm.Settings.DB_SCHEMA, " ")
//
//	for TableName, Table1 := range MapTable {
//		//текст запроса
//		NameID, TypeGo := FindNameType_from_PrimaryKey(Table1)
//		if NameID == "" {
//			continue
//		}
//		TextSQL := ""
//		Is_UUID_Type := create_files.Is_UUID_Type(TypeGo)
//		if Is_UUID_Type == false {
//			DefaultValueSQL := create_files.FindText_DefaultValueSQL(TypeGo)
//			TextSQL = `SELECT
//				Min("` + NameID + `") as id_minimum
//				FROM
//					"` + Schema + `"."` + TableName + `"
//				WHERE
//					"` + NameID + `" <> ` + DefaultValueSQL
//		} else {
//			TextSQL = `SELECT "` + NameID + `" as id_minimum
//				FROM "` + Schema + `"."` + TableName + `"
//				WHERE "` + NameID + `" is not null
//				ORDER BY ` + NameID + `
//				LIMIT 1`
//		}
//		ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
//		defer ctxCancelFunc()
//		db.WithContext(ctx)
//
//		//запрос
//		tx := db.Raw(TextSQL)
//		err = tx.Error
//		if err != nil {
//			log.Panic("Wrong SQL query: ", TextSQL, " error: ", err)
//		}
//
//		var IDMinimum sql.NullString
//		tx = tx.Scan(&IDMinimum)
//		err = tx.Error
//		if err != nil {
//			log.Panic("Wrong SQL Scan(): ", TextSQL, " error: ", err)
//		}
//
//		//
//		ColumnPK := create_files.Find_PrimaryKeyColumn(Table1)
//		ColumnPK.IDMinimum = IDMinimum.String
//	}
//
//	return err
//}

// FindNameType_from_PrimaryKey - возвращает наименование и тип БД для колонки PrimaryKey (ID)
func FindNameType_from_PrimaryKey(Table1 *types.Table) (string, string) {
	Otvet := ""
	Type := ""

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
		if Column1.IsPrimaryKey == true {
			return Column1.NameGo, Column1.TypeGo
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

// FillIDMinimum_ManyPK - находим минимальный ID, для тестов с этим ID, для многих Primary Key
func FillIDMinimum_ManyPK(MapTable map[string]*types.Table) error {
	var err error

	//соединение
	db := postgres_gorm.GetConnection()
	ctxMain := contextmain.GetContext()

	Schema := strings.Trim(postgres_gorm.Settings.DB_SCHEMA, " ")

	for TableName, Table1 := range MapTable {
		ColumnsPK := create_files.Find_PrimaryKeyColumns(Table1)

		Is_UUID_Type := false
		for _, Column1 := range ColumnsPK {
			Is_UUID_Type1 := create_files.Is_UUID_Type(Column1.TypeGo)
			Is_UUID_Type = Is_UUID_Type || Is_UUID_Type1
		}

		//текст запроса
		TextSQL := ""
		if Is_UUID_Type == false {
			TextSQL = `SELECT 
				`
			Comma := ""
			for _, Column1 := range ColumnsPK {
				TextSQL += Comma + `Min("` + Column1.Name + `") as "` + Column1.Name + `"`
				Comma = ","
			}

			TextSQL = TextSQL + ` 
				FROM
					"` + Schema + `"."` + TableName + `" 
				WHERE 1=1
`

			for _, Column1 := range ColumnsPK {
				DefaultValueSQL := create_files.FindText_DefaultValueSQL(Column1.TypeGo)
				TextSQL += `and ` + Column1.Name + ` <> ` + DefaultValueSQL + "\n"
			}
		} else {
			TextSQL = `SELECT 
				`
			Comma := ""
			for _, Column1 := range ColumnsPK {
				TextSQL += Comma + `"` + Column1.Name + `" as "` + Column1.Name + `"`
				Comma = ","
			}

			TextSQL = TextSQL + ` 
				FROM
					"` + Schema + `"."` + TableName + `" 
				WHERE 1=1
`

			for _, Column1 := range ColumnsPK {
				TextSQL += `and ` + Column1.Name + ` is not null ` + "\n"
			}
			TextSQL = TextSQL + `
			ORDER BY
`
			Comma = ""
			for _, Column1 := range ColumnsPK {
				TextSQL += Comma + `"` + Column1.Name + `"`
				Comma = ","
			}
			//TextSQL = `SELECT "` + NameID + `" as id_minimum
			//	FROM "` + Schema + `"."` + TableName + `"
			//	WHERE "` + NameID + `" is not null
			//	ORDER BY ` + NameID + `
			//	LIMIT 1`
		}
		ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
		defer ctxCancelFunc()
		db.WithContext(ctx)

		//запрос с разным количеством колонок
		tx := db.Raw(TextSQL)
		err = tx.Error
		if err != nil {
			log.Panic("Raw() Wrong SQL query: ", TextSQL, " error: ", err)
		}

		rows, err := tx.Rows()
		if err != nil {
			log.Panic("Rows() Wrong SQL query: ", TextSQL, " error: ", err)
		}
		has_next := rows.Next()
		if has_next == false { //нет строк
			//log.Panic("Next() Wrong SQL query: ", TextSQL, " error: ", err)
			continue
		}
		ColumnsGorm, err := rows.Columns()

		MapIDMinimum := make(map[string]string)

		values := make([]interface{}, len(ColumnsGorm))
		for i := range values {
			values[i] = new(interface{})
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}

		rows.Close()

		for i, colName := range ColumnsGorm {
			value1 := ""
			value1 = fmt.Sprint(*values[i].(*interface{}))
			if value1 == "<nil>" {
				value1 = ""
			}
			MapIDMinimum[colName] = value1
		}

		//
		for _, Column1 := range ColumnsPK {
			value := MapIDMinimum[Column1.Name]
			Column1.IDMinimum = value
		}
	}

	return err
}

// FillRowsCount - находим количество строк в таблице, для кэша
func FillRowsCount(MapTable map[string]*types.Table) error {
	var err error

	//соединение
	db := postgres_gorm.GetConnection()
	ctxMain := contextmain.GetContext()

	Schema := strings.Trim(postgres_gorm.Settings.DB_SCHEMA, " ")

	for TableName, Table1 := range MapTable {
		//текст запроса
		//только Postgres SQL
		TextSQL := `
SELECT 
	reltuples::bigint AS rows_count
FROM
	pg_class
WHERE  
	oid = '` + Schema + `."` + TableName + `"'::regclass
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

	return err
}

// FillTypeGo - заполняет тип golang из типа postgres
func FillTypeGo(SettingsFill types.SettingsFillFromDatabase, Column1 *types.Column) {
	var err error

	//SQLMapping1, ok := dbmeta.GetMappings()[Column1.Type]
	//if ok == false {
	//	log.Panic("GetMappings() ", Column1.Type, " error: not found")
	//}

	SQLMapping1, ok := SettingsFill.MapDBTypes[Column1.Type]
	if ok == false {
		log.Panic("FillTypeGo() Column1.Type: ", Column1.Type, ", error: not found")
	}

	ColumnName := Column1.Name
	Type_go := SQLMapping1.GoType
	Column1.TypeGo = Type_go
	if Type_go == "" {
		err = errors.New("FillTypeGo() error: Column: " + ColumnName + ", Type: " + Column1.Type + ", TypeGo= \"\"")
		log.Panic(err)
	}

	return
}
