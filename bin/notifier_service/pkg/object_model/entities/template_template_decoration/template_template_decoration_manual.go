package template_template_decoration

// Crud_manual_TemplateTemplateDecoration - объект контроллер crud операций
var Crud_manual_TemplateTemplateDecoration ICrud_manual_TemplateTemplateDecoration

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_TemplateTemplateDecoration interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m TemplateTemplateDecoration) SetCrudManualInterface(crud ICrud_manual_TemplateTemplateDecoration) {
	Crud_manual_TemplateTemplateDecoration = crud

	return
}
