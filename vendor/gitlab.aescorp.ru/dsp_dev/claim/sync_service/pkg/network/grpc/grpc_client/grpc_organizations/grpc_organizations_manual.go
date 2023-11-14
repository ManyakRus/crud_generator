package grpc_organizations

import (
	"context"
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organizations"
	"log"
	"time"
)

// FindByInnKpp - находит модель в БД по ИНН и КПП
func (crud Crud_GRPC) Find_ByInnKpp(o *organizations.Organization) error {
	// var Otvet model.Organization

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestInnKpp{}
	Request.Inn = o.INN
	Request.Kpp = o.KPP
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.Organization_FindByInnKpp(ctx, Request)
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
