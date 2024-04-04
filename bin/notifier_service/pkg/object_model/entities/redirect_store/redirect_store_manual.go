package redirect_store

// Crud_manual_RedirectStore - объект контроллер crud операций
var Crud_manual_RedirectStore ICrud_manual_RedirectStore

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_RedirectStore interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m RedirectStore) SetCrudManualInterface(crud ICrud_manual_RedirectStore) {
	Crud_manual_RedirectStore = crud

	return
}
