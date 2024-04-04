//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package u_link_store

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionULinkStore - версия структуры модели, с учётом имен и типов полей
var versionULinkStore uint32

// Crud_ULinkStore - объект контроллер crud операций
var Crud_ULinkStore ICrud_ULinkStore

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ULinkStore interface {
	Read(*ULinkStore) error
	Save(*ULinkStore) error
	Update(*ULinkStore) error
	Create(*ULinkStore) error
	ReadFromCache(ID int64) (ULinkStore, error)
	UpdateManyFields(*ULinkStore, []string) error
	Update_ULink(*ULinkStore) error
	Update_Use(*ULinkStore) error
}

// TableName - возвращает имя таблицы в БД
func (m ULinkStore) TableNameDB() string {
	return "u_link_store"
}

// NewULinkStore - возвращает новый	объект
func NewULinkStore() ULinkStore {
	return ULinkStore{}
}

// AsULinkStore - создаёт объект из упакованного объекта в массиве байтов
func AsULinkStore(b []byte) (ULinkStore, error) {
	c := NewULinkStore()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewULinkStore(), err
	}
	return c, nil
}

// ULinkStoreAsBytes - упаковывает объект в массив байтов
func ULinkStoreAsBytes(m *ULinkStore) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ULinkStore) GetStructVersion() uint32 {
	if versionULinkStore == 0 {
		versionULinkStore = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionULinkStore
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ULinkStore) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ULinkStore) GetJSON() (string, error) {
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
func (m *ULinkStore) Read() error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ULinkStore.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ULinkStore) Save() error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ULinkStore.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ULinkStore) Update() error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ULinkStore.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ULinkStore) Create() error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ULinkStore.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *ULinkStore) ReadFromCache(ID int64) (ULinkStore, error) {
	Otvet := ULinkStore{}
	var err error

	if Crud_ULinkStore == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_ULinkStore.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ULinkStore) SetCrudInterface(crud ICrud_ULinkStore) {
	Crud_ULinkStore = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *ULinkStore) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ULinkStore.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
