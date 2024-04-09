package message

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_message"
)

// Message - модель для таблицы message: сообщенияЦелевая БД сервиса уведомлений
type Message struct {
	table_message.Table_Message	
}
