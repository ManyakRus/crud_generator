//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package meter_list

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionMeterList - версия структуры модели, с учётом имен и типов полей
var versionMeterList uint32

// Crud_MeterList - объект контроллер crud операций
var Crud_MeterList ICrud_MeterList

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_MeterList interface {
	Read(*MeterList) error
	Save(*MeterList) error
	Update(*MeterList) error
	Create(*MeterList) error
	ReadFromCache(ID int64) (MeterList, error)
	UpdateManyFields(*MeterList, []string) error
	Update_ContactInfo(*MeterList) error
	Update_Flat(*MeterList) error
	Update_House(*MeterList) error
	Update_Locality(*MeterList) error
	Update_Ls(*MeterList) error
	Update_Plot(*MeterList) error
	Update_Region(*MeterList) error
	Update_Street(*MeterList) error
}

// TableName - возвращает имя таблицы в БД
func (m MeterList) TableNameDB() string {
	return "meter_list"
}

// NewMeterList - возвращает новый	объект
func NewMeterList() MeterList {
	return MeterList{}
}

// AsMeterList - создаёт объект из упакованного объекта в массиве байтов
func AsMeterList(b []byte) (MeterList, error) {
	c := NewMeterList()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMeterList(), err
	}
	return c, nil
}

// MeterListAsBytes - упаковывает объект в массив байтов
func MeterListAsBytes(m *MeterList) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m MeterList) GetStructVersion() uint32 {
	if versionMeterList == 0 {
		versionMeterList = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMeterList
}

// GetModelFromJSON - создаёт модель из строки json
func (m *MeterList) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m MeterList) GetJSON() (string, error) {
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
func (m *MeterList) Read() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MeterList.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *MeterList) Save() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MeterList.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *MeterList) Update() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MeterList.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *MeterList) Create() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MeterList.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *MeterList) ReadFromCache(ID int64) (MeterList, error) {
	Otvet := MeterList{}
	var err error

	if Crud_MeterList == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_MeterList.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m MeterList) SetCrudInterface(crud ICrud_MeterList) {
	Crud_MeterList = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *MeterList) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
