package mailing_stats

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_mailing_stats"
)

// MailingStat - модель для таблицы mailing_stats: 
type MailingStat struct {
	table_mailing_stats.Table_MailingStat	
}
