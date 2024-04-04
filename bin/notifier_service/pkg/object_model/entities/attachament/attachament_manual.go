package attachament

// Crud_manual_Attachament - объект контроллер crud операций
var Crud_manual_Attachament ICrud_manual_Attachament

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Attachament interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Attachament) SetCrudManualInterface(crud ICrud_manual_Attachament) {
	Crud_manual_Attachament = crud

	return
}
