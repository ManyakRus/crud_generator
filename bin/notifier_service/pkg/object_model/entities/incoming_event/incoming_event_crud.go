//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package incoming_event

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionIncomingEvent - версия структуры модели, с учётом имен и типов полей
var versionIncomingEvent uint32

// Crud_IncomingEvent - объект контроллер crud операций
var Crud_IncomingEvent ICrud_IncomingEvent

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_IncomingEvent interface {
	Read(*IncomingEvent) error
	Save(*IncomingEvent) error
	Update(*IncomingEvent) error
	Create(*IncomingEvent) error
	ReadFromCache(ID int64) (IncomingEvent, error)
	UpdateManyFields(*IncomingEvent, []string) error
	Update_IncomingChannel(*IncomingEvent) error
	Update_IsParsed(*IncomingEvent) error
	Update_IsProcessed(*IncomingEvent) error
	Update_ParseError(*IncomingEvent) error
	Update_ReceivedAt(*IncomingEvent) error
	Update_SrcMsgJson(*IncomingEvent) error
	Update_SysID(*IncomingEvent) error
}

// TableName - возвращает имя таблицы в БД
func (m IncomingEvent) TableNameDB() string {
	return "incoming_event"
}

// NewIncomingEvent - возвращает новый	объект
func NewIncomingEvent() IncomingEvent {
	return IncomingEvent{}
}

// AsIncomingEvent - создаёт объект из упакованного объекта в массиве байтов
func AsIncomingEvent(b []byte) (IncomingEvent, error) {
	c := NewIncomingEvent()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewIncomingEvent(), err
	}
	return c, nil
}

// IncomingEventAsBytes - упаковывает объект в массив байтов
func IncomingEventAsBytes(m *IncomingEvent) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m IncomingEvent) GetStructVersion() uint32 {
	if versionIncomingEvent == 0 {
		versionIncomingEvent = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionIncomingEvent
}

// GetModelFromJSON - создаёт модель из строки json
func (m *IncomingEvent) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m IncomingEvent) GetJSON() (string, error) {
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
func (m *IncomingEvent) Read() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_IncomingEvent.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *IncomingEvent) Save() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_IncomingEvent.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *IncomingEvent) Update() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_IncomingEvent.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *IncomingEvent) Create() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_IncomingEvent.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *IncomingEvent) ReadFromCache(ID int64) (IncomingEvent, error) {
	Otvet := IncomingEvent{}
	var err error

	if Crud_IncomingEvent == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_IncomingEvent.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m IncomingEvent) SetCrudInterface(crud ICrud_IncomingEvent) {
	Crud_IncomingEvent = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *IncomingEvent) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
