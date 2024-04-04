package delivery_status

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Code - изменяет объект в БД по ID, присваивает Code
func (m *DeliveryStatus) Update_Code() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.Update_Code(m)

	return err
}

// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *DeliveryStatus) Update_Description() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.Update_Description(m)

	return err
}

// Update_FormalName - изменяет объект в БД по ID, присваивает FormalName
func (m *DeliveryStatus) Update_FormalName() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.Update_FormalName(m)

	return err
}

// Update_IsActive - изменяет объект в БД по ID, присваивает IsActive
func (m *DeliveryStatus) Update_IsActive() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.Update_IsActive(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *DeliveryStatus) Update_Name() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.Update_Name(m)

	return err
}

// Update_Note - изменяет объект в БД по ID, присваивает Note
func (m *DeliveryStatus) Update_Note() error {
	if Crud_DeliveryStatus == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_DeliveryStatus.Update_Note(m)

	return err
}
