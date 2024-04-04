package w_options

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *WOption) Update_Description() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_Description(m)

	return err
}

// Update_IsActiv - изменяет объект в БД по ID, присваивает IsActiv
func (m *WOption) Update_IsActiv() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_IsActiv(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *WOption) Update_Name() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_Name(m)

	return err
}

// Update_ValBool - изменяет объект в БД по ID, присваивает ValBool
func (m *WOption) Update_ValBool() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_ValBool(m)

	return err
}

// Update_ValDt - изменяет объект в БД по ID, присваивает ValDt
func (m *WOption) Update_ValDt() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_ValDt(m)

	return err
}

// Update_ValInt - изменяет объект в БД по ID, присваивает ValInt
func (m *WOption) Update_ValInt() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_ValInt(m)

	return err
}

// Update_ValStr - изменяет объект в БД по ID, присваивает ValStr
func (m *WOption) Update_ValStr() error {
	if Crud_WOption == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_WOption.Update_ValStr(m)

	return err
}
