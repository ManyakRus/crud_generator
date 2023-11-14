//File generated automatic with crud_generator app
//Do not change anything here.

package accrual_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionAccrualType - версия структуры модели, с учётом имен и типов полей
var versionAccrualType uint32

// crud_AccrualType - объект контроллер crud операций
var crud_AccrualType ICrud_AccrualType

type ICrud_AccrualType interface {
	Read(*AccrualType) error
	Save(*AccrualType) error
	Update(*AccrualType) error
	Create(*AccrualType) error
	Delete(*AccrualType) error
	Restore(*AccrualType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m AccrualType) TableNameDB() string {
	return "accrual_types"
}

// NewAccrualType - возвращает новый	объект
func NewAccrualType() AccrualType {
	return AccrualType{}
}

// AsAccrualType - создаёт объект из упакованного объекта в массиве байтов
func AsAccrualType(b []byte) (AccrualType, error) {
	c := NewAccrualType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewAccrualType(), err
	}
	return c, nil
}

// AccrualTypeAsBytes - упаковывает объект в массив байтов
func AccrualTypeAsBytes(m *AccrualType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m AccrualType) GetStructVersion() uint32 {
	if versionAccrualType == 0 {
		versionAccrualType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionAccrualType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *AccrualType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m AccrualType) GetJSON() (string, error) {
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
func (m *AccrualType) Read() error {
	err := crud_AccrualType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *AccrualType) Save() error {
	err := crud_AccrualType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *AccrualType) Update() error {
	err := crud_AccrualType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *AccrualType) Create() error {
	err := crud_AccrualType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *AccrualType) Delete() error {
	err := crud_AccrualType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *AccrualType) Restore() error {
	err := crud_AccrualType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m AccrualType) SetCrudInterface(crud ICrud_AccrualType) {
	crud_AccrualType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
