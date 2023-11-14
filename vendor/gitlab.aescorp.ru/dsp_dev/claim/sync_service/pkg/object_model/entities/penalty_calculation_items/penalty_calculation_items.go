package penalty_calculation_items

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/accrual_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/contracts"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/documents"
	"time"
)

// PenaltyCalculationItem - model from table penalty_calculation_items: Расчет пени. (СТЕК Расчет пени по договорам)
type PenaltyCalculationItem struct {
	entities.CommonStruct
	AccrualSum               float64                   `json:"accrual_sum"                 gorm:"column:accrual_sum;default:0"`                    // СуммаНачислено
	AccrualTypeID            int64                     `json:"accrual_type_id"             gorm:"column:accrual_type_id;default:0"`                // ТипНачислений
	AccrualType              accrual_types.AccrualType `json:"accrual_type"                gorm:"-:all"`                                           // AccrualType
	Analytics                string                    `json:"analytics"                   gorm:"column:analytics;default:\"\""`                   // АналитикаПени
	Bid                      float64                   `json:"bid"                         gorm:"column:bid;default:0"`                            // Ставка
	BillingMonth             time.Time                 `json:"billing_month"               gorm:"column:billing_month;default:null"`               // РасчМесяц
	CalculationFormula       string                    `json:"calculation_formula"         gorm:"column:calculation_formula;default:\"\""`         // ФормулаРасчета
	ConnectionID             int64                     `json:"connection_id"               gorm:"column:connection_id;default:null"`               // Database
	Contract                 contracts.Contract        `json:"contract"                    gorm:"-:all"`                                           // Contract
	ContractID               int64                     `json:"contract_id"                 gorm:"column:contract_id;default:null"`                 // ContractID
	DateFrom                 time.Time                 `json:"date_from"                   gorm:"column:date_from;default:null"`                   // ДатНач
	DateTo                   time.Time                 `json:"date_to"                     gorm:"column:date_to;default:null"`                     // ДатКнц
	DaysCount                int                       `json:"days_count"                  gorm:"column:days_count;default:0"`                     // КолДней
	DebtSum                  float64                   `json:"debt_sum"                    gorm:"column:debt_sum;default:0"`                       // Задолженность
	DocumentInvoiceID        int64                     `json:"document_invoice_id"         gorm:"column:document_invoice_id;default:null"`         // Пени-ДокументЗадолженности
	DocumentInvoice          documents.Document        `json:"document_invoice"            gorm:"-:all"`                                           // DocumentInvoice
	DocumentPenaltyInvoiceID int64                     `json:"document_penalty_invoice_id" gorm:"column:document_penalty_invoice_id;default:null"` // Пени-ДокументРеализации
	DocumentPenaltyInvoice   documents.Document        `json:"document_penalty_invoice"    gorm:"-:all"`                                           // DocumentPenaltyInvoice
	Note                     string                    `json:"note"                        gorm:"column:note;default:\"\""`                        // Примечание
	Percent                  float64                   `json:"percent"                     gorm:"column:percent;default:0"`                        // Процент
	UsedMonth                time.Time                 `json:"used_month"                  gorm:"column:used_month;default:null"`                  // ЗаМесяц
}
