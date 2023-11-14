//File generated automatic with crud_generator app
//Do not change anything here.

package white_list_reason_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionWhiteListReasonType - версия структуры модели, с учётом имен и типов полей
var versionWhiteListReasonType uint32

// crud_WhiteListReasonType - объект контроллер crud операций
var crud_WhiteListReasonType ICrud_WhiteListReasonType

type ICrud_WhiteListReasonType interface {
	Read(*WhiteListReasonType) error
	Save(*WhiteListReasonType) error
	Update(*WhiteListReasonType) error
	Create(*WhiteListReasonType) error
	Delete(*WhiteListReasonType) error
	Restore(*WhiteListReasonType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m WhiteListReasonType) TableNameDB() string {
	return "white_list_reason_types"
}

// NewWhiteListReasonType - возвращает новый	объект
func NewWhiteListReasonType() WhiteListReasonType {
	return WhiteListReasonType{}
}

// AsWhiteListReasonType - создаёт объект из упакованного объекта в массиве байтов
func AsWhiteListReasonType(b []byte) (WhiteListReasonType, error) {
	c := NewWhiteListReasonType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewWhiteListReasonType(), err
	}
	return c, nil
}

// WhiteListReasonTypeAsBytes - упаковывает объект в массив байтов
func WhiteListReasonTypeAsBytes(m *WhiteListReasonType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m WhiteListReasonType) GetStructVersion() uint32 {
	if versionWhiteListReasonType == 0 {
		versionWhiteListReasonType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionWhiteListReasonType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *WhiteListReasonType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m WhiteListReasonType) GetJSON() (string, error) {
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
func (m *WhiteListReasonType) Read() error {
	err := crud_WhiteListReasonType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *WhiteListReasonType) Save() error {
	err := crud_WhiteListReasonType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *WhiteListReasonType) Update() error {
	err := crud_WhiteListReasonType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *WhiteListReasonType) Create() error {
	err := crud_WhiteListReasonType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *WhiteListReasonType) Delete() error {
	err := crud_WhiteListReasonType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *WhiteListReasonType) Restore() error {
	err := crud_WhiteListReasonType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WhiteListReasonType) SetCrudInterface(crud ICrud_WhiteListReasonType) {
	crud_WhiteListReasonType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
