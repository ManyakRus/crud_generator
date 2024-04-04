package telegram_users

// Crud_manual_TelegramUser - объект контроллер crud операций
var Crud_manual_TelegramUser ICrud_manual_TelegramUser

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_TelegramUser interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TelegramUser) SetCrudManualInterface(crud ICrud_manual_TelegramUser) {
	Crud_manual_TelegramUser = crud

	return
}
