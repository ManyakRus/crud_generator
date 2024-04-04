//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package mailing_stats

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionMailingStat - версия структуры модели, с учётом имен и типов полей
var versionMailingStat uint32

// Crud_MailingStat - объект контроллер crud операций
var Crud_MailingStat ICrud_MailingStat

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_MailingStat interface {
	Read(*MailingStat) error
	Save(*MailingStat) error
	Update(*MailingStat) error
	Create(*MailingStat) error
	ReadFromCache(ID int64) (MailingStat, error)
	UpdateManyFields(*MailingStat, []string) error
	Update_ErrorsCount(*MailingStat) error
	Update_MailingID(*MailingStat) error
	Update_MessagesTotalCount(*MailingStat) error
	Update_RedirectCount(*MailingStat) error
	Update_SentMessagesCount(*MailingStat) error
}

// TableName - возвращает имя таблицы в БД
func (m MailingStat) TableNameDB() string {
	return "mailing_stats"
}

// NewMailingStat - возвращает новый	объект
func NewMailingStat() MailingStat {
	return MailingStat{}
}

// AsMailingStat - создаёт объект из упакованного объекта в массиве байтов
func AsMailingStat(b []byte) (MailingStat, error) {
	c := NewMailingStat()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewMailingStat(), err
	}
	return c, nil
}

// MailingStatAsBytes - упаковывает объект в массив байтов
func MailingStatAsBytes(m *MailingStat) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m MailingStat) GetStructVersion() uint32 {
	if versionMailingStat == 0 {
		versionMailingStat = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionMailingStat
}

// GetModelFromJSON - создаёт модель из строки json
func (m *MailingStat) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m MailingStat) GetJSON() (string, error) {
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
func (m *MailingStat) Read() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MailingStat.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *MailingStat) Save() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MailingStat.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *MailingStat) Update() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MailingStat.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *MailingStat) Create() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_MailingStat.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *MailingStat) ReadFromCache(ID int64) (MailingStat, error) {
	Otvet := MailingStat{}
	var err error

	if Crud_MailingStat == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_MailingStat.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m MailingStat) SetCrudInterface(crud ICrud_MailingStat) {
	Crud_MailingStat = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *MailingStat) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MailingStat.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
