package meter_list

// Crud_manual_MeterList - объект контроллер crud операций
var Crud_manual_MeterList ICrud_manual_MeterList

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_MeterList interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m MeterList) SetCrudManualInterface(crud ICrud_manual_MeterList) {
	Crud_manual_MeterList = crud

	return
}
