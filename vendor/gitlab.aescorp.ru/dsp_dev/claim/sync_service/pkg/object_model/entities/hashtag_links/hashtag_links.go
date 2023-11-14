package hashtag_links

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// HashtagLink - модель для таблицы hashtag_links:
type HashtagLink struct {
	entities.CommonStruct
	entities.ExtLinkStruct
	HashtagID int64 `json:"hashtag_id" gorm:"column:hashtag_id;default:0" db:"hashtag_id"` //

}
