//File generated automatic with crud_generator app
//Do not change anything here.

package documents

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDocument - версия структуры модели, с учётом имен и типов полей
var versionDocument uint32

// crud_Document - объект контроллер crud операций
var crud_Document ICrud_Document

type ICrud_Document interface {
	Read(*Document) error
	Save(*Document) error
	Update(*Document) error
	Create(*Document) error
	Delete(*Document) error
	Restore(*Document) error
	Find_ByExtID(*Document) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Document) TableNameDB() string {
	return "documents"
}

// NewDocument - возвращает новый	объект
func NewDocument() Document {
	return Document{}
}

// AsDocument - создаёт объект из упакованного объекта в массиве байтов
func AsDocument(b []byte) (Document, error) {
	c := NewDocument()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDocument(), err
	}
	return c, nil
}

// DocumentAsBytes - упаковывает объект в массив байтов
func DocumentAsBytes(m *Document) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Document) GetStructVersion() uint32 {
	if versionDocument == 0 {
		versionDocument = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDocument
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Document) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Document) GetJSON() (string, error) {
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
func (m *Document) Read() error {
	err := crud_Document.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Document) Save() error {
	err := crud_Document.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Document) Update() error {
	err := crud_Document.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Document) Create() error {
	err := crud_Document.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Document) Delete() error {
	err := crud_Document.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Document) Restore() error {
	err := crud_Document.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Document) Find_ByExtID() error {
	err := crud_Document.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Document) SetCrudInterface(crud ICrud_Document) {
	crud_Document = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
