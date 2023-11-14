//File generated automatic with crud_generator app
//Do not change anything here.

package branches

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionBranch - версия структуры модели, с учётом имен и типов полей
var versionBranch uint32

// crud_Branch - объект контроллер crud операций
var crud_Branch ICrud_Branch

type ICrud_Branch interface {
	Read(*Branch) error
	Save(*Branch) error
	Update(*Branch) error
	Create(*Branch) error
	Delete(*Branch) error
	Restore(*Branch) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Branch) TableNameDB() string {
	return "branches"
}

// NewBranch - возвращает новый	объект
func NewBranch() Branch {
	return Branch{}
}

// AsBranch - создаёт объект из упакованного объекта в массиве байтов
func AsBranch(b []byte) (Branch, error) {
	c := NewBranch()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewBranch(), err
	}
	return c, nil
}

// BranchAsBytes - упаковывает объект в массив байтов
func BranchAsBytes(m *Branch) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Branch) GetStructVersion() uint32 {
	if versionBranch == 0 {
		versionBranch = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionBranch
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Branch) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Branch) GetJSON() (string, error) {
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
func (m *Branch) Read() error {
	err := crud_Branch.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Branch) Save() error {
	err := crud_Branch.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Branch) Update() error {
	err := crud_Branch.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Branch) Create() error {
	err := crud_Branch.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Branch) Delete() error {
	err := crud_Branch.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Branch) Restore() error {
	err := crud_Branch.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Branch) SetCrudInterface(crud ICrud_Branch) {
	crud_Branch = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
