package state_duties

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// StateDuty - модель для таблицы state_duties: Пошлины.
type StateDuty struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	CourtID       int64     `json:"court_id" gorm:"column:court_id;default:0" db:"court_id"`                      //Суд (ИД)
	LawsuitID     int64     `json:"lawsuit_id" gorm:"column:lawsuit_id;default:0" db:"lawsuit_id"`                //Претензия (ИД)
	RequestDate   time.Time `json:"request_date" gorm:"column:request_date;default:null" db:"request_date"`       //Дата запроса
	RequestNumber string    `json:"request_number" gorm:"column:request_number;default:\"\"" db:"request_number"` //Номер запроса
	Sum           float64   `json:"sum" gorm:"column:sum;default:0" db:"sum"`                                     //Сумма

}
