// File generated automatic with crud_generator app
// Do not change anything here.
package grpc_contract_black_items

import (
	"context"
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/contract_black_items"
	"log"
	"time"
)

// VersionModel - хранит версию структуры модели
var VersionModel uint32

// TableName - имя таблицы в БД Postgres
const TableName string = "contract_black_items"

// объект для CRUD операций через GRPC
type Crud_GRPC struct {
}

// GetVersionModel - возвращает хэш версии структуры модели
func (crud Crud_GRPC) GetVersionModel() uint32 {
	if VersionModel == 0 {
		VersionModel = contract_black_items.ContractBlackItem{}.GetStructVersion()
	}
	return VersionModel
}

// Read - возвращает модель из БД
func (crud Crud_GRPC) Read(m *contract_black_items.ContractBlackItem) error {
	// var Otvet contract_black_items.ContractBlackItem

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = int64(m.ID)
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.ContractBlackItem_Read(ctx, Request)
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
func (crud Crud_GRPC) Create(m *contract_black_items.ContractBlackItem) error {
	// var Otvet contract_black_items.ContractBlackItem

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
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

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.ContractBlackItem_Create(ctx, Request)
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
func (crud Crud_GRPC) Update(m *contract_black_items.ContractBlackItem) error {
	// var Otvet contract_black_items.ContractBlackItem

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
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

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.ContractBlackItem_Update(ctx, Request)
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
func (crud Crud_GRPC) Save(m *contract_black_items.ContractBlackItem) error {
	// var Otvet contract_black_items.ContractBlackItem

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
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

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.ContractBlackItem_Save(ctx, Request)
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
func (crud Crud_GRPC) Delete(m *contract_black_items.ContractBlackItem) error {
	// var Otvet contract_black_items.ContractBlackItem

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = int64(m.ID)
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.ContractBlackItem_Delete(ctx, Request)
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
func (crud Crud_GRPC) Restore(m *contract_black_items.ContractBlackItem) error {
	// var Otvet contract_black_items.ContractBlackItem

	// подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestId{}
	Request.Id = int64(m.ID)
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	// запрос
	Response, err := grpc_client.Client.ContractBlackItem_Restore(ctx, Request)
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
