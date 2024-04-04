//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package channel

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionChannel - версия структуры модели, с учётом имен и типов полей
var versionChannel uint32

// Crud_Channel - объект контроллер crud операций
var Crud_Channel ICrud_Channel

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Channel interface {
	Read(*Channel) error
	Save(*Channel) error
	Update(*Channel) error
	Create(*Channel) error
	ReadFromCache(ID int64) (Channel, error)
	UpdateManyFields(*Channel, []string) error
	Update_Code(*Channel) error
	Update_Description(*Channel) error
	Update_IsActive(*Channel) error
	Update_LatName(*Channel) error
	Update_Name(*Channel) error
	Update_StekCode(*Channel) error
}

// TableName - возвращает имя таблицы в БД
func (m Channel) TableNameDB() string {
	return "channel"
}

// NewChannel - возвращает новый	объект
func NewChannel() Channel {
	return Channel{}
}

// AsChannel - создаёт объект из упакованного объекта в массиве байтов
func AsChannel(b []byte) (Channel, error) {
	c := NewChannel()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewChannel(), err
	}
	return c, nil
}

// ChannelAsBytes - упаковывает объект в массив байтов
func ChannelAsBytes(m *Channel) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Channel) GetStructVersion() uint32 {
	if versionChannel == 0 {
		versionChannel = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionChannel
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Channel) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Channel) GetJSON() (string, error) {
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
func (m *Channel) Read() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Channel.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Channel) Save() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Channel.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Channel) Update() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Channel.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Channel) Create() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Channel.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Channel) ReadFromCache(ID int64) (Channel, error) {
	Otvet := Channel{}
	var err error

	if Crud_Channel == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Channel.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Channel) SetCrudInterface(crud ICrud_Channel) {
	Crud_Channel = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Channel) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
