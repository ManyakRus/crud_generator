//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package mailing

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionMailing - версия структуры модели, с учётом имен и типов полей
var versionMailing uint32

// Crud_Mailing - объект контроллер crud операций
var Crud_Mailing ICrud_Mailing

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Mailing interface {
	Read(*Mailing) error
	Save(*Mailing) error
	Update(*Mailing) error
	Create(*Mailing) error
	ReadFromCache(ID int64) (Mailing, error)
	UpdateManyFields(*Mailing, []string) error
	Update_Code(*Mailing) error
	Update_EndAt(*Mailing) error
	Update_IncomingEventID(*Mailing) error
	Update_IsCanceled(*Mailing) error
	Update_IsFinished(*Mailing) error
	Update_ProcessCode(*Mailing) error
	Update_ReceivedAt(*Mailing) error
	Update_StartAt(*Mailing) error
	Update_Subject(*Mailing) error
	Update_TimeZone(*Mailing) error
}

// TableName - возвращает имя таблицы в БД
func (m Mailing) TableNameDB() string {
	return "mailing"
}

// NewMailing - возвращает новый	объект
func NewMailing() Mailing {
	return Mailing{}
}

// AsMailing - создаёт объект из упакованного объекта в массиве байтов
func AsMailing(b []byte) (Mailing, error) {
	c := NewMailing()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMailing(), err
	}
	return c, nil
}

// MailingAsBytes - упаковывает объект в массив байтов
func MailingAsBytes(m *Mailing) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Mailing) GetStructVersion() uint32 {
	if versionMailing == 0 {
		versionMailing = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMailing
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Mailing) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Mailing) GetJSON() (string, error) {
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
func (m *Mailing) Read() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Mailing.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Mailing) Save() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Mailing.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Mailing) Update() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Mailing.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Mailing) Create() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Mailing.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Mailing) ReadFromCache(ID int64) (Mailing, error) {
	Otvet := Mailing{}
	var err error

	if Crud_Mailing == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Mailing.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Mailing) SetCrudInterface(crud ICrud_Mailing) {
	Crud_Mailing = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Mailing) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
