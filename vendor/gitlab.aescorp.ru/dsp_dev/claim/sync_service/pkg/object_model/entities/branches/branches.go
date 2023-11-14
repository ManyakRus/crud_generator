package branches

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Branch - модель для таблицы branches: Филиалы (справочник).
type Branch struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	Code             int64  `json:"code" gorm:"column:code;default:0" db:"code"`                                              //Код
	OrganizationID   int64  `json:"organization_id" gorm:"column:organization_id;default:0" db:"organization_id"`             //Организация (ИД)
	PersonalAreaLink string `json:"personal_area_link" gorm:"column:personal_area_link;default:\"\"" db:"personal_area_link"` //

}
