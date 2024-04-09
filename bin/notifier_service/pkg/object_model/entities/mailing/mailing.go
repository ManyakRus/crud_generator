package mailing

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_mailing"
)

// Mailing - модель для таблицы mailing: РассылкиЦелевая БД сервиса уведомлений
type Mailing struct {
	table_mailing.Table_Mailing	
}
