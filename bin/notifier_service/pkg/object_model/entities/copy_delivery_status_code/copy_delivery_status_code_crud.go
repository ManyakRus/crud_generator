//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package copy_delivery_status_code

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionCopyDeliveryStatusCode - версия структуры модели, с учётом имен и типов полей
var versionCopyDeliveryStatusCode uint32

// Crud_CopyDeliveryStatusCode - объект контроллер crud операций
var Crud_CopyDeliveryStatusCode ICrud_CopyDeliveryStatusCode

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_CopyDeliveryStatusCode interface {
	Read(*CopyDeliveryStatusCode) error
	Save(*CopyDeliveryStatusCode) error
	Update(*CopyDeliveryStatusCode) error
	Create(*CopyDeliveryStatusCode) error
	ReadFromCache(ID int64) (CopyDeliveryStatusCode, error)
	UpdateManyFields(*CopyDeliveryStatusCode, []string) error
	Update_Code(*CopyDeliveryStatusCode) error
}

// TableName - возвращает имя таблицы в БД
func (m CopyDeliveryStatusCode) TableNameDB() string {
	return "copy_delivery_status_code"
}

// NewCopyDeliveryStatusCode - возвращает новый	объект
func NewCopyDeliveryStatusCode() CopyDeliveryStatusCode {
	return CopyDeliveryStatusCode{}
}

// AsCopyDeliveryStatusCode - создаёт объект из упакованного объекта в массиве байтов
func AsCopyDeliveryStatusCode(b []byte) (CopyDeliveryStatusCode, error) {
	c := NewCopyDeliveryStatusCode()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewCopyDeliveryStatusCode(), err
	}
	return c, nil
}

// CopyDeliveryStatusCodeAsBytes - упаковывает объект в массив байтов
func CopyDeliveryStatusCodeAsBytes(m *CopyDeliveryStatusCode) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m CopyDeliveryStatusCode) GetStructVersion() uint32 {
	if versionCopyDeliveryStatusCode == 0 {
		versionCopyDeliveryStatusCode = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionCopyDeliveryStatusCode
}

// GetModelFromJSON - создаёт модель из строки json
func (m *CopyDeliveryStatusCode) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m CopyDeliveryStatusCode) GetJSON() (string, error) {
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
func (m *CopyDeliveryStatusCode) Read() error {
	if Crud_CopyDeliveryStatusCode == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_CopyDeliveryStatusCode.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *CopyDeliveryStatusCode) Save() error {
	if Crud_CopyDeliveryStatusCode == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_CopyDeliveryStatusCode.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *CopyDeliveryStatusCode) Update() error {
	if Crud_CopyDeliveryStatusCode == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_CopyDeliveryStatusCode.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *CopyDeliveryStatusCode) Create() error {
	if Crud_CopyDeliveryStatusCode == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_CopyDeliveryStatusCode.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *CopyDeliveryStatusCode) ReadFromCache(ID int64) (CopyDeliveryStatusCode, error) {
	Otvet := CopyDeliveryStatusCode{}
	var err error

	if Crud_CopyDeliveryStatusCode == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_CopyDeliveryStatusCode.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m CopyDeliveryStatusCode) SetCrudInterface(crud ICrud_CopyDeliveryStatusCode) {
	Crud_CopyDeliveryStatusCode = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *CopyDeliveryStatusCode) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_CopyDeliveryStatusCode == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_CopyDeliveryStatusCode.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
