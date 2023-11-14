//File generated automatic with crud_generator app
//Do not change anything here.

package table_names

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionTableName - версия структуры модели, с учётом имен и типов полей
var versionTableName uint32

// crud_TableName - объект контроллер crud операций
var crud_TableName ICrud_TableName

type ICrud_TableName interface {
	Read(*TableName) error
	Save(*TableName) error
	Update(*TableName) error
	Create(*TableName) error
	Delete(*TableName) error
	Restore(*TableName) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m TableName) TableNameDB() string {
	return "table_names"
}

// NewTableName - возвращает новый	объект
func NewTableName() TableName {
	return TableName{}
}

// AsTableName - создаёт объект из упакованного объекта в массиве байтов
func AsTableName(b []byte) (TableName, error) {
	c := NewTableName()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewTableName(), err
	}
	return c, nil
}

// TableNameAsBytes - упаковывает объект в массив байтов
func TableNameAsBytes(m *TableName) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m TableName) GetStructVersion() uint32 {
	if versionTableName == 0 {
		versionTableName = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionTableName
}

// GetModelFromJSON - создаёт модель из строки json
func (m *TableName) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m TableName) GetJSON() (string, error) {
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
func (m *TableName) Read() error {
	err := crud_TableName.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *TableName) Save() error {
	err := crud_TableName.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *TableName) Update() error {
	err := crud_TableName.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *TableName) Create() error {
	err := crud_TableName.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *TableName) Delete() error {
	err := crud_TableName.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *TableName) Restore() error {
	err := crud_TableName.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TableName) SetCrudInterface(crud ICrud_TableName) {
	crud_TableName = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
