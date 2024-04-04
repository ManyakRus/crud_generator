//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package template_decoration

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionTemplateDecoration - версия структуры модели, с учётом имен и типов полей
var versionTemplateDecoration uint32

// Crud_TemplateDecoration - объект контроллер crud операций
var Crud_TemplateDecoration ICrud_TemplateDecoration

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_TemplateDecoration interface {
	Read(*TemplateDecoration) error
	Save(*TemplateDecoration) error
	Update(*TemplateDecoration) error
	Create(*TemplateDecoration) error
	ReadFromCache(ID int64) (TemplateDecoration, error)
	UpdateManyFields(*TemplateDecoration, []string) error
	Update_Data(*TemplateDecoration) error
	Update_Description(*TemplateDecoration) error
	Update_Filename(*TemplateDecoration) error
	Update_SavedOnServerAt(*TemplateDecoration) error
}

// TableName - возвращает имя таблицы в БД
func (m TemplateDecoration) TableNameDB() string {
	return "template_decoration"
}

// NewTemplateDecoration - возвращает новый	объект
func NewTemplateDecoration() TemplateDecoration {
	return TemplateDecoration{}
}

// AsTemplateDecoration - создаёт объект из упакованного объекта в массиве байтов
func AsTemplateDecoration(b []byte) (TemplateDecoration, error) {
	c := NewTemplateDecoration()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewTemplateDecoration(), err
	}
	return c, nil
}

// TemplateDecorationAsBytes - упаковывает объект в массив байтов
func TemplateDecorationAsBytes(m *TemplateDecoration) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m TemplateDecoration) GetStructVersion() uint32 {
	if versionTemplateDecoration == 0 {
		versionTemplateDecoration = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionTemplateDecoration
}

// GetModelFromJSON - создаёт модель из строки json
func (m *TemplateDecoration) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m TemplateDecoration) GetJSON() (string, error) {
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
func (m *TemplateDecoration) Read() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateDecoration.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *TemplateDecoration) Save() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateDecoration.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *TemplateDecoration) Update() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateDecoration.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *TemplateDecoration) Create() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_TemplateDecoration.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *TemplateDecoration) ReadFromCache(ID int64) (TemplateDecoration, error) {
	Otvet := TemplateDecoration{}
	var err error

	if Crud_TemplateDecoration == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_TemplateDecoration.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TemplateDecoration) SetCrudInterface(crud ICrud_TemplateDecoration) {
	Crud_TemplateDecoration = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *TemplateDecoration) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateDecoration.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
