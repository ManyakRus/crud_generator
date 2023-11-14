//File generated automatic with crud_generator app
//Do not change anything here.

package events

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionEvent - версия структуры модели, с учётом имен и типов полей
var versionEvent uint32

// crud_Event - объект контроллер crud операций
var crud_Event ICrud_Event

type ICrud_Event interface {
	Read(*Event) error
	Save(*Event) error
	Update(*Event) error
	Create(*Event) error
	Delete(*Event) error
	Restore(*Event) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Event) TableNameDB() string {
	return "events"
}

// NewEvent - возвращает новый	объект
func NewEvent() Event {
	return Event{}
}

// AsEvent - создаёт объект из упакованного объекта в массиве байтов
func AsEvent(b []byte) (Event, error) {
	c := NewEvent()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewEvent(), err
	}
	return c, nil
}

// EventAsBytes - упаковывает объект в массив байтов
func EventAsBytes(m *Event) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Event) GetStructVersion() uint32 {
	if versionEvent == 0 {
		versionEvent = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionEvent
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Event) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Event) GetJSON() (string, error) {
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
func (m *Event) Read() error {
	err := crud_Event.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Event) Save() error {
	err := crud_Event.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Event) Update() error {
	err := crud_Event.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Event) Create() error {
	err := crud_Event.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Event) Delete() error {
	err := crud_Event.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Event) Restore() error {
	err := crud_Event.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Event) SetCrudInterface(crud ICrud_Event) {
	crud_Event = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
