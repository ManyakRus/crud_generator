//File generated automatic with crud_generator app
//Do not change anything here.

package calendars

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionCalendar - версия структуры модели, с учётом имен и типов полей
var versionCalendar uint32

// crud_Calendar - объект контроллер crud операций
var crud_Calendar ICrud_Calendar

type ICrud_Calendar interface {
	Read(*Calendar) error
	Save(*Calendar) error
	Update(*Calendar) error
	Create(*Calendar) error
	Delete(*Calendar) error
	Restore(*Calendar) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Calendar) TableNameDB() string {
	return "calendars"
}

// NewCalendar - возвращает новый	объект
func NewCalendar() Calendar {
	return Calendar{}
}

// AsCalendar - создаёт объект из упакованного объекта в массиве байтов
func AsCalendar(b []byte) (Calendar, error) {
	c := NewCalendar()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewCalendar(), err
	}
	return c, nil
}

// CalendarAsBytes - упаковывает объект в массив байтов
func CalendarAsBytes(m *Calendar) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Calendar) GetStructVersion() uint32 {
	if versionCalendar == 0 {
		versionCalendar = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionCalendar
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Calendar) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Calendar) GetJSON() (string, error) {
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
func (m *Calendar) Read() error {
	err := crud_Calendar.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Calendar) Save() error {
	err := crud_Calendar.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Calendar) Update() error {
	err := crud_Calendar.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Calendar) Create() error {
	err := crud_Calendar.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Calendar) Delete() error {
	err := crud_Calendar.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Calendar) Restore() error {
	err := crud_Calendar.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Calendar) SetCrudInterface(crud ICrud_Calendar) {
	crud_Calendar = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
