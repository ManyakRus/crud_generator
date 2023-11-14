//File generated automatic with crud_generator app
//Do not change anything here.

package message_send_statuses

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionMessageSendStatus - версия структуры модели, с учётом имен и типов полей
var versionMessageSendStatus uint32

// crud_MessageSendStatus - объект контроллер crud операций
var crud_MessageSendStatus ICrud_MessageSendStatus

type ICrud_MessageSendStatus interface {
	Read(*MessageSendStatus) error
	Save(*MessageSendStatus) error
	Update(*MessageSendStatus) error
	Create(*MessageSendStatus) error
	Delete(*MessageSendStatus) error
	Restore(*MessageSendStatus) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m MessageSendStatus) TableNameDB() string {
	return "message_send_statuses"
}

// NewMessageSendStatus - возвращает новый	объект
func NewMessageSendStatus() MessageSendStatus {
	return MessageSendStatus{}
}

// AsMessageSendStatus - создаёт объект из упакованного объекта в массиве байтов
func AsMessageSendStatus(b []byte) (MessageSendStatus, error) {
	c := NewMessageSendStatus()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMessageSendStatus(), err
	}
	return c, nil
}

// MessageSendStatusAsBytes - упаковывает объект в массив байтов
func MessageSendStatusAsBytes(m *MessageSendStatus) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m MessageSendStatus) GetStructVersion() uint32 {
	if versionMessageSendStatus == 0 {
		versionMessageSendStatus = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMessageSendStatus
}

// GetModelFromJSON - создаёт модель из строки json
func (m *MessageSendStatus) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m MessageSendStatus) GetJSON() (string, error) {
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
func (m *MessageSendStatus) Read() error {
	err := crud_MessageSendStatus.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *MessageSendStatus) Save() error {
	err := crud_MessageSendStatus.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *MessageSendStatus) Update() error {
	err := crud_MessageSendStatus.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *MessageSendStatus) Create() error {
	err := crud_MessageSendStatus.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *MessageSendStatus) Delete() error {
	err := crud_MessageSendStatus.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *MessageSendStatus) Restore() error {
	err := crud_MessageSendStatus.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m MessageSendStatus) SetCrudInterface(crud ICrud_MessageSendStatus) {
	crud_MessageSendStatus = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
