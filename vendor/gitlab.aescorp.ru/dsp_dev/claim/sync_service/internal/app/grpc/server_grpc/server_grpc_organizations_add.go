package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organizations"
)

// Organization_FindByInnKpp - находит организацию по ИНН и КПП
func (s *ServerGRPC) Organization_FindByInnKpp(ctx context.Context, Request *grpc_proto.RequestInnKpp) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := organizations.Organization{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(organizations.Organization{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &organizations.Organization{}
	Model.INN = Request.Inn
	Model.KPP = Request.Kpp
	err = Model.Find_ByInnKpp()
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
