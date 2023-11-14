package lawsuit_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LawsuitType - модель для таблицы lawsuit_types: Типы исков.
type LawsuitType struct {
	entities.CommonStruct
	entities.NameStruct
	Code string `json:"code" gorm:"column:code;default:\"\"" db:"code"` //Код типа искового заявления

}
