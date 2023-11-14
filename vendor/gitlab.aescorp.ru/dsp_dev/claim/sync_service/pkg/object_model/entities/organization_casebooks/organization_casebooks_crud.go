//File generated automatic with crud_generator app
//Do not change anything here.

package organization_casebooks

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionOrganizationCasebook - версия структуры модели, с учётом имен и типов полей
var versionOrganizationCasebook uint32

// crud_OrganizationCasebook - объект контроллер crud операций
var crud_OrganizationCasebook ICrud_OrganizationCasebook

type ICrud_OrganizationCasebook interface {
	Read(o *OrganizationCasebook) error
	Save(o *OrganizationCasebook) error
	Update(o *OrganizationCasebook) error
	Create(o *OrganizationCasebook) error
	Delete(o *OrganizationCasebook) error
	Restore(o *OrganizationCasebook) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m OrganizationCasebook) TableNameDB() string {
	return "organization_casebooks"
}

// NewOrganizationCasebook - возвращает новый	объект
func NewOrganizationCasebook() OrganizationCasebook {
	return OrganizationCasebook{}
}

// AsOrganizationCasebook - создаёт объект из упакованного объекта в массиве байтов
func AsOrganizationCasebook(b []byte) (OrganizationCasebook, error) {
	c := NewOrganizationCasebook()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewOrganizationCasebook(), err
	}
	return c, nil
}

// OrganizationCasebookAsBytes - упаковывает объект в массив байтов
func OrganizationCasebookAsBytes(m *OrganizationCasebook) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m OrganizationCasebook) GetStructVersion() uint32 {
	if versionOrganizationCasebook == 0 {
		versionOrganizationCasebook = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionOrganizationCasebook
}

// GetModelFromJSON - создаёт модель из строки json
func (m *OrganizationCasebook) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m OrganizationCasebook) GetJSON() (string, error) {
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
func (m *OrganizationCasebook) Read() error {
	err := crud_OrganizationCasebook.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *OrganizationCasebook) Save() error {
	err := crud_OrganizationCasebook.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *OrganizationCasebook) Update() error {
	err := crud_OrganizationCasebook.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *OrganizationCasebook) Create() error {
	err := crud_OrganizationCasebook.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *OrganizationCasebook) Delete() error {
	err := crud_OrganizationCasebook.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *OrganizationCasebook) Restore() error {
	err := crud_OrganizationCasebook.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m OrganizationCasebook) SetCrudInterface(crud ICrud_OrganizationCasebook) {
	crud_OrganizationCasebook = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
