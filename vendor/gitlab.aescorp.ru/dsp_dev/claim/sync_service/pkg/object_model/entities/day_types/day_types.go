package day_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// DayType - модель для таблицы day_types:
type DayType struct {
	entities.CommonStruct
	entities.NameStruct
	IsWorkDay bool   `json:"is_work_day" gorm:"column:is_work_day" db:"is_work_day"`           //
	ShortName string `json:"short_name" gorm:"column:short_name;default:\"\"" db:"short_name"` //

}
