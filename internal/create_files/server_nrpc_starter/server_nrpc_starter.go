package server_nrpc_starter

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/starter/log"
)

// CreateAllFiles - создаёт все файлы в папке grpc proto
func CreateAllFiles() error {
	var err error

	if config.Settings.NEED_CREATE_NRPC == false {
		return err
	}

	//
	err = CreateFile_ServerGRPCStarter()
	if err != nil {
		log.Error("CreateFile_ServerGRPCStarter() error: ", err)
		return err
	}

	//
	err = CreateFile_ServerGRPCStarter_Test()
	if err != nil {
		log.Error("CreateFile_ServerGRPCStarter_Test() error: ", err)
		return err
	}

	return err
}
