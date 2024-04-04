package delivery_status

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_delivery_status"
)

// DeliveryStatus - модель для таблицы delivery_status: Справочник статусов доставки из НСИ
type DeliveryStatus struct {
	table_delivery_status.Table_DeliveryStatus	
}
