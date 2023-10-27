package types

type Column struct {
	Name        string `json:"name"   gorm:"column:name;default:''"`
	Type        string `json:"type_name"   gorm:"column:type_name;default:''"`
	IsIdentity  bool   `json:"is_identity"   gorm:"column:is_identity;default:false"`
	IsNullable  bool   `json:"is_nullable"   gorm:"column:is_nullable;default:''"`
	Description string `json:"description"   gorm:"column:description;default:''"`
	OrderNumber int
	TableKey    string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnKey   string `json:"column_key"   gorm:"column:column_key;default:''"`
	NameGo      string `gorm:-`
	TypeGo      string `gorm:-`
}

type Table struct {
	Name string `json:"name"   gorm:"column:name;default:''"`
	//Element     *etree.Element
	MapColumns map[string]Column
	//Columns []Column
	OrderNumber int
	NameGo      string
	IDMinimum   string
}

var MapReplaceName = make(map[string]string, 0)
