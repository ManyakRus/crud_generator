package claim_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// ClaimType - модель для таблицы claim_types: Типы исков.
type ClaimType struct {
	entities.CommonStruct
	entities.NameStruct
	Code string `json:"code" gorm:"column:code;default:\"\"" db:"code"` //Код типа искового заявления

}
