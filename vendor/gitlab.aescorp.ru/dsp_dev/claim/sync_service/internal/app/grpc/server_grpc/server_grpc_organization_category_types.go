//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_category_types"
)

// OrganizationCategoryType_Read - читает и возвращает модель из БД
func (s *ServerGRPC) OrganizationCategoryType_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &organization_category_types.OrganizationCategoryType{}
	Model.ID = Request.Id
	err = Model.Read()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// OrganizationCategoryType_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) OrganizationCategoryType_Delete(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &organization_category_types.OrganizationCategoryType{}
	Model.ID = Request.Id
	err = Model.Delete()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// OrganizationCategoryType_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) OrganizationCategoryType_Restore(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &organization_category_types.OrganizationCategoryType{}
	Model.ID = Request.Id
	err = Model.Restore()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// OrganizationCategoryType_Create - создаёт новую запись в БД
func (s *ServerGRPC) OrganizationCategoryType_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &organization_category_types.OrganizationCategoryType{}
	err = Model.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	// запрос в БД
	err = Model.Create()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// OrganizationCategoryType_Update - обновляет новую запись в БД
func (s *ServerGRPC) OrganizationCategoryType_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &organization_category_types.OrganizationCategoryType{}
	err = Model.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	// запрос в БД
	err = Model.Update()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// OrganizationCategoryType_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) OrganizationCategoryType_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := organization_category_types.OrganizationCategoryType{}
	err = Model.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	// запрос в БД
	err = Model.Save()
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// OrganizationCategoryType_FindByExtID - возвращает запись из БД по ext_id и connection_id
func (s *ServerGRPC) OrganizationCategoryType_FindByExtID(ctx context.Context, Request *grpc_proto.RequestExtId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим совпадения версии модели
	VersionServer := organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_category_types.OrganizationCategoryType{})
		return &Otvet, err
	}

	//запрос в БД
	Model := &organization_category_types.OrganizationCategoryType{}
	Model.ExtID = Request.ExtId
	Model.ConnectionID = Request.ConnectionId
	err = Model.Find_ByExtID()
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}
