package server_grpc

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
)

// Employee_FindByLogin - возвращает запись из БД по login
func (s *ServerGRPC) Employee_FindByLogin(ctx context.Context, Request *grpc_proto.RequestString) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := employees.Employee{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(employees.Employee{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &employees.Employee{}
	Model.Login = Request.StringFind
	err = Model.Find_ByLogin()
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

// Employee_FindByEMail - возвращает запись из БД по EMail
func (s *ServerGRPC) Employee_FindByEMail(ctx context.Context, Request *grpc_proto.RequestString) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := employees.Employee{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(employees.Employee{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &employees.Employee{}
	Model.Email = Request.StringFind
	err = Model.Find_ByEMail()
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

// Employee_FindByFIO - возвращает запись из БД по second_name + name + parent_name
func (s *ServerGRPC) Employee_FindByFIO(ctx context.Context, Request *grpc_proto.RequestString3) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	// проверим совпадения версии модели
	VersionServer := employees.Employee{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(employees.Employee{})
		return &Otvet, err
	}

	// запрос в БД
	Model := &employees.Employee{}
	Model.SecondName = Request.StringFind1
	Model.Name = Request.StringFind2
	Model.ParentName = Request.StringFind3
	err = Model.Find_ByFIO()
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
