package nrpc_files

import (
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/files"
	"log"
)

// FindByFileId - находит модель в БД по ИНН и КПП
func (crud Crud_NRPC) Find_ByFileId(f *files.File) error {
	// var Otvet model.File

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestString{}
	Request.StringFind = f.FileID
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.File_FindByFileID(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), f)
	if err != nil {
		return err
	}

	return err
}

// FindByExtId - находит модель в БД по ext_id
func (crud Crud_NRPC) Find_ByFullName(f *files.File) error {
	// var Otvet model.File

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestString{}
	Request.StringFind = f.FullName
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.File_FindByFullName(Request)
	if err != nil {
		sError := err.Error()
		if len(sError) >= len(constants.TEXT_ERROR_MODEL_VERSION) && sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), f)
	if err != nil {
		return err
	}

	return err
}
