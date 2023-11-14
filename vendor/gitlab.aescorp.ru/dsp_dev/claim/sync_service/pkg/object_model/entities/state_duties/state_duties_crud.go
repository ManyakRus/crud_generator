//File generated automatic with crud_generator app
//Do not change anything here.

package state_duties

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionStateDuty - версия структуры модели, с учётом имен и типов полей
var versionStateDuty uint32

// crud_StateDuty - объект контроллер crud операций
var crud_StateDuty ICrud_StateDuty

type ICrud_StateDuty interface {
	Read(*StateDuty) error
	Save(*StateDuty) error
	Update(*StateDuty) error
	Create(*StateDuty) error
	Delete(*StateDuty) error
	Restore(*StateDuty) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m StateDuty) TableNameDB() string {
	return "state_duties"
}

// NewStateDuty - возвращает новый	объект
func NewStateDuty() StateDuty {
	return StateDuty{}
}

// AsStateDuty - создаёт объект из упакованного объекта в массиве байтов
func AsStateDuty(b []byte) (StateDuty, error) {
	c := NewStateDuty()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewStateDuty(), err
	}
	return c, nil
}

// StateDutyAsBytes - упаковывает объект в массив байтов
func StateDutyAsBytes(m *StateDuty) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m StateDuty) GetStructVersion() uint32 {
	if versionStateDuty == 0 {
		versionStateDuty = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionStateDuty
}

// GetModelFromJSON - создаёт модель из строки json
func (m *StateDuty) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m StateDuty) GetJSON() (string, error) {
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
func (m *StateDuty) Read() error {
	err := crud_StateDuty.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *StateDuty) Save() error {
	err := crud_StateDuty.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *StateDuty) Update() error {
	err := crud_StateDuty.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *StateDuty) Create() error {
	err := crud_StateDuty.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *StateDuty) Delete() error {
	err := crud_StateDuty.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *StateDuty) Restore() error {
	err := crud_StateDuty.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m StateDuty) SetCrudInterface(crud ICrud_StateDuty) {
	crud_StateDuty = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
