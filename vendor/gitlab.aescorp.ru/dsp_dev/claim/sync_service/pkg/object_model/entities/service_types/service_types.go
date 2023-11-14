package service_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// ServiceType - модель для таблицы service_types: Типы услуг (справочник).
type ServiceType struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	Code              int32  `json:"code" gorm:"column:code;default:0" db:"code"`                                              //Номер услуги
	ConnectionID      int64  `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`                   //Соединение к БД СТЕК (ИД)
	FullName          string `json:"full_name" gorm:"column:full_name;default:\"\"" db:"full_name"`                            //Полное наименование
	Measure           string `json:"measure" gorm:"column:measure;default:\"\"" db:"measure"`                                  //Единица измерения
	ServiceProviderID int64  `json:"service_provider_id" gorm:"column:service_provider_id;default:0" db:"service_provider_id"` //Поставщик (ИД)

}
