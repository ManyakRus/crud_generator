//File generated automatic with crud_generator app
//Do not change anything here.

package employees

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionEmployee - версия структуры модели, с учётом имен и типов полей
var versionEmployee uint32

// crud_Employee - объект контроллер crud операций
var crud_Employee ICrud_Employee

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Employee interface {
	Read(e *Employee) error
	Save(e *Employee) error
	Update(e *Employee) error
	Create(e *Employee) error
	Delete(e *Employee) error
	Restore(e *Employee) error
	Find_ByExtID(e *Employee) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Employee) TableNameDB() string {
	return "employees"
}

// NewEmployee - возвращает новый	объект
func NewEmployee() Employee {
	return Employee{}
}

// AsEmployee - создаёт объект из упакованного объекта в массиве байтов
func AsEmployee(b []byte) (Employee, error) {
	c := NewEmployee()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewEmployee(), err
	}
	return c, nil
}

// EmployeeAsBytes - упаковывает объект в массив байтов
func EmployeeAsBytes(m *Employee) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Employee) GetStructVersion() uint32 {
	if versionEmployee == 0 {
		versionEmployee = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionEmployee
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Employee) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Employee) GetJSON() (string, error) {
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
func (m *Employee) Read() error {
	err := crud_Employee.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Employee) Save() error {
	err := crud_Employee.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Employee) Update() error {
	err := crud_Employee.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Employee) Create() error {
	err := crud_Employee.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Employee) Delete() error {
	err := crud_Employee.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Employee) Restore() error {
	err := crud_Employee.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Employee) Find_ByExtID() error {
	err := crud_Employee.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Employee) SetCrudInterface(crud ICrud_Employee) {
	crud_Employee = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
