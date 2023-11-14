package lawsuit_reason_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LawsuitReasonType - модель для таблицы lawsuit_reason_types: Причина отбора для претензии (Справочник).
type LawsuitReasonType struct {
	entities.CommonStruct
	entities.NameStruct
	Code string `json:"code" gorm:"column:code;default:\"\"" db:"code"` //Код причины

}
