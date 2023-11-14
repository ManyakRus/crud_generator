package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_states"
)

// LawsuitStatusState_FillFromLawsuit - заполняет debt_sum такой же как в LawSuit
func (s *ServerGRPC) LawsuitStatusState_FillFromLawsuit(ctx context.Context, Request *grpc_proto.RequestIdId) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &lawsuit_status_states.LawsuitStatusState{}
	Model.LawsuitID = Request.Id1
	Model.StatusID = Request.Id2
	err = Model.Fill_from_Lawsuit(Request.Id1, Request.Id2)
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ

	return &Otvet, err

}

// LawsuitStatusState_FillFromLawsuit - заполняет debt_sum такой же как в LawSuit
func (s *ServerGRPC) LawsuitStatusState_FindDebtSum(ctx context.Context, Request *grpc_proto.RequestIdId) (*grpc_proto.ResponseFloat64, error) {
	var Otvet grpc_proto.ResponseFloat64
	var err error

	// проверим совпадения версии модели
	VersionServer := lawsuit_status_states.LawsuitStatusState{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(lawsuit_status_states.LawsuitStatusState{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &lawsuit_status_states.LawsuitStatusState{}
	Model.LawsuitID = Request.Id1
	Model.StatusID = Request.Id2
	debt_sum, err := Model.FindDebtSum(Request.Id1, Request.Id2)
	if err != nil {
		return &Otvet, err
	}

	// заполяем ответ
	Otvet.Otvet = debt_sum

	return &Otvet, err

}
