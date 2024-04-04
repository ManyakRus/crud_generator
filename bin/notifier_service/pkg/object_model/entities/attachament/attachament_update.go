package attachament

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Dataset - изменяет объект в БД по ID, присваивает Dataset
func (m *Attachament) Update_Dataset() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Attachament.Update_Dataset(m)

	return err
}

// Update_Filename - изменяет объект в БД по ID, присваивает Filename
func (m *Attachament) Update_Filename() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Attachament.Update_Filename(m)

	return err
}

// Update_Qrutm - изменяет объект в БД по ID, присваивает Qrutm
func (m *Attachament) Update_Qrutm() error {
	if Crud_Attachament == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Attachament.Update_Qrutm(m)

	return err
}
