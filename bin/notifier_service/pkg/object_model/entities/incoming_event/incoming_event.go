package incoming_event

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_incoming_event"
)

// IncomingEvent - модель для таблицы incoming_event: Входяшие события из шины или API, как есть.
Целевая БД сервиса уведомлений
type IncomingEvent struct {
	table_incoming_event.Table_IncomingEvent	
}
