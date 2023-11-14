package accrual_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// AccrualType - модель для таблицы accrual_types: ТипНачислений (Справочник).
type AccrualType struct {
	entities.CommonStruct
	entities.NameStruct
}
