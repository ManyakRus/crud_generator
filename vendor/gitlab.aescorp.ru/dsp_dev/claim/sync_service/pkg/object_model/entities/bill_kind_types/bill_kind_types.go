package bill_kind_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// BillKindType - модель для таблицы bill_kind_types: Виды платежей (справочник).
type BillKindType struct {
	entities.CommonStruct
	entities.NameStruct
	Code int32 `json:"code" gorm:"column:code;default:0" db:"code"` //Код

}
