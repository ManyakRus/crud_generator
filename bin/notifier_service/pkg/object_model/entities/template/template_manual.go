package template

// Crud_manual_Template - объект контроллер crud операций
var Crud_manual_Template ICrud_manual_Template

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_Template interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Template) SetCrudManualInterface(crud ICrud_manual_Template) {
	Crud_manual_Template = crud

	return
}
