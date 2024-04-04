package statistic

// Crud_manual_Statistic - объект контроллер crud операций
var Crud_manual_Statistic ICrud_manual_Statistic

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Statistic interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Statistic) SetCrudManualInterface(crud ICrud_manual_Statistic) {
	Crud_manual_Statistic = crud

	return
}
