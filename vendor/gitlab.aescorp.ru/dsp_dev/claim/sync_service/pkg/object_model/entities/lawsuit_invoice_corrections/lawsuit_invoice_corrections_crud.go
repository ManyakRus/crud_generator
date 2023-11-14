//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_invoice_corrections

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitInvoiceCorrection - версия структуры модели, с учётом имен и типов полей
var versionLawsuitInvoiceCorrection uint32

// crud_LawsuitInvoiceCorrection - объект контроллер crud операций
var crud_LawsuitInvoiceCorrection ICrud_LawsuitInvoiceCorrection

type ICrud_LawsuitInvoiceCorrection interface {
	Read(*LawsuitInvoiceCorrection) error
	Save(*LawsuitInvoiceCorrection) error
	Update(*LawsuitInvoiceCorrection) error
	Create(*LawsuitInvoiceCorrection) error
	Delete(*LawsuitInvoiceCorrection) error
	Restore(*LawsuitInvoiceCorrection) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitInvoiceCorrection) TableNameDB() string {
	return "lawsuit_invoice_corrections"
}

// NewLawsuitInvoiceCorrection - возвращает новый	объект
func NewLawsuitInvoiceCorrection() LawsuitInvoiceCorrection {
	return LawsuitInvoiceCorrection{}
}

// AsLawsuitInvoiceCorrection - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitInvoiceCorrection(b []byte) (LawsuitInvoiceCorrection, error) {
	c := NewLawsuitInvoiceCorrection()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitInvoiceCorrection(), err
	}
	return c, nil
}

// LawsuitInvoiceCorrectionAsBytes - упаковывает объект в массив байтов
func LawsuitInvoiceCorrectionAsBytes(m *LawsuitInvoiceCorrection) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitInvoiceCorrection) GetStructVersion() uint32 {
	if versionLawsuitInvoiceCorrection == 0 {
		versionLawsuitInvoiceCorrection = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitInvoiceCorrection
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitInvoiceCorrection) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitInvoiceCorrection) GetJSON() (string, error) {
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
func (m *LawsuitInvoiceCorrection) Read() error {
	err := crud_LawsuitInvoiceCorrection.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitInvoiceCorrection) Save() error {
	err := crud_LawsuitInvoiceCorrection.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitInvoiceCorrection) Update() error {
	err := crud_LawsuitInvoiceCorrection.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitInvoiceCorrection) Create() error {
	err := crud_LawsuitInvoiceCorrection.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitInvoiceCorrection) Delete() error {
	err := crud_LawsuitInvoiceCorrection.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitInvoiceCorrection) Restore() error {
	err := crud_LawsuitInvoiceCorrection.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitInvoiceCorrection) SetCrudInterface(crud ICrud_LawsuitInvoiceCorrection) {
	crud_LawsuitInvoiceCorrection = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
