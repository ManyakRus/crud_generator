//File generated automatic with crud_generator app
//Do not change anything here.

package claim_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionClaimType - версия структуры модели, с учётом имен и типов полей
var versionClaimType uint32

// crud_ClaimType - объект контроллер crud операций
var crud_ClaimType ICrud_ClaimType

type ICrud_ClaimType interface {
	Read(*ClaimType) error
	Save(*ClaimType) error
	Update(*ClaimType) error
	Create(*ClaimType) error
	Delete(*ClaimType) error
	Restore(*ClaimType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m ClaimType) TableNameDB() string {
	return "claim_types"
}

// NewClaimType - возвращает новый	объект
func NewClaimType() ClaimType {
	return ClaimType{}
}

// AsClaimType - создаёт объект из упакованного объекта в массиве байтов
func AsClaimType(b []byte) (ClaimType, error) {
	c := NewClaimType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewClaimType(), err
	}
	return c, nil
}

// ClaimTypeAsBytes - упаковывает объект в массив байтов
func ClaimTypeAsBytes(m *ClaimType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ClaimType) GetStructVersion() uint32 {
	if versionClaimType == 0 {
		versionClaimType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionClaimType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ClaimType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ClaimType) GetJSON() (string, error) {
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
func (m *ClaimType) Read() error {
	err := crud_ClaimType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ClaimType) Save() error {
	err := crud_ClaimType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ClaimType) Update() error {
	err := crud_ClaimType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ClaimType) Create() error {
	err := crud_ClaimType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *ClaimType) Delete() error {
	err := crud_ClaimType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *ClaimType) Restore() error {
	err := crud_ClaimType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ClaimType) SetCrudInterface(crud ICrud_ClaimType) {
	crud_ClaimType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
