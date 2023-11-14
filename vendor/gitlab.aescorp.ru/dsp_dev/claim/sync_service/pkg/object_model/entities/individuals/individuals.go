package individuals

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// Individual - модель для таблицы individuals: Физические лица (справочник).
type Individual struct {
	entities.CommonStruct
	entities.NameStruct
	BirthDate    time.Time `json:"birth_date" gorm:"column:birth_date;default:null" db:"birth_date"`       //Дата рождения
	ConnectionID int64     `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //Соединение к БД СТЕК (ИД)
	DeathDate    time.Time `json:"death_date" gorm:"column:death_date;default:null" db:"death_date"`       //Дата смерти
	Email        string    `json:"email" gorm:"column:email;default:\"\"" db:"email"`                      //Е-майл
	GenderID     int64     `json:"gender_id" gorm:"column:gender_id;default:0" db:"gender_id"`             //Пол физического лица (ИД)
	INN          string    `json:"inn" gorm:"column:inn;default:\"\"" db:"inn"`                            //ИНН
	ParentName   string    `json:"parent_name" gorm:"column:parent_name;default:\"\"" db:"parent_name"`    //Отчество
	Phone        string    `json:"phone" gorm:"column:phone;default:\"\"" db:"phone"`                      //Номер телефона
	SecondName   string    `json:"second_name" gorm:"column:second_name;default:\"\"" db:"second_name"`    //Фамилия
	Snils        string    `json:"snils" gorm:"column:snils;default:\"\"" db:"snils"`                      //СНИЛС

}
