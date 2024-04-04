package statistic

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_FilialIds - изменяет объект в БД по ID, присваивает FilialIds
func (m *Statistic) Update_FilialIds() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_FilialIds(m)

	return err
}

// Update_LastUpdate - изменяет объект в БД по ID, присваивает LastUpdate
func (m *Statistic) Update_LastUpdate() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_LastUpdate(m)

	return err
}

// Update_MailingChannelID - изменяет объект в БД по ID, присваивает MailingChannelID
func (m *Statistic) Update_MailingChannelID() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MailingChannelID(m)

	return err
}

// Update_MailingCode - изменяет объект в БД по ID, присваивает MailingCode
func (m *Statistic) Update_MailingCode() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MailingCode(m)

	return err
}

// Update_MailingEnd - изменяет объект в БД по ID, присваивает MailingEnd
func (m *Statistic) Update_MailingEnd() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MailingEnd(m)

	return err
}

// Update_MailingID - изменяет объект в БД по ID, присваивает MailingID
func (m *Statistic) Update_MailingID() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MailingID(m)

	return err
}

// Update_MailingStart - изменяет объект в БД по ID, присваивает MailingStart
func (m *Statistic) Update_MailingStart() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MailingStart(m)

	return err
}

// Update_MailingStatus - изменяет объект в БД по ID, присваивает MailingStatus
func (m *Statistic) Update_MailingStatus() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MailingStatus(m)

	return err
}

// Update_MessagesDelivered - изменяет объект в БД по ID, присваивает MessagesDelivered
func (m *Statistic) Update_MessagesDelivered() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MessagesDelivered(m)

	return err
}

// Update_MessagesSent - изменяет объект в БД по ID, присваивает MessagesSent
func (m *Statistic) Update_MessagesSent() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MessagesSent(m)

	return err
}

// Update_MessagesTotal - изменяет объект в БД по ID, присваивает MessagesTotal
func (m *Statistic) Update_MessagesTotal() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MessagesTotal(m)

	return err
}

// Update_MessagesUndelivered - изменяет объект в БД по ID, присваивает MessagesUndelivered
func (m *Statistic) Update_MessagesUndelivered() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MessagesUndelivered(m)

	return err
}

// Update_MessagesUnsent - изменяет объект в БД по ID, присваивает MessagesUnsent
func (m *Statistic) Update_MessagesUnsent() error {
	if Crud_Statistic == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Statistic.Update_MessagesUnsent(m)

	return err
}
