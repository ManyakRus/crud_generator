package channel

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_channel"
)

// Channel - модель для таблицы channel: справочник каналов доставки
type Channel struct {
	table_channel.Table_Channel	
}
