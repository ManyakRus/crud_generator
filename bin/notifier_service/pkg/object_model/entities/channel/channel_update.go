package channel

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Code - изменяет объект в БД по ID, присваивает Code
func (m *Channel) Update_Code() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.Update_Code(m)

	return err
}

// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *Channel) Update_Description() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.Update_Description(m)

	return err
}

// Update_IsActive - изменяет объект в БД по ID, присваивает IsActive
func (m *Channel) Update_IsActive() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.Update_IsActive(m)

	return err
}

// Update_LatName - изменяет объект в БД по ID, присваивает LatName
func (m *Channel) Update_LatName() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.Update_LatName(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *Channel) Update_Name() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.Update_Name(m)

	return err
}

// Update_StekCode - изменяет объект в БД по ID, присваивает StekCode
func (m *Channel) Update_StekCode() error {
	if Crud_Channel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Channel.Update_StekCode(m)

	return err
}
