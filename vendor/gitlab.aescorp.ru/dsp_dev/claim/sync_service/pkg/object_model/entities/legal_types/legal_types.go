package legal_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LegalType - модель для таблицы legal_types: Направления (справочник).
type LegalType struct {
	entities.CommonStruct
	entities.NameStruct
	IsIndividual bool `json:"is_individual" gorm:"column:is_individual" db:"is_individual"` //

}
