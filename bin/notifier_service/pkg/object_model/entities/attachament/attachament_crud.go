//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package attachament

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionAttachament - версия структуры модели, с учётом имен и типов полей
var versionAttachament uint32

// Crud_Attachament - объект контроллер crud операций
var Crud_Attachament ICrud_Attachament

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Attachament interface {
	Read(*Attachament) error
	Save(*Attachament) error
	Update(*Attachament) error
	Create(*Attachament) error
	ReadFromCache(ID int64) (Attachament, error)
	UpdateManyFields(*Attachament, []string) error
	Update_Dataset(*Attachament) error
	Update_Filename(*Attachament) error
	Update_Qrutm(*Attachament) error
}

// TableName - возвращает имя таблицы в БД
func (m Attachament) TableNameDB() string {
	return "attachament"
}

// NewAttachament - возвращает новый	объект
func NewAttachament() Attachament {
	return Attachament{}
}

// AsAttachament - создаёт объект из упакованного объекта в массиве байтов
func AsAttachament(b []byte) (Attachament, error) {
	c := NewAttachament()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewAttachament(), err
	}
	return c, nil
}

// AttachamentAsBytes - упаковывает объект в массив байтов
func AttachamentAsBytes(m *Attachament) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Attachament) GetStructVersion() uint32 {
	if versionAttachament == 0 {
		versionAttachament = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionAttachament
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Attachament) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Attachament) GetJSON() (string, error) {
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
func (m *Attachament) Read() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Attachament.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Attachament) Save() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Attachament.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Attachament) Update() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Attachament.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Attachament) Create() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Attachament.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Attachament) ReadFromCache(ID int64) (Attachament, error) {
	Otvet := Attachament{}
	var err error

	if Crud_Attachament == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Attachament.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Attachament) SetCrudInterface(crud ICrud_Attachament) {
	Crud_Attachament = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Attachament) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Attachament.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
