//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package unique_link

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionUniqueLink - версия структуры модели, с учётом имен и типов полей
var versionUniqueLink uint32

// Crud_UniqueLink - объект контроллер crud операций
var Crud_UniqueLink ICrud_UniqueLink

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_UniqueLink interface {
	Read(*UniqueLink) error
	Save(*UniqueLink) error
	Update(*UniqueLink) error
	Create(*UniqueLink) error
	ReadFromCache(ID int64) (UniqueLink, error)
	UpdateManyFields(*UniqueLink, []string) error
	Update_CountRedirect(*UniqueLink) error
	Update_FirstRedirectedAt(*UniqueLink) error
	Update_IsActive(*UniqueLink) error
	Update_LastRedirectedAt(*UniqueLink) error
	Update_LinkOriginal(*UniqueLink) error
	Update_LinkTypeID(*UniqueLink) error
	Update_LinkUnique(*UniqueLink) error
	Update_PersonalAcc(*UniqueLink) error
	Update_Utm(*UniqueLink) error
}

// TableName - возвращает имя таблицы в БД
func (m UniqueLink) TableNameDB() string {
	return "unique_link"
}

// NewUniqueLink - возвращает новый	объект
func NewUniqueLink() UniqueLink {
	return UniqueLink{}
}

// AsUniqueLink - создаёт объект из упакованного объекта в массиве байтов
func AsUniqueLink(b []byte) (UniqueLink, error) {
	c := NewUniqueLink()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewUniqueLink(), err
	}
	return c, nil
}

// UniqueLinkAsBytes - упаковывает объект в массив байтов
func UniqueLinkAsBytes(m *UniqueLink) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m UniqueLink) GetStructVersion() uint32 {
	if versionUniqueLink == 0 {
		versionUniqueLink = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionUniqueLink
}

// GetModelFromJSON - создаёт модель из строки json
func (m *UniqueLink) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m UniqueLink) GetJSON() (string, error) {
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
func (m *UniqueLink) Read() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_UniqueLink.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *UniqueLink) Save() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_UniqueLink.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *UniqueLink) Update() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_UniqueLink.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *UniqueLink) Create() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_UniqueLink.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *UniqueLink) ReadFromCache(ID int64) (UniqueLink, error) {
	Otvet := UniqueLink{}
	var err error

	if Crud_UniqueLink == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_UniqueLink.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m UniqueLink) SetCrudInterface(crud ICrud_UniqueLink) {
	Crud_UniqueLink = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *UniqueLink) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
