//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/connections"
)

// Connection_Read - читает и возвращает модель из БД
func (s *ServerGRPC) Connection_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := connections.Connection{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(connections.Connection{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &connections.Connection{}
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

// Connection_Create - создаёт новую запись в БД
func (s *ServerGRPC) Connection_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := connections.Connection{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(connections.Connection{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &connections.Connection{}
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

// Connection_Update - обновляет новую запись в БД
func (s *ServerGRPC) Connection_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := connections.Connection{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(connections.Connection{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &connections.Connection{}
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

// Connection_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) Connection_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := connections.Connection{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(connections.Connection{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := connections.Connection{}
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