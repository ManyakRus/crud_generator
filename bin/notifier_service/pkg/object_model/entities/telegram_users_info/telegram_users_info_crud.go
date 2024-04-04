//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package telegram_users_info

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionTelegramUsersInfo - версия структуры модели, с учётом имен и типов полей
var versionTelegramUsersInfo uint32

// Crud_TelegramUsersInfo - объект контроллер crud операций
var Crud_TelegramUsersInfo ICrud_TelegramUsersInfo

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_TelegramUsersInfo interface {
	Read(*TelegramUsersInfo) error
	Save(*TelegramUsersInfo) error
	Update(*TelegramUsersInfo) error
	Create(*TelegramUsersInfo) error
	ReadFromCache(ID int64) (TelegramUsersInfo, error)
	UpdateManyFields(*TelegramUsersInfo, []string) error
	Update_Address(*TelegramUsersInfo) error
	Update_ChatID(*TelegramUsersInfo) error
	Update_IsChecked(*TelegramUsersInfo) error
	Update_PersAcc(*TelegramUsersInfo) error
}

// TableName - возвращает имя таблицы в БД
func (m TelegramUsersInfo) TableNameDB() string {
	return "telegram_users_info"
}

// NewTelegramUsersInfo - возвращает новый	объект
func NewTelegramUsersInfo() TelegramUsersInfo {
	return TelegramUsersInfo{}
}

// AsTelegramUsersInfo - создаёт объект из упакованного объекта в массиве байтов
func AsTelegramUsersInfo(b []byte) (TelegramUsersInfo, error) {
	c := NewTelegramUsersInfo()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewTelegramUsersInfo(), err
	}
	return c, nil
}

// TelegramUsersInfoAsBytes - упаковывает объект в массив байтов
func TelegramUsersInfoAsBytes(m *TelegramUsersInfo) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m TelegramUsersInfo) GetStructVersion() uint32 {
	if versionTelegramUsersInfo == 0 {
		versionTelegramUsersInfo = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionTelegramUsersInfo
}

// GetModelFromJSON - создаёт модель из строки json
func (m *TelegramUsersInfo) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m TelegramUsersInfo) GetJSON() (string, error) {
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
func (m *TelegramUsersInfo) Read() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUsersInfo.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *TelegramUsersInfo) Save() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUsersInfo.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *TelegramUsersInfo) Update() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUsersInfo.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *TelegramUsersInfo) Create() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TelegramUsersInfo.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *TelegramUsersInfo) ReadFromCache(ID int64) (TelegramUsersInfo, error) {
	Otvet := TelegramUsersInfo{}
	var err error

	if Crud_TelegramUsersInfo == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_TelegramUsersInfo.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TelegramUsersInfo) SetCrudInterface(crud ICrud_TelegramUsersInfo) {
	Crud_TelegramUsersInfo = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *TelegramUsersInfo) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUsersInfo.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
