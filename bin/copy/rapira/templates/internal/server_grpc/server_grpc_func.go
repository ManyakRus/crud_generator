package server_grpc

import (
	"errors"
	"github.com/ManyakRus/starter/micro"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/constants"
)

// ErrorModelVersion
func ErrorModelVersion(Model interface{}) error {
	var err error

	TypeName := micro.GetType(Model)

	s := constants.TEXT_ERROR_MODEL_VERSION + " " + TypeName
	err = errors.New(s)
	return err
}
