//File generated automatic with crud_generator app
//Do not change anything here.

package payment_schedules

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionPaymentSchedule - версия структуры модели, с учётом имен и типов полей
var versionPaymentSchedule uint32

// crud_PaymentSchedule - объект контроллер crud операций
var crud_PaymentSchedule ICrud_PaymentSchedule

type ICrud_PaymentSchedule interface {
	Read(*PaymentSchedule) error
	Save(*PaymentSchedule) error
	Update(*PaymentSchedule) error
	Create(*PaymentSchedule) error
	Delete(*PaymentSchedule) error
	Restore(*PaymentSchedule) error
	Find_ByExtID(*PaymentSchedule) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m PaymentSchedule) TableNameDB() string {
	return "payment_schedules"
}

// NewPaymentSchedule - возвращает новый	объект
func NewPaymentSchedule() PaymentSchedule {
	return PaymentSchedule{}
}

// AsPaymentSchedule - создаёт объект из упакованного объекта в массиве байтов
func AsPaymentSchedule(b []byte) (PaymentSchedule, error) {
	c := NewPaymentSchedule()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPaymentSchedule(), err
	}
	return c, nil
}

// PaymentScheduleAsBytes - упаковывает объект в массив байтов
func PaymentScheduleAsBytes(m *PaymentSchedule) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PaymentSchedule) GetStructVersion() uint32 {
	if versionPaymentSchedule == 0 {
		versionPaymentSchedule = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPaymentSchedule
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PaymentSchedule) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PaymentSchedule) GetJSON() (string, error) {
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
func (m *PaymentSchedule) Read() error {
	err := crud_PaymentSchedule.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PaymentSchedule) Save() error {
	err := crud_PaymentSchedule.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PaymentSchedule) Update() error {
	err := crud_PaymentSchedule.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PaymentSchedule) Create() error {
	err := crud_PaymentSchedule.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PaymentSchedule) Delete() error {
	err := crud_PaymentSchedule.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PaymentSchedule) Restore() error {
	err := crud_PaymentSchedule.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *PaymentSchedule) Find_ByExtID() error {
	err := crud_PaymentSchedule.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PaymentSchedule) SetCrudInterface(crud ICrud_PaymentSchedule) {
	crud_PaymentSchedule = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
