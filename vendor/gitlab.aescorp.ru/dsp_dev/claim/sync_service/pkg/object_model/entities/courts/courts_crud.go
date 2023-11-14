//File generated automatic with crud_generator app
//Do not change anything here.

package courts

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionCourt - версия структуры модели, с учётом имен и типов полей
var versionCourt uint32

// crud_Court - объект контроллер crud операций
var crud_Court ICrud_Court

type ICrud_Court interface {
	Read(*Court) error
	Save(*Court) error
	Update(*Court) error
	Create(*Court) error
	Delete(*Court) error
	Restore(*Court) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Court) TableNameDB() string {
	return "courts"
}

// NewCourt - возвращает новый	объект
func NewCourt() Court {
	return Court{}
}

// AsCourt - создаёт объект из упакованного объекта в массиве байтов
func AsCourt(b []byte) (Court, error) {
	c := NewCourt()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewCourt(), err
	}
	return c, nil
}

// CourtAsBytes - упаковывает объект в массив байтов
func CourtAsBytes(m *Court) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Court) GetStructVersion() uint32 {
	if versionCourt == 0 {
		versionCourt = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionCourt
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Court) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Court) GetJSON() (string, error) {
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
func (m *Court) Read() error {
	err := crud_Court.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Court) Save() error {
	err := crud_Court.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Court) Update() error {
	err := crud_Court.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Court) Create() error {
	err := crud_Court.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Court) Delete() error {
	err := crud_Court.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Court) Restore() error {
	err := crud_Court.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Court) SetCrudInterface(crud ICrud_Court) {
	crud_Court = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
