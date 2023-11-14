//File generated automatic with crud_generator app
//Do not change anything here.

package invoice_debt_type_stats

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionInvoiceDebtTypeStat - версия структуры модели, с учётом имен и типов полей
var versionInvoiceDebtTypeStat uint32

// crud_InvoiceDebtTypeStat - объект контроллер crud операций
var crud_InvoiceDebtTypeStat ICrud_InvoiceDebtTypeStat

type ICrud_InvoiceDebtTypeStat interface {
	Read(*InvoiceDebtTypeStat) error
	Save(*InvoiceDebtTypeStat) error
	Update(*InvoiceDebtTypeStat) error
	Create(*InvoiceDebtTypeStat) error
	Delete(*InvoiceDebtTypeStat) error
	Restore(*InvoiceDebtTypeStat) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m InvoiceDebtTypeStat) TableNameDB() string {
	return "invoice_debt_type_stats"
}

// NewInvoiceDebtTypeStat - возвращает новый	объект
func NewInvoiceDebtTypeStat() InvoiceDebtTypeStat {
	return InvoiceDebtTypeStat{}
}

// AsInvoiceDebtTypeStat - создаёт объект из упакованного объекта в массиве байтов
func AsInvoiceDebtTypeStat(b []byte) (InvoiceDebtTypeStat, error) {
	c := NewInvoiceDebtTypeStat()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewInvoiceDebtTypeStat(), err
	}
	return c, nil
}

// InvoiceDebtTypeStatAsBytes - упаковывает объект в массив байтов
func InvoiceDebtTypeStatAsBytes(m *InvoiceDebtTypeStat) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m InvoiceDebtTypeStat) GetStructVersion() uint32 {
	if versionInvoiceDebtTypeStat == 0 {
		versionInvoiceDebtTypeStat = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionInvoiceDebtTypeStat
}

// GetModelFromJSON - создаёт модель из строки json
func (m *InvoiceDebtTypeStat) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m InvoiceDebtTypeStat) GetJSON() (string, error) {
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
func (m *InvoiceDebtTypeStat) Read() error {
	err := crud_InvoiceDebtTypeStat.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *InvoiceDebtTypeStat) Save() error {
	err := crud_InvoiceDebtTypeStat.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *InvoiceDebtTypeStat) Update() error {
	err := crud_InvoiceDebtTypeStat.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *InvoiceDebtTypeStat) Create() error {
	err := crud_InvoiceDebtTypeStat.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *InvoiceDebtTypeStat) Delete() error {
	err := crud_InvoiceDebtTypeStat.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *InvoiceDebtTypeStat) Restore() error {
	err := crud_InvoiceDebtTypeStat.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m InvoiceDebtTypeStat) SetCrudInterface(crud ICrud_InvoiceDebtTypeStat) {
	crud_InvoiceDebtTypeStat = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
