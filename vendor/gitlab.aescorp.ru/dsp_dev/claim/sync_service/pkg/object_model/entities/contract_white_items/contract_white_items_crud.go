//File generated automatic with crud_generator app
//Do not change anything here.

package contract_white_items

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionContractWhiteItem - версия структуры модели, с учётом имен и типов полей
var versionContractWhiteItem uint32

// crud_ContractWhiteItem - объект контроллер crud операций
var crud_ContractWhiteItem ICrud_ContractWhiteItem

type ICrud_ContractWhiteItem interface {
	Read(*ContractWhiteItem) error
	Save(*ContractWhiteItem) error
	Update(*ContractWhiteItem) error
	Create(*ContractWhiteItem) error
	Delete(*ContractWhiteItem) error
	Restore(*ContractWhiteItem) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ContractWhiteItem) TableNameDB() string {
	return "contract_white_items"
}

// NewContractWhiteItem - возвращает новый	объект
func NewContractWhiteItem() ContractWhiteItem {
	return ContractWhiteItem{}
}

// AsContractWhiteItem - создаёт объект из упакованного объекта в массиве байтов
func AsContractWhiteItem(b []byte) (ContractWhiteItem, error) {
	c := NewContractWhiteItem()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewContractWhiteItem(), err
	}
	return c, nil
}

// ContractWhiteItemAsBytes - упаковывает объект в массив байтов
func ContractWhiteItemAsBytes(m *ContractWhiteItem) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ContractWhiteItem) GetStructVersion() uint32 {
	if versionContractWhiteItem == 0 {
		versionContractWhiteItem = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionContractWhiteItem
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ContractWhiteItem) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ContractWhiteItem) GetJSON() (string, error) {
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
func (m *ContractWhiteItem) Read() error {
	err := crud_ContractWhiteItem.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ContractWhiteItem) Save() error {
	err := crud_ContractWhiteItem.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ContractWhiteItem) Update() error {
	err := crud_ContractWhiteItem.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ContractWhiteItem) Create() error {
	err := crud_ContractWhiteItem.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ContractWhiteItem) Delete() error {
	err := crud_ContractWhiteItem.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ContractWhiteItem) Restore() error {
	err := crud_ContractWhiteItem.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ContractWhiteItem) SetCrudInterface(crud ICrud_ContractWhiteItem) {
	crud_ContractWhiteItem = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
