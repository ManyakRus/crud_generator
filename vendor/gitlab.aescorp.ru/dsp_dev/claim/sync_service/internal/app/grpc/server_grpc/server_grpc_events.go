//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/events"
)

// Event_Read - читает и возвращает модель из БД
func (s *ServerGRPC) Event_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := events.Event{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(events.Event{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &events.Event{}
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

// Event_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) Event_Delete(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := events.Event{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(events.Event{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &events.Event{}
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

// Event_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) Event_Restore(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := events.Event{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(events.Event{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &events.Event{}
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

// Event_Create - создаёт новую запись в БД
func (s *ServerGRPC) Event_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := events.Event{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(events.Event{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &events.Event{}
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

// Event_Update - обновляет новую запись в БД
func (s *ServerGRPC) Event_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := events.Event{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(events.Event{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &events.Event{}
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

// Event_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) Event_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := events.Event{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(events.Event{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := events.Event{}
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