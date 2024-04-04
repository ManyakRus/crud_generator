package mailing_stats

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_ErrorsCount - изменяет объект в БД по ID, присваивает ErrorsCount
func (m *MailingStat) Update_ErrorsCount() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MailingStat.Update_ErrorsCount(m)

	return err
}

// Update_MailingID - изменяет объект в БД по ID, присваивает MailingID
func (m *MailingStat) Update_MailingID() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MailingStat.Update_MailingID(m)

	return err
}

// Update_MessagesTotalCount - изменяет объект в БД по ID, присваивает MessagesTotalCount
func (m *MailingStat) Update_MessagesTotalCount() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MailingStat.Update_MessagesTotalCount(m)

	return err
}

// Update_RedirectCount - изменяет объект в БД по ID, присваивает RedirectCount
func (m *MailingStat) Update_RedirectCount() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MailingStat.Update_RedirectCount(m)

	return err
}

// Update_SentMessagesCount - изменяет объект в БД по ID, присваивает SentMessagesCount
func (m *MailingStat) Update_SentMessagesCount() error {
	if Crud_MailingStat == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_MailingStat.Update_SentMessagesCount(m)

	return err
}
