package template_decoration

// Crud_manual_TemplateDecoration - объект контроллер crud операций
var Crud_manual_TemplateDecoration ICrud_manual_TemplateDecoration

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_TemplateDecoration interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TemplateDecoration) SetCrudManualInterface(crud ICrud_manual_TemplateDecoration) {
	Crud_manual_TemplateDecoration = crud

	return
}
