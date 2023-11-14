//File generated automatic with crud_generator app
//Do not change anything here.

package facsimiles

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionFacsimile - версия структуры модели, с учётом имен и типов полей
var versionFacsimile uint32

// crud_Facsimile - объект контроллер crud операций
var crud_Facsimile ICrud_Facsimile

type ICrud_Facsimile interface {
	Read(*Facsimile) error
	Save(*Facsimile) error
	Update(*Facsimile) error
	Create(*Facsimile) error
	Delete(*Facsimile) error
	Restore(*Facsimile) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Facsimile) TableNameDB() string {
	return "facsimiles"
}

// NewFacsimile - возвращает новый	объект
func NewFacsimile() Facsimile {
	return Facsimile{}
}

// AsFacsimile - создаёт объект из упакованного объекта в массиве байтов
func AsFacsimile(b []byte) (Facsimile, error) {
	c := NewFacsimile()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewFacsimile(), err
	}
	return c, nil
}

// FacsimileAsBytes - упаковывает объект в массив байтов
func FacsimileAsBytes(m *Facsimile) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Facsimile) GetStructVersion() uint32 {
	if versionFacsimile == 0 {
		versionFacsimile = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionFacsimile
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Facsimile) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Facsimile) GetJSON() (string, error) {
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
func (m *Facsimile) Read() error {
	err := crud_Facsimile.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Facsimile) Save() error {
	err := crud_Facsimile.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Facsimile) Update() error {
	err := crud_Facsimile.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Facsimile) Create() error {
	err := crud_Facsimile.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Facsimile) Delete() error {
	err := crud_Facsimile.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Facsimile) Restore() error {
	err := crud_Facsimile.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Facsimile) SetCrudInterface(crud ICrud_Facsimile) {
	crud_Facsimile = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
