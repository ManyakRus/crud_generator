//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package w_options

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionWOption - версия структуры модели, с учётом имен и типов полей
var versionWOption uint32

// Crud_WOption - объект контроллер crud операций
var Crud_WOption ICrud_WOption

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_WOption interface {
	Read(*WOption) error
	Save(*WOption) error
	Update(*WOption) error
	Create(*WOption) error
	ReadFromCache(ID int64) (WOption, error)
	UpdateManyFields(*WOption, []string) error
	Update_Description(*WOption) error
	Update_IsActiv(*WOption) error
	Update_Name(*WOption) error
	Update_ValBool(*WOption) error
	Update_ValDt(*WOption) error
	Update_ValInt(*WOption) error
	Update_ValStr(*WOption) error
}

// TableName - возвращает имя таблицы в БД
func (m WOption) TableNameDB() string {
	return "w_options"
}

// NewWOption - возвращает новый	объект
func NewWOption() WOption {
	return WOption{}
}

// AsWOption - создаёт объект из упакованного объекта в массиве байтов
func AsWOption(b []byte) (WOption, error) {
	c := NewWOption()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewWOption(), err
	}
	return c, nil
}

// WOptionAsBytes - упаковывает объект в массив байтов
func WOptionAsBytes(m *WOption) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m WOption) GetStructVersion() uint32 {
	if versionWOption == 0 {
		versionWOption = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionWOption
}

// GetModelFromJSON - создаёт модель из строки json
func (m *WOption) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m WOption) GetJSON() (string, error) {
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
func (m *WOption) Read() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WOption.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *WOption) Save() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WOption.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *WOption) Update() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WOption.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *WOption) Create() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_WOption.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *WOption) ReadFromCache(ID int64) (WOption, error) {
	Otvet := WOption{}
	var err error

	if Crud_WOption == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_WOption.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WOption) SetCrudInterface(crud ICrud_WOption) {
	Crud_WOption = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *WOption) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
