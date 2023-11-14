package debt_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// DebtType - модель для таблицы debt_types: Виды задолженности (Справочник).
type DebtType struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	CodeNSI      int32 `json:"code_nsi" gorm:"column:code_nsi;default:0" db:"code_nsi"`                //Код НСИ
	ConnectionID int64 `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"` //Соединение к БД СТЕК (ИД)
	ExtCode      int64 `json:"ext_code" gorm:"column:ext_code;default:0" db:"ext_code"`                //Код

}
