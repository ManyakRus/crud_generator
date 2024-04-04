package filial

// Crud_manual_Filial - объект контроллер crud операций
var Crud_manual_Filial ICrud_manual_Filial

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Filial interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Filial) SetCrudManualInterface(crud ICrud_manual_Filial) {
	Crud_manual_Filial = crud

	return
}
