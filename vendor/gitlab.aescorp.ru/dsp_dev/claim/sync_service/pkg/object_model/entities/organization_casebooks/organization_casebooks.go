package organization_casebooks

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"time"
)

// OrganizationCasebook - модель для таблицы organization_casebooks:
type OrganizationCasebook struct {
	entities.CommonStruct
	INN            string    `json:"inn" gorm:"column:inn;default:\"\"" db:"inn"`                                     //ИНН организации
	JSONFileID     int64     `json:"json_file_id" gorm:"column:json_file_id;default:0" db:"json_file_id"`             //Файл типа .json (ИД)
	JSONUpdatedAt  time.Time `json:"json_updated_at" gorm:"column:json_updated_at;default:null" db:"json_updated_at"` //Время последнего обновления файла .json
	OrganizationID int64     `json:"organization_id" gorm:"column:organization_id;default:0" db:"organization_id"`    //Организация (ИД)
	PDFFileID      int64     `json:"pdf_file_id" gorm:"column:pdf_file_id;default:0" db:"pdf_file_id"`                //Файл типа .pdf (ИД)
	PDFUpdatedAt   time.Time `json:"pdf_updated_at" gorm:"column:pdf_updated_at;default:null" db:"pdf_updated_at"`    //Время последнего обновления файла .pdf

}
