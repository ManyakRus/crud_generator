package user_roles

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities"
)

// UserRole - модель для таблицы user_roles: Роли.
type UserRole struct {
	entities.CommonStruct
	entities.NameStruct
}
