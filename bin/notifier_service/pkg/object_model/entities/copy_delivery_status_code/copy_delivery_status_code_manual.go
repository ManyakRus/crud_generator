package copy_delivery_status_code

// Crud_manual_CopyDeliveryStatusCode - объект контроллер crud операций
var Crud_manual_CopyDeliveryStatusCode ICrud_manual_CopyDeliveryStatusCode

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_CopyDeliveryStatusCode interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m CopyDeliveryStatusCode) SetCrudManualInterface(crud ICrud_manual_CopyDeliveryStatusCode) {
	Crud_manual_CopyDeliveryStatusCode = crud

	return
}
