package telegram_users_info

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Address - изменяет объект в БД по ID, присваивает Address
func (m *TelegramUsersInfo) Update_Address() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUsersInfo.Update_Address(m)

	return err
}

// Update_ChatID - изменяет объект в БД по ID, присваивает ChatID
func (m *TelegramUsersInfo) Update_ChatID() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUsersInfo.Update_ChatID(m)

	return err
}

// Update_IsChecked - изменяет объект в БД по ID, присваивает IsChecked
func (m *TelegramUsersInfo) Update_IsChecked() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUsersInfo.Update_IsChecked(m)

	return err
}

// Update_PersAcc - изменяет объект в БД по ID, присваивает PersAcc
func (m *TelegramUsersInfo) Update_PersAcc() error {
	if Crud_TelegramUsersInfo == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUsersInfo.Update_PersAcc(m)

	return err
}
