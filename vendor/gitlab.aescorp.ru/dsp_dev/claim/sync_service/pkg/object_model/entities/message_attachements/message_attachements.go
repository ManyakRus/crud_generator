package message_attachements

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// MessageAttachement - модель для таблицы message_attachements: Вложения файлов в (емайл) сообщения
type MessageAttachement struct {
	entities.CommonStruct
	FilesID    int64 `json:"files_id" gorm:"column:files_id;default:0" db:"files_id"`          //Файл (ИД)
	MessagesID int64 `json:"messages_id" gorm:"column:messages_id;default:0" db:"messages_id"` //Сообщение (ИД)

}
