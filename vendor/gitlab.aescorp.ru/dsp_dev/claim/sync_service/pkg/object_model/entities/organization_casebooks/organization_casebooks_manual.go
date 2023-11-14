package organization_casebooks

// crud_OrganizationCasebook - объект контроллер crud операций
var crud_manual_OrganizationCasebook ICrudManual_OrganizationCasebook

type ICrudManual_OrganizationCasebook interface {
	Find_ByInn(o *OrganizationCasebook) error
	Find_ByOrganizationId(o *OrganizationCasebook) error
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m OrganizationCasebook) SetCrudManualInterface(crud ICrudManual_OrganizationCasebook) {
	crud_manual_OrganizationCasebook = crud

	return
}

// Find_ByInnKpp - находит запись по ИНН и КПП
func (o *OrganizationCasebook) Find_ByInn() error {
	err := crud_manual_OrganizationCasebook.Find_ByInn(o)

	return err
}

// Find_ByOrganizationId - находит запись по OrganizationId
func (o *OrganizationCasebook) Find_ByOrganizationId() error {
	err := crud_manual_OrganizationCasebook.Find_ByOrganizationId(o)

	return err
}
