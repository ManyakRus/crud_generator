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
		FileMode1 = os.FileMode(0777)
	}

	if _, err := os.Stat(FilenameFull); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(FilenameFull, FileMode1)
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

// CreateAllFolders - создаёт все нужные каталоги Ready
func CreateAllFolders() {
	var err error

	dir := micro.ProgramDir_bin()

	//
	Filename := dir + config.Settings.SERVICE_NAME
	ok, err := micro.FileExists(Filename)
	if ok == false || err != nil {
		err = CreateFolder(Filename, 0777)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)
	}

	//
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "internal"
	ok, err = micro.FileExists(Filename)
	if ok == false || err != nil {
		err = CreateFolder(Filename, 0777)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)
	}

	//
	//Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg"
	//ok, err = micro.FileExists(Filename)
	//if ok == false || err != nil {
	//	err = CreateFolder(Filename, 0777)
	//	if err != nil {
	//		log.Panic("CreateFolder() ", Filename, " error: ", err)
	//	}
	//	log.Info("CreateFolder() ", Filename)
	//}

	//model
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_MODEL
	ok, err = micro.FileExists(Filename)
	if ok == false || err != nil {
		err = CreateFolder(Filename, 0777)
		if err != nil {
			log.Panic("CreateFolder() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder() ", Filename)
	}

	if config.Settings.NEED_CREATE_DB == true {
		//db
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_DB
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//crud_starter
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

	}

	if config.Settings.NEED_CREATE_GRPC == true {
		//grpc
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//grpc_server
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//grpc client
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//grpc_proto
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC + micro.SeparatorFile() + "grpc_proto"
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}
	}

	if config.Settings.NEED_CREATE_NRPC == true {
		//nrpc
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_NRPC
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//server_nrpc
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//nrpc client
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}

		//grpc_proto
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC + micro.SeparatorFile() + "grpc_proto"
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder() ", Filename)
		}
	}

	//return err
}
