package lawsuit_status_states

// crud_LawsuitStatusState - объект контроллер crud операций
var crud_manual_LawsuitStatusState ICrudManual_LawsuitStatusState

type ICrudManual_LawsuitStatusState interface {
	Fill_from_Lawsuit(Lawsuit_id int64, Status_id int64) error
	FindDebtSum(Lawsuit_id int64, Status_id int64) (float64, error)
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m LawsuitStatusState) SetCrudManualInterface(crud ICrudManual_LawsuitStatusState) {
	crud_manual_LawsuitStatusState = crud

	return
}

func (l *LawsuitStatusState) Fill_from_Lawsuit(Lawsuit_id int64, Status_id int64) error {
	err := crud_manual_LawsuitStatusState.Fill_from_Lawsuit(Lawsuit_id, Status_id)
	return err
}

func (l *LawsuitStatusState) FindDebtSum(Lawsuit_id int64, Status_id int64) (float64, error) {
	Otvet, err := crud_manual_LawsuitStatusState.FindDebtSum(Lawsuit_id, Status_id)
	return Otvet, err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
