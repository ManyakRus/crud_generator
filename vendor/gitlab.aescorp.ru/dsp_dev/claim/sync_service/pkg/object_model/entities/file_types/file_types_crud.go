//File generated automatic with crud_generator app
//Do not change anything here.

package file_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionFileType - версия структуры модели, с учётом имен и типов полей
var versionFileType uint32

// crud_FileType - объект контроллер crud операций
var crud_FileType ICrud_FileType

type ICrud_FileType interface {
	Read(*FileType) error
	Save(*FileType) error
	Update(*FileType) error
	Create(*FileType) error
	Delete(*FileType) error
	Restore(*FileType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m FileType) TableNameDB() string {
	return "file_types"
}

// NewFileType - возвращает новый	объект
func NewFileType() FileType {
	return FileType{}
}

// AsFileType - создаёт объект из упакованного объекта в массиве байтов
func AsFileType(b []byte) (FileType, error) {
	c := NewFileType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewFileType(), err
	}
	return c, nil
}

// FileTypeAsBytes - упаковывает объект в массив байтов
func FileTypeAsBytes(m *FileType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m FileType) GetStructVersion() uint32 {
	if versionFileType == 0 {
		versionFileType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionFileType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *FileType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m FileType) GetJSON() (string, error) {
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
func (m *FileType) Read() error {
	err := crud_FileType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *FileType) Save() error {
	err := crud_FileType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *FileType) Update() error {
	err := crud_FileType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *FileType) Create() error {
	err := crud_FileType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *FileType) Delete() error {
	err := crud_FileType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *FileType) Restore() error {
	err := crud_FileType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m FileType) SetCrudInterface(crud ICrud_FileType) {
	crud_FileType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
