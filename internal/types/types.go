package types

import "github.com/ManyakRus/crud_generator/pkg/dbmeta"

type Column struct {
	Name            string `json:"name"   gorm:"column:name;default:''"`
	Type            string `json:"type_name"   gorm:"column:type_name;default:''"`
	IsIdentity      bool   `json:"is_identity"   gorm:"column:is_identity;default:false"`
	IsNullable      bool   `json:"is_nullable"   gorm:"column:is_nullable;default:''"`
	Description     string `json:"description"   gorm:"column:description;default:''"`
	OrderNumber     int
	TableKey        string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnKey       string `json:"column_key"   gorm:"column:column_key;default:''"`
	NameGo          string `gorm:-`
	NameGo_translit string `gorm:-` //имя для golang, русские буквы заменены на английские
	TypeGo          string `gorm:-`
	IsPrimaryKey    bool   `json:"is_primary_key"   gorm:"column:is_primary_key;default:false"`
	IDMinimum       string
	IsGenerated     bool
}

type Table struct {
	Name string `json:"name"   gorm:"column:name;default:''"`
	//Element     *etree.Element
	MapColumns map[string]*Column
	//Column []Column
	OrderNumber            int
	NameGo                 string
	NameGo_translit        string `gorm:-` //имя для golang, русские буквы заменены на английские
	Comment                string `json:"table_comment"   gorm:"column:table_comment;default:''"`
	RowsCount              int64
	PrimaryKeyColumnsCount int
	IsView                 bool
}

type ReplaceStruct struct {
	Old string `json:"old"`
	New string `json:"new"`
}

//type FunctionsReplace struct {
//	Functions []ReplaceStruct `json:"functions"`
//}

// MapReplaceName - карта замены имени поля на другое
var MapReplaceName = make(map[string]string, 0)

// MapNullableFileds - карта полей которые могут быть null
var MapNullableFileds = make(map[string]bool, 0)

// MapConvertID - карта ИмяПоля:Тип
var MapConvertID = make(map[string]string, 0)

// MapModelCrudDeleteFunctions - карта функций которые будут удалены из файл model crud, ИмяТаблицыБД:ИмяФункцииГоу
var MapModelCrudDeleteFunctions = make(map[string]string, 0)

// MapRenameFunctions - карта функций которые будут переименованый в файлах crud
var MapRenameFunctions = make(map[string][]ReplaceStruct, 0)

// TableColumns_String - структура строк для загрузки из JSON
type TableColumns_String struct {
	TableName   string   `json:"Table"`
	ColumnNames []string `json:"Columns"`
}

// MassFindBy_String - карта функций которые будут созданы для поиска 1 строки в таблице
var MassFindBy_String = make([]TableColumns_String, 0)

// MassFindMassBy_String - карта функций которые будут созданы для поиска много строк в таблице
var MassFindMassBy_String = make([]TableColumns_String, 0)

// TableColumns - структура таблица + колонки
type TableColumns struct {
	Table   *Table
	Columns []*Column
}

// TableColumn - структура таблица + колонка
type TableColumn struct {
	Table  *Table
	Column *Column
}

// TableColumn_string - структура таблица + колонка, только строки
type TableColumn_string struct {
	TableName  string
	ColumnName string
}

// MassFindBy - карта функций которые будут созданы для поиска 1 строки в таблице
var MassFindBy = make([]TableColumns, 0)

// MassFindMassBy - карта функций которые будут созданы для поиска много строк в таблице
var MassFindMassBy = make([]TableColumns, 0)

// MapReadAll - таблицы, для которых нужна функция ReadAll()
var MapReadAll = make(map[*Table]bool, 0)

// SettingsFillFromDatabase - настройки для заполнения данных из базы данных
type SettingsFillFromDatabase struct {
	INCLUDE_TABLES       string
	EXCLUDE_TABLES       string
	NEED_USE_DB_VIEWS    bool
	SINGULAR_TABLE_NAMES bool
	MapDBTypes           map[string]*dbmeta.SQLMapping //карта соответсвий типов в базе данных и типов в golang
	//MapPrimaryKeys    map[string][]string
}

// MapPrimaryKeys - карта первичных ключей, добавленных вручную в файл primary_keys.json
var MapPrimaryKeys = make(map[string][]string, 0)

// MassFindModelBy - массив функций которые будут созданы для поиска 1 строки в связанной таблицы из foreign key, фильтр по колонке
var MassFindModelBy = make([]TableColumn, 0)
