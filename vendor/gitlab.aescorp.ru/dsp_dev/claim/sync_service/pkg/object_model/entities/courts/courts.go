package courts

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Court - модель для таблицы courts: Суды (справочник).
type Court struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	City           string `json:"city_name" gorm:"column:city_name;default:\"\"" db:"city_name"`                //Город
	OrganizationID int64  `json:"organization_id" gorm:"column:organization_id;default:0" db:"organization_id"` //Организация (ИД)

}
