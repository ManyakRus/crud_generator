package file_templates

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// FileTemplate - модель для таблицы file_templates: Шаблоны файлов (справочник).
type FileTemplate struct {
	entities.CommonStruct
	entities.NameStruct
	FileID int64 `json:"file_id" gorm:"column:file_id;default:0" db:"file_id"` //Указатель на файл.

}
