//File generated automatic with crud_generator app
//Do not change anything here.

package hashtag_links

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionHashtagLink - версия структуры модели, с учётом имен и типов полей
var versionHashtagLink uint32

// crud_HashtagLink - объект контроллер crud операций
var crud_HashtagLink ICrud_HashtagLink

type ICrud_HashtagLink interface {
	Read(*HashtagLink) error
	Save(*HashtagLink) error
	Update(*HashtagLink) error
	Create(*HashtagLink) error
	Delete(*HashtagLink) error
	Restore(*HashtagLink) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m HashtagLink) TableNameDB() string {
	return "hashtag_links"
}

// NewHashtagLink - возвращает новый	объект
func NewHashtagLink() HashtagLink {
	return HashtagLink{}
}

// AsHashtagLink - создаёт объект из упакованного объекта в массиве байтов
func AsHashtagLink(b []byte) (HashtagLink, error) {
	c := NewHashtagLink()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewHashtagLink(), err
	}
	return c, nil
}

// HashtagLinkAsBytes - упаковывает объект в массив байтов
func HashtagLinkAsBytes(m *HashtagLink) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m HashtagLink) GetStructVersion() uint32 {
	if versionHashtagLink == 0 {
		versionHashtagLink = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionHashtagLink
}

// GetModelFromJSON - создаёт модель из строки json
func (m *HashtagLink) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m HashtagLink) GetJSON() (string, error) {
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
func (m *HashtagLink) Read() error {
	err := crud_HashtagLink.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *HashtagLink) Save() error {
	err := crud_HashtagLink.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *HashtagLink) Update() error {
	err := crud_HashtagLink.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *HashtagLink) Create() error {
	err := crud_HashtagLink.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *HashtagLink) Delete() error {
	err := crud_HashtagLink.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *HashtagLink) Restore() error {
	err := crud_HashtagLink.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m HashtagLink) SetCrudInterface(crud ICrud_HashtagLink) {
	crud_HashtagLink = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
