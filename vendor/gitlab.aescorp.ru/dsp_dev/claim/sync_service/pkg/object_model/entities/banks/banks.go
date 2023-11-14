package banks

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Bank - модель для таблицы banks: Банки (справочник).
type Bank struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	BIK                  string `json:"bik" gorm:"column:bik;default:\"\"" db:"bik"`                                                       //БИК (Банковский идентификатор)
	City                 string `json:"city_name" gorm:"column:city_name;default:\"\"" db:"city_name"`                                     //Город
	ConnectionID         int64  `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`                            //Соединение к БД СТЕК (ИД)
	CorrespondentAccount string `json:"correspondent_account" gorm:"column:correspondent_account;default:\"\"" db:"correspondent_account"` //Корреспондентский счёт
	OrganizationID       int64  `json:"organization_id" gorm:"column:organization_id;default:0" db:"organization_id"`                      //Организация (ИД)

}
