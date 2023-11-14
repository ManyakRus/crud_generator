//File generated automatic with crud_generator app
//Do not change anything here.

package files

// crud_manual_File - объект контроллер crud операций
var crud_manual_File ICrudManual_File

type ICrudManual_File interface {
	Find_ByFileId(f *File) error
	Find_ByFullName(f *File) error
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m File) SetCrudManualInterface(crud ICrudManual_File) {
	crud_manual_File = crud

	return
}

// Find_ByFileId - находит запись по FileID
func (f *File) Find_ByFileId() error {
	err := crud_manual_File.Find_ByFileId(f)

	return err
}

// Find_ByFull_name - находит запись по FullName
func (f *File) Find_ByFull_name() error {
	err := crud_manual_File.Find_ByFullName(f)

	return err
}
