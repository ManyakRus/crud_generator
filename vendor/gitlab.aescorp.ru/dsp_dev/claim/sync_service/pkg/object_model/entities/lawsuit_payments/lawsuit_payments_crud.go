//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_payments

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitPayment - версия структуры модели, с учётом имен и типов полей
var versionLawsuitPayment uint32

// crud_LawsuitPayment - объект контроллер crud операций
var crud_LawsuitPayment ICrud_LawsuitPayment

type ICrud_LawsuitPayment interface {
	Read(*LawsuitPayment) error
	Save(*LawsuitPayment) error
	Update(*LawsuitPayment) error
	Create(*LawsuitPayment) error
	Delete(*LawsuitPayment) error
	Restore(*LawsuitPayment) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitPayment) TableNameDB() string {
	return "lawsuit_payments"
}

// NewLawsuitPayment - возвращает новый	объект
func NewLawsuitPayment() LawsuitPayment {
	return LawsuitPayment{}
}

// AsLawsuitPayment - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitPayment(b []byte) (LawsuitPayment, error) {
	c := NewLawsuitPayment()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitPayment(), err
	}
	return c, nil
}

// LawsuitPaymentAsBytes - упаковывает объект в массив байтов
func LawsuitPaymentAsBytes(m *LawsuitPayment) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitPayment) GetStructVersion() uint32 {
	if versionLawsuitPayment == 0 {
		versionLawsuitPayment = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitPayment
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitPayment) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitPayment) GetJSON() (string, error) {
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
func (m *LawsuitPayment) Read() error {
	err := crud_LawsuitPayment.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitPayment) Save() error {
	err := crud_LawsuitPayment.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitPayment) Update() error {
	err := crud_LawsuitPayment.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitPayment) Create() error {
	err := crud_LawsuitPayment.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitPayment) Delete() error {
	err := crud_LawsuitPayment.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitPayment) Restore() error {
	err := crud_LawsuitPayment.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitPayment) SetCrudInterface(crud ICrud_LawsuitPayment) {
	crud_LawsuitPayment = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
