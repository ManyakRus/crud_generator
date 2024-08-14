package grpc_client_func

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/api/grpc_proto"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/db_constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_constants"
	"strings"
)

// Client - подключение к клиенту GRPC
var Client grpc_proto.SyncServiceClient

// интерфейс для запуска функции IConnect_GRPC_NRPC
type IConnect_GRPC_NRPC interface {
	Connect_GRPC_NRPC()
}

// функция для запуска функции Connect_GRPC_NRPC()
var Func_Connect_GRPC_NRPC IConnect_GRPC_NRPC

// IsErrorModelVersion - возвращает true если ошибка TEXT_ERROR_MODEL_VERSION
func IsErrorModelVersion(err error) bool {
	Otvet := false

	sError := err.Error()
	if strings.Contains(sError, grpc_constants.TEXT_ERROR_MODEL_VERSION) == true {
		Otvet = true
	}

	return Otvet
}

// IsRecordNotFound - возвращает true если ошибка = "record not found"
func IsRecordNotFound(err error) bool {
	Otvet := false

	if err == nil {
		return Otvet
	}

	TextErr := err.Error()
	pos1 := strings.Index(TextErr, db_constants.TEXT_RECORD_NOT_FOUND)
	if pos1 >= 0 {
		Otvet = true
	}

	return Otvet
}
