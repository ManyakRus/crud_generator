//File generated automatic with crud_generator app
//Do not change anything here.

package bill_kind_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionBillKindType - версия структуры модели, с учётом имен и типов полей
var versionBillKindType uint32

// crud_BillKindType - объект контроллер crud операций
var crud_BillKindType ICrud_BillKindType

type ICrud_BillKindType interface {
	Read(*BillKindType) error
	Save(*BillKindType) error
	Update(*BillKindType) error
	Create(*BillKindType) error
	Delete(*BillKindType) error
	Restore(*BillKindType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m BillKindType) TableNameDB() string {
	return "bill_kind_types"
}

// NewBillKindType - возвращает новый	объект
func NewBillKindType() BillKindType {
	return BillKindType{}
}

// AsBillKindType - создаёт объект из упакованного объекта в массиве байтов
func AsBillKindType(b []byte) (BillKindType, error) {
	c := NewBillKindType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewBillKindType(), err
	}
	return c, nil
}

// BillKindTypeAsBytes - упаковывает объект в массив байтов
func BillKindTypeAsBytes(m *BillKindType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m BillKindType) GetStructVersion() uint32 {
	if versionBillKindType == 0 {
		versionBillKindType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionBillKindType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *BillKindType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m BillKindType) GetJSON() (string, error) {
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
func (m *BillKindType) Read() error {
	err := crud_BillKindType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *BillKindType) Save() error {
	err := crud_BillKindType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *BillKindType) Update() error {
	err := crud_BillKindType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *BillKindType) Create() error {
	err := crud_BillKindType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *BillKindType) Delete() error {
	err := crud_BillKindType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *BillKindType) Restore() error {
	err := crud_BillKindType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m BillKindType) SetCrudInterface(crud ICrud_BillKindType) {
	crud_BillKindType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
