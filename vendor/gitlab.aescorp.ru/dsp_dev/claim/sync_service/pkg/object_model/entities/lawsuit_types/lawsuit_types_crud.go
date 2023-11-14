//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitType - версия структуры модели, с учётом имен и типов полей
var versionLawsuitType uint32

// crud_LawsuitType - объект контроллер crud операций
var crud_LawsuitType ICrud_LawsuitType

type ICrud_LawsuitType interface {
	Read(*LawsuitType) error
	Save(*LawsuitType) error
	Update(*LawsuitType) error
	Create(*LawsuitType) error
	Delete(*LawsuitType) error
	Restore(*LawsuitType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitType) TableNameDB() string {
	return "lawsuit_types"
}

// NewLawsuitType - возвращает новый	объект
func NewLawsuitType() LawsuitType {
	return LawsuitType{}
}

// AsLawsuitType - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitType(b []byte) (LawsuitType, error) {
	c := NewLawsuitType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitType(), err
	}
	return c, nil
}

// LawsuitTypeAsBytes - упаковывает объект в массив байтов
func LawsuitTypeAsBytes(m *LawsuitType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitType) GetStructVersion() uint32 {
	if versionLawsuitType == 0 {
		versionLawsuitType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitType) GetJSON() (string, error) {
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
func (m *LawsuitType) Read() error {
	err := crud_LawsuitType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitType) Save() error {
	err := crud_LawsuitType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitType) Update() error {
	err := crud_LawsuitType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitType) Create() error {
	err := crud_LawsuitType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitType) Delete() error {
	err := crud_LawsuitType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitType) Restore() error {
	err := crud_LawsuitType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitType) SetCrudInterface(crud ICrud_LawsuitType) {
	crud_LawsuitType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
