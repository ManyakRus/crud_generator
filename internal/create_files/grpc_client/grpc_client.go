package grpc_client

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

	err = CreateGRPCClient()
	if err != nil {
		log.Error("CreateGRPCClient() error: ", err)
		return err
	}

	err = CreateGRPCClientTest()
	if err != nil {
		log.Error("CreateGRPCClientTest() error: ", err)
		return err
	}

	return err
}
