//File generated automatic with crud_generator app
//Do not change anything here.

package banks

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionBank - версия структуры модели, с учётом имен и типов полей
var versionBank uint32

// crud_Bank - объект контроллер crud операций
var crud_Bank ICrud_Bank

type ICrud_Bank interface {
	Read(*Bank) error
	Save(*Bank) error
	Update(*Bank) error
	Create(*Bank) error
	Delete(*Bank) error
	Restore(*Bank) error
	Find_ByExtID(*Bank) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Bank) TableNameDB() string {
	return "banks"
}

// NewBank - возвращает новый	объект
func NewBank() Bank {
	return Bank{}
}

// AsBank - создаёт объект из упакованного объекта в массиве байтов
func AsBank(b []byte) (Bank, error) {
	c := NewBank()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewBank(), err
	}
	return c, nil
}

// BankAsBytes - упаковывает объект в массив байтов
func BankAsBytes(m *Bank) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Bank) GetStructVersion() uint32 {
	if versionBank == 0 {
		versionBank = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionBank
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Bank) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Bank) GetJSON() (string, error) {
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
func (m *Bank) Read() error {
	err := crud_Bank.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Bank) Save() error {
	err := crud_Bank.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Bank) Update() error {
	err := crud_Bank.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Bank) Create() error {
	err := crud_Bank.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Bank) Delete() error {
	err := crud_Bank.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Bank) Restore() error {
	err := crud_Bank.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Bank) Find_ByExtID() error {
	err := crud_Bank.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Bank) SetCrudInterface(crud ICrud_Bank) {
	crud_Bank = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
