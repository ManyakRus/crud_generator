//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package redirect_store

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionRedirectStore - версия структуры модели, с учётом имен и типов полей
var versionRedirectStore uint32

// Crud_RedirectStore - объект контроллер crud операций
var Crud_RedirectStore ICrud_RedirectStore

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_RedirectStore interface {
	Read(*RedirectStore) error
	Save(*RedirectStore) error
	Update(*RedirectStore) error
	Create(*RedirectStore) error
	ReadFromCache(ID int64) (RedirectStore, error)
	UpdateManyFields(*RedirectStore, []string) error
	Update_Utm(*RedirectStore) error
	Update_Store(*RedirectStore) error
}

// TableName - возвращает имя таблицы в БД
func (m RedirectStore) TableNameDB() string {
	return "redirect_store"
}

// NewRedirectStore - возвращает новый	объект
func NewRedirectStore() RedirectStore {
	return RedirectStore{}
}

// AsRedirectStore - создаёт объект из упакованного объекта в массиве байтов
func AsRedirectStore(b []byte) (RedirectStore, error) {
	c := NewRedirectStore()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewRedirectStore(), err
	}
	return c, nil
}

// RedirectStoreAsBytes - упаковывает объект в массив байтов
func RedirectStoreAsBytes(m *RedirectStore) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m RedirectStore) GetStructVersion() uint32 {
	if versionRedirectStore == 0 {
		versionRedirectStore = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionRedirectStore
}

// GetModelFromJSON - создаёт модель из строки json
func (m *RedirectStore) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m RedirectStore) GetJSON() (string, error) {
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
func (m *RedirectStore) Read() error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_RedirectStore.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *RedirectStore) Save() error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_RedirectStore.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *RedirectStore) Update() error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_RedirectStore.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *RedirectStore) Create() error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_RedirectStore.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *RedirectStore) ReadFromCache(ID int64) (RedirectStore, error) {
	Otvet := RedirectStore{}
	var err error

	if Crud_RedirectStore == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_RedirectStore.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m RedirectStore) SetCrudInterface(crud ICrud_RedirectStore) {
	Crud_RedirectStore = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *RedirectStore) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_RedirectStore.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
