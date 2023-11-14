package lawsuits

import (
	"time"

	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/branches"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/claim_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/contracts"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/courts"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_reason_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_stage_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_states"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
)

// Lawsuit Дело. Объединяет весь набор данных по конкретному должнику.
type Lawsuit struct {
	entities.CommonStruct
	entities.GroupStruct
	entities.NameStruct
	Branch                    branches.Branch                            `json:"branch"                  gorm:"-:all"`
	BranchID                  int64                                      `json:"branch_id"               gorm:"column:branch_id;default:null"`
	Chance                    string                                     `json:"chance"                  gorm:"column:chance;default:\"\""`
	ClaimAt                   time.Time                                  `json:"claim_at"                gorm:"column:claim_at;default:null"` // Уведомление о задолженности. Оплатить до.
	ClaimPeriodStr            string                                     `json:"claim_period_str"        gorm:"column:claim_period_str;default:\"\""`
	ClaimType                 claim_types.ClaimType                      `json:"claim_type"              gorm:"-:all"` // Тип задолженности
	ClaimTypeID               int64                                      `json:"claim_type_id"           gorm:"column:claim_type_id;default:null"`
	ClosedAt                  time.Time                                  `json:"closed_at"               gorm:"column:closed_at;default:null"`
	ConnectionID              int64                                      `json:"connection_id"           gorm:"column:connection_id;default:null"`
	Contract                  contracts.Contract                         `json:"contract"                gorm:"-:all"` // Договор
	ContractID                int64                                      `json:"contract_id"             gorm:"column:contract_id;default:null"`
	ControlledAt              time.Time                                  `json:"controlled_at"           gorm:"column:controlled_at;default:null"`
	Court                     courts.Court                               `json:"court"                   gorm:"-:all"`
	CourtID                   int64                                      `json:"court_id"                gorm:"column:court_id;default:null"`
	CuratorClaim              employees.Employee                         `json:"curator_claim"           gorm:"-:all"`
	CuratorClaimID            int64                                      `json:"curator_claim_id"        gorm:"column:curator_claim_id;default:null"`
	CuratorContract           employees.Employee                         `json:"curator_contract"        gorm:"-:all"`
	CuratorContractID         int64                                      `json:"curator_contract_id"     gorm:"column:curator_contract_id;default:null"`
	CuratorLegal              employees.Employee                         `json:"curator_legal"           gorm:"-:all"`
	CuratorLegalID            int64                                      `json:"curator_legal_id"        gorm:"column:curator_legal_id;default:null"`
	CuratorPayment            employees.Employee                         `json:"curator_payment"         gorm:"-:all"`
	CuratorPaymentID          int64                                      `json:"curator_payment_id"      gorm:"column:curator_payment_id;default:null"`
	CuratorTechAudit          employees.Employee                         `json:"curator_tech_audit"      gorm:"-:all"`
	CuratorTechAuditID        int64                                      `json:"curator_tech_audit_id"   gorm:"column:curator_tech_audit_id;default:null"`
	DateFrom                  time.Time                                  `json:"date_from"               gorm:"column:date_from;default:null"`
	DateTo                    time.Time                                  `json:"date_to"                 gorm:"column:date_to;default:null"`
	DebtSum                   float64                                    `json:"debt_sum"                gorm:"column:debt_sum;default:0"`    // Текущая сумма долга, руб.
	InvoiceSum                float64                                    `json:"invoice_sum"             gorm:"column:invoice_sum;default:0"` // Сумма долга за период
	IsClosed                  bool                                       `json:"is_closed"               gorm:"column:is_closed;default:false"`
	MainSum                   float64                                    `json:"main_sum"                gorm:"column:main_sum;default:0"`                             // Сумма долга по основному виду деятельности
	NotifyClaimAt             time.Time                                  `json:"notify_claim_at"         gorm:"column:notify_claim_at;default:null"`                   // Уведомление о задолженности. Дата отправки.
	NotifyClaimChannel        int                                        `json:"notify_claim_channel"    gorm:"column:notify_claim_channel;default:null"`              // Уведомление о задолженности. Канал отправки.
	NotifyClaimCode           int                                        `json:"notify_claim_code"       gorm:"column:notify_claim_code;default:null"`                 // Уведомление о задолженности. Код доставки из НСИ.
	NotifyClaimDone           bool                                       `json:"notify_claim_done"       gorm:"column:notify_claim_done;default:false"`                // Уведомление о задолженности. Факт отправки.
	NotifyClaimMailingCode    string                                     `json:"notify_claim_mailing_code" gorm:"column:notify_claim_mailing_code;default:null"`       // Уведомление о задолженности. Уникальный код отправки.
	NotifyPretrialAt          time.Time                                  `json:"notify_pretrial_at"      gorm:"column:notify_pretrial_at;default:null"`                // Досудебная претензия. Дата отправки.
	NotifyPretrialChannel     int                                        `json:"notify_pretrial_channel" gorm:"column:notify_pretrial_channel;default:null"`           // Досудебная претензия. Канал отправки.
	NotifyPretrialCode        int                                        `json:"notify_pretrial_code"    gorm:"column:notify_pretrial_code;default:null"`              // Досудебная претензия. Код доставки из НСИ.
	NotifyPretrialDone        bool                                       `json:"notify_pretrial_done"    gorm:"column:notify_pretrial_done;default:false"`             // Досудебная претензия. Факт отправки.
	NotifyPretrialMailingCode string                                     `json:"notify_pretrial_mailing_code" gorm:"column:notify_pretrial_mailing_code;default:null"` // Досудебная претензия. Уникальный код отправки.
	Number                    alias.LawsuitNumber                        `json:"number"                  gorm:"column:number;default:\"\""`
	NumberClaim               alias.ClaimNumber                          `json:"number_claim"            gorm:"column:number_claim;default:\"\""`
	NumberTrial               string                                     `json:"number_trial"            gorm:"column:number_trial;default:\"\""`
	PaySum                    float64                                    `json:"pay_sum"                 gorm:"column:pay_sum;default:0"` // Платежи
	Penalty                   float64                                    `json:"penalty"                 gorm:"column:penalty;default:0"`
	Penny                     float64                                    `json:"penny"                   gorm:"column:penny;default:0"`
	Percent317                float64                                    `json:"percent_317"             gorm:"column:percent_317;default:0"`
	Percent395                float64                                    `json:"percent_395"             gorm:"column:percent_395;default:0"`
	PretrialAt                time.Time                                  `json:"pretrial_at"             gorm:"column:pretrial_at;default:null"` // Досудебная претензия. Оплатить до.
	ProcessKey                string                                     `json:"process_key"             gorm:"column:process_key;default:\"\""`
	ProcessStartedAt          time.Time                                  `json:"process_started_at"      gorm:"column:process_started_at;default:null"`
	Reason                    lawsuit_reason_types.LawsuitReasonType     `json:"reason"                  gorm:"-:all"`
	ReasonID                  int64                                      `json:"reason_id"               gorm:"column:reason_id;default:null"`
	RestrictSum               float64                                    `json:"restrict_sum"            gorm:"column:restrict_sum;default:0"`
	Stage                     lawsuit_stage_types.LawsuitStageType       `json:"stage"                   gorm:"-:all"` // Этап
	StageAt                   time.Time                                  `json:"stage_at"                gorm:"column:stage_at;default:null"`
	StageID                   int64                                      `json:"stage_id"                gorm:"column:stage_id;default:null"`
	StateDuty                 float64                                    `json:"state_duty"              gorm:"column:state_duty;default:0"` // Пошлина
	Status                    lawsuit_status_types.LawsuitStatusType     `json:"status"                  gorm:"-:all"`                       // Статус
	StatusAt                  time.Time                                  `json:"status_at"               gorm:"column:status_at;default:null"`
	StatusID                  int64                                      `json:"status_id"               gorm:"column:status_id;default:null"`
	StatusStates              []lawsuit_status_states.LawsuitStatusState `json:"status_states"           gorm:"-:all"` // TODO Перенести Суммы на разных статусах дела
	Tag                       string                                     `json:"tag"                     gorm:"column:tag;default:\"\""`
	Type                      lawsuit_types.LawsuitType                  `json:"type"                    gorm:"-:all"` // Тип претензии
	TypeID                    int64                                      `json:"type_id"                 gorm:"column:type_id;default:null"`
	UnknownPayments           bool                                       `json:"unknown_payments"        gorm:"column:unknown_payments;default:false"` // "С не разнесёнными платежами"
}
