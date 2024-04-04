package unique_link

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_CountRedirect - изменяет объект в БД по ID, присваивает CountRedirect
func (m *UniqueLink) Update_CountRedirect() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_CountRedirect(m)

	return err
}

// Update_FirstRedirectedAt - изменяет объект в БД по ID, присваивает FirstRedirectedAt
func (m *UniqueLink) Update_FirstRedirectedAt() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_FirstRedirectedAt(m)

	return err
}

// Update_IsActive - изменяет объект в БД по ID, присваивает IsActive
func (m *UniqueLink) Update_IsActive() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_IsActive(m)

	return err
}

// Update_LastRedirectedAt - изменяет объект в БД по ID, присваивает LastRedirectedAt
func (m *UniqueLink) Update_LastRedirectedAt() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_LastRedirectedAt(m)

	return err
}

// Update_LinkOriginal - изменяет объект в БД по ID, присваивает LinkOriginal
func (m *UniqueLink) Update_LinkOriginal() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_LinkOriginal(m)

	return err
}

// Update_LinkTypeID - изменяет объект в БД по ID, присваивает LinkTypeID
func (m *UniqueLink) Update_LinkTypeID() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_LinkTypeID(m)

	return err
}

// Update_LinkUnique - изменяет объект в БД по ID, присваивает LinkUnique
func (m *UniqueLink) Update_LinkUnique() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_LinkUnique(m)

	return err
}

// Update_PersonalAcc - изменяет объект в БД по ID, присваивает PersonalAcc
func (m *UniqueLink) Update_PersonalAcc() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_PersonalAcc(m)

	return err
}

// Update_Utm - изменяет объект в БД по ID, присваивает Utm
func (m *UniqueLink) Update_Utm() error {
	if Crud_UniqueLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_UniqueLink.Update_Utm(m)

	return err
}
