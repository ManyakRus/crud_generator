package crud

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

	err = CreateCrud()
	if err != nil {
		log.Error("CreateCrud() error: ", err)
		return err
	}

	err = CreateCrudTest()
	if err != nil {
		log.Error("CreateCrudTest() error: ", err)
		return err
	}

	return err
}
