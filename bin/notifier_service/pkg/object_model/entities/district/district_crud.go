//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package district

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionDistrict - версия структуры модели, с учётом имен и типов полей
var versionDistrict uint32

// Crud_District - объект контроллер crud операций
var Crud_District ICrud_District

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_District interface {
	Read(*District) error
	Save(*District) error
	Update(*District) error
	Create(*District) error
	ReadFromCache(ID int64) (District, error)
	UpdateManyFields(*District, []string) error
	Update_DepartmentCode(*District) error
	Update_DepartmentName(*District) error
	Update_DistrictCode(*District) error
	Update_DistrictName(*District) error
	Update_FilialID(*District) error
	Update_RegionCode(*District) error
}

// TableName - возвращает имя таблицы в БД
func (m District) TableNameDB() string {
	return "district"
}

// NewDistrict - возвращает новый	объект
func NewDistrict() District {
	return District{}
}

// AsDistrict - создаёт объект из упакованного объекта в массиве байтов
func AsDistrict(b []byte) (District, error) {
	c := NewDistrict()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDistrict(), err
	}
	return c, nil
}

// DistrictAsBytes - упаковывает объект в массив байтов
func DistrictAsBytes(m *District) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m District) GetStructVersion() uint32 {
	if versionDistrict == 0 {
		versionDistrict = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDistrict
}

// GetModelFromJSON - создаёт модель из строки json
func (m *District) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m District) GetJSON() (string, error) {
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
func (m *District) Read() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_District.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *District) Save() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_District.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *District) Update() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_District.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *District) Create() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_District.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *District) ReadFromCache(ID int64) (District, error) {
	Otvet := District{}
	var err error

	if Crud_District == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_District.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m District) SetCrudInterface(crud ICrud_District) {
	Crud_District = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *District) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
