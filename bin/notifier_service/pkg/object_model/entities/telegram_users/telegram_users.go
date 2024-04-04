package telegram_users

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_telegram_users"
)

// TelegramUser - модель для таблицы telegram_users: Table for users chatID and phone numbers
type TelegramUser struct {
	table_telegram_users.Table_TelegramUser	
}
