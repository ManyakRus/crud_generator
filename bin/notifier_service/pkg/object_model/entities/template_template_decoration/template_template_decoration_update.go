package template_template_decoration

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/constants"
	)


// Update_TemplateDecorationID - изменяет объект в БД по ID, присваивает TemplateDecorationID
func (m *TemplateTemplateDecoration) Update_TemplateDecorationID() error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateTemplateDecoration.Update_TemplateDecorationID(m)

	return err
}

// Update_TemplateID - изменяет объект в БД по ID, присваивает TemplateID
func (m *TemplateTemplateDecoration) Update_TemplateID() error {
	if Crud_TemplateTemplateDecoration == nil {
		return constants.ErrorCrudIsNotInit
	}

	err := Crud_TemplateTemplateDecoration.Update_TemplateID(m)

	return err
}
