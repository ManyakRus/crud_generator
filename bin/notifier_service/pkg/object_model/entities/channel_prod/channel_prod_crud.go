//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package channel_prod

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionChannelProd - версия структуры модели, с учётом имен и типов полей
var versionChannelProd uint32

// Crud_ChannelProd - объект контроллер crud операций
var Crud_ChannelProd ICrud_ChannelProd

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ChannelProd interface {
	Read(*ChannelProd) error
	Save(*ChannelProd) error
	Update(*ChannelProd) error
	Create(*ChannelProd) error
	ReadFromCache(ID int64) (ChannelProd, error)
	UpdateManyFields(*ChannelProd, []string) error
	Update_Code(*ChannelProd) error
	Update_Description(*ChannelProd) error
	Update_IsActive(*ChannelProd) error
	Update_Name(*ChannelProd) error
}

// TableName - возвращает имя таблицы в БД
func (m ChannelProd) TableNameDB() string {
	return "channel_prod"
}

// NewChannelProd - возвращает новый	объект
func NewChannelProd() ChannelProd {
	return ChannelProd{}
}

// AsChannelProd - создаёт объект из упакованного объекта в массиве байтов
func AsChannelProd(b []byte) (ChannelProd, error) {
	c := NewChannelProd()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewChannelProd(), err
	}
	return c, nil
}

// ChannelProdAsBytes - упаковывает объект в массив байтов
func ChannelProdAsBytes(m *ChannelProd) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m ChannelProd) GetStructVersion() uint32 {
	if versionChannelProd == 0 {
		versionChannelProd = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionChannelProd
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ChannelProd) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ChannelProd) GetJSON() (string, error) {
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
func (m *ChannelProd) Read() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ChannelProd.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *ChannelProd) Save() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ChannelProd.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *ChannelProd) Update() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ChannelProd.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *ChannelProd) Create() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_ChannelProd.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *ChannelProd) ReadFromCache(ID int64) (ChannelProd, error) {
	Otvet := ChannelProd{}
	var err error

	if Crud_ChannelProd == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_ChannelProd.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ChannelProd) SetCrudInterface(crud ICrud_ChannelProd) {
	Crud_ChannelProd = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *ChannelProd) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ChannelProd.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
