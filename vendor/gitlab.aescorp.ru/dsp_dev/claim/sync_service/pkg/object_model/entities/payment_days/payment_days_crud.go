//File generated automatic with crud_generator app
//Do not change anything here.

package payment_days

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionPaymentDay - версия структуры модели, с учётом имен и типов полей
var versionPaymentDay uint32

// crud_PaymentDay - объект контроллер crud операций
var crud_PaymentDay ICrud_PaymentDay

type ICrud_PaymentDay interface {
	Read(*PaymentDay) error
	Save(*PaymentDay) error
	Update(*PaymentDay) error
	Create(*PaymentDay) error
	Delete(*PaymentDay) error
	Restore(*PaymentDay) error
	Find_ByExtID(*PaymentDay) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m PaymentDay) TableNameDB() string {
	return "payment_days"
}

// NewPaymentDay - возвращает новый	объект
func NewPaymentDay() PaymentDay {
	return PaymentDay{}
}

// AsPaymentDay - создаёт объект из упакованного объекта в массиве байтов
func AsPaymentDay(b []byte) (PaymentDay, error) {
	c := NewPaymentDay()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPaymentDay(), err
	}
	return c, nil
}

// PaymentDayAsBytes - упаковывает объект в массив байтов
func PaymentDayAsBytes(m *PaymentDay) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PaymentDay) GetStructVersion() uint32 {
	if versionPaymentDay == 0 {
		versionPaymentDay = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPaymentDay
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PaymentDay) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PaymentDay) GetJSON() (string, error) {
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
func (m *PaymentDay) Read() error {
	err := crud_PaymentDay.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PaymentDay) Save() error {
	err := crud_PaymentDay.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PaymentDay) Update() error {
	err := crud_PaymentDay.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PaymentDay) Create() error {
	err := crud_PaymentDay.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PaymentDay) Delete() error {
	err := crud_PaymentDay.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PaymentDay) Restore() error {
	err := crud_PaymentDay.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *PaymentDay) Find_ByExtID() error {
	err := crud_PaymentDay.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PaymentDay) SetCrudInterface(crud ICrud_PaymentDay) {
	crud_PaymentDay = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
