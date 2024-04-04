//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package attachament_message

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionAttachamentMessage - версия структуры модели, с учётом имен и типов полей
var versionAttachamentMessage uint32

// Crud_AttachamentMessage - объект контроллер crud операций
var Crud_AttachamentMessage ICrud_AttachamentMessage

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_AttachamentMessage interface {
	Read(*AttachamentMessage) error
	Save(*AttachamentMessage) error
	Update(*AttachamentMessage) error
	Create(*AttachamentMessage) error
	ReadFromCache(ID int64) (AttachamentMessage, error)
	UpdateManyFields(*AttachamentMessage, []string) error
	Update_AttachamentID(*AttachamentMessage) error
	Update_MessageID(*AttachamentMessage) error
}

// TableName - возвращает имя таблицы в БД
func (m AttachamentMessage) TableNameDB() string {
	return "attachament_message"
}

// NewAttachamentMessage - возвращает новый	объект
func NewAttachamentMessage() AttachamentMessage {
	return AttachamentMessage{}
}

// AsAttachamentMessage - создаёт объект из упакованного объекта в массиве байтов
func AsAttachamentMessage(b []byte) (AttachamentMessage, error) {
	c := NewAttachamentMessage()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewAttachamentMessage(), err
	}
	return c, nil
}

// AttachamentMessageAsBytes - упаковывает объект в массив байтов
func AttachamentMessageAsBytes(m *AttachamentMessage) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m AttachamentMessage) GetStructVersion() uint32 {
	if versionAttachamentMessage == 0 {
		versionAttachamentMessage = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionAttachamentMessage
}

// GetModelFromJSON - создаёт модель из строки json
func (m *AttachamentMessage) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m AttachamentMessage) GetJSON() (string, error) {
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
func (m *AttachamentMessage) Read() error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_AttachamentMessage.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *AttachamentMessage) Save() error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_AttachamentMessage.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *AttachamentMessage) Update() error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_AttachamentMessage.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *AttachamentMessage) Create() error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_AttachamentMessage.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *AttachamentMessage) ReadFromCache(ID int64) (AttachamentMessage, error) {
	Otvet := AttachamentMessage{}
	var err error

	if Crud_AttachamentMessage == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_AttachamentMessage.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m AttachamentMessage) SetCrudInterface(crud ICrud_AttachamentMessage) {
	Crud_AttachamentMessage = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *AttachamentMessage) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_AttachamentMessage.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
