package debt_list

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Accrual - изменяет объект в БД по ID, присваивает Accrual
func (m *DebtList) Update_Accrual() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Accrual(m)

	return err
}

// Update_ChannelCode - изменяет объект в БД по ID, присваивает ChannelCode
func (m *DebtList) Update_ChannelCode() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_ChannelCode(m)

	return err
}

// Update_Deb - изменяет объект в БД по ID, присваивает Deb
func (m *DebtList) Update_Deb() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Deb(m)

	return err
}

// Update_Email - изменяет объект в БД по ID, присваивает Email
func (m *DebtList) Update_Email() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Email(m)

	return err
}

// Update_Fio - изменяет объект в БД по ID, присваивает Fio
func (m *DebtList) Update_Fio() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Fio(m)

	return err
}

// Update_Flat - изменяет объект в БД по ID, присваивает Flat
func (m *DebtList) Update_Flat() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Flat(m)

	return err
}

// Update_House - изменяет объект в БД по ID, присваивает House
func (m *DebtList) Update_House() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_House(m)

	return err
}

// Update_KcTel - изменяет объект в БД по ID, присваивает KcTel
func (m *DebtList) Update_KcTel() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_KcTel(m)

	return err
}

// Update_Locality - изменяет объект в БД по ID, присваивает Locality
func (m *DebtList) Update_Locality() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Locality(m)

	return err
}

// Update_PersAcc - изменяет объект в БД по ID, присваивает PersAcc
func (m *DebtList) Update_PersAcc() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_PersAcc(m)

	return err
}

// Update_Plot - изменяет объект в БД по ID, присваивает Plot
func (m *DebtList) Update_Plot() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Plot(m)

	return err
}

// Update_Region - изменяет объект в БД по ID, присваивает Region
func (m *DebtList) Update_Region() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Region(m)

	return err
}

// Update_Street - изменяет объект в БД по ID, присваивает Street
func (m *DebtList) Update_Street() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Street(m)

	return err
}

// Update_Tel - изменяет объект в БД по ID, присваивает Tel
func (m *DebtList) Update_Tel() error {
	if Crud_DebtList == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DebtList.Update_Tel(m)

	return err
}
