package short_links

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_short_links"
)

// ShortLink - модель для таблицы short_links: Таблица для сервиса коротких ссылок. Хранит связь между короткой ссылкой и длинной, дополнительно ведёт статистику по кол-ву переходов и времени последнего перехода.
type ShortLink struct {
	table_short_links.Table_ShortLink	
}
