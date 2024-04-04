//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package message

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionMessage - версия структуры модели, с учётом имен и типов полей
var versionMessage uint32

// Crud_Message - объект контроллер crud операций
var Crud_Message ICrud_Message

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Message interface {
	Read(*Message) error
	Save(*Message) error
	Update(*Message) error
	Create(*Message) error
	ReadFromCache(ID int64) (Message, error)
	UpdateManyFields(*Message, []string) error
	Update_Attachments(*Message) error
	Update_CanceledMsgID(*Message) error
	Update_ChannelID(*Message) error
	Update_ContactAddress(*Message) error
	Update_ContactInfo(*Message) error
	Update_CountTryCheck(*Message) error
	Update_CountTrySent(*Message) error
	Update_DeliveredAt(*Message) error
	Update_DeliveryStatusID(*Message) error
	Update_ExternalID(*Message) error
	Update_FilialID(*Message) error
	Update_GateName(*Message) error
	Update_IncomingEventID(*Message) error
	Update_IsChecked(*Message) error
	Update_IsOrganisationAcc(*Message) error
	Update_IsSent(*Message) error
	Update_MailingID(*Message) error
	Update_Msg(*Message) error
	Update_PersonalAcc(*Message) error
	Update_SendingStatus(*Message) error
	Update_Seq(*Message) error
	Update_TryCheckAt(*Message) error
	Update_TrySendAt(*Message) error
	Update_UserAddress(*Message) error
	Update_UserID(*Message) error
	Update_Utm(*Message) error
}

// TableName - возвращает имя таблицы в БД
func (m Message) TableNameDB() string {
	return "message"
}

// NewMessage - возвращает новый	объект
func NewMessage() Message {
	return Message{}
}

// AsMessage - создаёт объект из упакованного объекта в массиве байтов
func AsMessage(b []byte) (Message, error) {
	c := NewMessage()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMessage(), err
	}
	return c, nil
}

// MessageAsBytes - упаковывает объект в массив байтов
func MessageAsBytes(m *Message) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Message) GetStructVersion() uint32 {
	if versionMessage == 0 {
		versionMessage = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMessage
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Message) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Message) GetJSON() (string, error) {
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
func (m *Message) Read() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Message.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Message) Save() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Message.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Message) Update() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Message.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Message) Create() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Message.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Message) ReadFromCache(ID int64) (Message, error) {
	Otvet := Message{}
	var err error

	if Crud_Message == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Message.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Message) SetCrudInterface(crud ICrud_Message) {
	Crud_Message = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Message) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
