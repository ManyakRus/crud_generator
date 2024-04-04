package mailing

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Code - изменяет объект в БД по ID, присваивает Code
func (m *Mailing) Update_Code() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_Code(m)

	return err
}

// Update_EndAt - изменяет объект в БД по ID, присваивает EndAt
func (m *Mailing) Update_EndAt() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_EndAt(m)

	return err
}

// Update_IncomingEventID - изменяет объект в БД по ID, присваивает IncomingEventID
func (m *Mailing) Update_IncomingEventID() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_IncomingEventID(m)

	return err
}

// Update_IsCanceled - изменяет объект в БД по ID, присваивает IsCanceled
func (m *Mailing) Update_IsCanceled() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_IsCanceled(m)

	return err
}

// Update_IsFinished - изменяет объект в БД по ID, присваивает IsFinished
func (m *Mailing) Update_IsFinished() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_IsFinished(m)

	return err
}

// Update_ProcessCode - изменяет объект в БД по ID, присваивает ProcessCode
func (m *Mailing) Update_ProcessCode() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_ProcessCode(m)

	return err
}

// Update_ReceivedAt - изменяет объект в БД по ID, присваивает ReceivedAt
func (m *Mailing) Update_ReceivedAt() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_ReceivedAt(m)

	return err
}

// Update_StartAt - изменяет объект в БД по ID, присваивает StartAt
func (m *Mailing) Update_StartAt() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_StartAt(m)

	return err
}

// Update_Subject - изменяет объект в БД по ID, присваивает Subject
func (m *Mailing) Update_Subject() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_Subject(m)

	return err
}

// Update_TimeZone - изменяет объект в БД по ID, присваивает TimeZone
func (m *Mailing) Update_TimeZone() error {
	if Crud_Mailing == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Mailing.Update_TimeZone(m)

	return err
}
