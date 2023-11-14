package bank_account_organizations

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// BankAccountOrganization - модель для таблицы bank_account_organizations: Соответствие банка - лицевого счёта - Юр.Лица. У Юр.Лица может быть несколько счетов в разных банках.
type BankAccountOrganization struct {
	entities.CommonStruct
	AccountNumber  string `json:"account_number" gorm:"column:account_number;default:\"\"" db:"account_number"` //
	BankID         int64  `json:"bank_id" gorm:"column:bank_id;default:0" db:"bank_id"`                         //Банк (ИД)
	ConnectionID   int64  `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`       //Соединение к БД СТЕК (ИД)
	OrganizationID int64  `json:"organization_id" gorm:"column:organization_id;default:0" db:"organization_id"` //Организация (ИД)

}
