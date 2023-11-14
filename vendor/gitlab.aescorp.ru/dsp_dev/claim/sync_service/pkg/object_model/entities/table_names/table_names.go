package table_names

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// TableName - модель для таблицы table_names: Таблицы (Справочник). Некоторые объекты могут быть привязаны к разным таблицам (объектам). Например файлы, теги, комменты.
type TableName struct {
	entities.CommonStruct
	entities.NameStruct
}
