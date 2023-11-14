package file_changes

import (
	"time"
)

// FileChange - модель для таблицы file_changes: Журнал изменений файлов.
type FileChange struct {
	Description string    `json:"description" gorm:"column:description;default:\"\"" db:"description"`                //Описание
	EmployeeID  int64     `json:"employee_id" gorm:"column:employee_id;default:0" db:"employee_id"`                   //Сотрудник (ИД)
	FileID      int64     `json:"file_id" gorm:"column:file_id;default:0" db:"file_id"`                               //Файл (ИД)
	ID          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;default:0" db:"id"`                //Уникальный технический идентификатор
	ModifiedAt  time.Time `json:"modified_at" gorm:"column:modified_at;default:null;autoUpdateTime" db:"modified_at"` //Дата изменения элемента

}
