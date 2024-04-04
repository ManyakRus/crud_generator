package message

// Crud_manual_Message - объект контроллер crud операций
var Crud_manual_Message ICrud_manual_Message

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Message interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Message) SetCrudManualInterface(crud ICrud_manual_Message) {
	Crud_manual_Message = crud

	return
}
