//File generated automatic with crud_generator app
//Do not change anything here.

package file_templates

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionFileTemplate - версия структуры модели, с учётом имен и типов полей
var versionFileTemplate uint32

// crud_FileTemplate - объект контроллер crud операций
var crud_FileTemplate ICrud_FileTemplate

type ICrud_FileTemplate interface {
	Read(*FileTemplate) error
	Save(*FileTemplate) error
	Update(*FileTemplate) error
	Create(*FileTemplate) error
	Delete(*FileTemplate) error
	Restore(*FileTemplate) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m FileTemplate) TableNameDB() string {
	return "file_templates"
}

// NewFileTemplate - возвращает новый	объект
func NewFileTemplate() FileTemplate {
	return FileTemplate{}
}

// AsFileTemplate - создаёт объект из упакованного объекта в массиве байтов
func AsFileTemplate(b []byte) (FileTemplate, error) {
	c := NewFileTemplate()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewFileTemplate(), err
	}
	return c, nil
}

// FileTemplateAsBytes - упаковывает объект в массив байтов
func FileTemplateAsBytes(m *FileTemplate) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m FileTemplate) GetStructVersion() uint32 {
	if versionFileTemplate == 0 {
		versionFileTemplate = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionFileTemplate
}

// GetModelFromJSON - создаёт модель из строки json
func (m *FileTemplate) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m FileTemplate) GetJSON() (string, error) {
	var ReturnVar string
	var err error

	bytes, err := json.Marshal(m)
	if err != nil {
		return ReturnVar, err
	}
	ReturnVar = string(bytes)
	return ReturnVar, err
}

// ---------------------------- CRUD операции ------------------------------------------------------------

// Read - находит запись в БД по ID, и заполняет в объект
func (m *FileTemplate) Read() error {
	err := crud_FileTemplate.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *FileTemplate) Save() error {
	err := crud_FileTemplate.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *FileTemplate) Update() error {
	err := crud_FileTemplate.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *FileTemplate) Create() error {
	err := crud_FileTemplate.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *FileTemplate) Delete() error {
	err := crud_FileTemplate.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *FileTemplate) Restore() error {
	err := crud_FileTemplate.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m FileTemplate) SetCrudInterface(crud ICrud_FileTemplate) {
	crud_FileTemplate = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
