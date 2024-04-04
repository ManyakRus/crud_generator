package u_link_store

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_ULink - изменяет объект в БД по ID, присваивает ULink
func (m *ULinkStore) Update_ULink() error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ULinkStore.Update_ULink(m)

	return err
}

// Update_Use - изменяет объект в БД по ID, присваивает Use
func (m *ULinkStore) Update_Use() error {
	if Crud_ULinkStore == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_ULinkStore.Update_Use(m)

	return err
}
