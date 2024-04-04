package delivery_error

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_DeliveryStatusID - изменяет объект в БД по ID, присваивает DeliveryStatusID
func (m *DeliveryError) Update_DeliveryStatusID() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryError.Update_DeliveryStatusID(m)

	return err
}

// Update_IsActive - изменяет объект в БД по ID, присваивает IsActive
func (m *DeliveryError) Update_IsActive() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryError.Update_IsActive(m)

	return err
}

// Update_TextError - изменяет объект в БД по ID, присваивает TextError
func (m *DeliveryError) Update_TextError() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryError.Update_TextError(m)

	return err
}

// Update_UpdatedAt - изменяет объект в БД по ID, присваивает UpdatedAt
func (m *DeliveryError) Update_UpdatedAt() error {
	if Crud_DeliveryError == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryError.Update_UpdatedAt(m)

	return err
}
