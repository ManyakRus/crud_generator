//File generated automatic with crud_generator app
//Do not change anything here.

package service_providers

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionServiceProvider - версия структуры модели, с учётом имен и типов полей
var versionServiceProvider uint32

// crud_ServiceProvider - объект контроллер crud операций
var crud_ServiceProvider ICrud_ServiceProvider

type ICrud_ServiceProvider interface {
	Read(*ServiceProvider) error
	Save(*ServiceProvider) error
	Update(*ServiceProvider) error
	Create(*ServiceProvider) error
	Delete(*ServiceProvider) error
	Restore(*ServiceProvider) error
	Find_ByExtID(*ServiceProvider) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ServiceProvider) TableNameDB() string {
	return "service_providers"
}

// NewServiceProvider - возвращает новый	объект
func NewServiceProvider() ServiceProvider {
	return ServiceProvider{}
}

// AsServiceProvider - создаёт объект из упакованного объекта в массиве байтов
func AsServiceProvider(b []byte) (ServiceProvider, error) {
	c := NewServiceProvider()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewServiceProvider(), err
	}
	return c, nil
}

// ServiceProviderAsBytes - упаковывает объект в массив байтов
func ServiceProviderAsBytes(m *ServiceProvider) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ServiceProvider) GetStructVersion() uint32 {
	if versionServiceProvider == 0 {
		versionServiceProvider = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionServiceProvider
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ServiceProvider) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ServiceProvider) GetJSON() (string, error) {
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
func (m *ServiceProvider) Read() error {
	err := crud_ServiceProvider.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ServiceProvider) Save() error {
	err := crud_ServiceProvider.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ServiceProvider) Update() error {
	err := crud_ServiceProvider.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ServiceProvider) Create() error {
	err := crud_ServiceProvider.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ServiceProvider) Delete() error {
	err := crud_ServiceProvider.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ServiceProvider) Restore() error {
	err := crud_ServiceProvider.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *ServiceProvider) Find_ByExtID() error {
	err := crud_ServiceProvider.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ServiceProvider) SetCrudInterface(crud ICrud_ServiceProvider) {
	crud_ServiceProvider = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
