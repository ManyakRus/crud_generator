//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package delivery_status

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionDeliveryStatus - версия структуры модели, с учётом имен и типов полей
var versionDeliveryStatus uint32

// Crud_DeliveryStatus - объект контроллер crud операций
var Crud_DeliveryStatus ICrud_DeliveryStatus

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_DeliveryStatus interface {
	Read(*DeliveryStatus) error
	Save(*DeliveryStatus) error
	Update(*DeliveryStatus) error
	Create(*DeliveryStatus) error
	ReadFromCache(ID int64) (DeliveryStatus, error)
	UpdateManyFields(*DeliveryStatus, []string) error
	Update_Code(*DeliveryStatus) error
	Update_Description(*DeliveryStatus) error
	Update_FormalName(*DeliveryStatus) error
	Update_IsActive(*DeliveryStatus) error
	Update_Name(*DeliveryStatus) error
	Update_Note(*DeliveryStatus) error
}

// TableName - возвращает имя таблицы в БД
func (m DeliveryStatus) TableNameDB() string {
	return "delivery_status"
}

// NewDeliveryStatus - возвращает новый	объект
func NewDeliveryStatus() DeliveryStatus {
	return DeliveryStatus{}
}

// AsDeliveryStatus - создаёт объект из упакованного объекта в массиве байтов
func AsDeliveryStatus(b []byte) (DeliveryStatus, error) {
	c := NewDeliveryStatus()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDeliveryStatus(), err
	}
	return c, nil
}

// DeliveryStatusAsBytes - упаковывает объект в массив байтов
func DeliveryStatusAsBytes(m *DeliveryStatus) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DeliveryStatus) GetStructVersion() uint32 {
	if versionDeliveryStatus == 0 {
		versionDeliveryStatus = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDeliveryStatus
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DeliveryStatus) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DeliveryStatus) GetJSON() (string, error) {
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
func (m *DeliveryStatus) Read() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryStatus.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DeliveryStatus) Save() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryStatus.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DeliveryStatus) Update() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryStatus.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DeliveryStatus) Create() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DeliveryStatus.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *DeliveryStatus) ReadFromCache(ID int64) (DeliveryStatus, error) {
	Otvet := DeliveryStatus{}
	var err error

	if Crud_DeliveryStatus == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_DeliveryStatus.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DeliveryStatus) SetCrudInterface(crud ICrud_DeliveryStatus) {
	Crud_DeliveryStatus = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *DeliveryStatus) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
