package w_log

// Crud_manual_WLog - объект контроллер crud операций
var Crud_manual_WLog ICrud_manual_WLog

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_WLog interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WLog) SetCrudManualInterface(crud ICrud_manual_WLog) {
	Crud_manual_WLog = crud

	return
}
