//File generated automatic with crud_generator app
//Do not change anything here.

package accounting_areas

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionAccountingArea - версия структуры модели, с учётом имен и типов полей
var versionAccountingArea uint32

// crud_AccountingArea - объект контроллер crud операций
var crud_AccountingArea ICrud_AccountingArea

type ICrud_AccountingArea interface {
	Read(*AccountingArea) error
	Save(*AccountingArea) error
	Update(*AccountingArea) error
	Create(*AccountingArea) error
	Delete(*AccountingArea) error
	Restore(*AccountingArea) error
	Find_ByExtID(*AccountingArea) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m AccountingArea) TableNameDB() string {
	return "accounting_areas"
}

// NewAccountingArea - возвращает новый	объект
func NewAccountingArea() AccountingArea {
	return AccountingArea{}
}

// AsAccountingArea - создаёт объект из упакованного объекта в массиве байтов
func AsAccountingArea(b []byte) (AccountingArea, error) {
	c := NewAccountingArea()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewAccountingArea(), err
	}
	return c, nil
}

// AccountingAreaAsBytes - упаковывает объект в массив байтов
func AccountingAreaAsBytes(m *AccountingArea) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m AccountingArea) GetStructVersion() uint32 {
	if versionAccountingArea == 0 {
		versionAccountingArea = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionAccountingArea
}

// GetModelFromJSON - создаёт модель из строки json
func (m *AccountingArea) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m AccountingArea) GetJSON() (string, error) {
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
func (m *AccountingArea) Read() error {
	err := crud_AccountingArea.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *AccountingArea) Save() error {
	err := crud_AccountingArea.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *AccountingArea) Update() error {
	err := crud_AccountingArea.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *AccountingArea) Create() error {
	err := crud_AccountingArea.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *AccountingArea) Delete() error {
	err := crud_AccountingArea.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *AccountingArea) Restore() error {
	err := crud_AccountingArea.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *AccountingArea) Find_ByExtID() error {
	err := crud_AccountingArea.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m AccountingArea) SetCrudInterface(crud ICrud_AccountingArea) {
	crud_AccountingArea = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
