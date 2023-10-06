package folders

import (
	"errors"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
)

// CreateFolder - создаёт папку на диске
func CreateFolder(FilenameFull string, FilePermissions uint32) error {
	var err error

	FileMode1 := os.FileMode(FilePermissions)
	if FilePermissions == 0 {
		FileMode1 = os.FileMode(0700)
	}

	if _, err := os.Stat(FilenameFull); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(FilenameFull, FileMode1)
		if err != nil {
			return err
		}
	}

	return err
}

// DeleteFolder - создаёт папку на диске
func DeleteFolder(FilenameFull string) error {
	var err error

	if _, err := os.Stat(FilenameFull); errors.Is(err, os.ErrNotExist) {
		return err
	}

	err = os.RemoveAll(FilenameFull)
	if err != nil {
		return err
	}

	return err
}

func CreateAllFolders() {
	var err error

	dir := micro.ProgramDir()

	//
	Filename := dir + config.Settings.SERVICE_NAME
	err = CreateFolder(Filename, 0)
	if err != nil {
		log.Panic("CreateFolder() ", Filename, " error: ", err)
	}

	//
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "internal"
	err = CreateFolder(Filename, 0)
	if err != nil {
		log.Panic("CreateFolder() ", Filename, " error: ", err)
	}
	log.Info("CreateFolder() ", Filename)

	//
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg"
	err = CreateFolder(Filename, 0)
	if err != nil {
		log.Panic("CreateFolder() ", Filename, " error: ", err)
	}
	log.Info("CreateFolder() ", Filename)

	//
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "model"
	err = CreateFolder(Filename, 0)
	if err != nil {
		log.Panic("CreateFolder() ", Filename, " error: ", err)
	}
	log.Info("CreateFolder() ", Filename)

	if config.Settings.NEED_CRUD == true {
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "db"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)
	}

	if config.Settings.NEED_GRPC == true {
		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)

		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc" + micro.SeparatorFile() + "grpc_server"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)

		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc" + micro.SeparatorFile() + "grpc_client"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)

		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc" + micro.SeparatorFile() + "grpc_proto"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)
	}

	if config.Settings.NEED_NRPC == true {
		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "nrpc"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)

		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc" + micro.SeparatorFile() + "nrpc_server"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)

		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc" + micro.SeparatorFile() + "nrpc_client"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)

		//
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg" + micro.SeparatorFile() + "grpc" + micro.SeparatorFile() + "grpc_proto"
		err = CreateFolder(Filename, 0)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)
	}

	//return err
}
