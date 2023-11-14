package invoice_debt_type_stats

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
	"time"
)

// InvoiceDebtTypeStat - модель для таблицы invoice_debt_type_stats: История передачи Вид задолженности по СФ в СТЕК
type InvoiceDebtTypeStat struct {
	entities.CommonStruct
	DebtTypeID int64           `json:"debt_type_id" gorm:"column:debt_type_id;default:0" db:"debt_type_id"` //Вид задолженности (ИД)
	InvoiceID  alias.InvoiceId `json:"invoice_id" gorm:"column:invoice_id" db:"invoice_id"`                 //Документ СФ (ИД)
	StateAt    time.Time       `json:"state_at" gorm:"column:state_at;default:null" db:"state_at"`          //Дата установки Вид задолженности

}
