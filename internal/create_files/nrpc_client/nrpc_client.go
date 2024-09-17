package nrpc_client

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/starter/log"
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles() error {
	var err error

	if config.Settings.NEED_CREATE_GRPC == false {
		return err
	}

	err = CreateNRPCClient()
	if err != nil {
		log.Error("CreateNRPCClient() error: ", err)
		return err
	}

	err = CreateNRPCClient_Test()
	if err != nil {
		log.Error("CreateNRPCClient_Test() error: ", err)
		return err
	}

	return err
}
