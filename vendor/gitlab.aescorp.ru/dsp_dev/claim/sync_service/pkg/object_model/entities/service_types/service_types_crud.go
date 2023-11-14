//File generated automatic with crud_generator app
//Do not change anything here.

package service_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionServiceType - версия структуры модели, с учётом имен и типов полей
var versionServiceType uint32

// crud_ServiceType - объект контроллер crud операций
var crud_ServiceType ICrud_ServiceType

type ICrud_ServiceType interface {
	Read(*ServiceType) error
	Save(*ServiceType) error
	Update(*ServiceType) error
	Create(*ServiceType) error
	Delete(*ServiceType) error
	Restore(*ServiceType) error
	Find_ByExtID(*ServiceType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ServiceType) TableNameDB() string {
	return "service_types"
}

// NewServiceType - возвращает новый	объект
func NewServiceType() ServiceType {
	return ServiceType{}
}

// AsServiceType - создаёт объект из упакованного объекта в массиве байтов
func AsServiceType(b []byte) (ServiceType, error) {
	c := NewServiceType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewServiceType(), err
	}
	return c, nil
}

// ServiceTypeAsBytes - упаковывает объект в массив байтов
func ServiceTypeAsBytes(m *ServiceType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ServiceType) GetStructVersion() uint32 {
	if versionServiceType == 0 {
		versionServiceType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionServiceType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ServiceType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ServiceType) GetJSON() (string, error) {
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
func (m *ServiceType) Read() error {
	err := crud_ServiceType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ServiceType) Save() error {
	err := crud_ServiceType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ServiceType) Update() error {
	err := crud_ServiceType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ServiceType) Create() error {
	err := crud_ServiceType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ServiceType) Delete() error {
	err := crud_ServiceType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ServiceType) Restore() error {
	err := crud_ServiceType.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *ServiceType) Find_ByExtID() error {
	err := crud_ServiceType.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ServiceType) SetCrudInterface(crud ICrud_ServiceType) {
	crud_ServiceType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
