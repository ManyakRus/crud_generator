package message

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Attachments - изменяет объект в БД по ID, присваивает Attachments
func (m *Message) Update_Attachments() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_Attachments(m)

	return err
}

// Update_CanceledMsgID - изменяет объект в БД по ID, присваивает CanceledMsgID
func (m *Message) Update_CanceledMsgID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_CanceledMsgID(m)

	return err
}

// Update_ChannelID - изменяет объект в БД по ID, присваивает ChannelID
func (m *Message) Update_ChannelID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_ChannelID(m)

	return err
}

// Update_ContactAddress - изменяет объект в БД по ID, присваивает ContactAddress
func (m *Message) Update_ContactAddress() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_ContactAddress(m)

	return err
}

// Update_ContactInfo - изменяет объект в БД по ID, присваивает ContactInfo
func (m *Message) Update_ContactInfo() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_ContactInfo(m)

	return err
}

// Update_CountTryCheck - изменяет объект в БД по ID, присваивает CountTryCheck
func (m *Message) Update_CountTryCheck() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_CountTryCheck(m)

	return err
}

// Update_CountTrySent - изменяет объект в БД по ID, присваивает CountTrySent
func (m *Message) Update_CountTrySent() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_CountTrySent(m)

	return err
}

// Update_DeliveredAt - изменяет объект в БД по ID, присваивает DeliveredAt
func (m *Message) Update_DeliveredAt() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_DeliveredAt(m)

	return err
}

// Update_DeliveryStatusID - изменяет объект в БД по ID, присваивает DeliveryStatusID
func (m *Message) Update_DeliveryStatusID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_DeliveryStatusID(m)

	return err
}

// Update_ExternalID - изменяет объект в БД по ID, присваивает ExternalID
func (m *Message) Update_ExternalID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_ExternalID(m)

	return err
}

// Update_FilialID - изменяет объект в БД по ID, присваивает FilialID
func (m *Message) Update_FilialID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_FilialID(m)

	return err
}

// Update_GateName - изменяет объект в БД по ID, присваивает GateName
func (m *Message) Update_GateName() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_GateName(m)

	return err
}

// Update_IncomingEventID - изменяет объект в БД по ID, присваивает IncomingEventID
func (m *Message) Update_IncomingEventID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_IncomingEventID(m)

	return err
}

// Update_IsChecked - изменяет объект в БД по ID, присваивает IsChecked
func (m *Message) Update_IsChecked() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_IsChecked(m)

	return err
}

// Update_IsOrganisationAcc - изменяет объект в БД по ID, присваивает IsOrganisationAcc
func (m *Message) Update_IsOrganisationAcc() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_IsOrganisationAcc(m)

	return err
}

// Update_IsSent - изменяет объект в БД по ID, присваивает IsSent
func (m *Message) Update_IsSent() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_IsSent(m)

	return err
}

// Update_MailingID - изменяет объект в БД по ID, присваивает MailingID
func (m *Message) Update_MailingID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_MailingID(m)

	return err
}

// Update_Msg - изменяет объект в БД по ID, присваивает Msg
func (m *Message) Update_Msg() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_Msg(m)

	return err
}

// Update_PersonalAcc - изменяет объект в БД по ID, присваивает PersonalAcc
func (m *Message) Update_PersonalAcc() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_PersonalAcc(m)

	return err
}

// Update_SendingStatus - изменяет объект в БД по ID, присваивает SendingStatus
func (m *Message) Update_SendingStatus() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_SendingStatus(m)

	return err
}

// Update_Seq - изменяет объект в БД по ID, присваивает Seq
func (m *Message) Update_Seq() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_Seq(m)

	return err
}

// Update_TryCheckAt - изменяет объект в БД по ID, присваивает TryCheckAt
func (m *Message) Update_TryCheckAt() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_TryCheckAt(m)

	return err
}

// Update_TrySendAt - изменяет объект в БД по ID, присваивает TrySendAt
func (m *Message) Update_TrySendAt() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_TrySendAt(m)

	return err
}

// Update_UserAddress - изменяет объект в БД по ID, присваивает UserAddress
func (m *Message) Update_UserAddress() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_UserAddress(m)

	return err
}

// Update_UserID - изменяет объект в БД по ID, присваивает UserID
func (m *Message) Update_UserID() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_UserID(m)

	return err
}

// Update_Utm - изменяет объект в БД по ID, присваивает Utm
func (m *Message) Update_Utm() error {
	if Crud_Message == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Message.Update_Utm(m)

	return err
}
