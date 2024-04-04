package mailing

// Crud_manual_Mailing - объект контроллер crud операций
var Crud_manual_Mailing ICrud_manual_Mailing

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Mailing interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Mailing) SetCrudManualInterface(crud ICrud_manual_Mailing) {
	Crud_manual_Mailing = crud

	return
}
