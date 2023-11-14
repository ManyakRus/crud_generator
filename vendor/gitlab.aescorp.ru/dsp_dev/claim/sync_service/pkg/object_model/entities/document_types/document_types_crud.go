//File generated automatic with crud_generator app
//Do not change anything here.

package document_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDocumentType - версия структуры модели, с учётом имен и типов полей
var versionDocumentType uint32

// crud_DocumentType - объект контроллер crud операций
var crud_DocumentType ICrud_DocumentType

type ICrud_DocumentType interface {
	Read(*DocumentType) error
	Save(*DocumentType) error
	Update(*DocumentType) error
	Create(*DocumentType) error
	Delete(*DocumentType) error
	Restore(*DocumentType) error
	Find_ByExtID(*DocumentType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m DocumentType) TableNameDB() string {
	return "document_types"
}

// NewDocumentType - возвращает новый	объект
func NewDocumentType() DocumentType {
	return DocumentType{}
}

// AsDocumentType - создаёт объект из упакованного объекта в массиве байтов
func AsDocumentType(b []byte) (DocumentType, error) {
	c := NewDocumentType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDocumentType(), err
	}
	return c, nil
}

// DocumentTypeAsBytes - упаковывает объект в массив байтов
func DocumentTypeAsBytes(m *DocumentType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DocumentType) GetStructVersion() uint32 {
	if versionDocumentType == 0 {
		versionDocumentType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDocumentType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DocumentType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DocumentType) GetJSON() (string, error) {
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
func (m *DocumentType) Read() error {
	err := crud_DocumentType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DocumentType) Save() error {
	err := crud_DocumentType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DocumentType) Update() error {
	err := crud_DocumentType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DocumentType) Create() error {
	err := crud_DocumentType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *DocumentType) Delete() error {
	err := crud_DocumentType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *DocumentType) Restore() error {
	err := crud_DocumentType.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *DocumentType) Find_ByExtID() error {
	err := crud_DocumentType.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DocumentType) SetCrudInterface(crud ICrud_DocumentType) {
	crud_DocumentType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
