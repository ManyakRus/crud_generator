package document_link_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// DocumentLinkType - модель для таблицы document_link_types: Типы связей документов (справочник).
type DocumentLinkType struct {
	entities.CommonStruct
	entities.NameStruct
	Code int32 `json:"code" gorm:"column:code;default:0" db:"code"` //Код

}
