//File generated automatic with crud_generator app
//Do not change anything here.

package file_changes

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionFileChange - версия структуры модели, с учётом имен и типов полей
var versionFileChange uint32

// crud_FileChange - объект контроллер crud операций
var crud_FileChange ICrud_FileChange

type ICrud_FileChange interface {
	Read(*FileChange) error
	Save(*FileChange) error
	Update(*FileChange) error
	Create(*FileChange) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m FileChange) TableNameDB() string {
	return "file_changes"
}

// NewFileChange - возвращает новый	объект
func NewFileChange() FileChange {
	return FileChange{}
}

// AsFileChange - создаёт объект из упакованного объекта в массиве байтов
func AsFileChange(b []byte) (FileChange, error) {
	c := NewFileChange()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewFileChange(), err
	}
	return c, nil
}

// FileChangeAsBytes - упаковывает объект в массив байтов
func FileChangeAsBytes(m *FileChange) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m FileChange) GetStructVersion() uint32 {
	if versionFileChange == 0 {
		versionFileChange = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionFileChange
}

// GetModelFromJSON - создаёт модель из строки json
func (m *FileChange) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m FileChange) GetJSON() (string, error) {
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
func (m *FileChange) Read() error {
	err := crud_FileChange.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *FileChange) Save() error {
	err := crud_FileChange.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *FileChange) Update() error {
	err := crud_FileChange.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *FileChange) Create() error {
	err := crud_FileChange.Create(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m FileChange) SetCrudInterface(crud ICrud_FileChange) {
	crud_FileChange = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
