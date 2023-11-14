//File generated automatic with crud_generator app
//Do not change anything here.

package document_link_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionDocumentLinkType - версия структуры модели, с учётом имен и типов полей
var versionDocumentLinkType uint32

// crud_DocumentLinkType - объект контроллер crud операций
var crud_DocumentLinkType ICrud_DocumentLinkType

type ICrud_DocumentLinkType interface {
	Read(*DocumentLinkType) error
	Save(*DocumentLinkType) error
	Update(*DocumentLinkType) error
	Create(*DocumentLinkType) error
	Delete(*DocumentLinkType) error
	Restore(*DocumentLinkType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m DocumentLinkType) TableNameDB() string {
	return "document_link_types"
}

// NewDocumentLinkType - возвращает новый	объект
func NewDocumentLinkType() DocumentLinkType {
	return DocumentLinkType{}
}

// AsDocumentLinkType - создаёт объект из упакованного объекта в массиве байтов
func AsDocumentLinkType(b []byte) (DocumentLinkType, error) {
	c := NewDocumentLinkType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDocumentLinkType(), err
	}
	return c, nil
}

// DocumentLinkTypeAsBytes - упаковывает объект в массив байтов
func DocumentLinkTypeAsBytes(m *DocumentLinkType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DocumentLinkType) GetStructVersion() uint32 {
	if versionDocumentLinkType == 0 {
		versionDocumentLinkType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDocumentLinkType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DocumentLinkType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DocumentLinkType) GetJSON() (string, error) {
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
func (m *DocumentLinkType) Read() error {
	err := crud_DocumentLinkType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DocumentLinkType) Save() error {
	err := crud_DocumentLinkType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DocumentLinkType) Update() error {
	err := crud_DocumentLinkType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DocumentLinkType) Create() error {
	err := crud_DocumentLinkType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *DocumentLinkType) Delete() error {
	err := crud_DocumentLinkType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *DocumentLinkType) Restore() error {
	err := crud_DocumentLinkType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DocumentLinkType) SetCrudInterface(crud ICrud_DocumentLinkType) {
	crud_DocumentLinkType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
