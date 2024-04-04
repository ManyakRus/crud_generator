package telegram_users_info

// Crud_manual_TelegramUsersInfo - объект контроллер crud операций
var Crud_manual_TelegramUsersInfo ICrud_manual_TelegramUsersInfo

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_TelegramUsersInfo interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TelegramUsersInfo) SetCrudManualInterface(crud ICrud_manual_TelegramUsersInfo) {
	Crud_manual_TelegramUsersInfo = crud

	return
}
