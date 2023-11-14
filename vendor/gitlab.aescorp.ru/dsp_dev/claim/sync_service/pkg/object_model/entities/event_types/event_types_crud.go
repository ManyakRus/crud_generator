//File generated automatic with crud_generator app
//Do not change anything here.

package event_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionEventType - версия структуры модели, с учётом имен и типов полей
var versionEventType uint32

// crud_EventType - объект контроллер crud операций
var crud_EventType ICrud_EventType

type ICrud_EventType interface {
	Read(*EventType) error
	Save(*EventType) error
	Update(*EventType) error
	Create(*EventType) error
	Delete(*EventType) error
	Restore(*EventType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m EventType) TableNameDB() string {
	return "event_types"
}

// NewEventType - возвращает новый	объект
func NewEventType() EventType {
	return EventType{}
}

// AsEventType - создаёт объект из упакованного объекта в массиве байтов
func AsEventType(b []byte) (EventType, error) {
	c := NewEventType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewEventType(), err
	}
	return c, nil
}

// EventTypeAsBytes - упаковывает объект в массив байтов
func EventTypeAsBytes(m *EventType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m EventType) GetStructVersion() uint32 {
	if versionEventType == 0 {
		versionEventType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionEventType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *EventType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m EventType) GetJSON() (string, error) {
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
func (m *EventType) Read() error {
	err := crud_EventType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *EventType) Save() error {
	err := crud_EventType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *EventType) Update() error {
	err := crud_EventType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *EventType) Create() error {
	err := crud_EventType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *EventType) Delete() error {
	err := crud_EventType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *EventType) Restore() error {
	err := crud_EventType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m EventType) SetCrudInterface(crud ICrud_EventType) {
	crud_EventType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
