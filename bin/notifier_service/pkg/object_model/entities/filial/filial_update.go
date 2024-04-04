package filial

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_AccPrefix - изменяет объект в БД по ID, присваивает AccPrefix
func (m *Filial) Update_AccPrefix() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_AccPrefix(m)

	return err
}

// Update_DivisionName - изменяет объект в БД по ID, присваивает DivisionName
func (m *Filial) Update_DivisionName() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_DivisionName(m)

	return err
}

// Update_LongFormalName - изменяет объект в БД по ID, присваивает LongFormalName
func (m *Filial) Update_LongFormalName() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_LongFormalName(m)

	return err
}

// Update_OperatorEmail - изменяет объект в БД по ID, присваивает OperatorEmail
func (m *Filial) Update_OperatorEmail() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_OperatorEmail(m)

	return err
}

// Update_RegionFormalName - изменяет объект в БД по ID, присваивает RegionFormalName
func (m *Filial) Update_RegionFormalName() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_RegionFormalName(m)

	return err
}

// Update_RegionSmallLatName - изменяет объект в БД по ID, присваивает RegionSmallLatName
func (m *Filial) Update_RegionSmallLatName() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_RegionSmallLatName(m)

	return err
}

// Update_RegionSmallRusName - изменяет объект в БД по ID, присваивает RegionSmallRusName
func (m *Filial) Update_RegionSmallRusName() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_RegionSmallRusName(m)

	return err
}

// Update_ShortFormalName - изменяет объект в БД по ID, присваивает ShortFormalName
func (m *Filial) Update_ShortFormalName() error {
	if Crud_Filial == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_Filial.Update_ShortFormalName(m)

	return err
}
