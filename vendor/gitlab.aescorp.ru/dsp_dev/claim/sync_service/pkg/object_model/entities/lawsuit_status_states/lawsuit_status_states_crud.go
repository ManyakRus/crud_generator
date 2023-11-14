//File generated automatic with crud_generator app
//Do not change anything here.

package lawsuit_status_states

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionLawsuitStatusState - версия структуры модели, с учётом имен и типов полей
var versionLawsuitStatusState uint32

// crud_LawsuitStatusState - объект контроллер crud операций
var crud_LawsuitStatusState ICrud_LawsuitStatusState

type ICrud_LawsuitStatusState interface {
	Read(l *LawsuitStatusState) error
	Save(l *LawsuitStatusState) error
	Update(l *LawsuitStatusState) error
	Create(l *LawsuitStatusState) error
	Delete(l *LawsuitStatusState) error
	Restore(l *LawsuitStatusState) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m LawsuitStatusState) TableNameDB() string {
	return "lawsuit_status_states"
}

// NewLawsuitStatusState - возвращает новый	объект
func NewLawsuitStatusState() LawsuitStatusState {
	return LawsuitStatusState{}
}

// AsLawsuitStatusState - создаёт объект из упакованного объекта в массиве байтов
func AsLawsuitStatusState(b []byte) (LawsuitStatusState, error) {
	c := NewLawsuitStatusState()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewLawsuitStatusState(), err
	}
	return c, nil
}

// LawsuitStatusStateAsBytes - упаковывает объект в массив байтов
func LawsuitStatusStateAsBytes(m *LawsuitStatusState) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m LawsuitStatusState) GetStructVersion() uint32 {
	if versionLawsuitStatusState == 0 {
		versionLawsuitStatusState = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionLawsuitStatusState
}

// GetModelFromJSON - создаёт модель из строки json
func (m *LawsuitStatusState) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m LawsuitStatusState) GetJSON() (string, error) {
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
func (m *LawsuitStatusState) Read() error {
	err := crud_LawsuitStatusState.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *LawsuitStatusState) Save() error {
	err := crud_LawsuitStatusState.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *LawsuitStatusState) Update() error {
	err := crud_LawsuitStatusState.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *LawsuitStatusState) Create() error {
	err := crud_LawsuitStatusState.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *LawsuitStatusState) Delete() error {
	err := crud_LawsuitStatusState.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *LawsuitStatusState) Restore() error {
	err := crud_LawsuitStatusState.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitStatusState) SetCrudInterface(crud ICrud_LawsuitStatusState) {
	crud_LawsuitStatusState = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
