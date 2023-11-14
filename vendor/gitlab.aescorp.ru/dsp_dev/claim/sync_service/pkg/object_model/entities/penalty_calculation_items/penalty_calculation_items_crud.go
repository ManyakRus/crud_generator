//File generated automatic with crud_generator app
//Do not change anything here.

package penalty_calculation_items

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionPenaltyCalculationItem - версия структуры модели, с учётом имен и типов полей
var versionPenaltyCalculationItem uint32

// crud_PenaltyCalculationItem - объект контроллер crud операций
var crud_PenaltyCalculationItem ICrud_PenaltyCalculationItem

type ICrud_PenaltyCalculationItem interface {
	Read(*PenaltyCalculationItem) error
	Save(*PenaltyCalculationItem) error
	Update(*PenaltyCalculationItem) error
	Create(*PenaltyCalculationItem) error
	Delete(*PenaltyCalculationItem) error
	Restore(*PenaltyCalculationItem) error
	Find_ByExtID(*PenaltyCalculationItem) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m PenaltyCalculationItem) TableNameDB() string {
	return "penalty_calculation_items"
}

// NewPenaltyCalculationItem - возвращает новый	объект
func NewPenaltyCalculationItem() PenaltyCalculationItem {
	return PenaltyCalculationItem{}
}

// AsPenaltyCalculationItem - создаёт объект из упакованного объекта в массиве байтов
func AsPenaltyCalculationItem(b []byte) (PenaltyCalculationItem, error) {
	c := NewPenaltyCalculationItem()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPenaltyCalculationItem(), err
	}
	return c, nil
}

// PenaltyCalculationItemAsBytes - упаковывает объект в массив байтов
func PenaltyCalculationItemAsBytes(m *PenaltyCalculationItem) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PenaltyCalculationItem) GetStructVersion() uint32 {
	if versionPenaltyCalculationItem == 0 {
		versionPenaltyCalculationItem = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPenaltyCalculationItem
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PenaltyCalculationItem) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PenaltyCalculationItem) GetJSON() (string, error) {
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
func (m *PenaltyCalculationItem) Read() error {
	err := crud_PenaltyCalculationItem.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PenaltyCalculationItem) Save() error {
	err := crud_PenaltyCalculationItem.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PenaltyCalculationItem) Update() error {
	err := crud_PenaltyCalculationItem.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PenaltyCalculationItem) Create() error {
	err := crud_PenaltyCalculationItem.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PenaltyCalculationItem) Delete() error {
	err := crud_PenaltyCalculationItem.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PenaltyCalculationItem) Restore() error {
	err := crud_PenaltyCalculationItem.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *PenaltyCalculationItem) Find_ByExtID() error {
	err := crud_PenaltyCalculationItem.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PenaltyCalculationItem) SetCrudInterface(crud ICrud_PenaltyCalculationItem) {
	crud_PenaltyCalculationItem = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
