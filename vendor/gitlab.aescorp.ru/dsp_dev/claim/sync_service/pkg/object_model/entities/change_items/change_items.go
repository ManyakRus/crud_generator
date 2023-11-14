package change_items

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// ChangeItem - модель для таблицы change_items:
type ChangeItem struct {
	entities.CommonStruct
	entities.ExtLinkStruct
	Key   string `json:"key" gorm:"column:key;default:\"\"" db:"key"`       //
	Prev  string `json:"prev" gorm:"column:prev;default:\"\"" db:"prev"`    //
	Value string `json:"value" gorm:"column:value;default:\"\"" db:"value"` //

}
