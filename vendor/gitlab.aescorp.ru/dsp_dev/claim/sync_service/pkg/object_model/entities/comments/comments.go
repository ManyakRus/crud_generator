package comments

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Comment - модель для таблицы comments: Комментарии.
type Comment struct {
	entities.CommonStruct
	entities.ExtLinkStruct
	Message string `json:"message" gorm:"column:message;default:\"\"" db:"message"` //Текст комментария

}
