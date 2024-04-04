package redirect_store

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_redirect_store"
)

// RedirectStore - модель для таблицы redirect_store: Таблица нужна для отслеживания переходов по магазинам мобильных приложений. Сервис при переходе в конкретный магазин добавляет запись в таблицу.
type RedirectStore struct {
	table_redirect_store.Table_RedirectStore	
}
