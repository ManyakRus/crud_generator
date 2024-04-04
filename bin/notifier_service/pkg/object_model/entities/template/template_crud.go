//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package template

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionTemplate - версия структуры модели, с учётом имен и типов полей
var versionTemplate uint32

// Crud_Template - объект контроллер crud операций
var Crud_Template ICrud_Template

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Template interface {
	Read(*Template) error
	Save(*Template) error
	Update(*Template) error
	Create(*Template) error
	ReadFromCache(ID int64) (Template, error)
	UpdateManyFields(*Template, []string) error
	Update_ChannelID(*Template) error
	Update_Code(*Template) error
	Update_IncomingEventID(*Template) error
	Update_Name(*Template) error
	Update_Template(*Template) error
	Update_UpdatedAt(*Template) error
}

// TableName - возвращает имя таблицы в БД
func (m Template) TableNameDB() string {
	return "template"
}

// NewTemplate - возвращает новый	объект
func NewTemplate() Template {
	return Template{}
}

// AsTemplate - создаёт объект из упакованного объекта в массиве байтов
func AsTemplate(b []byte) (Template, error) {
	c := NewTemplate()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewTemplate(), err
	}
	return c, nil
}

// TemplateAsBytes - упаковывает объект в массив байтов
func TemplateAsBytes(m *Template) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Template) GetStructVersion() uint32 {
	if versionTemplate == 0 {
		versionTemplate = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionTemplate
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Template) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Template) GetJSON() (string, error) {
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
func (m *Template) Read() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Template.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Template) Save() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Template.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Template) Update() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Template.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Template) Create() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Template.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Template) ReadFromCache(ID int64) (Template, error) {
	Otvet := Template{}
	var err error

	if Crud_Template == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Template.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Template) SetCrudInterface(crud ICrud_Template) {
	Crud_Template = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Template) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
