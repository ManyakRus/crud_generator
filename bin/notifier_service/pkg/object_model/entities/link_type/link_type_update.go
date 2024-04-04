package link_type

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *LinkType) Update_Description() error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_LinkType.Update_Description(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *LinkType) Update_Name() error {
	if Crud_LinkType == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_LinkType.Update_Name(m)

	return err
}
