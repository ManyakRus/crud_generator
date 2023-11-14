//File generated automatic with crud_generator app
//Do not change anything here.

package organization_state_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionOrganizationStateType - версия структуры модели, с учётом имен и типов полей
var versionOrganizationStateType uint32

// crud_OrganizationStateType - объект контроллер crud операций
var crud_OrganizationStateType ICrud_OrganizationStateType

type ICrud_OrganizationStateType interface {
	Read(*OrganizationStateType) error
	Save(*OrganizationStateType) error
	Update(*OrganizationStateType) error
	Create(*OrganizationStateType) error
	Delete(*OrganizationStateType) error
	Restore(*OrganizationStateType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m OrganizationStateType) TableNameDB() string {
	return "organization_state_types"
}

// NewOrganizationStateType - возвращает новый	объект
func NewOrganizationStateType() OrganizationStateType {
	return OrganizationStateType{}
}

// AsOrganizationStateType - создаёт объект из упакованного объекта в массиве байтов
func AsOrganizationStateType(b []byte) (OrganizationStateType, error) {
	c := NewOrganizationStateType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewOrganizationStateType(), err
	}
	return c, nil
}

// OrganizationStateTypeAsBytes - упаковывает объект в массив байтов
func OrganizationStateTypeAsBytes(m *OrganizationStateType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m OrganizationStateType) GetStructVersion() uint32 {
	if versionOrganizationStateType == 0 {
		versionOrganizationStateType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionOrganizationStateType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *OrganizationStateType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m OrganizationStateType) GetJSON() (string, error) {
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
func (m *OrganizationStateType) Read() error {
	err := crud_OrganizationStateType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *OrganizationStateType) Save() error {
	err := crud_OrganizationStateType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *OrganizationStateType) Update() error {
	err := crud_OrganizationStateType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *OrganizationStateType) Create() error {
	err := crud_OrganizationStateType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *OrganizationStateType) Delete() error {
	err := crud_OrganizationStateType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *OrganizationStateType) Restore() error {
	err := crud_OrganizationStateType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m OrganizationStateType) SetCrudInterface(crud ICrud_OrganizationStateType) {
	crud_OrganizationStateType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
