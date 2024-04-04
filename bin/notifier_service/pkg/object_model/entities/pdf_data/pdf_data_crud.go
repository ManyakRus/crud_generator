//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pdf_data

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPdfDatum - версия структуры модели, с учётом имен и типов полей
var versionPdfDatum uint32

// Crud_PdfDatum - объект контроллер crud операций
var Crud_PdfDatum ICrud_PdfDatum

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PdfDatum interface {
	Read(*PdfDatum) error
	Save(*PdfDatum) error
	Update(*PdfDatum) error
	Create(*PdfDatum) error
	ReadFromCache(ID int64) (PdfDatum, error)
	UpdateManyFields(*PdfDatum, []string) error
	Update_Msg(*PdfDatum) error
	Update_PersonalAcc(*PdfDatum) error
}

// TableName - возвращает имя таблицы в БД
func (m PdfDatum) TableNameDB() string {
	return "pdf_data"
}

// NewPdfDatum - возвращает новый	объект
func NewPdfDatum() PdfDatum {
	return PdfDatum{}
}

// AsPdfDatum - создаёт объект из упакованного объекта в массиве байтов
func AsPdfDatum(b []byte) (PdfDatum, error) {
	c := NewPdfDatum()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPdfDatum(), err
	}
	return c, nil
}

// PdfDatumAsBytes - упаковывает объект в массив байтов
func PdfDatumAsBytes(m *PdfDatum) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PdfDatum) GetStructVersion() uint32 {
	if versionPdfDatum == 0 {
		versionPdfDatum = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPdfDatum
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PdfDatum) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PdfDatum) GetJSON() (string, error) {
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
func (m *PdfDatum) Read() error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PdfDatum.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PdfDatum) Save() error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PdfDatum.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PdfDatum) Update() error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PdfDatum.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PdfDatum) Create() error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PdfDatum.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PdfDatum) ReadFromCache(ID int64) (PdfDatum, error) {
	Otvet := PdfDatum{}
	var err error

	if Crud_PdfDatum == nil {
		return Otvet, constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PdfDatum.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PdfDatum) SetCrudInterface(crud ICrud_PdfDatum) {
	Crud_PdfDatum = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PdfDatum) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_PdfDatum.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
