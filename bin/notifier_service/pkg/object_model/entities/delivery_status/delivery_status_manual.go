package delivery_status

// Crud_manual_DeliveryStatus - объект контроллер crud операций
var Crud_manual_DeliveryStatus ICrud_manual_DeliveryStatus

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_DeliveryStatus interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DeliveryStatus) SetCrudManualInterface(crud ICrud_manual_DeliveryStatus) {
	Crud_manual_DeliveryStatus = crud

	return
}
