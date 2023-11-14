package lawsuit_status_states

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// LawsuitStatusState - модель для таблицы lawsuit_status_states: История статусов дела.
type LawsuitStatusState struct {
	entities.CommonStruct
	CommentID    int64     `json:"comment_id" gorm:"column:comment_id;default:0" db:"comment_id"`             //ИД комментария comments
	InvoiceSum   float64   `json:"invoice_sum" gorm:"column:invoice_sum;default:0" db:"invoice_sum"`          //Сумма счетов фактур
	LawsuitID    int64     `json:"lawsuit_id" gorm:"column:lawsuit_id;default:0" db:"lawsuit_id"`             //ИД претензии lawsuits
	MainSum      float64   `json:"main_sum" gorm:"column:main_sum;default:0" db:"main_sum"`                   //Сумма долга по основному виду деятельности
	PaySum       float64   `json:"pay_sum" gorm:"column:pay_sum;default:0" db:"pay_sum"`                      //Сумма плптежей
	PenaltySum   float64   `json:"penalty_sum" gorm:"column:penalty_sum;default:0" db:"penalty_sum"`          //Сумма госпошлины
	PennySum     float64   `json:"penny_sum" gorm:"column:penny_sum;default:0" db:"penny_sum"`                //Сумма ПЕНИ
	RestrictSum  float64   `json:"restrict_sum" gorm:"column:restrict_sum;default:0" db:"restrict_sum"`       //Сумма ограничений
	StateDutySum float64   `json:"state_duty_sum" gorm:"column:state_duty_sum;default:0" db:"state_duty_sum"` //Сумма пошлины
	StatusAt     time.Time `json:"status_at" gorm:"column:status_at;default:null" db:"status_at"`             //Дата установки статуса
	StatusID     int64     `json:"status_id" gorm:"column:status_id;default:0" db:"status_id"`                //Статус (ИД)
	Tag          string    `json:"tag" gorm:"column:tag;default:\"\"" db:"tag"`                               //Тэг для поиска
	TotalDebt    float64   `json:"total_debt" gorm:"column:total_debt;default:0" db:"total_debt"`             //Сумма долга

}
