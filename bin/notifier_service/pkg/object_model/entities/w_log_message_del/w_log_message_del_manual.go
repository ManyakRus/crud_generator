package w_log_message_del

// Crud_manual_WLogMessageDel - объект контроллер crud операций
var Crud_manual_WLogMessageDel ICrud_manual_WLogMessageDel

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_WLogMessageDel interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WLogMessageDel) SetCrudManualInterface(crud ICrud_manual_WLogMessageDel) {
	Crud_manual_WLogMessageDel = crud

	return
}
