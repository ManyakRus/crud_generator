package district

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_DepartmentCode - изменяет объект в БД по ID, присваивает DepartmentCode
func (m *District) Update_DepartmentCode() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.Update_DepartmentCode(m)

	return err
}

// Update_DepartmentName - изменяет объект в БД по ID, присваивает DepartmentName
func (m *District) Update_DepartmentName() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.Update_DepartmentName(m)

	return err
}

// Update_DistrictCode - изменяет объект в БД по ID, присваивает DistrictCode
func (m *District) Update_DistrictCode() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.Update_DistrictCode(m)

	return err
}

// Update_DistrictName - изменяет объект в БД по ID, присваивает DistrictName
func (m *District) Update_DistrictName() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.Update_DistrictName(m)

	return err
}

// Update_FilialID - изменяет объект в БД по ID, присваивает FilialID
func (m *District) Update_FilialID() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.Update_FilialID(m)

	return err
}

// Update_RegionCode - изменяет объект в БД по ID, присваивает RegionCode
func (m *District) Update_RegionCode() error {
	if Crud_District == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_District.Update_RegionCode(m)

	return err
}
