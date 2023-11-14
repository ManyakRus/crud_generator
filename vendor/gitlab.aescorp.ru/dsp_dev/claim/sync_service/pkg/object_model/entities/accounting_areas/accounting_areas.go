package accounting_areas

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// AccountingArea - модель для таблицы accounting_areas: Области учёта.
type AccountingArea struct {
	entities.CommonStruct
	entities.NameStruct
	Code         int32 `json:"code" gorm:"column:code;default:0" db:"code"`                            //Код в СТЕК
	ConnectionID int64 `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //Соединение к БД СТЕК (ИД)

}
