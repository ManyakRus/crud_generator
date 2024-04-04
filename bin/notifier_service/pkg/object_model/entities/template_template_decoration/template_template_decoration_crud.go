//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package template_template_decoration

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionTemplateTemplateDecoration - версия структуры модели, с учётом имен и типов полей
var versionTemplateTemplateDecoration uint32

// Crud_TemplateTemplateDecoration - объект контроллер crud операций
var Crud_TemplateTemplateDecoration ICrud_TemplateTemplateDecoration

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_TemplateTemplateDecoration interface {
	Read(*TemplateTemplateDecoration) error
	Save(*TemplateTemplateDecoration) error
	Update(*TemplateTemplateDecoration) error
	Create(*TemplateTemplateDecoration) error
	ReadFromCache(ID int64) (TemplateTemplateDecoration, error)
	UpdateManyFields(*TemplateTemplateDecoration, []string) error
	Update_TemplateDecorationID(*TemplateTemplateDecoration) error
	Update_TemplateID(*TemplateTemplateDecoration) error
}

// TableName - возвращает имя таблицы в БД
func (m TemplateTemplateDecoration) TableNameDB() string {
	return "template_template_decoration"
}

// NewTemplateTemplateDecoration - возвращает новый	объект
func NewTemplateTemplateDecoration() TemplateTemplateDecoration {
	return TemplateTemplateDecoration{}
}

// AsTemplateTemplateDecoration - создаёт объект из упакованного объекта в массиве байтов
func AsTemplateTemplateDecoration(b []byte) (TemplateTemplateDecoration, error) {
	c := NewTemplateTemplateDecoration()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewTemplateTemplateDecoration(), err
	}
	return c, nil
}

// TemplateTemplateDecorationAsBytes - упаковывает объект в массив байтов
func TemplateTemplateDecorationAsBytes(m *TemplateTemplateDecoration) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m TemplateTemplateDecoration) GetStructVersion() uint32 {
	if versionTemplateTemplateDecoration == 0 {
		versionTemplateTemplateDecoration = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionTemplateTemplateDecoration
}

// GetModelFromJSON - создаёт модель из строки json
func (m *TemplateTemplateDecoration) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m TemplateTemplateDecoration) GetJSON() (string, error) {
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
func (m *TemplateTemplateDecoration) Read() error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateTemplateDecoration.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *TemplateTemplateDecoration) Save() error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateTemplateDecoration.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *TemplateTemplateDecoration) Update() error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateTemplateDecoration.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *TemplateTemplateDecoration) Create() error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateTemplateDecoration.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *TemplateTemplateDecoration) ReadFromCache(ID int64) (TemplateTemplateDecoration, error) {
	Otvet := TemplateTemplateDecoration{}
	var err error

	if Crud_TemplateTemplateDecoration == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_TemplateTemplateDecoration.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TemplateTemplateDecoration) SetCrudInterface(crud ICrud_TemplateTemplateDecoration) {
	Crud_TemplateTemplateDecoration = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *TemplateTemplateDecoration) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateTemplateDecoration.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
