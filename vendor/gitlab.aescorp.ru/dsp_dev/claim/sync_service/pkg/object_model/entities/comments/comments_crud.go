//File generated automatic with crud_generator app
//Do not change anything here.

package comments

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/calc_struct_version"
	"reflect"
)

// versionComment - версия структуры модели, с учётом имен и типов полей
var versionComment uint32

// crud_Comment - объект контроллер crud операций
var crud_Comment ICrud_Comment

type ICrud_Comment interface {
	Read(*Comment) error
	Save(*Comment) error
	Update(*Comment) error
	Create(*Comment) error
	Delete(*Comment) error
	Restore(*Comment) error
}

// TableName - возвращает имя таблицы в БД, нужен для gorm
func (m Comment) TableNameDB() string {
	return "comments"
}

// NewComment - возвращает новый	объект
func NewComment() Comment {
	return Comment{}
}

// AsComment - создаёт объект из упакованного объекта в массиве байтов
func AsComment(b []byte) (Comment, error) {
	c := NewComment()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewComment(), err
	}
	return c, nil
}

// CommentAsBytes - упаковывает объект в массив байтов
func CommentAsBytes(m *Comment) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Comment) GetStructVersion() uint32 {
	if versionComment == 0 {
		versionComment = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionComment
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Comment) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Comment) GetJSON() (string, error) {
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
func (m *Comment) Read() error {
	err := crud_Comment.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Comment) Save() error {
	err := crud_Comment.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Comment) Update() error {
	err := crud_Comment.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Comment) Create() error {
	err := crud_Comment.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *Comment) Delete() error {
	err := crud_Comment.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *Comment) Restore() error {
	err := crud_Comment.Restore(m)

	return err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Comment) SetCrudInterface(crud ICrud_Comment) {
	crud_Comment = crud

	return
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
