//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package statistic

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionStatistic - версия структуры модели, с учётом имен и типов полей
var versionStatistic uint32

// Crud_Statistic - объект контроллер crud операций
var Crud_Statistic ICrud_Statistic

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Statistic interface {
	Read(*Statistic) error
	Save(*Statistic) error
	Update(*Statistic) error
	Create(*Statistic) error
	ReadFromCache(ID int64) (Statistic, error)
	UpdateManyFields(*Statistic, []string) error
	Update_FilialIds(*Statistic) error
	Update_LastUpdate(*Statistic) error
	Update_MailingChannelID(*Statistic) error
	Update_MailingCode(*Statistic) error
	Update_MailingEnd(*Statistic) error
	Update_MailingID(*Statistic) error
	Update_MailingStart(*Statistic) error
	Update_MailingStatus(*Statistic) error
	Update_MessagesDelivered(*Statistic) error
	Update_MessagesSent(*Statistic) error
	Update_MessagesTotal(*Statistic) error
	Update_MessagesUndelivered(*Statistic) error
	Update_MessagesUnsent(*Statistic) error
}

// TableName - возвращает имя таблицы в БД
func (m Statistic) TableNameDB() string {
	return "statistic"
}

// NewStatistic - возвращает новый	объект
func NewStatistic() Statistic {
	return Statistic{}
}

// AsStatistic - создаёт объект из упакованного объекта в массиве байтов
func AsStatistic(b []byte) (Statistic, error) {
	c := NewStatistic()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewStatistic(), err
	}
	return c, nil
}

// StatisticAsBytes - упаковывает объект в массив байтов
func StatisticAsBytes(m *Statistic) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Statistic) GetStructVersion() uint32 {
	if versionStatistic == 0 {
		versionStatistic = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionStatistic
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Statistic) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Statistic) GetJSON() (string, error) {
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
func (m *Statistic) Read() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Statistic.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Statistic) Save() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Statistic.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Statistic) Update() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Statistic.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Statistic) Create() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Statistic.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Statistic) ReadFromCache(ID int64) (Statistic, error) {
	Otvet := Statistic{}
	var err error

	if Crud_Statistic == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Statistic.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Statistic) SetCrudInterface(crud ICrud_Statistic) {
	Crud_Statistic = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Statistic) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
