//File generated automatic with crud_generator app
//Do not change anything here.

package completed_months

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionCompletedMonth - версия структуры модели, с учётом имен и типов полей
var versionCompletedMonth uint32

// crud_CompletedMonth - объект контроллер crud операций
var crud_CompletedMonth ICrud_CompletedMonth

type ICrud_CompletedMonth interface {
	Read(*CompletedMonth) error
	Save(*CompletedMonth) error
	Update(*CompletedMonth) error
	Create(*CompletedMonth) error
	Delete(*CompletedMonth) error
	Restore(*CompletedMonth) error
	Find_ByExtID(*CompletedMonth) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m CompletedMonth) TableNameDB() string {
	return "completed_months"
}

// NewCompletedMonth - возвращает новый	объект
func NewCompletedMonth() CompletedMonth {
	return CompletedMonth{}
}

// AsCompletedMonth - создаёт объект из упакованного объекта в массиве байтов
func AsCompletedMonth(b []byte) (CompletedMonth, error) {
	c := NewCompletedMonth()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewCompletedMonth(), err
	}
	return c, nil
}

// CompletedMonthAsBytes - упаковывает объект в массив байтов
func CompletedMonthAsBytes(m *CompletedMonth) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m CompletedMonth) GetStructVersion() uint32 {
	if versionCompletedMonth == 0 {
		versionCompletedMonth = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionCompletedMonth
}

// GetModelFromJSON - создаёт модель из строки json
func (m *CompletedMonth) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m CompletedMonth) GetJSON() (string, error) {
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
func (m *CompletedMonth) Read() error {
	err := crud_CompletedMonth.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *CompletedMonth) Save() error {
	err := crud_CompletedMonth.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *CompletedMonth) Update() error {
	err := crud_CompletedMonth.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *CompletedMonth) Create() error {
	err := crud_CompletedMonth.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *CompletedMonth) Delete() error {
	err := crud_CompletedMonth.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *CompletedMonth) Restore() error {
	err := crud_CompletedMonth.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *CompletedMonth) Find_ByExtID() error {
	err := crud_CompletedMonth.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m CompletedMonth) SetCrudInterface(crud ICrud_CompletedMonth) {
	crud_CompletedMonth = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
