//File generated automatic with crud_generator app
//Do not change anything here.

package gender_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionGenderType - версия структуры модели, с учётом имен и типов полей
var versionGenderType uint32

// crud_GenderType - объект контроллер crud операций
var crud_GenderType ICrud_GenderType

type ICrud_GenderType interface {
	Read(*GenderType) error
	Save(*GenderType) error
	Update(*GenderType) error
	Create(*GenderType) error
	Delete(*GenderType) error
	Restore(*GenderType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m GenderType) TableNameDB() string {
	return "gender_types"
}

// NewGenderType - возвращает новый	объект
func NewGenderType() GenderType {
	return GenderType{}
}

// AsGenderType - создаёт объект из упакованного объекта в массиве байтов
func AsGenderType(b []byte) (GenderType, error) {
	c := NewGenderType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewGenderType(), err
	}
	return c, nil
}

// GenderTypeAsBytes - упаковывает объект в массив байтов
func GenderTypeAsBytes(m *GenderType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m GenderType) GetStructVersion() uint32 {
	if versionGenderType == 0 {
		versionGenderType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionGenderType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *GenderType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m GenderType) GetJSON() (string, error) {
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
func (m *GenderType) Read() error {
	err := crud_GenderType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *GenderType) Save() error {
	err := crud_GenderType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *GenderType) Update() error {
	err := crud_GenderType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *GenderType) Create() error {
	err := crud_GenderType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *GenderType) Delete() error {
	err := crud_GenderType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *GenderType) Restore() error {
	err := crud_GenderType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m GenderType) SetCrudInterface(crud ICrud_GenderType) {
	crud_GenderType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
