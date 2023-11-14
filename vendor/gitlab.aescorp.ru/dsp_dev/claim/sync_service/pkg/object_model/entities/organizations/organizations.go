package organizations

import (
	"time"

	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_state_types"
)

// Organization Юридическое лицо (справочник).
type Organization struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	BankruptAt     time.Time                                      `json:"bankrupt_at"     gorm:"column:bankrupt_at"`
	BookkeeperName string                                         `json:"bookkeeper_name" gorm:"column:bookkeeper_name;default:\"\""`
	CategoryID     int64                                          `json:"category_id"     gorm:"column:category_id;default:null"`
	ConnectionID   int64                                          `json:"connection_id"   gorm:"column:connection_id;default:null"`
	Email          string                                         `json:"email"           gorm:"column:email;default:\"\""`
	FullName       string                                         `json:"full_name"       gorm:"column:full_name;default:\"\""`
	INN            string                                         `json:"inn"             gorm:"column:inn;default:\"\""`
	IsActive       bool                                           `json:"is_active"       gorm:"column:is_active;default:false"`
	IsBankrupt     bool                                           `json:"is_bankrupt"     gorm:"column:is_bankrupt;default:false"`
	IsLiquidated   bool                                           `json:"is_liquidated"   gorm:"column:is_liquidated;default:false"`
	KPP            string                                         `json:"kpp"             gorm:"column:kpp;default:\"\""`
	LegalAddress   string                                         `json:"legal_address"   gorm:"column:legal_address;default:\"\""`
	LegalTypeID    int64                                          `json:"legal_type_id"   gorm:"column:legal_type_id;default:0"`
	LiquidateAt    time.Time                                      `json:"liquidate_at"    gorm:"column:liquidate_at"`
	ManagerName    string                                         `json:"manager_name"    gorm:"column:manager_name;default:\"\""`
	NSIFlat        string                                         `json:"nsi_flat"        gorm:"column:nsi_flat;default:\"\""` // Значение квартиры из НСИ
	NSIFlatID      int64                                          `json:"nsi_flat_id"     gorm:"column:nsi_flat_id;default:0"` // ИД типа квартиры из НСИ
	NSIID          int64                                          `json:"nsi_id"          gorm:"column:nsi_id;default:0"`      // ИД адреса из НСИ
	OGRN           string                                         `json:"ogrn"            gorm:"column:ogrn;default:\"\""`
	OKATO          string                                         `json:"okato"           gorm:"column:okato;default:\"\""`
	OKPO           string                                         `json:"okpo"            gorm:"column:okpo;default:\"\""`
	Phone          string                                         `json:"phone"           gorm:"column:phone;default:\"\""`
	PostAddress    string                                         `json:"post_address"    gorm:"column:post_address;default:\"\""`
	RegistrationAt time.Time                                      `json:"registration_at" gorm:"column:registration_at;default:null"`
	State          organization_state_types.OrganizationStateType `json:"state"           gorm:"-:all"`                          // Статус организации из НСИ.
	StateCode      string                                         `json:"state_code"      gorm:"column:state_code;default:\"\""` // Код статуса организации из НСИ.
	StateID        int64                                          `json:"state_id"        gorm:"column:state_id;default:null"`   // ID статуса организации из НСИ.
	WWW            string                                         `json:"www"             gorm:"column:www;default:\"\""`
}
