//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package telegram_users

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionTelegramUser - версия структуры модели, с учётом имен и типов полей
var versionTelegramUser uint32

// Crud_TelegramUser - объект контроллер crud операций
var Crud_TelegramUser ICrud_TelegramUser

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_TelegramUser interface {
	Read(*TelegramUser) error
	Save(*TelegramUser) error
	Update(*TelegramUser) error
	Create(*TelegramUser) error
	ReadFromCache(ID int64) (TelegramUser, error)
	UpdateManyFields(*TelegramUser, []string) error
	Update_AllowedBillNotify(*TelegramUser) error
	Update_AllowedDebtNotify(*TelegramUser) error
	Update_AllowedMeterNotify(*TelegramUser) error
	Update_AllowedMiscNotify(*TelegramUser) error
	Update_BlockedByUser(*TelegramUser) error
	Update_ChatID(*TelegramUser) error
	Update_ContactInfo(*TelegramUser) error
	Update_DateStatusChanged(*TelegramUser) error
	Update_IsTester(*TelegramUser) error
}

// TableName - возвращает имя таблицы в БД
func (m TelegramUser) TableNameDB() string {
	return "telegram_users"
}

// NewTelegramUser - возвращает новый	объект
func NewTelegramUser() TelegramUser {
	return TelegramUser{}
}

// AsTelegramUser - создаёт объект из упакованного объекта в массиве байтов
func AsTelegramUser(b []byte) (TelegramUser, error) {
	c := NewTelegramUser()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewTelegramUser(), err
	}
	return c, nil
}

// TelegramUserAsBytes - упаковывает объект в массив байтов
func TelegramUserAsBytes(m *TelegramUser) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m TelegramUser) GetStructVersion() uint32 {
	if versionTelegramUser == 0 {
		versionTelegramUser = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionTelegramUser
}

// GetModelFromJSON - создаёт модель из строки json
func (m *TelegramUser) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m TelegramUser) GetJSON() (string, error) {
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
func (m *TelegramUser) Read() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUser.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *TelegramUser) Save() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUser.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *TelegramUser) Update() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUser.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *TelegramUser) Create() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUser.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *TelegramUser) ReadFromCache(ID int64) (TelegramUser, error) {
	Otvet := TelegramUser{}
	var err error

	if Crud_TelegramUser == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_TelegramUser.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TelegramUser) SetCrudInterface(crud ICrud_TelegramUser) {
	Crud_TelegramUser = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *TelegramUser) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
