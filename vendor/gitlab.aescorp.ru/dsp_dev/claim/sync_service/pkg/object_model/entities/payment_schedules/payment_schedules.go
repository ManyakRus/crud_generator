package payment_schedules

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// PaymentSchedule - модель для таблицы payment_schedules: График платежей по договору.
type PaymentSchedule struct {
	entities.CommonStruct
	ConnectionID int64     `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //Соединение к БД СТЕК (ИД)
	ContractID   int64     `json:"contract_id" gorm:"column:contract_id;default:0" db:"contract_id"`       //Договор (ИД)
	DateFrom     time.Time `json:"date_from" gorm:"column:date_from;default:null" db:"date_from"`          //
	DateTo       time.Time `json:"date_to" gorm:"column:date_to;default:null" db:"date_to"`                //
	Day          int32     `json:"day" gorm:"column:day;default:0" db:"day"`                               //День месяца оплаты
	Percent      int32     `json:"percent" gorm:"column:percent;default:0" db:"percent"`                   //Процент оплаты

}
