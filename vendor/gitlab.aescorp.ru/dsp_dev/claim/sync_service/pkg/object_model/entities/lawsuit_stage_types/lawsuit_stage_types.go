package lawsuit_stage_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LawsuitStageType - модель для таблицы lawsuit_stage_types: Этапы дел (справочник).
type LawsuitStageType struct {
	entities.CommonStruct
	entities.NameStruct
	Code string `json:"code" gorm:"column:code;default:\"\"" db:"code"` //Код

}
