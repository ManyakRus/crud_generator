//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_invoices

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitInvoice - версия структуры модели, с учётом имен и типов полей
var versionLawsuitInvoice uint32

// crud_LawsuitInvoice - объект контроллер crud операций
var crud_LawsuitInvoice ICrud_LawsuitInvoice

type ICrud_LawsuitInvoice interface {
	Read(*LawsuitInvoice) error
	Save(*LawsuitInvoice) error
	Update(*LawsuitInvoice) error
	Create(*LawsuitInvoice) error
	Delete(*LawsuitInvoice) error
	Restore(*LawsuitInvoice) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitInvoice) TableNameDB() string {
	return "lawsuit_invoices"
}

// NewLawsuitInvoice - возвращает новый	объект
func NewLawsuitInvoice() LawsuitInvoice {
	return LawsuitInvoice{}
}

// AsLawsuitInvoice - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitInvoice(b []byte) (LawsuitInvoice, error) {
	c := NewLawsuitInvoice()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitInvoice(), err
	}
	return c, nil
}

// LawsuitInvoiceAsBytes - упаковывает объект в массив байтов
func LawsuitInvoiceAsBytes(m *LawsuitInvoice) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitInvoice) GetStructVersion() uint32 {
	if versionLawsuitInvoice == 0 {
		versionLawsuitInvoice = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitInvoice
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitInvoice) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitInvoice) GetJSON() (string, error) {
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
func (m *LawsuitInvoice) Read() error {
	err := crud_LawsuitInvoice.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitInvoice) Save() error {
	err := crud_LawsuitInvoice.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitInvoice) Update() error {
	err := crud_LawsuitInvoice.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitInvoice) Create() error {
	err := crud_LawsuitInvoice.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitInvoice) Delete() error {
	err := crud_LawsuitInvoice.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitInvoice) Restore() error {
	err := crud_LawsuitInvoice.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitInvoice) SetCrudInterface(crud ICrud_LawsuitInvoice) {
	crud_LawsuitInvoice = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
