// File generated automatic with crud_generator app
// Do not change anything here.
package grpc_organization_category_types

import (
	"context"
	"encoding/json"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_category_types"
	"log"
	"time"
)

// VersionModel - хранит версию структуры модели
var VersionModel uint32

// TableName - имя таблицы в БД Postgres
const TableName string = "organization_category_types"

// объект для CRUD операций через GRPC
type Crud_GRPC struct {
}

// GetVersionModel - возвращает хэш версии структуры модели
func (crud Crud_GRPC) GetVersionModel() uint32 {
	if VersionModel == 0 {
		VersionModel = organization_category_types.OrganizationCategoryType{}.GetStructVersion()
	}
	return VersionModel
}

// Read - возвращает модель из БД
func (crud Crud_GRPC) Read(m *organization_category_types.OrganizationCategoryType) error {
	// var Otvet organization_category_types.OrganizationCategoryType

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
	Response, err := grpc_client.Client.OrganizationCategoryType_Read(ctx, Request)
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
func (crud Crud_GRPC) Create(m *organization_category_types.OrganizationCategoryType) error {
	// var Otvet organization_category_types.OrganizationCategoryType

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
	Response, err := grpc_client.Client.OrganizationCategoryType_Create(ctx, Request)
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
func (crud Crud_GRPC) Update(m *organization_category_types.OrganizationCategoryType) error {
	// var Otvet organization_category_types.OrganizationCategoryType

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
	Response, err := grpc_client.Client.OrganizationCategoryType_Update(ctx, Request)
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
func (crud Crud_GRPC) Save(m *organization_category_types.OrganizationCategoryType) error {
	// var Otvet organization_category_types.OrganizationCategoryType

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
	Response, err := grpc_client.Client.OrganizationCategoryType_Save(ctx, Request)
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
func (crud Crud_GRPC) Delete(m *organization_category_types.OrganizationCategoryType) error {
	// var Otvet organization_category_types.OrganizationCategoryType

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
	Response, err := grpc_client.Client.OrganizationCategoryType_Delete(ctx, Request)
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
func (crud Crud_GRPC) Restore(m *organization_category_types.OrganizationCategoryType) error {
	// var Otvet organization_category_types.OrganizationCategoryType

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
	Response, err := grpc_client.Client.OrganizationCategoryType_Restore(ctx, Request)
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

// Find_ByExtID - находит модель в БД по ext_id и connection_id
func (crud Crud_GRPC) Find_ByExtID(m *organization_category_types.OrganizationCategoryType) error {
	//подключение
	if grpc_client.Client == nil {
		grpc_client.Connect()
	}

	//подготовка запроса
	var VersionModel = crud.GetVersionModel()

	Request := &grpc_proto.RequestExtId{}
	Request.ExtId = m.ExtID
	Request.ConnectionId = m.ConnectionID
	Request.VersionModel = VersionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_SECONDS))
	defer ctxCancelFunc()

	//запрос
	Response, err := grpc_client.Client.OrganizationCategoryType_FindByExtID(ctx, Request)
	if err != nil {
		sError := err.Error()
		if sError[0:len(constants.TEXT_ERROR_MODEL_VERSION)] == constants.TEXT_ERROR_MODEL_VERSION {
			log.Panic("table: ", TableName, " error: ", err)
		}
		return err
	}

	//ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), m)
	if err != nil {
		return err
	}

	return err
}