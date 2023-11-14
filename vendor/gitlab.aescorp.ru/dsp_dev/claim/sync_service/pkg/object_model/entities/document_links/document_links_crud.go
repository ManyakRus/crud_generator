//File generated automatic with crud_generator app
//Do not change anything here.

package document_links

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDocumentLink - версия структуры модели, с учётом имен и типов полей
var versionDocumentLink uint32

// crud_DocumentLink - объект контроллер crud операций
var crud_DocumentLink ICrud_DocumentLink

type ICrud_DocumentLink interface {
	Read(*DocumentLink) error
	Save(*DocumentLink) error
	Update(*DocumentLink) error
	Create(*DocumentLink) error
	Delete(*DocumentLink) error
	Restore(*DocumentLink) error
	Find_ByExtID(*DocumentLink) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m DocumentLink) TableNameDB() string {
	return "document_links"
}

// NewDocumentLink - возвращает новый	объект
func NewDocumentLink() DocumentLink {
	return DocumentLink{}
}

// AsDocumentLink - создаёт объект из упакованного объекта в массиве байтов
func AsDocumentLink(b []byte) (DocumentLink, error) {
	c := NewDocumentLink()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDocumentLink(), err
	}
	return c, nil
}

// DocumentLinkAsBytes - упаковывает объект в массив байтов
func DocumentLinkAsBytes(m *DocumentLink) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DocumentLink) GetStructVersion() uint32 {
	if versionDocumentLink == 0 {
		versionDocumentLink = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDocumentLink
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DocumentLink) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DocumentLink) GetJSON() (string, error) {
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
func (m *DocumentLink) Read() error {
	err := crud_DocumentLink.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DocumentLink) Save() error {
	err := crud_DocumentLink.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DocumentLink) Update() error {
	err := crud_DocumentLink.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DocumentLink) Create() error {
	err := crud_DocumentLink.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *DocumentLink) Delete() error {
	err := crud_DocumentLink.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *DocumentLink) Restore() error {
	err := crud_DocumentLink.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *DocumentLink) Find_ByExtID() error {
	err := crud_DocumentLink.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DocumentLink) SetCrudInterface(crud ICrud_DocumentLink) {
	crud_DocumentLink = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
