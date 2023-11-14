//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_payment_corrections

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitPaymentCorrection - версия структуры модели, с учётом имен и типов полей
var versionLawsuitPaymentCorrection uint32

// crud_LawsuitPaymentCorrection - объект контроллер crud операций
var crud_LawsuitPaymentCorrection ICrud_LawsuitPaymentCorrection

type ICrud_LawsuitPaymentCorrection interface {
	Read(*LawsuitPaymentCorrection) error
	Save(*LawsuitPaymentCorrection) error
	Update(*LawsuitPaymentCorrection) error
	Create(*LawsuitPaymentCorrection) error
	Delete(*LawsuitPaymentCorrection) error
	Restore(*LawsuitPaymentCorrection) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitPaymentCorrection) TableNameDB() string {
	return "lawsuit_payment_corrections"
}

// NewLawsuitPaymentCorrection - возвращает новый	объект
func NewLawsuitPaymentCorrection() LawsuitPaymentCorrection {
	return LawsuitPaymentCorrection{}
}

// AsLawsuitPaymentCorrection - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitPaymentCorrection(b []byte) (LawsuitPaymentCorrection, error) {
	c := NewLawsuitPaymentCorrection()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitPaymentCorrection(), err
	}
	return c, nil
}

// LawsuitPaymentCorrectionAsBytes - упаковывает объект в массив байтов
func LawsuitPaymentCorrectionAsBytes(m *LawsuitPaymentCorrection) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitPaymentCorrection) GetStructVersion() uint32 {
	if versionLawsuitPaymentCorrection == 0 {
		versionLawsuitPaymentCorrection = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitPaymentCorrection
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitPaymentCorrection) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitPaymentCorrection) GetJSON() (string, error) {
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
func (m *LawsuitPaymentCorrection) Read() error {
	err := crud_LawsuitPaymentCorrection.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitPaymentCorrection) Save() error {
	err := crud_LawsuitPaymentCorrection.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitPaymentCorrection) Update() error {
	err := crud_LawsuitPaymentCorrection.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitPaymentCorrection) Create() error {
	err := crud_LawsuitPaymentCorrection.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitPaymentCorrection) Delete() error {
	err := crud_LawsuitPaymentCorrection.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitPaymentCorrection) Restore() error {
	err := crud_LawsuitPaymentCorrection.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitPaymentCorrection) SetCrudInterface(crud ICrud_LawsuitPaymentCorrection) {
	crud_LawsuitPaymentCorrection = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
