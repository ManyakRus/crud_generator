package nrpc_lawsuit_status_states

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/constants"
	"log"
)

func (crud Crud_NRPC) Fill_from_Lawsuit(Lawsuit_id int64, Status_id int64) error {
	//подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	//подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestIdId{}
	Request.Id1 = Lawsuit_id
	Request.Id2 = Status_id
	Request.VersionModel = VersionModel

	//запрос
	_, err := nrpc_client.Client.LawsuitStatusState_FillFromLawsuit(Request)
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

func (crud Crud_NRPC) FindDebtSum(Lawsuit_id int64, Status_id int64) (float64, error) {
	var Otvet float64

	//подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	//подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestIdId{}
	Request.Id1 = Lawsuit_id
	Request.Id2 = Status_id
	Request.VersionModel = VersionModel

	//запрос
	ResponseFloat64, err := nrpc_client.Client.LawsuitStatusState_FindDebtSum(Request)
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
