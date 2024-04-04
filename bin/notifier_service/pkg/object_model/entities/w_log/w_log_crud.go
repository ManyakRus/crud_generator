//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package w_log

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionWLog - версия структуры модели, с учётом имен и типов полей
var versionWLog uint32

// Crud_WLog - объект контроллер crud операций
var Crud_WLog ICrud_WLog

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_WLog interface {
	Read(*WLog) error
	Save(*WLog) error
	Update(*WLog) error
	Create(*WLog) error
	ReadFromCache(ID int64) (WLog, error)
	UpdateManyFields(*WLog, []string) error
	Update_Addr(*WLog) error
	Update_DtTz(*WLog) error
	Update_DtUtc(*WLog) error
	Update_Msg(*WLog) error
	Update_MsgByte(*WLog) error
	Update_MsgJsonb(*WLog) error
	Update_Pid(*WLog) error
	Update_Proc(*WLog) error
	Update_Usr(*WLog) error
	Update_Var(*WLog) error
}

// TableName - возвращает имя таблицы в БД
func (m WLog) TableNameDB() string {
	return "w_log"
}

// NewWLog - возвращает новый	объект
func NewWLog() WLog {
	return WLog{}
}

// AsWLog - создаёт объект из упакованного объекта в массиве байтов
func AsWLog(b []byte) (WLog, error) {
	c := NewWLog()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewWLog(), err
	}
	return c, nil
}

// WLogAsBytes - упаковывает объект в массив байтов
func WLogAsBytes(m *WLog) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m WLog) GetStructVersion() uint32 {
	if versionWLog == 0 {
		versionWLog = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionWLog
}

// GetModelFromJSON - создаёт модель из строки json
func (m *WLog) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m WLog) GetJSON() (string, error) {
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
func (m *WLog) Read() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLog.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *WLog) Save() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLog.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *WLog) Update() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLog.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *WLog) Create() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLog.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *WLog) ReadFromCache(ID int64) (WLog, error) {
	Otvet := WLog{}
	var err error

	if Crud_WLog == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_WLog.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WLog) SetCrudInterface(crud ICrud_WLog) {
	Crud_WLog = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *WLog) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
