package files

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// File - модель для таблицы files: Файлы.
type File struct {
	entities.CommonStruct
	entities.NameStruct
	entities.GroupStruct
	entities.ExtLinkStruct
	BranchID   int64  `json:"branch_id" gorm:"column:branch_id;default:0" db:"branch_id"`          //Филиал (ИД)
	EmployeeID int64  `json:"employee_id" gorm:"column:employee_id;default:0" db:"employee_id"`    //Сотрудник (ИД)
	Extension  string `json:"extension" gorm:"column:extension;default:\"\"" db:"extension"`       //Расширение файла
	FileID     string `json:"file_id" gorm:"column:file_id;default:\"\"" db:"file_id"`             //ИД в хранилище файлов
	FileName   string `json:"file_name" gorm:"column:file_name;default:\"\"" db:"file_name"`       //Краткое имя файла
	FileTypeID int64  `json:"file_type_id" gorm:"column:file_type_id;default:0" db:"file_type_id"` //Вид файла (ИД)
	FullName   string `json:"full_name" gorm:"column:full_name;default:\"\"" db:"full_name"`       //Полное имя файла с путём, откуда загружен
	Size       int64  `json:"size" gorm:"column:size;default:0" db:"size"`                         //Размер файла
	TemplateID int64  `json:"template_id" gorm:"column:template_id;default:0" db:"template_id"`    //Шаблон файла (ИД)
	Version    int32  `json:"version" gorm:"column:version;default:0" db:"version"`                //Версия изменения (должно увеличиваться само)

}
