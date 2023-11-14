package organization_state_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// OrganizationStateType - модель для таблицы organization_state_types:
type OrganizationStateType struct {
	entities.CommonStruct
	entities.NameStruct
	ActionIndividual   string `json:"action_individual" gorm:"column:action_individual;default:\"\"" db:"action_individual"`       //
	ActionOrganization string `json:"action_organization" gorm:"column:action_organization;default:\"\"" db:"action_organization"` //
	Code               string `json:"code" gorm:"column:code;default:\"\"" db:"code"`                                              //
	Color              string `json:"color" gorm:"column:color;default:\"\"" db:"color"`                                           //

}
