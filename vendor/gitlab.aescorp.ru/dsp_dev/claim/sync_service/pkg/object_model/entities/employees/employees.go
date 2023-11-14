package employees

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Employee - модель для таблицы employees: Сотрудники (Справочник).
type Employee struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	BranchID     int64  `json:"branch_id" gorm:"column:branch_id;default:0" db:"branch_id"`             //Филиал (ИД)
	ConnectionID int64  `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //Соединение к БД СТЕК (ИД)
	Email        string `json:"email" gorm:"column:email;default:\"\"" db:"email"`                      //ЕМайл пользователя
	IsActive     bool   `json:"is_active" gorm:"column:is_active" db:"is_active"`                       //Активно (не отключен)
	Login        string `json:"login" gorm:"column:login;default:\"\"" db:"login"`                      //Логин в систему ПИР, он же логин в домене
	ParentName   string `json:"parent_name" gorm:"column:parent_name;default:\"\"" db:"parent_name"`    //Отчество пользователя
	Phone        string `json:"phone" gorm:"column:phone;default:\"\"" db:"phone"`                      //Номер телефона пользователя
	Photo        string `json:"photo" gorm:"column:photo;default:\"\"" db:"photo"`                      //Фотография пользователя
	Position     string `json:"position" gorm:"column:position;default:\"\"" db:"position"`             //Должность
	SecondName   string `json:"second_name" gorm:"column:second_name;default:\"\"" db:"second_name"`    //Фамилия пользователя
	Tag          string `json:"tag" gorm:"column:tag;default:\"\"" db:"tag"`                            //Тэг для поиска

}
