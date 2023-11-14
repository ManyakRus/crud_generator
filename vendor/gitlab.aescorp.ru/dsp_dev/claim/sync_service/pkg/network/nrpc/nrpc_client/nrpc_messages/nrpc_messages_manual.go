package nrpc_messages

import (
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/messages"
	"log"
)

// FindBy_LawsuitID_MessageTypeID - находит сообщение по LawsuitID и MessageTypeID
func (crud Crud_NRPC) FindBy_LawsuitID_MessageTypeID(m *messages.Message) error {
	// var Otvet model.Message

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestIdId{}
	Request.Id1 = m.LawsuitID
	Request.Id2 = m.MessageTypeID
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.Message_FindBy_LawsuitID_MessageTypeID(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), m)
	if err != nil {
		return err
	}

	return err
}
