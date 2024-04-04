package delivery_error

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_delivery_error"
)

// DeliveryError - модель для таблицы delivery_error: Таблица маппинга автоответов серверов и кодов ошибок НСИ
type DeliveryError struct {
	table_delivery_error.Table_DeliveryError	
}
