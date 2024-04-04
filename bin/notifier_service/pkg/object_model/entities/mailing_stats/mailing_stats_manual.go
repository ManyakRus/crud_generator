package mailing_stats

// Crud_manual_MailingStat - объект контроллер crud операций
var Crud_manual_MailingStat ICrud_manual_MailingStat

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_MailingStat interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m MailingStat) SetCrudManualInterface(crud ICrud_manual_MailingStat) {
	Crud_manual_MailingStat = crud

	return
}
