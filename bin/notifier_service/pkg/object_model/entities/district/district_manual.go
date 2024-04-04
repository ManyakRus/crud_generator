package district

// Crud_manual_District - объект контроллер crud операций
var Crud_manual_District ICrud_manual_District

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_District interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m District) SetCrudManualInterface(crud ICrud_manual_District) {
	Crud_manual_District = crud

	return
}
