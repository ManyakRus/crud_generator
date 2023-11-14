//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_reason_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitReasonType - версия структуры модели, с учётом имен и типов полей
var versionLawsuitReasonType uint32

// crud_LawsuitReasonType - объект контроллер crud операций
var crud_LawsuitReasonType ICrud_LawsuitReasonType

type ICrud_LawsuitReasonType interface {
	Read(*LawsuitReasonType) error
	Save(*LawsuitReasonType) error
	Update(*LawsuitReasonType) error
	Create(*LawsuitReasonType) error
	Delete(*LawsuitReasonType) error
	Restore(*LawsuitReasonType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitReasonType) TableNameDB() string {
	return "lawsuit_reason_types"
}

// NewLawsuitReasonType - возвращает новый	объект
func NewLawsuitReasonType() LawsuitReasonType {
	return LawsuitReasonType{}
}

// AsLawsuitReasonType - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitReasonType(b []byte) (LawsuitReasonType, error) {
	c := NewLawsuitReasonType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitReasonType(), err
	}
	return c, nil
}

// LawsuitReasonTypeAsBytes - упаковывает объект в массив байтов
func LawsuitReasonTypeAsBytes(m *LawsuitReasonType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitReasonType) GetStructVersion() uint32 {
	if versionLawsuitReasonType == 0 {
		versionLawsuitReasonType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitReasonType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitReasonType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitReasonType) GetJSON() (string, error) {
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
func (m *LawsuitReasonType) Read() error {
	err := crud_LawsuitReasonType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitReasonType) Save() error {
	err := crud_LawsuitReasonType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitReasonType) Update() error {
	err := crud_LawsuitReasonType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitReasonType) Create() error {
	err := crud_LawsuitReasonType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitReasonType) Delete() error {
	err := crud_LawsuitReasonType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitReasonType) Restore() error {
	err := crud_LawsuitReasonType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitReasonType) SetCrudInterface(crud ICrud_LawsuitReasonType) {
	crud_LawsuitReasonType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
