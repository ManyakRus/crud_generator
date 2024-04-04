package pdf_data

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/tables/table_pdf_data"
)

// PdfDatum - модель для таблицы pdf_data: Данные из pdf файлов квитанций
type PdfDatum struct {
	table_pdf_data.Table_PdfDatum	
}
