//File generated automatic with crud_generator app
//Do not change anything here.

package contract_category_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionContractCategoryType - версия структуры модели, с учётом имен и типов полей
var versionContractCategoryType uint32

// crud_ContractCategoryType - объект контроллер crud операций
var crud_ContractCategoryType ICrud_ContractCategoryType

type ICrud_ContractCategoryType interface {
	Read(*ContractCategoryType) error
	Save(*ContractCategoryType) error
	Update(*ContractCategoryType) error
	Create(*ContractCategoryType) error
	Delete(*ContractCategoryType) error
	Restore(*ContractCategoryType) error
	Find_ByExtID(*ContractCategoryType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ContractCategoryType) TableNameDB() string {
	return "contract_category_types"
}

// NewContractCategoryType - возвращает новый	объект
func NewContractCategoryType() ContractCategoryType {
	return ContractCategoryType{}
}

// AsContractCategoryType - создаёт объект из упакованного объекта в массиве байтов
func AsContractCategoryType(b []byte) (ContractCategoryType, error) {
	c := NewContractCategoryType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewContractCategoryType(), err
	}
	return c, nil
}

// ContractCategoryTypeAsBytes - упаковывает объект в массив байтов
func ContractCategoryTypeAsBytes(m *ContractCategoryType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ContractCategoryType) GetStructVersion() uint32 {
	if versionContractCategoryType == 0 {
		versionContractCategoryType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionContractCategoryType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ContractCategoryType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ContractCategoryType) GetJSON() (string, error) {
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
func (m *ContractCategoryType) Read() error {
	err := crud_ContractCategoryType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ContractCategoryType) Save() error {
	err := crud_ContractCategoryType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ContractCategoryType) Update() error {
	err := crud_ContractCategoryType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ContractCategoryType) Create() error {
	err := crud_ContractCategoryType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ContractCategoryType) Delete() error {
	err := crud_ContractCategoryType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ContractCategoryType) Restore() error {
	err := crud_ContractCategoryType.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *ContractCategoryType) Find_ByExtID() error {
	err := crud_ContractCategoryType.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ContractCategoryType) SetCrudInterface(crud ICrud_ContractCategoryType) {
	crud_ContractCategoryType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
