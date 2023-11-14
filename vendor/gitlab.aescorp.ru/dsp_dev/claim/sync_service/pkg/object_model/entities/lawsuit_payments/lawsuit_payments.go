package lawsuit_payments

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/documents"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_invoices"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
)

// LawsuitPayment -- платежи относящиеся к делу
type LawsuitPayment struct {
	entities.CommonStruct
	ID           alias.PaymentId                 `json:"id"          gorm:"column:id;primaryKey;autoIncrement:true"`
	Document     documents.Document              `json:"document"     gorm:"-:all"`
	DocumentID   int64                           `json:"document_id"  gorm:"column:document_id;default:null"`        // Document
	DocumentSum  float64                         `json:"document_sum" gorm:"column:document_sum;not null;default:0"` // Сумма указанная в платёжном документе
	Invoice      lawsuit_invoices.LawsuitInvoice `json:"invoice"      gorm:"-:all"`
	InvoiceID    alias.InvoiceId                 `json:"invoice_id"   gorm:"column:invoice_id;default:null"` // LawsuitInvoice
	IsCorrective bool                            `json:"is_corrective"    gorm:"column:is_corrective;default:false"`
	LawsuitID    alias.LawsuitId                 `json:"lawsuit_id"   gorm:"column:lawsuit_id;default:null"` // Lawsuit
	Sum          float64                         `json:"sum"          gorm:"column:sum;not null;default:0"`  // Сумма погашения после коррекции
}
