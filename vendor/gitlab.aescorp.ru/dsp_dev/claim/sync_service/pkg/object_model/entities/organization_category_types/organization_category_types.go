package organization_category_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// OrganizationCategoryType - модель для таблицы organization_category_types: Категории организаций
type OrganizationCategoryType struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	ConnectionID int64 `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //

}
