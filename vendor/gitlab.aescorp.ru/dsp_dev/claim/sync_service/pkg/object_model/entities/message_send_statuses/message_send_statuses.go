package message_send_statuses

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// MessageSendStatus - модель для таблицы message_send_statuses: Статусы отправки сообщений
type MessageSendStatus struct {
	entities.CommonStruct
	entities.NameStruct
	Code        int32  `json:"code" gorm:"column:code;default:0" db:"code"`                         //Код
	FormalName  string `json:"formal_name" gorm:"column:formal_name;default:\"\"" db:"formal_name"` //
	IsDelivered bool   `json:"is_delivered" gorm:"column:is_delivered" db:"is_delivered"`           //
	NotifierID  string `json:"notifier_id" gorm:"column:notifier_id;default:\"\"" db:"notifier_id"` //

}
