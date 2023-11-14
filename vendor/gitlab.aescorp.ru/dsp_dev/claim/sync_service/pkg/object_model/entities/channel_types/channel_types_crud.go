//File generated automatic with crud_generator app
//Do not change anything here.

package channel_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionChannelType - версия структуры модели, с учётом имен и типов полей
var versionChannelType uint32

// crud_ChannelType - объект контроллер crud операций
var crud_ChannelType ICrud_ChannelType

type ICrud_ChannelType interface {
	Read(*ChannelType) error
	Save(*ChannelType) error
	Update(*ChannelType) error
	Create(*ChannelType) error
	Delete(*ChannelType) error
	Restore(*ChannelType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ChannelType) TableNameDB() string {
	return "channel_types"
}

// NewChannelType - возвращает новый	объект
func NewChannelType() ChannelType {
	return ChannelType{}
}

// AsChannelType - создаёт объект из упакованного объекта в массиве байтов
func AsChannelType(b []byte) (ChannelType, error) {
	c := NewChannelType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewChannelType(), err
	}
	return c, nil
}

// ChannelTypeAsBytes - упаковывает объект в массив байтов
func ChannelTypeAsBytes(m *ChannelType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ChannelType) GetStructVersion() uint32 {
	if versionChannelType == 0 {
		versionChannelType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionChannelType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ChannelType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ChannelType) GetJSON() (string, error) {
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
func (m *ChannelType) Read() error {
	err := crud_ChannelType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ChannelType) Save() error {
	err := crud_ChannelType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ChannelType) Update() error {
	err := crud_ChannelType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ChannelType) Create() error {
	err := crud_ChannelType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ChannelType) Delete() error {
	err := crud_ChannelType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ChannelType) Restore() error {
	err := crud_ChannelType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ChannelType) SetCrudInterface(crud ICrud_ChannelType) {
	crud_ChannelType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
