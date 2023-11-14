//File generated automatic with crud_generator app
//Do not change anything here.

package legal_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLegalType - версия структуры модели, с учётом имен и типов полей
var versionLegalType uint32

// crud_LegalType - объект контроллер crud операций
var crud_LegalType ICrud_LegalType

type ICrud_LegalType interface {
	Read(*LegalType) error
	Save(*LegalType) error
	Update(*LegalType) error
	Create(*LegalType) error
	Delete(*LegalType) error
	Restore(*LegalType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LegalType) TableNameDB() string {
	return "legal_types"
}

// NewLegalType - возвращает новый	объект
func NewLegalType() LegalType {
	return LegalType{}
}

// AsLegalType - создаёт объект из упакованного объекта в массиве байтов
func AsLegalType(b []byte) (LegalType, error) {
	c := NewLegalType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLegalType(), err
	}
	return c, nil
}

// LegalTypeAsBytes - упаковывает объект в массив байтов
func LegalTypeAsBytes(m *LegalType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LegalType) GetStructVersion() uint32 {
	if versionLegalType == 0 {
		versionLegalType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLegalType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LegalType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LegalType) GetJSON() (string, error) {
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
func (m *LegalType) Read() error {
	err := crud_LegalType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LegalType) Save() error {
	err := crud_LegalType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LegalType) Update() error {
	err := crud_LegalType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LegalType) Create() error {
	err := crud_LegalType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LegalType) Delete() error {
	err := crud_LegalType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LegalType) Restore() error {
	err := crud_LegalType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LegalType) SetCrudInterface(crud ICrud_LegalType) {
	crud_LegalType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
