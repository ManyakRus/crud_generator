package unique_link

// Crud_manual_UniqueLink - объект контроллер crud операций
var Crud_manual_UniqueLink ICrud_manual_UniqueLink

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_UniqueLink interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m UniqueLink) SetCrudManualInterface(crud ICrud_manual_UniqueLink) {
	Crud_manual_UniqueLink = crud

	return
}
