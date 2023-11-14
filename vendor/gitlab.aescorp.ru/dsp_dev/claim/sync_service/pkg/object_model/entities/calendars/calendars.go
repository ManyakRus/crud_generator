package calendars

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// Calendar - модель для таблицы calendars: Производственный календарь
type Calendar struct {
	entities.CommonStruct
	Comment   string    `json:"comment" gorm:"column:comment;default:\"\"" db:"comment"`          //
	Date      time.Time `json:"date" gorm:"column:date;default:null" db:"date"`                   //
	DayTypeID int64     `json:"day_type_id" gorm:"column:day_type_id;default:0" db:"day_type_id"` //
	Days      int32     `json:"days" gorm:"column:days;default:0" db:"days"`                      //
	Hours     int32     `json:"hours" gorm:"column:hours;default:0" db:"hours"`                   //

}
