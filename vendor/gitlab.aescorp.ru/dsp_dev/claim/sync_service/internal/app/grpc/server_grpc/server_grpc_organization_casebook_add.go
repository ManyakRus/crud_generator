package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_casebooks"
)

// Organization_FindByInnKpp - находит организацию по ИНН
func (s *ServerGRPC) OrganizationCasebook_FindByInn(ctx context.Context, Request *grpc_proto.RequestString) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_casebooks.OrganizationCasebook{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_casebooks.OrganizationCasebook{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &organization_casebooks.OrganizationCasebook{}
	Model.INN = Request.StringFind
	err = Model.Find_ByInn()
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

// Organization_FindByOrganizationId - находит организацию по ИНН и КПП
func (s *ServerGRPC) OrganizationCasebook_FindByOrganizationId(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organization_casebooks.OrganizationCasebook{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organization_casebooks.OrganizationCasebook{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &organization_casebooks.OrganizationCasebook{}
	Model.OrganizationID = Request.Id
	err = Model.Find_ByOrganizationId()
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
