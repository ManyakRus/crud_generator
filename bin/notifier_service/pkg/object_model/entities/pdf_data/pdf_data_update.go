package pdf_data

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Msg - изменяет объект в БД по ID, присваивает Msg
func (m *PdfDatum) Update_Msg() error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_PdfDatum.Update_Msg(m)

	return err
}

// Update_PersonalAcc - изменяет объект в БД по ID, присваивает PersonalAcc
func (m *PdfDatum) Update_PersonalAcc() error {
	if Crud_PdfDatum == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_PdfDatum.Update_PersonalAcc(m)

	return err
}
