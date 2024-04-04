package meter_list

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_ContactInfo - изменяет объект в БД по ID, присваивает ContactInfo
func (m *MeterList) Update_ContactInfo() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_ContactInfo(m)

	return err
}

// Update_Flat - изменяет объект в БД по ID, присваивает Flat
func (m *MeterList) Update_Flat() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_Flat(m)

	return err
}

// Update_House - изменяет объект в БД по ID, присваивает House
func (m *MeterList) Update_House() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_House(m)

	return err
}

// Update_Locality - изменяет объект в БД по ID, присваивает Locality
func (m *MeterList) Update_Locality() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_Locality(m)

	return err
}

// Update_Ls - изменяет объект в БД по ID, присваивает Ls
func (m *MeterList) Update_Ls() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_Ls(m)

	return err
}

// Update_Plot - изменяет объект в БД по ID, присваивает Plot
func (m *MeterList) Update_Plot() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_Plot(m)

	return err
}

// Update_Region - изменяет объект в БД по ID, присваивает Region
func (m *MeterList) Update_Region() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_Region(m)

	return err
}

// Update_Street - изменяет объект в БД по ID, присваивает Street
func (m *MeterList) Update_Street() error {
	if Crud_MeterList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MeterList.Update_Street(m)

	return err
}
