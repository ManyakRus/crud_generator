//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package link_type

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionLinkType - версия структуры модели, с учётом имен и типов полей
var versionLinkType uint32

// Crud_LinkType - объект контроллер crud операций
var Crud_LinkType ICrud_LinkType

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_LinkType interface {
	Read(*LinkType) error
	Save(*LinkType) error
	Update(*LinkType) error
	Create(*LinkType) error
	ReadFromCache(ID int64) (LinkType, error)
	UpdateManyFields(*LinkType, []string) error
	Update_Description(*LinkType) error
	Update_Name(*LinkType) error
}

// TableName - возвращает имя таблицы в БД
func (m LinkType) TableNameDB() string {
	return "link_type"
}

// NewLinkType - возвращает новый	объект
func NewLinkType() LinkType {
	return LinkType{}
}

// AsLinkType - создаёт объект из упакованного объекта в массиве байтов
func AsLinkType(b []byte) (LinkType, error) {
	c := NewLinkType()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLinkType(), err
	}
	return c, nil
}

// LinkTypeAsBytes - упаковывает объект в массив байтов
func LinkTypeAsBytes(m *LinkType) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LinkType) GetStructVersion() uint32 {
	if versionLinkType == 0 {
		versionLinkType = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLinkType
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LinkType) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LinkType) GetJSON() (string, error) {
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
func (m *LinkType) Read() error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_LinkType.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LinkType) Save() error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_LinkType.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LinkType) Update() error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_LinkType.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LinkType) Create() error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_LinkType.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *LinkType) ReadFromCache(ID int64) (LinkType, error) {
	Otvet := LinkType{}
	var err error

	if Crud_LinkType == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_LinkType.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LinkType) SetCrudInterface(crud ICrud_LinkType) {
	Crud_LinkType = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *LinkType) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_LinkType.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
