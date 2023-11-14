//File generated automatic with crud_generator app
//Do not change anything here.

package organizations

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionOrganization - версия структуры модели, с учётом имен и типов полей
var versionOrganization uint32

// crud_Organization - объект контроллер crud операций
var crud_Organization ICrud_Organization

type ICrud_Organization interface {
	Read(o *Organization) error
	Save(o *Organization) error
	Update(o *Organization) error
	Create(o *Organization) error
	Delete(o *Organization) error
	Restore(o *Organization) error
	Find_ByExtID(o *Organization) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Organization) TableNameDB() string {
	return "organizations"
}

// NewOrganization - возвращает новый	объект
func NewOrganization() Organization {
	return Organization{}
}

// AsOrganization - создаёт объект из упакованного объекта в массиве байтов
func AsOrganization(b []byte) (Organization, error) {
	c := NewOrganization()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewOrganization(), err
	}
	return c, nil
}

// OrganizationAsBytes - упаковывает объект в массив байтов
func OrganizationAsBytes(m *Organization) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Organization) GetStructVersion() uint32 {
	if versionOrganization == 0 {
		versionOrganization = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionOrganization
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Organization) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Organization) GetJSON() (string, error) {
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
func (m *Organization) Read() error {
	err := crud_Organization.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Organization) Save() error {
	err := crud_Organization.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Organization) Update() error {
	err := crud_Organization.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Organization) Create() error {
	err := crud_Organization.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Organization) Delete() error {
	err := crud_Organization.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Organization) Restore() error {
	err := crud_Organization.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Organization) Find_ByExtID() error {
	err := crud_Organization.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Organization) SetCrudInterface(crud ICrud_Organization) {
	crud_Organization = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
