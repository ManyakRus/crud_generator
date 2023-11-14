package hashtags

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Hashtag - модель для таблицы hashtags:
type Hashtag struct {
	entities.CommonStruct
	entities.NameStruct
}
