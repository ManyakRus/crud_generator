//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package debt_list

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionDebtList - версия структуры модели, с учётом имен и типов полей
var versionDebtList uint32

// Crud_DebtList - объект контроллер crud операций
var Crud_DebtList ICrud_DebtList

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_DebtList interface {
	Read(*DebtList) error
	Save(*DebtList) error
	Update(*DebtList) error
	Create(*DebtList) error
	ReadFromCache(ID int64) (DebtList, error)
	UpdateManyFields(*DebtList, []string) error
	Update_Accrual(*DebtList) error
	Update_ChannelCode(*DebtList) error
	Update_Deb(*DebtList) error
	Update_Email(*DebtList) error
	Update_Fio(*DebtList) error
	Update_Flat(*DebtList) error
	Update_House(*DebtList) error
	Update_KcTel(*DebtList) error
	Update_Locality(*DebtList) error
	Update_PersAcc(*DebtList) error
	Update_Plot(*DebtList) error
	Update_Region(*DebtList) error
	Update_Street(*DebtList) error
	Update_Tel(*DebtList) error
}

// TableName - возвращает имя таблицы в БД
func (m DebtList) TableNameDB() string {
	return "debt_list"
}

// NewDebtList - возвращает новый	объект
func NewDebtList() DebtList {
	return DebtList{}
}

// AsDebtList - создаёт объект из упакованного объекта в массиве байтов
func AsDebtList(b []byte) (DebtList, error) {
	c := NewDebtList()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDebtList(), err
	}
	return c, nil
}

// DebtListAsBytes - упаковывает объект в массив байтов
func DebtListAsBytes(m *DebtList) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DebtList) GetStructVersion() uint32 {
	if versionDebtList == 0 {
		versionDebtList = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDebtList
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DebtList) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DebtList) GetJSON() (string, error) {
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
func (m *DebtList) Read() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DebtList) Save() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DebtList) Update() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DebtList) Create() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *DebtList) ReadFromCache(ID int64) (DebtList, error) {
	Otvet := DebtList{}
	var err error

	if Crud_DebtList == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_DebtList.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DebtList) SetCrudInterface(crud ICrud_DebtList) {
	Crud_DebtList = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *DebtList) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
