package event_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// EventType - модель для таблицы event_types: Типы событий.
type EventType struct {
	entities.CommonStruct
	entities.NameStruct
	Code int32 `json:"code" gorm:"column:code;default:0" db:"code"` //Код типа события

}
