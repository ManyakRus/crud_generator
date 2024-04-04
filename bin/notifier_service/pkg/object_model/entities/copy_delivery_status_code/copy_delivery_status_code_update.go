package copy_delivery_status_code

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Code - изменяет объект в БД по ID, присваивает Code
func (m *CopyDeliveryStatusCode) Update_Code() error {
	if Crud_CopyDeliveryStatusCode == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_CopyDeliveryStatusCode.Update_Code(m)

	return err
}
