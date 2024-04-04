package link_type

// Crud_manual_LinkType - объект контроллер crud операций
var Crud_manual_LinkType ICrud_manual_LinkType

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_LinkType interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LinkType) SetCrudManualInterface(crud ICrud_manual_LinkType) {
	Crud_manual_LinkType = crud

	return
}
