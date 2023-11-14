//File generated automatic with crud_generator app
//Do not change anything here.

package user_roles

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionUserRole - версия структуры модели, с учётом имен и типов полей
var versionUserRole uint32

// crud_UserRole - объект контроллер crud операций
var crud_UserRole ICrud_UserRole

type ICrud_UserRole interface {
	Read(*UserRole) error
	Save(*UserRole) error
	Update(*UserRole) error
	Create(*UserRole) error
	Delete(*UserRole) error
	Restore(*UserRole) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m UserRole) TableNameDB() string {
	return "user_roles"
}

// NewUserRole - возвращает новый	объект
func NewUserRole() UserRole {
	return UserRole{}
}

// AsUserRole - создаёт объект из упакованного объекта в массиве байтов
func AsUserRole(b []byte) (UserRole, error) {
	c := NewUserRole()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewUserRole(), err
	}
	return c, nil
}

// UserRoleAsBytes - упаковывает объект в массив байтов
func UserRoleAsBytes(m *UserRole) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m UserRole) GetStructVersion() uint32 {
	if versionUserRole == 0 {
		versionUserRole = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionUserRole
}

// GetModelFromJSON - создаёт модель из строки json
func (m *UserRole) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m UserRole) GetJSON() (string, error) {
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
func (m *UserRole) Read() error {
	err := crud_UserRole.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *UserRole) Save() error {
	err := crud_UserRole.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *UserRole) Update() error {
	err := crud_UserRole.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *UserRole) Create() error {
	err := crud_UserRole.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *UserRole) Delete() error {
	err := crud_UserRole.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *UserRole) Restore() error {
	err := crud_UserRole.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m UserRole) SetCrudInterface(crud ICrud_UserRole) {
	crud_UserRole = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
