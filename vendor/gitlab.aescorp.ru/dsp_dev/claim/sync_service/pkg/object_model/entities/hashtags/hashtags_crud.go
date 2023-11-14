//File generated automatic with crud_generator app
//Do not change anything here.

package hashtags

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionHashtag - версия структуры модели, с учётом имен и типов полей
var versionHashtag uint32

// crud_Hashtag - объект контроллер crud операций
var crud_Hashtag ICrud_Hashtag

type ICrud_Hashtag interface {
	Read(*Hashtag) error
	Save(*Hashtag) error
	Update(*Hashtag) error
	Create(*Hashtag) error
	Delete(*Hashtag) error
	Restore(*Hashtag) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Hashtag) TableNameDB() string {
	return "hashtags"
}

// NewHashtag - возвращает новый	объект
func NewHashtag() Hashtag {
	return Hashtag{}
}

// AsHashtag - создаёт объект из упакованного объекта в массиве байтов
func AsHashtag(b []byte) (Hashtag, error) {
	c := NewHashtag()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewHashtag(), err
	}
	return c, nil
}

// HashtagAsBytes - упаковывает объект в массив байтов
func HashtagAsBytes(m *Hashtag) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Hashtag) GetStructVersion() uint32 {
	if versionHashtag == 0 {
		versionHashtag = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionHashtag
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Hashtag) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Hashtag) GetJSON() (string, error) {
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
func (m *Hashtag) Read() error {
	err := crud_Hashtag.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Hashtag) Save() error {
	err := crud_Hashtag.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Hashtag) Update() error {
	err := crud_Hashtag.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Hashtag) Create() error {
	err := crud_Hashtag.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Hashtag) Delete() error {
	err := crud_Hashtag.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Hashtag) Restore() error {
	err := crud_Hashtag.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Hashtag) SetCrudInterface(crud ICrud_Hashtag) {
	crud_Hashtag = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
