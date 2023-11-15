//File generated automatic with crud_generator app
//Do not change anything here.

package message_attachements

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionMessageAttachement - версия структуры модели, с учётом имен и типов полей
var versionMessageAttachement uint32

// crud_MessageAttachement - объект контроллер crud операций
var crud_MessageAttachement ICrud_MessageAttachement

type ICrud_MessageAttachement interface {
	Read(*MessageAttachement) error
	Save(*MessageAttachement) error
	Update(*MessageAttachement) error
	Create(*MessageAttachement) error
	Delete(*MessageAttachement) error
	Restore(*MessageAttachement) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m MessageAttachement) TableNameDB() string {
	return "message_attachements"
}

// NewMessageAttachement - возвращает новый	объект
func NewMessageAttachement() MessageAttachement {
	return MessageAttachement{}
}

// AsMessageAttachement - создаёт объект из упакованного объекта в массиве байтов
func AsMessageAttachement(b []byte) (MessageAttachement, error) {
	c := NewMessageAttachement()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMessageAttachement(), err
	}
	return c, nil
}

// MessageAttachementAsBytes - упаковывает объект в массив байтов
func MessageAttachementAsBytes(m *MessageAttachement) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m MessageAttachement) GetStructVersion() uint32 {
	if versionMessageAttachement == 0 {
		versionMessageAttachement = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMessageAttachement
}

// GetModelFromJSON - создаёт модель из строки json
func (m *MessageAttachement) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m MessageAttachement) GetJSON() (string, error) {
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
func (m *MessageAttachement) Read() error {
	err := crud_MessageAttachement.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *MessageAttachement) Save() error {
	err := crud_MessageAttachement.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *MessageAttachement) Update() error {
	err := crud_MessageAttachement.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *MessageAttachement) Create() error {
	err := crud_MessageAttachement.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *MessageAttachement) Delete() error {
	err := crud_MessageAttachement.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *MessageAttachement) Restore() error {
	err := crud_MessageAttachement.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m MessageAttachement) SetCrudInterface(crud ICrud_MessageAttachement) {
	crud_MessageAttachement = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------