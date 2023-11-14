package documents

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// Document - модель для таблицы documents: Документы.
type Document struct {
	entities.CommonStruct
	entities.GroupStruct
	Analytics      string    `json:"analytics" gorm:"column:analytics;default:\"\"" db:"analytics"`                   //Аналитика (ПЕНИ и др.)
	Balance        float64   `json:"balance" gorm:"column:balance;default:0" db:"balance"`                            //удалить
	BillKindID     int64     `json:"bill_kind_id" gorm:"column:bill_kind_id;default:0" db:"bill_kind_id"`             //Вид счёта (ИД)
	BillingMonth   time.Time `json:"billing_month" gorm:"column:billing_month;default:null" db:"billing_month"`       //Месяц расчета (дата начала месяца)
	ConnectionID   int64     `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`          //Соединение к БД СТЕК (ИД)
	ContractID     int64     `json:"contract_id" gorm:"column:contract_id;default:0" db:"contract_id"`                //Договор (ИД)
	Count          int64     `json:"count" gorm:"column:count;default:0" db:"count"`                                  //Количество (КВт)
	DebtSum        float64   `json:"debt_sum" gorm:"column:debt_sum;default:0" db:"debt_sum"`                         //удалить
	DocumentAt     time.Time `json:"document_at" gorm:"column:document_at;default:null" db:"document_at"`             //Время документа
	DocumentSum    float64   `json:"document_sum" gorm:"column:document_sum;default:0" db:"document_sum"`             //Сумма документа с НДС
	DocumentTypeID int64     `json:"document_type_id" gorm:"column:document_type_id;default:0" db:"document_type_id"` //Тип документа (ИД)
	Note           string    `json:"note" gorm:"column:note;default:\"\"" db:"note"`                                  //Примечание
	Number         string    `json:"number" gorm:"column:number;default:\"\"" db:"number"`                            //Номер документа
	NumberFull     string    `json:"number_full" gorm:"column:number_full;default:\"\"" db:"number_full"`             //Полный номер документа
	PayDeadline    time.Time `json:"pay_deadline" gorm:"column:pay_deadline;default:null" db:"pay_deadline"`          //Дата необходимой оплаты по Счет-фактуре, с учётом срока оплаты
	PayFrom        time.Time `json:"pay_from" gorm:"column:pay_from;default:null" db:"pay_from"`                      //Платеж с
	PayTo          time.Time `json:"pay_to" gorm:"column:pay_to;default:null" db:"pay_to"`                            //Платеж по
	Payment        float64   `json:"payment" gorm:"column:payment;default:0" db:"payment"`                            //удалить
	Reason         string    `json:"reason" gorm:"column:reason;default:\"\"" db:"reason"`                            //Основание
	ReversalID     int64     `json:"reversal_id" gorm:"column:reversal_id;default:0" db:"reversal_id"`                //Документ сторно (ИД)

}
