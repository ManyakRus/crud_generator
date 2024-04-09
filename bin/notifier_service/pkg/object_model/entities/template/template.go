package template

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_template"
)

// Template - модель для таблицы template: Шаблоны уведомленийЦелевая БД сервиса уведомлений
type Template struct {
	table_template.Table_Template	
}
