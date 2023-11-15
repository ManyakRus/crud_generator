//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/claim_types"
)

// ClaimType_Read - читает и возвращает модель из БД
func (s *ServerGRPC) ClaimType_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := claim_types.ClaimType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(claim_types.ClaimType{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &claim_types.ClaimType{}
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

// ClaimType_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) ClaimType_Delete(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := claim_types.ClaimType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(claim_types.ClaimType{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &claim_types.ClaimType{}
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

// ClaimType_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) ClaimType_Restore(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := claim_types.ClaimType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(claim_types.ClaimType{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &claim_types.ClaimType{}
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

// ClaimType_Create - создаёт новую запись в БД
func (s *ServerGRPC) ClaimType_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := claim_types.ClaimType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(claim_types.ClaimType{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &claim_types.ClaimType{}
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

// ClaimType_Update - обновляет новую запись в БД
func (s *ServerGRPC) ClaimType_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := claim_types.ClaimType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(claim_types.ClaimType{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &claim_types.ClaimType{}
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

// ClaimType_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) ClaimType_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := claim_types.ClaimType{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(claim_types.ClaimType{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := claim_types.ClaimType{}
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