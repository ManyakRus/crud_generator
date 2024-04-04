//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package w_log_message_del

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionWLogMessageDel - версия структуры модели, с учётом имен и типов полей
var versionWLogMessageDel uint32

// Crud_WLogMessageDel - объект контроллер crud операций
var Crud_WLogMessageDel ICrud_WLogMessageDel

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_WLogMessageDel interface {
	Read(*WLogMessageDel) error
	Save(*WLogMessageDel) error
	Update(*WLogMessageDel) error
	Create(*WLogMessageDel) error
	ReadFromCache(ID int64) (WLogMessageDel, error)
	UpdateManyFields(*WLogMessageDel, []string) error
	Update_Addr(*WLogMessageDel) error
	Update_Dt(*WLogMessageDel) error
	Update_Msg(*WLogMessageDel) error
	Update_Pid(*WLogMessageDel) error
	Update_Proc(*WLogMessageDel) error
	Update_Usr(*WLogMessageDel) error
	Update_Var(*WLogMessageDel) error
}

// TableName - возвращает имя таблицы в БД
func (m WLogMessageDel) TableNameDB() string {
	return "w_log_message_del"
}

// NewWLogMessageDel - возвращает новый	объект
func NewWLogMessageDel() WLogMessageDel {
	return WLogMessageDel{}
}

// AsWLogMessageDel - создаёт объект из упакованного объекта в массиве байтов
func AsWLogMessageDel(b []byte) (WLogMessageDel, error) {
	c := NewWLogMessageDel()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewWLogMessageDel(), err
	}
	return c, nil
}

// WLogMessageDelAsBytes - упаковывает объект в массив байтов
func WLogMessageDelAsBytes(m *WLogMessageDel) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m WLogMessageDel) GetStructVersion() uint32 {
	if versionWLogMessageDel == 0 {
		versionWLogMessageDel = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionWLogMessageDel
}

// GetModelFromJSON - создаёт модель из строки json
func (m *WLogMessageDel) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m WLogMessageDel) GetJSON() (string, error) {
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
func (m *WLogMessageDel) Read() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLogMessageDel.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *WLogMessageDel) Save() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLogMessageDel.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *WLogMessageDel) Update() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLogMessageDel.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *WLogMessageDel) Create() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WLogMessageDel.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *WLogMessageDel) ReadFromCache(ID int64) (WLogMessageDel, error) {
	Otvet := WLogMessageDel{}
	var err error

	if Crud_WLogMessageDel == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_WLogMessageDel.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WLogMessageDel) SetCrudInterface(crud ICrud_WLogMessageDel) {
	Crud_WLogMessageDel = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *WLogMessageDel) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
