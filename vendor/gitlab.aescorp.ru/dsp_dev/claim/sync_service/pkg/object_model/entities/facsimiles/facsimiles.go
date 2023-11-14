package facsimiles

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// Facsimile - модель для таблицы facsimiles:
type Facsimile struct {
	entities.CommonStruct
	Branch      string `json:"branch" gorm:"column:branch;default:\"\"" db:"branch"`                //
	Contract    string `json:"contract" gorm:"column:contract;default:\"\"" db:"contract"`          //
	Department  string `json:"department" gorm:"column:department;default:\"\"" db:"department"`    //
	Post        string `json:"post" gorm:"column:post;default:\"\"" db:"post"`                      //
	Responsible string `json:"responsible" gorm:"column:responsible;default:\"\"" db:"responsible"` //

}
