package debt_list

// Crud_manual_DebtList - объект контроллер crud операций
var Crud_manual_DebtList ICrud_manual_DebtList

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_DebtList interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DebtList) SetCrudManualInterface(crud ICrud_manual_DebtList) {
	Crud_manual_DebtList = crud

	return
}
