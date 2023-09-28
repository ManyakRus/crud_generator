package object_model

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionMessageType - версия структуры модели, с учётом имен и типов полей
var versionMessageType uint32

// MessageType - Типы сообщений
type MessageType struct {
	CommonStruct
	NameStruct
	Code int `json:"code"        gorm:"column:code;default:0"`
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m MessageType) TableName() string {
	return "message_types"
}

// GetID - возвращает ID объекта
func (m MessageType) GetID() int64 {
	return m.ID
}

// NewMessageType - возвращает новый	объект
func NewMessageType() MessageType {
	return MessageType{}
}

// AsMessageType - создаёт объект из упакованного объекта в массиве байтов
func AsMessageType(b []byte) (MessageType, error) {
	c := NewMessageType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMessageType(), err
	}
	return c, nil
}

// MessageTypeAsBytes - упаковывает объект в массив байтов
func MessageTypeAsBytes(m *MessageType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m MessageType) GetStructVersion() uint32 {
	if versionMessageType == 0 {
		versionMessageType = CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMessageType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *MessageType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m MessageType) GetJSON() (string, error) {
	var ReturnVar string
	var err error

	bytes, err := json.Marshal(m)
	if err != nil {
		return ReturnVar, err
	}
	ReturnVar = string(bytes)
	return ReturnVar, err
}

//---------------------------- CRUD операции ------------------------------------------------------------

// Read - находит запись в БД по ID, и заполняет в объект
func (m *MessageType) Read() error {
	err := m.read()

	return err
}

// Save - записывает объект в БД по ID
func (m *MessageType) Save() error {
	err := m.save()

	return err
}

// Update - обновляет объект в БД по ID
func (m *MessageType) Update() error {
	err := m.update()

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *MessageType) Create() error {
	err := m.create()

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *MessageType) Delete() error {
	err := m.delete()

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *MessageType) Restore() error {
	err := m.restore()

	return err
}

//---------------------------- конец CRUD операции ------------------------------------------------------------
