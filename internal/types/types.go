package types

type Column struct {
	Name         string `json:"name"   gorm:"column:name;default:''"`
	Type         string `json:"type_name"   gorm:"column:type_name;default:''"`
	IsIdentity   bool   `json:"is_identity"   gorm:"column:is_identity;default:false"`
	IsNullable   bool   `json:"is_nullable"   gorm:"column:is_nullable;default:''"`
	Description  string `json:"description"   gorm:"column:description;default:''"`
	OrderNumber  int
	TableKey     string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnKey    string `json:"column_key"   gorm:"column:column_key;default:''"`
	NameGo       string `gorm:-`
	TypeGo       string `gorm:-`
	IsPrimaryKey bool   `json:"is_primary_key"   gorm:"column:is_primary_key;default:false"`
}

type Table struct {
	Name string `json:"name"   gorm:"column:name;default:''"`
	//Element     *etree.Element
	MapColumns map[string]*Column
	//Columns []Column
	OrderNumber            int
	NameGo                 string
	IDMinimum              string
	Comment                string `json:"table_comment"   gorm:"column:table_comment;default:''"`
	RowsCount              int64
	PrimaryKeyColumnsCount int
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
