//File generated automatic with crud_generator app
//Do not change anything here.

package messages

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionMessage - версия структуры модели, с учётом имен и типов полей
var versionMessage uint32

// crud_Message - объект контроллер crud операций
var crud_Message ICrud_Message

type ICrud_Message interface {
	Read(*Message) error
	Save(*Message) error
	Update(*Message) error
	Create(*Message) error
	Delete(*Message) error
	Restore(*Message) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Message) TableNameDB() string {
	return "messages"
}

// NewMessage - возвращает новый	объект
func NewMessage() Message {
	return Message{}
}

// AsMessage - создаёт объект из упакованного объекта в массиве байтов
func AsMessage(b []byte) (Message, error) {
	c := NewMessage()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMessage(), err
	}
	return c, nil
}

// MessageAsBytes - упаковывает объект в массив байтов
func MessageAsBytes(m *Message) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Message) GetStructVersion() uint32 {
	if versionMessage == 0 {
		versionMessage = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMessage
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Message) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Message) GetJSON() (string, error) {
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
func (m *Message) Read() error {
	err := crud_Message.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Message) Save() error {
	err := crud_Message.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Message) Update() error {
	err := crud_Message.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Message) Create() error {
	err := crud_Message.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Message) Delete() error {
	err := crud_Message.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Message) Restore() error {
	err := crud_Message.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Message) SetCrudInterface(crud ICrud_Message) {
	crud_Message = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
