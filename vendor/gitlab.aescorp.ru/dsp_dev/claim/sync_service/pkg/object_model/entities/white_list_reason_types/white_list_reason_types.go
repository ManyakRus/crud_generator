package white_list_reason_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// WhiteListReasonType - модель для таблицы white_list_reason_types: Причина добавления договора в "белый" список (справочник).
type WhiteListReasonType struct {
	entities.CommonStruct
	entities.NameStruct
	Code int32 `json:"code" gorm:"column:code;default:0" db:"code"` //Код

}
