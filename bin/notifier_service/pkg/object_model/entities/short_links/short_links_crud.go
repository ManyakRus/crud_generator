//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package short_links

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionShortLink - версия структуры модели, с учётом имен и типов полей
var versionShortLink uint32

// Crud_ShortLink - объект контроллер crud операций
var Crud_ShortLink ICrud_ShortLink

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ShortLink interface {
	Read(*ShortLink) error
	Save(*ShortLink) error
	Update(*ShortLink) error
	Create(*ShortLink) error
	ReadFromCache(ID int64) (ShortLink, error)
	UpdateManyFields(*ShortLink, []string) error
	Update_CountRedirect(*ShortLink) error
	Update_IsActive(*ShortLink) error
	Update_LinkLong(*ShortLink) error
	Update_LinkShort(*ShortLink) error
	Update_Name(*ShortLink) error
	Update_RedirectedAt(*ShortLink) error
}

// TableName - возвращает имя таблицы в БД
func (m ShortLink) TableNameDB() string {
	return "short_links"
}

// NewShortLink - возвращает новый	объект
func NewShortLink() ShortLink {
	return ShortLink{}
}

// AsShortLink - создаёт объект из упакованного объекта в массиве байтов
func AsShortLink(b []byte) (ShortLink, error) {
	c := NewShortLink()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewShortLink(), err
	}
	return c, nil
}

// ShortLinkAsBytes - упаковывает объект в массив байтов
func ShortLinkAsBytes(m *ShortLink) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ShortLink) GetStructVersion() uint32 {
	if versionShortLink == 0 {
		versionShortLink = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionShortLink
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ShortLink) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ShortLink) GetJSON() (string, error) {
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
func (m *ShortLink) Read() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ShortLink.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ShortLink) Save() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ShortLink.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ShortLink) Update() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ShortLink.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ShortLink) Create() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ShortLink.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *ShortLink) ReadFromCache(ID int64) (ShortLink, error) {
	Otvet := ShortLink{}
	var err error

	if Crud_ShortLink == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_ShortLink.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ShortLink) SetCrudInterface(crud ICrud_ShortLink) {
	Crud_ShortLink = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *ShortLink) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
