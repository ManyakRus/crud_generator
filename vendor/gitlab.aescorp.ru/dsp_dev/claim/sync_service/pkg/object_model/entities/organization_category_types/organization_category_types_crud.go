//File generated automatic with crud_generator app
//Do not change anything here.

package organization_category_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionOrganizationCategoryType - версия структуры модели, с учётом имен и типов полей
var versionOrganizationCategoryType uint32

// crud_OrganizationCategoryType - объект контроллер crud операций
var crud_OrganizationCategoryType ICrud_OrganizationCategoryType

type ICrud_OrganizationCategoryType interface {
	Read(*OrganizationCategoryType) error
	Save(*OrganizationCategoryType) error
	Update(*OrganizationCategoryType) error
	Create(*OrganizationCategoryType) error
	Delete(*OrganizationCategoryType) error
	Restore(*OrganizationCategoryType) error
	Find_ByExtID(*OrganizationCategoryType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m OrganizationCategoryType) TableNameDB() string {
	return "organization_category_types"
}

// NewOrganizationCategoryType - возвращает новый	объект
func NewOrganizationCategoryType() OrganizationCategoryType {
	return OrganizationCategoryType{}
}

// AsOrganizationCategoryType - создаёт объект из упакованного объекта в массиве байтов
func AsOrganizationCategoryType(b []byte) (OrganizationCategoryType, error) {
	c := NewOrganizationCategoryType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewOrganizationCategoryType(), err
	}
	return c, nil
}

// OrganizationCategoryTypeAsBytes - упаковывает объект в массив байтов
func OrganizationCategoryTypeAsBytes(m *OrganizationCategoryType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m OrganizationCategoryType) GetStructVersion() uint32 {
	if versionOrganizationCategoryType == 0 {
		versionOrganizationCategoryType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionOrganizationCategoryType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *OrganizationCategoryType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m OrganizationCategoryType) GetJSON() (string, error) {
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
func (m *OrganizationCategoryType) Read() error {
	err := crud_OrganizationCategoryType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *OrganizationCategoryType) Save() error {
	err := crud_OrganizationCategoryType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *OrganizationCategoryType) Update() error {
	err := crud_OrganizationCategoryType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *OrganizationCategoryType) Create() error {
	err := crud_OrganizationCategoryType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *OrganizationCategoryType) Delete() error {
	err := crud_OrganizationCategoryType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *OrganizationCategoryType) Restore() error {
	err := crud_OrganizationCategoryType.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *OrganizationCategoryType) Find_ByExtID() error {
	err := crud_OrganizationCategoryType.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m OrganizationCategoryType) SetCrudInterface(crud ICrud_OrganizationCategoryType) {
	crud_OrganizationCategoryType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
