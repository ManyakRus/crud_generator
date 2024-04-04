//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package debt_list2

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionDebtList2 - версия структуры модели, с учётом имен и типов полей
var versionDebtList2 uint32

// Crud_DebtList2 - объект контроллер crud операций
var Crud_DebtList2 ICrud_DebtList2

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_DebtList2 interface {
	Read(*DebtList2) error
	Save(*DebtList2) error
	Update(*DebtList2) error
	Create(*DebtList2) error
	ReadFromCache(ID int64) (DebtList2, error)
	UpdateManyFields(*DebtList2, []string) error
	Update_Accrual(*DebtList2) error
	Update_ChannelCode(*DebtList2) error
	Update_Deb(*DebtList2) error
	Update_Email(*DebtList2) error
	Update_Fio(*DebtList2) error
	Update_Flat(*DebtList2) error
	Update_House(*DebtList2) error
	Update_KcTel(*DebtList2) error
	Update_Locality(*DebtList2) error
	Update_PersAcc(*DebtList2) error
	Update_Plot(*DebtList2) error
	Update_Region(*DebtList2) error
	Update_Street(*DebtList2) error
	Update_Tel(*DebtList2) error
}

// TableName - возвращает имя таблицы в БД
func (m DebtList2) TableNameDB() string {
	return "debt_list2"
}

// NewDebtList2 - возвращает новый	объект
func NewDebtList2() DebtList2 {
	return DebtList2{}
}

// AsDebtList2 - создаёт объект из упакованного объекта в массиве байтов
func AsDebtList2(b []byte) (DebtList2, error) {
	c := NewDebtList2()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewDebtList2(), err
	}
	return c, nil
}

// DebtList2AsBytes - упаковывает объект в массив байтов
func DebtList2AsBytes(m *DebtList2) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m DebtList2) GetStructVersion() uint32 {
	if versionDebtList2 == 0 {
		versionDebtList2 = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionDebtList2
}

// GetModelFromJSON - создаёт модель из строки json
func (m *DebtList2) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m DebtList2) GetJSON() (string, error) {
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
func (m *DebtList2) Read() error {
	if Crud_DebtList2 == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList2.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *DebtList2) Save() error {
	if Crud_DebtList2 == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList2.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *DebtList2) Update() error {
	if Crud_DebtList2 == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList2.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *DebtList2) Create() error {
	if Crud_DebtList2 == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_DebtList2.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *DebtList2) ReadFromCache(ID int64) (DebtList2, error) {
	Otvet := DebtList2{}
	var err error

	if Crud_DebtList2 == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_DebtList2.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DebtList2) SetCrudInterface(crud ICrud_DebtList2) {
	Crud_DebtList2 = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *DebtList2) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_DebtList2 == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList2.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
