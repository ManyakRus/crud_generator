package channel_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// ChannelType - модель для таблицы channel_types: Каналы (справочник).
type ChannelType struct {
	entities.CommonStruct
	entities.NameStruct
	Code       int32  `json:"code" gorm:"column:code;default:0" db:"code"`                         //
	NotifierID string `json:"notifier_id" gorm:"column:notifier_id;default:\"\"" db:"notifier_id"` //

}
