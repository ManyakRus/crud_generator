package u_link_store

// Crud_manual_ULinkStore - объект контроллер crud операций
var Crud_manual_ULinkStore ICrud_manual_ULinkStore

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_ULinkStore interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ULinkStore) SetCrudManualInterface(crud ICrud_manual_ULinkStore) {
	Crud_manual_ULinkStore = crud

	return
}
