//File generated automatic with crud_generator app
//Do not change anything here.

package individuals

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionIndividual - версия структуры модели, с учётом имен и типов полей
var versionIndividual uint32

// crud_Individual - объект контроллер crud операций
var crud_Individual ICrud_Individual

type ICrud_Individual interface {
	Read(*Individual) error
	Save(*Individual) error
	Update(*Individual) error
	Create(*Individual) error
	Delete(*Individual) error
	Restore(*Individual) error
	Find_ByExtID(*Individual) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Individual) TableNameDB() string {
	return "individuals"
}

// NewIndividual - возвращает новый	объект
func NewIndividual() Individual {
	return Individual{}
}

// AsIndividual - создаёт объект из упакованного объекта в массиве байтов
func AsIndividual(b []byte) (Individual, error) {
	c := NewIndividual()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewIndividual(), err
	}
	return c, nil
}

// IndividualAsBytes - упаковывает объект в массив байтов
func IndividualAsBytes(m *Individual) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Individual) GetStructVersion() uint32 {
	if versionIndividual == 0 {
		versionIndividual = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionIndividual
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Individual) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Individual) GetJSON() (string, error) {
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
func (m *Individual) Read() error {
	err := crud_Individual.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Individual) Save() error {
	err := crud_Individual.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Individual) Update() error {
	err := crud_Individual.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Individual) Create() error {
	err := crud_Individual.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Individual) Delete() error {
	err := crud_Individual.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Individual) Restore() error {
	err := crud_Individual.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Individual) Find_ByExtID() error {
	err := crud_Individual.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Individual) SetCrudInterface(crud ICrud_Individual) {
	crud_Individual = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
