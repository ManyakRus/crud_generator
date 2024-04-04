package channel

// Crud_manual_Channel - объект контроллер crud операций
var Crud_manual_Channel ICrud_manual_Channel

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Channel interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Channel) SetCrudManualInterface(crud ICrud_manual_Channel) {
	Crud_manual_Channel = crud

	return
}
