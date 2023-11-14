package contract_category_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// ContractCategoryType - модель для таблицы contract_category_types: Категории договоров (справочник).
type ContractCategoryType struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	Code         string `json:"code" gorm:"column:code;default:\"\"" db:"code"`                         //Код
	ConnectionID int64  `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //Соединение к БД СТЕК (ИД)

}
