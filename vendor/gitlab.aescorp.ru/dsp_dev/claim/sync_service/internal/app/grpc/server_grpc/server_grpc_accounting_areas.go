//File generated automatic with crud_generator app
//Do not change anything here.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/starter/micro"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/accounting_areas"
)

// AccountingArea_Read - читает и возвращает модель из БД
func (s *ServerGRPC) AccountingArea_Read(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &accounting_areas.AccountingArea{}
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

// AccountingArea_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) AccountingArea_Delete(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &accounting_areas.AccountingArea{}
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

// AccountingArea_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) AccountingArea_Restore(ctx context.Context, Request *grpc_proto.RequestId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &accounting_areas.AccountingArea{}
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

// AccountingArea_Create - создаёт новую запись в БД
func (s *ServerGRPC) AccountingArea_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &accounting_areas.AccountingArea{}
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

// AccountingArea_Update - обновляет новую запись в БД
func (s *ServerGRPC) AccountingArea_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := &accounting_areas.AccountingArea{}
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

// AccountingArea_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) AccountingArea_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	// получим модель из строки JSON
	Model := accounting_areas.AccountingArea{}
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

// AccountingArea_FindByExtID - возвращает запись из БД по ext_id и connection_id
func (s *ServerGRPC) AccountingArea_FindByExtID(ctx context.Context, Request *grpc_proto.RequestExtId) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := accounting_areas.AccountingArea{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(accounting_areas.AccountingArea{})
		return &Otvet, err
	}

	//запрос в БД
	Model := &accounting_areas.AccountingArea{}
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
