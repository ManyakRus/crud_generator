//File generated automatic with crud_generator app
//Do not change anything here.

package direction_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDirectionType - версия структуры модели, с учётом имен и типов полей
var versionDirectionType uint32

// crud_DirectionType - объект контроллер crud операций
var crud_DirectionType ICrud_DirectionType

type ICrud_DirectionType interface {
	Read(*DirectionType) error
	Save(*DirectionType) error
	Update(*DirectionType) error
	Create(*DirectionType) error
	Delete(*DirectionType) error
	Restore(*DirectionType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m DirectionType) TableNameDB() string {
	return "direction_types"
}

// NewDirectionType - возвращает новый	объект
func NewDirectionType() DirectionType {
	return DirectionType{}
}

// AsDirectionType - создаёт объект из упакованного объекта в массиве байтов
func AsDirectionType(b []byte) (DirectionType, error) {
	c := NewDirectionType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDirectionType(), err
	}
	return c, nil
}

// DirectionTypeAsBytes - упаковывает объект в массив байтов
func DirectionTypeAsBytes(m *DirectionType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DirectionType) GetStructVersion() uint32 {
	if versionDirectionType == 0 {
		versionDirectionType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDirectionType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DirectionType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DirectionType) GetJSON() (string, error) {
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
func (m *DirectionType) Read() error {
	err := crud_DirectionType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DirectionType) Save() error {
	err := crud_DirectionType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DirectionType) Update() error {
	err := crud_DirectionType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DirectionType) Create() error {
	err := crud_DirectionType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *DirectionType) Delete() error {
	err := crud_DirectionType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *DirectionType) Restore() error {
	err := crud_DirectionType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DirectionType) SetCrudInterface(crud ICrud_DirectionType) {
	crud_DirectionType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
