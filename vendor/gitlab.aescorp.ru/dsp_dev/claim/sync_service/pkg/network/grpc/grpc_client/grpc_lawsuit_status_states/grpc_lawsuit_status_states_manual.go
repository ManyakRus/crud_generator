package grpc_lawsuit_status_states

import (
	"context"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"log"
	"time"
)

func (crud Crud_GRPC) Fill_from_Lawsuit(Lawsuit_id int64, Status_id int64) error {
	//подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	//подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestIdId{}
	Request.Id1 = Lawsuit_id
	Request.Id2 = Status_id
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	//запрос
	_, err := grpc_client.Client.LawsuitStatusState_FillFromLawsuit(ctx, Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	//ответ
	//sModel := Response.ModelString
	//err = json.Unmarshal([]byte(sModel), &Otvet)
	//if err != nil {
	//	return Otvet, err
	//}

	return err

}

func (crud Crud_GRPC) FindDebtSum(Lawsuit_id int64, Status_id int64) (float64, error) {
	var Otvet float64

	//подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	//подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestIdId{}
	Request.Id1 = Lawsuit_id
	Request.Id2 = Status_id
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	//запрос
	ResponseFloat64, err := grpc_client.Client.LawsuitStatusState_FindDebtSum(ctx, Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return Otvet, err
	}

	//ответ
	//sModel := Response.ModelString
	//err = json.Unmarshal([]byte(sModel), &Otvet)
	//if err != nil {
	//	return Otvet, err
	//}

	Otvet = ResponseFloat64.Otvet

	return Otvet, err

}
