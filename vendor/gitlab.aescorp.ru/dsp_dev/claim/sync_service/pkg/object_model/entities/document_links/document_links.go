package document_links

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// DocumentLink - модель для таблицы document_links:
type DocumentLink struct {
	entities.CommonStruct
	ConnectionID  int64   `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`    //Соединение к БД СТЕК (ИД)
	ContractID    int64   `json:"contract_id" gorm:"column:contract_id;default:0" db:"contract_id"`          //Договор (ИД)
	CorrectionSum float64 `json:"correction_sum" gorm:"column:correction_sum;default:0" db:"correction_sum"` //
	Document1ID   int64   `json:"document1_id" gorm:"column:document1_id;default:0" db:"document1_id"`       //Документ реализации (ИД)
	Document2ID   int64   `json:"document2_id" gorm:"column:document2_id;default:0" db:"document2_id"`       //Документ оплаты (ИД)
	LinkTypeID    int64   `json:"link_type_id" gorm:"column:link_type_id;default:0" db:"link_type_id"`       //Вид связи документов (ИД)

}
