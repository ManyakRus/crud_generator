package short_links

// Crud_manual_ShortLink - объект контроллер crud операций
var Crud_manual_ShortLink ICrud_manual_ShortLink

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_ShortLink interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ShortLink) SetCrudManualInterface(crud ICrud_manual_ShortLink) {
	Crud_manual_ShortLink = crud

	return
}
