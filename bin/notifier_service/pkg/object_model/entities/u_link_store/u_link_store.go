package u_link_store

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_u_link_store"
)

// ULinkStore - модель для таблицы u_link_store: Склад уникальных 6-значных кодов для коротких ссылок
type ULinkStore struct {
	table_u_link_store.Table_ULinkStore	
}
