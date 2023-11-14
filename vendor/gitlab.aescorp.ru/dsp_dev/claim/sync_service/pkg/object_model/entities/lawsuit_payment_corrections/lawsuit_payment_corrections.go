package lawsuit_payment_corrections

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// LawsuitPaymentCorrection - модель для таблицы lawsuit_payment_corrections:
type LawsuitPaymentCorrection struct {
	entities.CommonStruct
	CorrectionDocumentID  int64   `json:"correction_document_id" gorm:"column:correction_document_id;default:0" db:"correction_document_id"`    //
	CorrectionDocumentSum float64 `json:"correction_document_sum" gorm:"column:correction_document_sum;default:0" db:"correction_document_sum"` //
	LawsuitID             int64   `json:"lawsuit_id" gorm:"column:lawsuit_id;default:0" db:"lawsuit_id"`                                        //
	PaymentDocumentID     int64   `json:"payment_document_id" gorm:"column:payment_document_id;default:0" db:"payment_document_id"`             //

}
