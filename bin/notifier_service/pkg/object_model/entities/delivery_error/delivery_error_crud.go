//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package delivery_error

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionDeliveryError - версия структуры модели, с учётом имен и типов полей
var versionDeliveryError uint32

// Crud_DeliveryError - объект контроллер crud операций
var Crud_DeliveryError ICrud_DeliveryError

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_DeliveryError interface {
	Read(*DeliveryError) error
	Save(*DeliveryError) error
	Update(*DeliveryError) error
	Create(*DeliveryError) error
	ReadFromCache(ID int64) (DeliveryError, error)
	UpdateManyFields(*DeliveryError, []string) error
	Update_DeliveryStatusID(*DeliveryError) error
	Update_IsActive(*DeliveryError) error
	Update_TextError(*DeliveryError) error
	Update_UpdatedAt(*DeliveryError) error
}

// TableName - возвращает имя таблицы в БД
func (m DeliveryError) TableNameDB() string {
	return "delivery_error"
}

// NewDeliveryError - возвращает новый	объект
func NewDeliveryError() DeliveryError {
	return DeliveryError{}
}

// AsDeliveryError - создаёт объект из упакованного объекта в массиве байтов
func AsDeliveryError(b []byte) (DeliveryError, error) {
	c := NewDeliveryError()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDeliveryError(), err
	}
	return c, nil
}

// DeliveryErrorAsBytes - упаковывает объект в массив байтов
func DeliveryErrorAsBytes(m *DeliveryError) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DeliveryError) GetStructVersion() uint32 {
	if versionDeliveryError == 0 {
		versionDeliveryError = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDeliveryError
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DeliveryError) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DeliveryError) GetJSON() (string, error) {
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
func (m *DeliveryError) Read() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryError.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DeliveryError) Save() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryError.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DeliveryError) Update() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryError.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DeliveryError) Create() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryError.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *DeliveryError) ReadFromCache(ID int64) (DeliveryError, error) {
	Otvet := DeliveryError{}
	var err error

	if Crud_DeliveryError == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_DeliveryError.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DeliveryError) SetCrudInterface(crud ICrud_DeliveryError) {
	Crud_DeliveryError = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *DeliveryError) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryError.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
