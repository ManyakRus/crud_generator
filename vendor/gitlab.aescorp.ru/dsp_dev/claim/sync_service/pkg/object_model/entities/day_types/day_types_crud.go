//File generated automatic with crud_generator app
//Do not change anything here.

package day_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDayType - версия структуры модели, с учётом имен и типов полей
var versionDayType uint32

// crud_DayType - объект контроллер crud операций
var crud_DayType ICrud_DayType

type ICrud_DayType interface {
	Read(*DayType) error
	Save(*DayType) error
	Update(*DayType) error
	Create(*DayType) error
	Delete(*DayType) error
	Restore(*DayType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m DayType) TableNameDB() string {
	return "day_types"
}

// NewDayType - возвращает новый	объект
func NewDayType() DayType {
	return DayType{}
}

// AsDayType - создаёт объект из упакованного объекта в массиве байтов
func AsDayType(b []byte) (DayType, error) {
	c := NewDayType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDayType(), err
	}
	return c, nil
}

// DayTypeAsBytes - упаковывает объект в массив байтов
func DayTypeAsBytes(m *DayType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DayType) GetStructVersion() uint32 {
	if versionDayType == 0 {
		versionDayType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDayType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DayType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DayType) GetJSON() (string, error) {
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
func (m *DayType) Read() error {
	err := crud_DayType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DayType) Save() error {
	err := crud_DayType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DayType) Update() error {
	err := crud_DayType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DayType) Create() error {
	err := crud_DayType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *DayType) Delete() error {
	err := crud_DayType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *DayType) Restore() error {
	err := crud_DayType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DayType) SetCrudInterface(crud ICrud_DayType) {
	crud_DayType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
