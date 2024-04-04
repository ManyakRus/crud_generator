package incoming_event

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_IncomingChannel - изменяет объект в БД по ID, присваивает IncomingChannel
func (m *IncomingEvent) Update_IncomingChannel() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_IncomingChannel(m)

	return err
}

// Update_IsParsed - изменяет объект в БД по ID, присваивает IsParsed
func (m *IncomingEvent) Update_IsParsed() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_IsParsed(m)

	return err
}

// Update_IsProcessed - изменяет объект в БД по ID, присваивает IsProcessed
func (m *IncomingEvent) Update_IsProcessed() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_IsProcessed(m)

	return err
}

// Update_ParseError - изменяет объект в БД по ID, присваивает ParseError
func (m *IncomingEvent) Update_ParseError() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_ParseError(m)

	return err
}

// Update_ReceivedAt - изменяет объект в БД по ID, присваивает ReceivedAt
func (m *IncomingEvent) Update_ReceivedAt() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_ReceivedAt(m)

	return err
}

// Update_SrcMsgJson - изменяет объект в БД по ID, присваивает SrcMsgJson
func (m *IncomingEvent) Update_SrcMsgJson() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_SrcMsgJson(m)

	return err
}

// Update_SysID - изменяет объект в БД по ID, присваивает SysID
func (m *IncomingEvent) Update_SysID() error {
	if Crud_IncomingEvent == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_IncomingEvent.Update_SysID(m)

	return err
}
