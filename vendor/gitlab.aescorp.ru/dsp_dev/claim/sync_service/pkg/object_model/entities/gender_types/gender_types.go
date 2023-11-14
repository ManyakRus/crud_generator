package gender_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// GenderType - модель для таблицы gender_types: Пол (справочник).
type GenderType struct {
	entities.CommonStruct
	Name string `json:"name" gorm:"column:name;default:\"\"" db:"name"` //Наименование

}
