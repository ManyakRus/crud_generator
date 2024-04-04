package channel_prod

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Code - изменяет объект в БД по ID, присваивает Code
func (m *ChannelProd) Update_Code() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ChannelProd.Update_Code(m)

	return err
}

// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *ChannelProd) Update_Description() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ChannelProd.Update_Description(m)

	return err
}

// Update_IsActive - изменяет объект в БД по ID, присваивает IsActive
func (m *ChannelProd) Update_IsActive() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ChannelProd.Update_IsActive(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *ChannelProd) Update_Name() error {
	if Crud_ChannelProd == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ChannelProd.Update_Name(m)

	return err
}
