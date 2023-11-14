package nrpc_organization_casebooks

import (
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_casebooks"
	"log"
)

// FindByInnKpp - находит модель в БД по ИНН
func (crud Crud_NRPC) Find_ByInn(o *organization_casebooks.OrganizationCasebook) error {
	// var Otvet organization_casebooks.OrganizationCasebook

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestString{}
	Request.StringFind = o.INN
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.OrganizationCasebook_FindByInn(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), o)
	if err != nil {
		return err
	}

	return err
}

// FindByOrganizationId - находит модель в БД по ИНН и КПП
func (crud Crud_NRPC) Find_ByOrganizationId(o *organization_casebooks.OrganizationCasebook) error {
	// var Otvet organization_casebooks.OrganizationCasebook

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = o.OrganizationID
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.OrganizationCasebook_FindByOrganizationId(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), o)
	if err != nil {
		return err
	}

	return err
}
