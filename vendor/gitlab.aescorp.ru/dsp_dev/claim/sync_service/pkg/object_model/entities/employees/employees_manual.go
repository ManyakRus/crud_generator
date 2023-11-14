//File generated automatic with crud_generator app
//Do not change anything here.

package employees

// crud_manual_Employee - объект контроллер crud операций
var crud_manual_Employee ICrudManual_Employee

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrudManual_Employee interface {
	Find_ByLogin(e *Employee) error
	Find_ByEMail(e *Employee) error
	Find_ByFIO(e *Employee) error
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Employee) SetCrudManualInterface(crud ICrudManual_Employee) {
	crud_manual_Employee = crud

	return
}

// Find_ByEMail - находит объект по email
func (e *Employee) Find_ByEMail() error {
	err := crud_manual_Employee.Find_ByEMail(e)

	return err
}

// Find_ByLogin - находит объект по Login
func (e *Employee) Find_ByLogin() error {
	err := crud_manual_Employee.Find_ByLogin(e)

	return err
}

// Find_ByFIO - находит объект по ФИО
func (e *Employee) Find_ByFIO() error {
	err := crud_manual_Employee.Find_ByFIO(e)

	return err
}
