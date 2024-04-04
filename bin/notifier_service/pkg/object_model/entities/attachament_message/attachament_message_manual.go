package attachament_message

// Crud_manual_AttachamentMessage - объект контроллер crud операций
var Crud_manual_AttachamentMessage ICrud_manual_AttachamentMessage

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_AttachamentMessage interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m AttachamentMessage) SetCrudManualInterface(crud ICrud_manual_AttachamentMessage) {
	Crud_manual_AttachamentMessage = crud

	return
}
