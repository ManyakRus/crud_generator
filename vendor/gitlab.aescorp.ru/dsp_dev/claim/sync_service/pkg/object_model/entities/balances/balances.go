package balances

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// Balance - модель для таблицы balances: Сальдо договора.
type Balance struct {
	entities.CommonStruct
	BillingMonth      time.Time `json:"billing_month" gorm:"column:billing_month;default:null" db:"billing_month"`                //Месяц задолженности
	ConnectionID      int64     `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`                   //Соединение к БД СТЕК (ИД)
	ContractID        int64     `json:"contract_id" gorm:"column:contract_id;default:0" db:"contract_id"`                         //Договор (ИД)
	DebtTypeID        int64     `json:"debt_type_id" gorm:"column:debt_type_id;default:0" db:"debt_type_id"`                      //Вид задолженности (ИД)
	DocumentAt        time.Time `json:"document_at" gorm:"column:document_at;default:null" db:"document_at"`                      //Дата документа реализации, или оплаты при переплате
	DocumentInvoiceID int64     `json:"document_invoice_id" gorm:"column:document_invoice_id;default:0" db:"document_invoice_id"` //Документ Счет-фактура (ИД)
	DocumentPaymentID int64     `json:"document_payment_id" gorm:"column:document_payment_id;default:0" db:"document_payment_id"` //Документ-оплаты (ИД)
	Sum               float64   `json:"sum" gorm:"column:sum;default:0" db:"sum"`                                                 //Сумма задолженности

}
