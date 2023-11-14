//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_stage_types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitStageType - версия структуры модели, с учётом имен и типов полей
var versionLawsuitStageType uint32

// crud_LawsuitStageType - объект контроллер crud операций
var crud_LawsuitStageType ICrud_LawsuitStageType

type ICrud_LawsuitStageType interface {
	Read(*LawsuitStageType) error
	Save(*LawsuitStageType) error
	Update(*LawsuitStageType) error
	Create(*LawsuitStageType) error
	Delete(*LawsuitStageType) error
	Restore(*LawsuitStageType) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitStageType) TableNameDB() string {
	return "lawsuit_stage_types"
}

// NewLawsuitStageType - возвращает новый	объект
func NewLawsuitStageType() LawsuitStageType {
	return LawsuitStageType{}
}

// AsLawsuitStageType - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitStageType(b []byte) (LawsuitStageType, error) {
	c := NewLawsuitStageType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitStageType(), err
	}
	return c, nil
}

// LawsuitStageTypeAsBytes - упаковывает объект в массив байтов
func LawsuitStageTypeAsBytes(m *LawsuitStageType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitStageType) GetStructVersion() uint32 {
	if versionLawsuitStageType == 0 {
		versionLawsuitStageType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitStageType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitStageType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitStageType) GetJSON() (string, error) {
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
func (m *LawsuitStageType) Read() error {
	err := crud_LawsuitStageType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitStageType) Save() error {
	err := crud_LawsuitStageType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitStageType) Update() error {
	err := crud_LawsuitStageType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitStageType) Create() error {
	err := crud_LawsuitStageType.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitStageType) Delete() error {
	err := crud_LawsuitStageType.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitStageType) Restore() error {
	err := crud_LawsuitStageType.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitStageType) SetCrudInterface(crud ICrud_LawsuitStageType) {
	crud_LawsuitStageType = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
