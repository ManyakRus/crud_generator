package pdf_data

// Crud_manual_PdfDatum - объект контроллер crud операций
var Crud_manual_PdfDatum ICrud_manual_PdfDatum

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_PdfDatum interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PdfDatum) SetCrudManualInterface(crud ICrud_manual_PdfDatum) {
	Crud_manual_PdfDatum = crud

	return
}
