//File generated automatic with crud_generator app
//Do not change anything here.

package debt_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDebtType - версия структуры модели, с учётом имен и типов полей
var versionDebtType uint32

// crud_DebtType - объект контроллер crud операций
var crud_DebtType ICrud_DebtType

type ICrud_DebtType interface {
	Read(*DebtType) error
	Save(*DebtType) error
	Update(*DebtType) error
	Create(*DebtType) error
	Delete(*DebtType) error
	Restore(*DebtType) error
	Find_ByExtID(*DebtType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m DebtType) TableNameDB() string {
	return "debt_types"
}

// NewDebtType - возвращает новый	объект
func NewDebtType() DebtType {
	return DebtType{}
}

// AsDebtType - создаёт объект из упакованного объекта в массиве байтов
func AsDebtType(b []byte) (DebtType, error) {
	c := NewDebtType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDebtType(), err
	}
	return c, nil
}

// DebtTypeAsBytes - упаковывает объект в массив байтов
func DebtTypeAsBytes(m *DebtType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DebtType) GetStructVersion() uint32 {
	if versionDebtType == 0 {
		versionDebtType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDebtType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DebtType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DebtType) GetJSON() (string, error) {
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
func (m *DebtType) Read() error {
	err := crud_DebtType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DebtType) Save() error {
	err := crud_DebtType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DebtType) Update() error {
	err := crud_DebtType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DebtType) Create() error {
	err := crud_DebtType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *DebtType) Delete() error {
	err := crud_DebtType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *DebtType) Restore() error {
	err := crud_DebtType.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *DebtType) Find_ByExtID() error {
	err := crud_DebtType.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DebtType) SetCrudInterface(crud ICrud_DebtType) {
	crud_DebtType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
