package organizations

// crud_Organization - объект контроллер crud операций
var crud_manual_Organization ICrudManual_Organization

type ICrudManual_Organization interface {
	Find_ByInnKpp(o *Organization) error
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Organization) SetCrudManualInterface(crud ICrudManual_Organization) {
	crud_manual_Organization = crud

	return
}

// Find_ByInnKpp - находит запись по ИНН и КПП
// если передаётся пустой КПП, то ищет без учёта КПП
func (o *Organization) Find_ByInnKpp() error {
	err := crud_manual_Organization.Find_ByInnKpp(o)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
