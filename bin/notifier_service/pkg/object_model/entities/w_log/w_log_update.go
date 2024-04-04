package w_log

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Addr - изменяет объект в БД по ID, присваивает Addr
func (m *WLog) Update_Addr() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_Addr(m)

	return err
}

// Update_DtTz - изменяет объект в БД по ID, присваивает DtTz
func (m *WLog) Update_DtTz() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_DtTz(m)

	return err
}

// Update_DtUtc - изменяет объект в БД по ID, присваивает DtUtc
func (m *WLog) Update_DtUtc() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_DtUtc(m)

	return err
}

// Update_Msg - изменяет объект в БД по ID, присваивает Msg
func (m *WLog) Update_Msg() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_Msg(m)

	return err
}

// Update_MsgByte - изменяет объект в БД по ID, присваивает MsgByte
func (m *WLog) Update_MsgByte() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_MsgByte(m)

	return err
}

// Update_MsgJsonb - изменяет объект в БД по ID, присваивает MsgJsonb
func (m *WLog) Update_MsgJsonb() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_MsgJsonb(m)

	return err
}

// Update_Pid - изменяет объект в БД по ID, присваивает Pid
func (m *WLog) Update_Pid() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_Pid(m)

	return err
}

// Update_Proc - изменяет объект в БД по ID, присваивает Proc
func (m *WLog) Update_Proc() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_Proc(m)

	return err
}

// Update_Usr - изменяет объект в БД по ID, присваивает Usr
func (m *WLog) Update_Usr() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_Usr(m)

	return err
}

// Update_Var - изменяет объект в БД по ID, присваивает Var
func (m *WLog) Update_Var() error {
	if Crud_WLog == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLog.Update_Var(m)

	return err
}
