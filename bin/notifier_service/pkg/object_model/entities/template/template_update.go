package template

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_ChannelID - изменяет объект в БД по ID, присваивает ChannelID
func (m *Template) Update_ChannelID() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.Update_ChannelID(m)

	return err
}

// Update_Code - изменяет объект в БД по ID, присваивает Code
func (m *Template) Update_Code() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.Update_Code(m)

	return err
}

// Update_IncomingEventID - изменяет объект в БД по ID, присваивает IncomingEventID
func (m *Template) Update_IncomingEventID() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.Update_IncomingEventID(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *Template) Update_Name() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.Update_Name(m)

	return err
}

// Update_Template - изменяет объект в БД по ID, присваивает Template
func (m *Template) Update_Template() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.Update_Template(m)

	return err
}

// Update_UpdatedAt - изменяет объект в БД по ID, присваивает UpdatedAt
func (m *Template) Update_UpdatedAt() error {
	if Crud_Template == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Template.Update_UpdatedAt(m)

	return err
}
