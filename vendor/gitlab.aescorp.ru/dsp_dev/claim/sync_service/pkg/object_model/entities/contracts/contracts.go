package contracts

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/contract_category_types"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organizations"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/payment_days"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/payment_schedules"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
	"time"
)

// Contract Договоры.
type Contract struct {
	entities.CommonStruct
	entities.GroupStruct
	BeginAt                 time.Time                                    `json:"begin_at"               gorm:"column:begin_at"`
	BranchID                int64                                        `json:"branch_id"              gorm:"column:branch_id;default:null"`
	Category                contract_category_types.ContractCategoryType `json:"category"               gorm:"-:all"`
	CategoryID              int64                                        `json:"category_id"            gorm:"column:category_id;default:null"`
	ConnectionID            int64                                        `json:"connection_id"          gorm:"column:connection_id;default:null"`
	CuratorClaim            employees.Employee                           `json:"curator_claim"          gorm:"-:all"`
	CuratorClaimID          int64                                        `json:"curator_claim_id"       gorm:"column:curator_claim_id;default:null"`
	CuratorContract         employees.Employee                           `json:"curator_contract"       gorm:"-:all"`
	CuratorContractID       int64                                        `json:"curator_contract_id"    gorm:"column:curator_contract_id;default:null"`
	CuratorLegal            employees.Employee                           `json:"curator_legal"          gorm:"-:all"`
	CuratorLegalID          int64                                        `json:"curator_legal_id"       gorm:"column:curator_legal_id;default:null"`
	CuratorPayment          employees.Employee                           `json:"curator_payment"        gorm:"-:all"`
	CuratorPaymentID        int64                                        `json:"curator_payment_id"     gorm:"column:curator_payment_id;default:null"`
	CuratorTechAudit        employees.Employee                           `json:"curator_tech_audit"     gorm:"-:all"`
	CuratorTechAuditID      int64                                        `json:"curator_tech_audit_id"  gorm:"column:curator_tech_audit_id;default:null"`
	DaysToResolveClaim      int                                          `json:"days_to_resolve_claim"  gorm:"column:days_to_resolve_claim"`
	Description             string                                       `json:"description"            gorm:"column:description;default:\"\""`
	Email                   string                                       `json:"email"                  gorm:"column:email;default:\"\""`
	EndAt                   time.Time                                    `json:"end_at"                 gorm:"column:end_at"`
	ErrorFromStackAt        time.Time                                    `json:"error_from_stack_at"    gorm:"column:error_from_stack_at"`
	IndividualID            int64                                        `json:"individual_id"          gorm:"column:individual_id;default:null"`
	IsErrorFromStack        bool                                         `json:"is_error_from_stack"    gorm:"column:is_error_from_stack;default:false"`
	IsIndOrganization       bool                                         `json:"is_ind_organization"    gorm:"column:is_ind_organization;default:false"`
	IsOrganization          bool                                         `json:"is_organization"        gorm:"column:is_organization;default:false"`
	IsValidEmail            bool                                         `json:"is_valid_email"         gorm:"column:is_valid_email;default:true"`
	Number                  alias.ContractNumber                         `json:"number"                 gorm:"column:number;default:\"\""`
	Organization            organizations.Organization                   `json:"organization"           gorm:"-:all"`
	OrganizationConsigneeID int64                                        `json:"organization_consignee_id" gorm:"column:organization_consignee_id;default:null"` // Грузополучатель - consignee
	OrganizationCustomerID  int64                                        `json:"organization_customer_id"  gorm:"column:organization_customer_id;default:null"`  // Заказчик - customer
	OrganizationPayerID     int64                                        `json:"organization_payer_id"     gorm:"column:organization_payer_id;default:null"`     // Плательщик - payer
	PostAddress             string                                       `json:"post_address"           gorm:"column:post_address;default:\"\""`
	SignAt                  time.Time                                    `json:"sign_at"                gorm:"column:sign_at"`
	Status                  string                                       `json:"status"                 gorm:"column:status;default:\"\""`
	TerminateAt             time.Time                                    `json:"terminate_at"           gorm:"column:terminate_at"`
	PaymentDays             []payment_days.PaymentDay                    `json:"payment_days"`      // Дни платежей
	PaymentSchedules        []payment_schedules.PaymentSchedule          `json:"payment_schedules"` // Графики платежей
}
