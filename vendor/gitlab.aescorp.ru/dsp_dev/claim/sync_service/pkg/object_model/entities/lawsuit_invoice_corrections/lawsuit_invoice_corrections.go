package lawsuit_invoice_corrections

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LawsuitInvoiceCorrection - модель для таблицы lawsuit_invoice_corrections:
type LawsuitInvoiceCorrection struct {
	entities.CommonStruct
	CorrectionDocumentID  int64   `json:"correction_document_id" gorm:"column:correction_document_id;default:0" db:"correction_document_id"`    //
	CorrectionDocumentSum float64 `json:"correction_document_sum" gorm:"column:correction_document_sum;default:0" db:"correction_document_sum"` //
	InvoiceDocumentID     int64   `json:"invoice_document_id" gorm:"column:invoice_document_id;default:0" db:"invoice_document_id"`             //
	LawsuitID             int64   `json:"lawsuit_id" gorm:"column:lawsuit_id;default:0" db:"lawsuit_id"`                                        //

}
