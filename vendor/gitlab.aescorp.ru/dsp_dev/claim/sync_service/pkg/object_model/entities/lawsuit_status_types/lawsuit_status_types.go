package lawsuit_status_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LawsuitStatusType Статусы дел (справочник).
type LawsuitStatusType struct {
	entities.CommonStruct
	entities.NameStruct
	Code string `json:"code" gorm:"column:code;default:0"`
}