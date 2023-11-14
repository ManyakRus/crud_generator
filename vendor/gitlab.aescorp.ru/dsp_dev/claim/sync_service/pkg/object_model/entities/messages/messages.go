package messages

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// Message - модель для таблицы messages: Сообщения (входящие и исходящие).
type Message struct {
	entities.CommonStruct
	ChannelTypeID   int64     `json:"channel_type_id" gorm:"column:channel_type_id;default:0" db:"channel_type_id"`       //Канал отправки сообщения (ИД)
	ContactFrom     string    `json:"contact_from" gorm:"column:contact_from;default:\"\"" db:"contact_from"`             //EMail от кого
	ContactTo       string    `json:"contact_to" gorm:"column:contact_to;default:\"\"" db:"contact_to"`                   //EMail кому
	DirectionTypeID int64     `json:"direction_type_id" gorm:"column:direction_type_id;default:0" db:"direction_type_id"` //ИД входящее или исходящее
	EmployeeIdFrom  int64     `json:"employee_id_from" gorm:"column:employee_id_from;default:0" db:"employee_id_from"`    //Сотрудник от кого сообщение (ИД)
	EmployeeIdTo    int64     `json:"employee_id_to" gorm:"column:employee_id_to;default:0" db:"employee_id_to"`          //Сотрудник от кого (ИД)
	ExtCode         string    `json:"ext_code" gorm:"column:ext_code;default:\"\"" db:"ext_code"`                         //ШПИ (штрихкод)
	LawsuitID       int64     `json:"lawsuit_id" gorm:"column:lawsuit_id;default:0" db:"lawsuit_id"`                      //Дело (ИД)
	MailingCode     string    `json:"mailing_code" gorm:"column:mailing_code;default:\"\"" db:"mailing_code"`             //mailing_code сервиса нотификации
	MessageFileID   int64     `json:"message_file_id" gorm:"column:message_file_id;default:0" db:"message_file_id"`       //Файл с текстом письма (ИД)
	MessageTypeID   int64     `json:"message_type_id" gorm:"column:message_type_id;default:0" db:"message_type_id"`       //Тип сообщения
	ReceiveResult   string    `json:"receive_result" gorm:"column:receive_result;default:\"\"" db:"receive_result"`       //Результат получения сообщения (текст ошибки)
	ReceiveStatusID int64     `json:"receive_status_id" gorm:"column:receive_status_id;default:0" db:"receive_status_id"` //
	ReceivedAt      time.Time `json:"received_at" gorm:"column:received_at;default:null" db:"received_at"`                //Дата получения сообщения
	SendResult      string    `json:"send_result" gorm:"column:send_result;default:\"\"" db:"send_result"`                //
	SendStatusID    int64     `json:"send_status_id" gorm:"column:send_status_id;default:0" db:"send_status_id"`          //Статус отправки (ИД)
	SentAt          time.Time `json:"sent_at" gorm:"column:sent_at;default:null" db:"sent_at"`                            //Время отправки сообщения
	Topic           string    `json:"topic" gorm:"column:topic;default:\"\"" db:"topic"`                                  //Тема письма
	NSICode         int       `json:"nsi_code"          gorm:"column:nsi_code;default:null"`                              // Статус доставки уведомлений

}
