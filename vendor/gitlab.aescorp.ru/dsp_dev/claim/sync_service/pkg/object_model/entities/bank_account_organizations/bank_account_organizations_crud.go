//File generated automatic with crud_generator app
//Do not change anything here.

package bank_account_organizations

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionBankAccountOrganization - версия структуры модели, с учётом имен и типов полей
var versionBankAccountOrganization uint32

// crud_BankAccountOrganization - объект контроллер crud операций
var crud_BankAccountOrganization ICrud_BankAccountOrganization

type ICrud_BankAccountOrganization interface {
	Read(*BankAccountOrganization) error
	Save(*BankAccountOrganization) error
	Update(*BankAccountOrganization) error
	Create(*BankAccountOrganization) error
	Delete(*BankAccountOrganization) error
	Restore(*BankAccountOrganization) error
	Find_ByExtID(*BankAccountOrganization) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m BankAccountOrganization) TableNameDB() string {
	return "bank_account_organizations"
}

// NewBankAccountOrganization - возвращает новый	объект
func NewBankAccountOrganization() BankAccountOrganization {
	return BankAccountOrganization{}
}

// AsBankAccountOrganization - создаёт объект из упакованного объекта в массиве байтов
func AsBankAccountOrganization(b []byte) (BankAccountOrganization, error) {
	c := NewBankAccountOrganization()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewBankAccountOrganization(), err
	}
	return c, nil
}

// BankAccountOrganizationAsBytes - упаковывает объект в массив байтов
func BankAccountOrganizationAsBytes(m *BankAccountOrganization) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m BankAccountOrganization) GetStructVersion() uint32 {
	if versionBankAccountOrganization == 0 {
		versionBankAccountOrganization = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionBankAccountOrganization
}

// GetModelFromJSON - создаёт модель из строки json
func (m *BankAccountOrganization) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m BankAccountOrganization) GetJSON() (string, error) {
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
func (m *BankAccountOrganization) Read() error {
	err := crud_BankAccountOrganization.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *BankAccountOrganization) Save() error {
	err := crud_BankAccountOrganization.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *BankAccountOrganization) Update() error {
	err := crud_BankAccountOrganization.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *BankAccountOrganization) Create() error {
	err := crud_BankAccountOrganization.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *BankAccountOrganization) Delete() error {
	err := crud_BankAccountOrganization.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *BankAccountOrganization) Restore() error {
	err := crud_BankAccountOrganization.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *BankAccountOrganization) Find_ByExtID() error {
	err := crud_BankAccountOrganization.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m BankAccountOrganization) SetCrudInterface(crud ICrud_BankAccountOrganization) {
	crud_BankAccountOrganization = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
