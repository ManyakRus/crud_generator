package debt_list2

// Crud_manual_DebtList2 - объект контроллер crud операций
var Crud_manual_DebtList2 ICrud_manual_DebtList2

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_DebtList2 interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m DebtList2) SetCrudManualInterface(crud ICrud_manual_DebtList2) {
	Crud_manual_DebtList2 = crud

	return
}
