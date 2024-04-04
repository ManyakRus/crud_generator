package template_decoration

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_Data - изменяет объект в БД по ID, присваивает Data
func (m *TemplateDecoration) Update_Data() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateDecoration.Update_Data(m)

	return err
}

// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *TemplateDecoration) Update_Description() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateDecoration.Update_Description(m)

	return err
}

// Update_Filename - изменяет объект в БД по ID, присваивает Filename
func (m *TemplateDecoration) Update_Filename() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateDecoration.Update_Filename(m)

	return err
}

// Update_SavedOnServerAt - изменяет объект в БД по ID, присваивает SavedOnServerAt
func (m *TemplateDecoration) Update_SavedOnServerAt() error {
	if Crud_TemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateDecoration.Update_SavedOnServerAt(m)

	return err
}
