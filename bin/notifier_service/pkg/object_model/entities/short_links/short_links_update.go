package short_links

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_CountRedirect - изменяет объект в БД по ID, присваивает CountRedirect
func (m *ShortLink) Update_CountRedirect() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.Update_CountRedirect(m)

	return err
}

// Update_IsActive - изменяет объект в БД по ID, присваивает IsActive
func (m *ShortLink) Update_IsActive() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.Update_IsActive(m)

	return err
}

// Update_LinkLong - изменяет объект в БД по ID, присваивает LinkLong
func (m *ShortLink) Update_LinkLong() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.Update_LinkLong(m)

	return err
}

// Update_LinkShort - изменяет объект в БД по ID, присваивает LinkShort
func (m *ShortLink) Update_LinkShort() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.Update_LinkShort(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *ShortLink) Update_Name() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.Update_Name(m)

	return err
}

// Update_RedirectedAt - изменяет объект в БД по ID, присваивает RedirectedAt
func (m *ShortLink) Update_RedirectedAt() error {
	if Crud_ShortLink == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ShortLink.Update_RedirectedAt(m)

	return err
}
