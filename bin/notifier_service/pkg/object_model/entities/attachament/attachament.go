package attachament

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_attachament"
)

// Attachament - модель для таблицы attachament: Вложения прикладываемые к сообщениюЦелевая БД сервиса уведомлений
type Attachament struct {
	table_attachament.Table_Attachament	
}
