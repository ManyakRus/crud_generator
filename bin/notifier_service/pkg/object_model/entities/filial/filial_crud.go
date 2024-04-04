//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package filial

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionFilial - версия структуры модели, с учётом имен и типов полей
var versionFilial uint32

// Crud_Filial - объект контроллер crud операций
var Crud_Filial ICrud_Filial

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Filial interface {
	Read(*Filial) error
	Save(*Filial) error
	Update(*Filial) error
	Create(*Filial) error
	ReadFromCache(ID int64) (Filial, error)
	UpdateManyFields(*Filial, []string) error
	Update_AccPrefix(*Filial) error
	Update_DivisionName(*Filial) error
	Update_LongFormalName(*Filial) error
	Update_OperatorEmail(*Filial) error
	Update_RegionFormalName(*Filial) error
	Update_RegionSmallLatName(*Filial) error
	Update_RegionSmallRusName(*Filial) error
	Update_ShortFormalName(*Filial) error
}

// TableName - возвращает имя таблицы в БД
func (m Filial) TableNameDB() string {
	return "filial"
}

// NewFilial - возвращает новый	объект
func NewFilial() Filial {
	return Filial{}
}

// AsFilial - создаёт объект из упакованного объекта в массиве байтов
func AsFilial(b []byte) (Filial, error) {
	c := NewFilial()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewFilial(), err
	}
	return c, nil
}

// FilialAsBytes - упаковывает объект в массив байтов
func FilialAsBytes(m *Filial) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Filial) GetStructVersion() uint32 {
	if versionFilial == 0 {
		versionFilial = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionFilial
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Filial) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Filial) GetJSON() (string, error) {
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
func (m *Filial) Read() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Filial.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Filial) Save() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Filial.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Filial) Update() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Filial.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Filial) Create() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Filial.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Filial) ReadFromCache(ID int64) (Filial, error) {
	Otvet := Filial{}
	var err error

	if Crud_Filial == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Filial.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Filial) SetCrudInterface(crud ICrud_Filial) {
	Crud_Filial = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Filial) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
