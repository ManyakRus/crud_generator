// File generated automatic with crud_generator app
// Do not change anything here.
package nrpc_lawsuit_stage_types

import (
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_stage_types"
	"log"
)

// VersionModel - хранит версию структуры модели
var VersionModel uint32

// TableName - имя таблицы в БД Postgres
const TableName string = "lawsuit_stage_types"

// объект для CRUD операций через GRPC
type Crud_NRPC struct {
}

// GetVersionModel - возвращает хэш версии структуры модели
func (crud Crud_NRPC) GetVersionModel() uint32 {
	if VersionModel == 0 {
		VersionModel = lawsuit_stage_types.LawsuitStageType{}.GetStructVersion()
	}
	return VersionModel
}

// Read - возвращает модель из БД
func (crud Crud_NRPC) Read(m *lawsuit_stage_types.LawsuitStageType) error {
	// var Otvet lawsuit_stage_types.LawsuitStageType

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = int64(m.ID)
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.LawsuitStageType_Read(Request)
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

// Create - записывает новую модель в БД
func (crud Crud_NRPC) Create(m *lawsuit_stage_types.LawsuitStageType) error {
	// var Otvet lawsuit_stage_types.LawsuitStageType

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	ModelString, err := m.GetJSON()
	if err != nil {
		return err
	}
	Request := &grpc_proto.RequestModel{}
	Request.ModelString = ModelString
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.LawsuitStageType_Create(Request)
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

// Update - обновляет модель в БД
func (crud Crud_NRPC) Update(m *lawsuit_stage_types.LawsuitStageType) error {
	// var Otvet lawsuit_stage_types.LawsuitStageType

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	ModelString, err := m.GetJSON()
	if err != nil {
		return err
	}
	Request := &grpc_proto.RequestModel{}
	Request.ModelString = ModelString
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.LawsuitStageType_Update(Request)
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

// Save - обновляет (или создаёт) модель в БД
func (crud Crud_NRPC) Save(m *lawsuit_stage_types.LawsuitStageType) error {
	// var Otvet lawsuit_stage_types.LawsuitStageType

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	ModelString, err := m.GetJSON()
	if err != nil {
		return err
	}
	Request := &grpc_proto.RequestModel{}
	Request.ModelString = ModelString
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.LawsuitStageType_Save(Request)
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

// Delete - устанавливает is_deleted = true в БД
func (crud Crud_NRPC) Delete(m *lawsuit_stage_types.LawsuitStageType) error {
	// var Otvet lawsuit_stage_types.LawsuitStageType

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = int64(m.ID)
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.LawsuitStageType_Delete(Request)
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

// Restore - устанавливает is_deleted = false в БД
func (crud Crud_NRPC) Restore(m *lawsuit_stage_types.LawsuitStageType) error {
	// var Otvet lawsuit_stage_types.LawsuitStageType

	// подключение
	if nrpc_client.Client == nil {
		nrpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = int64(m.ID)
	Request.VersionModel = VersionModel

	// запрос
	Response, err := nrpc_client.Client.LawsuitStageType_Restore(Request)
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
