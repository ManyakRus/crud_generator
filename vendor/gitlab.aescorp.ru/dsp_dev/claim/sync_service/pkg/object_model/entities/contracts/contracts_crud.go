//File generated automatic with crud_generator app
//Do not change anything here.

package contracts

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionContract - версия структуры модели, с учётом имен и типов полей
var versionContract uint32

// crud_Contract - объект контроллер crud операций
var crud_Contract ICrud_Contract

type ICrud_Contract interface {
	Read(*Contract) error
	Save(*Contract) error
	Update(*Contract) error
	Create(*Contract) error
	Delete(*Contract) error
	Restore(*Contract) error
	Find_ByExtID(*Contract) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Contract) TableNameDB() string {
	return "contracts"
}

// NewContract - возвращает новый	объект
func NewContract() Contract {
	return Contract{}
}

// AsContract - создаёт объект из упакованного объекта в массиве байтов
func AsContract(b []byte) (Contract, error) {
	c := NewContract()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewContract(), err
	}
	return c, nil
}

// ContractAsBytes - упаковывает объект в массив байтов
func ContractAsBytes(m *Contract) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Contract) GetStructVersion() uint32 {
	if versionContract == 0 {
		versionContract = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionContract
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Contract) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Contract) GetJSON() (string, error) {
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
func (m *Contract) Read() error {
	err := crud_Contract.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Contract) Save() error {
	err := crud_Contract.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Contract) Update() error {
	err := crud_Contract.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Contract) Create() error {
	err := crud_Contract.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Contract) Delete() error {
	err := crud_Contract.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Contract) Restore() error {
	err := crud_Contract.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Contract) Find_ByExtID() error {
	err := crud_Contract.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Contract) SetCrudInterface(crud ICrud_Contract) {
	crud_Contract = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
