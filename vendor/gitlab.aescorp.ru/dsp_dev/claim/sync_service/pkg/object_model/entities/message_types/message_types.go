package message_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// MessageType - модель для таблицы message_types: Типы сообщений
type MessageType struct {
	entities.CommonStruct
	entities.NameStruct
	Code int `json:"code" gorm:"column:code;default:0" db:"code"` //Код

}
