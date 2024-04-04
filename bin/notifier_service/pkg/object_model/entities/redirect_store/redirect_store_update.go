package redirect_store

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Utm - изменяет объект в БД по ID, присваивает Utm
func (m *RedirectStore) Update_Utm() error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_RedirectStore.Update_Utm(m)

	return err
}

// Update_Store - изменяет объект в БД по ID, присваивает Store
func (m *RedirectStore) Update_Store() error {
	if Crud_RedirectStore == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_RedirectStore.Update_Store(m)

	return err
}
