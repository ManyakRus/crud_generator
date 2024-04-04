package attachament_message

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_AttachamentID - изменяет объект в БД по ID, присваивает AttachamentID
func (m *AttachamentMessage) Update_AttachamentID() error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_AttachamentMessage.Update_AttachamentID(m)

	return err
}

// Update_MessageID - изменяет объект в БД по ID, присваивает MessageID
func (m *AttachamentMessage) Update_MessageID() error {
	if Crud_AttachamentMessage == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_AttachamentMessage.Update_MessageID(m)

	return err
}
