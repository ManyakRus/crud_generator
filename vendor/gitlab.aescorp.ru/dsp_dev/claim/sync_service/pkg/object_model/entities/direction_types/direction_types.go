package direction_types

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// DirectionType - модель для таблицы direction_types: Направления (справочник).
type DirectionType struct {
	entities.CommonStruct
	entities.NameStruct
}
