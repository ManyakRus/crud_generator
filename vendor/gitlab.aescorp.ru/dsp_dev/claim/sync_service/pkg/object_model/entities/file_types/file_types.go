package file_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// FileType Тип файла (справочник).
type FileType struct {
	entities.CommonStruct
	entities.NameStruct
	Code int `json:"code" gorm:"column:code;default:0"`
}
