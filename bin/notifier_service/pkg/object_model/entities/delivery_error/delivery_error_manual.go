package delivery_error

// Crud_manual_DeliveryError - объект контроллер crud операций
var Crud_manual_DeliveryError ICrud_manual_DeliveryError

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_DeliveryError interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DeliveryError) SetCrudManualInterface(crud ICrud_manual_DeliveryError) {
	Crud_manual_DeliveryError = crud

	return
}
