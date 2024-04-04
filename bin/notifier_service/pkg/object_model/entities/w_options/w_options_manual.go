package w_options

// Crud_manual_WOption - объект контроллер crud операций
var Crud_manual_WOption ICrud_manual_WOption

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_WOption interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m WOption) SetCrudManualInterface(crud ICrud_manual_WOption) {
	Crud_manual_WOption = crud

	return
}
