package document_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// DocumentType - модель для таблицы document_types: Типы документов (справочник).
type DocumentType struct {
	entities.CommonStruct
	entities.NameStruct
	ConnectionID  int64  `json:"connection_id" gorm:"column:connection_id;default:0" db:"connection_id"`    //Соединение к БД СТЕК (ИД)
	IncomeExpense int32  `json:"income_expense" gorm:"column:income_expense;default:0" db:"income_expense"` //ПриходРасход = +1 (нам платят) или -1 (нам должны) или 0 (сам придумал :-),
	IsService     bool   `json:"is_service" gorm:"column:is_service" db:"is_service"`                       //Служебный
	IsVisible     bool   `json:"is_visible" gorm:"column:is_visible" db:"is_visible"`                       //Скрытый
	ShortName     string `json:"short_name" gorm:"column:short_name;default:\"\"" db:"short_name"`          //Наименование краткое
	Type          int32  `json:"type" gorm:"column:type;default:0" db:"type"`                               //Тип (код)

}
