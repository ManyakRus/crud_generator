package service_providers

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// ServiceProvider - модель для таблицы service_providers: Поставщик услуг (справочник).
type ServiceProvider struct {
	entities.CommonStruct
	ConnectionID   int64     `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`       //Соединение к БД СТЕК (ИД)
	DateFrom       time.Time `json:"date_from" gorm:"column:date_from;default:null" db:"date_from"`                //Дата начала действия
	DateTo         time.Time `json:"date_to" gorm:"column:date_to;default:null" db:"date_to"`                      //Дата окончания действия
	OrganizationID int64     `json:"organization_id" gorm:"column:organization_id;default:0" db:"organization_id"` //Организация (ИД)

}
