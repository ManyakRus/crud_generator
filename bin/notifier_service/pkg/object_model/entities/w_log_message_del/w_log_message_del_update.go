package w_log_message_del

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Addr - изменяет объект в БД по ID, присваивает Addr
func (m *WLogMessageDel) Update_Addr() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Addr(m)

	return err
}

// Update_Dt - изменяет объект в БД по ID, присваивает Dt
func (m *WLogMessageDel) Update_Dt() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Dt(m)

	return err
}

// Update_Msg - изменяет объект в БД по ID, присваивает Msg
func (m *WLogMessageDel) Update_Msg() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Msg(m)

	return err
}

// Update_Pid - изменяет объект в БД по ID, присваивает Pid
func (m *WLogMessageDel) Update_Pid() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Pid(m)

	return err
}

// Update_Proc - изменяет объект в БД по ID, присваивает Proc
func (m *WLogMessageDel) Update_Proc() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Proc(m)

	return err
}

// Update_Usr - изменяет объект в БД по ID, присваивает Usr
func (m *WLogMessageDel) Update_Usr() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Usr(m)

	return err
}

// Update_Var - изменяет объект в БД по ID, присваивает Var
func (m *WLogMessageDel) Update_Var() error {
	if Crud_WLogMessageDel == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WLogMessageDel.Update_Var(m)

	return err
}
