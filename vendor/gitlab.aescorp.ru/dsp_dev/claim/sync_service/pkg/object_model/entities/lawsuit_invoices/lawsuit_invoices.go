package lawsuit_invoices

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/documents"
	"time"
)

// LawsuitInvoice - model from table lawsuit_invoices: Счета фактуры относящиеся к делу.
type LawsuitInvoice struct {
	ClosedAt     time.Time          `json:"closed_at" gorm:"column:closed_at;default:null" db:"closed_at"`                      //
	ClosedSum    float64            `json:"closed_sum" gorm:"column:closed_sum;default:0" db:"closed_sum"`                      //
	Count        int64              `json:"count" gorm:"column:count;default:0" db:"count"`                                     //Количество (КВт)
	CreatedAt    time.Time          `json:"created_at" gorm:"column:created_at;default:null;autoCreateTime" db:"created_at"`    //
	DeletedAt    time.Time          `json:"deleted_at" gorm:"column:deleted_at;default:null" db:"deleted_at"`                   //
	DocumentID   int64              `json:"document_id" gorm:"column:document_id;default:0" db:"document_id"`                   //Документ Счёт-фактура (ИД)
	DocumentSum  float64            `json:"document_sum" gorm:"column:document_sum;default:0" db:"document_sum"`                //
	ExtID        int64              `json:"ext_id" gorm:"column:ext_id;default:0" db:"ext_id"`                                  //
	ID           int64              `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;default:0" db:"id"`                //Уникальный технический идентификатор
	IsClosed     bool               `json:"is_closed" gorm:"column:is_closed" db:"is_closed"`                                   //
	IsCorrective bool               `json:"is_corrective" gorm:"column:is_corrective" db:"is_corrective"`                       //
	IsDeleted    bool               `json:"is_deleted" gorm:"column:is_deleted" db:"is_deleted"`                                //
	LawsuitID    int64              `json:"lawsuit_id" gorm:"column:lawsuit_id;default:0" db:"lawsuit_id"`                      //Претензия (ИД)
	ModifiedAt   time.Time          `json:"modified_at" gorm:"column:modified_at;default:null;autoUpdateTime" db:"modified_at"` //
	Sum          float64            `json:"sum" gorm:"column:sum;default:0" db:"sum"`                                           //Сумма документа
	Document     documents.Document `json:"document"      gorm:"-:all"`
}
