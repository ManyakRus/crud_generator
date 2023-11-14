//File generated automatic with crud_generator app
//Do not change anything here.

package contract_black_items

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionContractBlackItem - версия структуры модели, с учётом имен и типов полей
var versionContractBlackItem uint32

// crud_ContractBlackItem - объект контроллер crud операций
var crud_ContractBlackItem ICrud_ContractBlackItem

type ICrud_ContractBlackItem interface {
	Read(*ContractBlackItem) error
	Save(*ContractBlackItem) error
	Update(*ContractBlackItem) error
	Create(*ContractBlackItem) error
	Delete(*ContractBlackItem) error
	Restore(*ContractBlackItem) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ContractBlackItem) TableNameDB() string {
	return "contract_black_items"
}

// NewContractBlackItem - возвращает новый	объект
func NewContractBlackItem() ContractBlackItem {
	return ContractBlackItem{}
}

// AsContractBlackItem - создаёт объект из упакованного объекта в массиве байтов
func AsContractBlackItem(b []byte) (ContractBlackItem, error) {
	c := NewContractBlackItem()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewContractBlackItem(), err
	}
	return c, nil
}

// ContractBlackItemAsBytes - упаковывает объект в массив байтов
func ContractBlackItemAsBytes(m *ContractBlackItem) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ContractBlackItem) GetStructVersion() uint32 {
	if versionContractBlackItem == 0 {
		versionContractBlackItem = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionContractBlackItem
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ContractBlackItem) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ContractBlackItem) GetJSON() (string, error) {
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
func (m *ContractBlackItem) Read() error {
	err := crud_ContractBlackItem.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ContractBlackItem) Save() error {
	err := crud_ContractBlackItem.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ContractBlackItem) Update() error {
	err := crud_ContractBlackItem.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ContractBlackItem) Create() error {
	err := crud_ContractBlackItem.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ContractBlackItem) Delete() error {
	err := crud_ContractBlackItem.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ContractBlackItem) Restore() error {
	err := crud_ContractBlackItem.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ContractBlackItem) SetCrudInterface(crud ICrud_ContractBlackItem) {
	crud_ContractBlackItem = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
