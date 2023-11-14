package completed_months

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// CompletedMonth - модель для таблицы completed_months: Закрытые месяцы.
type CompletedMonth struct {
	entities.CommonStruct
	AccountingAreaID int64     `json:"accounting_area_id" gorm:"column:accounting_area_id;default:0" db:"accounting_area_id"` //Область учёта (ИД)
	BillingMonth     time.Time `json:"billing_month" gorm:"column:billing_month;default:null" db:"billing_month"`             //Дата начала последнего месяца, в котором запрещены изменения
	ConnectionID     int64     `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`                //Соединение к БД СТЕК (ИД)

}
