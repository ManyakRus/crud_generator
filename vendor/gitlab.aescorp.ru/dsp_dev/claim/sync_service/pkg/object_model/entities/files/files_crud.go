//File generated automatic with crud_generator app
//Do not change anything here.

package files

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionFile - версия структуры модели, с учётом имен и типов полей
var versionFile uint32

// crud_File - объект контроллер crud операций
var crud_File ICrud_File

type ICrud_File interface {
	Read(f *File) error
	Save(f *File) error
	Update(f *File) error
	Create(f *File) error
	Delete(f *File) error
	Restore(f *File) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m File) TableNameDB() string {
	return "files"
}

// NewFile - возвращает новый	объект
func NewFile() File {
	return File{}
}

// AsFile - создаёт объект из упакованного объекта в массиве байтов
func AsFile(b []byte) (File, error) {
	c := NewFile()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewFile(), err
	}
	return c, nil
}

// FileAsBytes - упаковывает объект в массив байтов
func FileAsBytes(m *File) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m File) GetStructVersion() uint32 {
	if versionFile == 0 {
		versionFile = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionFile
}

// GetModelFromJSON - создаёт модель из строки json
func (m *File) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m File) GetJSON() (string, error) {
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
func (m *File) Read() error {
	err := crud_File.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *File) Save() error {
	err := crud_File.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *File) Update() error {
	err := crud_File.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *File) Create() error {
	err := crud_File.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *File) Delete() error {
	err := crud_File.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *File) Restore() error {
	err := crud_File.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m File) SetCrudInterface(crud ICrud_File) {
	crud_File = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
