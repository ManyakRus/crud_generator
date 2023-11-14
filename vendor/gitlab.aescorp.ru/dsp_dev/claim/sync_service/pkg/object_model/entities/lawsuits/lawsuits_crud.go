//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuits

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuit - версия структуры модели, с учётом имен и типов полей
var versionLawsuit uint32

// crud_Lawsuit - объект контроллер crud операций
var crud_Lawsuit ICrud_Lawsuit

type ICrud_Lawsuit interface {
	Read(*Lawsuit) error
	Save(*Lawsuit) error
	Update(*Lawsuit) error
	Create(*Lawsuit) error
	Delete(*Lawsuit) error
	Restore(*Lawsuit) error
	Find_ByExtID(*Lawsuit) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Lawsuit) TableNameDB() string {
	return "lawsuits"
}

// LawsuitAsBytes - упаковывает объект в массив байтов
func LawsuitAsBytes(m *Lawsuit) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Lawsuit) GetStructVersion() uint32 {
	if versionLawsuit == 0 {
		versionLawsuit = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuit
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Lawsuit) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Lawsuit) GetJSON() (string, error) {
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
func (m *Lawsuit) Read() error {
	err := crud_Lawsuit.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Lawsuit) Save() error {
	err := crud_Lawsuit.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Lawsuit) Update() error {
	err := crud_Lawsuit.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Lawsuit) Create() error {
	err := crud_Lawsuit.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Lawsuit) Delete() error {
	err := crud_Lawsuit.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Lawsuit) Restore() error {
	err := crud_Lawsuit.Restore(m)

	return err
}

// Find_ByExtID - находит объект по ExtID
func (m *Lawsuit) Find_ByExtID() error {
	err := crud_Lawsuit.Find_ByExtID(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Lawsuit) SetCrudInterface(crud ICrud_Lawsuit) {
	crud_Lawsuit = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
