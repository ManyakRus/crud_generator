//File generated automatic with crud_generator app
//Do not change anything here.

package balances

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionBalance - версия структуры модели, с учётом имен и типов полей
var versionBalance uint32

// crud_Balance - объект контроллер crud операций
var crud_Balance ICrud_Balance

type ICrud_Balance interface {
	Read(*Balance) error
	Save(*Balance) error
	Update(*Balance) error
	Create(*Balance) error
	Delete(*Balance) error
	Restore(*Balance) error
	Find_ByExtID(*Balance) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Balance) TableNameDB() string {
	return "balances"
}

// NewBalance - возвращает новый	объект
func NewBalance() Balance {
	return Balance{}
}

// AsBalance - создаёт объект из упакованного объекта в массиве байтов
func AsBalance(b []byte) (Balance, error) {
	c := NewBalance()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewBalance(), err
	}
	return c, nil
}

// BalanceAsBytes - упаковывает объект в массив байтов
func BalanceAsBytes(m *Balance) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Balance) GetStructVersion() uint32 {
	if versionBalance == 0 {
		versionBalance = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionBalance
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Balance) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Balance) GetJSON() (string, error) {
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
func (m *Balance) Read() error {
	err := crud_Balance.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Balance) Save() error {
	err := crud_Balance.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Balance) Update() error {
	err := crud_Balance.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Balance) Create() error {
	err := crud_Balance.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Balance) Delete() error {
	err := crud_Balance.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Balance) Restore() error {
	err := crud_Balance.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Balance) Find_ByExtID() error {
	err := crud_Balance.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Balance) SetCrudInterface(crud ICrud_Balance) {
	crud_Balance = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
