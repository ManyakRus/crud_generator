package telegram_users

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_AllowedBillNotify - изменяет объект в БД по ID, присваивает AllowedBillNotify
func (m *TelegramUser) Update_AllowedBillNotify() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_AllowedBillNotify(m)

	return err
}

// Update_AllowedDebtNotify - изменяет объект в БД по ID, присваивает AllowedDebtNotify
func (m *TelegramUser) Update_AllowedDebtNotify() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_AllowedDebtNotify(m)

	return err
}

// Update_AllowedMeterNotify - изменяет объект в БД по ID, присваивает AllowedMeterNotify
func (m *TelegramUser) Update_AllowedMeterNotify() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_AllowedMeterNotify(m)

	return err
}

// Update_AllowedMiscNotify - изменяет объект в БД по ID, присваивает AllowedMiscNotify
func (m *TelegramUser) Update_AllowedMiscNotify() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_AllowedMiscNotify(m)

	return err
}

// Update_BlockedByUser - изменяет объект в БД по ID, присваивает BlockedByUser
func (m *TelegramUser) Update_BlockedByUser() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_BlockedByUser(m)

	return err
}

// Update_ChatID - изменяет объект в БД по ID, присваивает ChatID
func (m *TelegramUser) Update_ChatID() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_ChatID(m)

	return err
}

// Update_ContactInfo - изменяет объект в БД по ID, присваивает ContactInfo
func (m *TelegramUser) Update_ContactInfo() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_ContactInfo(m)

	return err
}

// Update_DateStatusChanged - изменяет объект в БД по ID, присваивает DateStatusChanged
func (m *TelegramUser) Update_DateStatusChanged() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_DateStatusChanged(m)

	return err
}

// Update_IsTester - изменяет объект в БД по ID, присваивает IsTester
func (m *TelegramUser) Update_IsTester() error {
	if Crud_TelegramUser == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TelegramUser.Update_IsTester(m)

	return err
}
