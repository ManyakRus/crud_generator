package contract_black_items

import (
	"time"

	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/contracts"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
)

// ContractBlackItem - model from table contract_black_items: "Чёрный" список договоров. Кому сразу предъявляется претензия.
type ContractBlackItem struct {
	entities.CommonStruct
	Contract       contracts.Contract   `json:"contract"        gorm:"-:all"`
	ContractID     int64                `json:"contract_id"     gorm:"column:contract_id;default:null"`
	ContractNumber alias.ContractNumber `json:"contract_number" gorm:"column:contract_number;default:null"`
	CreatedBy      employees.Employee   `json:"created_by"      gorm:"-:all"`
	CreatedByID    int64                `json:"created_by_id"   gorm:"column:created_by_id;default:null"`
	DateFrom       time.Time            `json:"date_from"       gorm:"column:date_from;default:null"`
	DateTo         time.Time            `json:"date_to"         gorm:"column:date_to;default:null"`
	EDMSLink       string               `json:"edms_link"       gorm:"column:edms_link;default:\"\""`
	ModifiedBy     employees.Employee   `json:"modified_by"     gorm:"-:all"`
	ModifiedByID   int64                `json:"modified_by_id"  gorm:"column:modified_by_id;default:null"`
	Note           string               `json:"note"            gorm:"column:note;default:\"\""`
	Reason         string               `json:"reason"          gorm:"column:reason;default:\"\""`
}
