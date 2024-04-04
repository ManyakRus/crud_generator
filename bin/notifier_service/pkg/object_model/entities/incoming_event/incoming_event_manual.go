package incoming_event

// Crud_manual_IncomingEvent - объект контроллер crud операций
var Crud_manual_IncomingEvent ICrud_manual_IncomingEvent

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_IncomingEvent interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m IncomingEvent) SetCrudManualInterface(crud ICrud_manual_IncomingEvent) {
	Crud_manual_IncomingEvent = crud

	return
}
