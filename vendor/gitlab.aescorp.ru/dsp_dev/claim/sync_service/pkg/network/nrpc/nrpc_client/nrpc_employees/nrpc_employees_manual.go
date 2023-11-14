package nrpc_employees

import (
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
	"log"
)

// FindByLogin - находит модель в БД по Login
func (crud Crud_NRPC) Find_ByLogin(e *employees.Employee) error {
	// var Otvet employees.Employee

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestString{}
	Request.StringFind = e.Login
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.Employee_FindByLogin(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), e)
	if err != nil {
		return err
	}

	return err
}

// FindByEMail - находит модель в БД по Login
func (crud Crud_NRPC) Find_ByEMail(e *employees.Employee) error {
	// var Otvet employees.Employee

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestString{}
	Request.StringFind = e.Email
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.Employee_FindByEMail(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), e)
	if err != nil {
		return err
	}

	return err
}

// FindByFIO - находит модель в БД по Second_name + Name + Parent_name
func (crud Crud_NRPC) Find_ByFIO(e *employees.Employee) error {
	// var Otvet employees.Employee

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestString3{}
	Request.StringFind1 = e.SecondName
	Request.StringFind2 = e.Name
	Request.StringFind3 = e.ParentName
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.Employee_FindByFIO(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), e)
	if err != nil {
		return err
	}

	return err
}
