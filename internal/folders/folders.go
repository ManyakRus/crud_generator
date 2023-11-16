package folders

import (
	"errors"
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	copy_files "github.com/otiai10/copy"
	"os"
	"path/filepath"
)

// CreateFolder - создаёт папку на диске
func CreateFolder(FoldernameFull string) {
	var err error

	ok, err := micro.FileExists(FoldernameFull)
	if ok == false || err != nil {
		err = CreateFolder_err(FoldernameFull, 0777)
		if err != nil {
			log.Panic("CreateFolder_err() ", FoldernameFull, " error: ", err)
		}
		//log.Info("CreateFolder_err() ", FoldernameFull)
	}

}

// CreateFolder_err - создаёт папку на диске
func CreateFolder_err(FilenameFull string, FilePermissions uint32) error {
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
		err = CreateFolder_err(Filename, 0777)
		if err != nil {
			log.Panic("CreateFolder_err() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder_err() ", Filename)
	}

	//
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "internal"
	ok, err = micro.FileExists(Filename)
	if ok == false || err != nil {
		err = CreateFolder_err(Filename, 0777)
		if err != nil {
			log.Panic("CreateFolder_err() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder_err() ", Filename)
	}

	//
	//Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + "pkg"
	//ok, err = micro.FileExists(Filename)
	//if ok == false || err != nil {
	//	err = CreateFolder_err(Filename, 0777)
	//	if err != nil {
	//		log.Panic("CreateFolder_err() ", Filename, " error: ", err)
	//	}
	//	log.Info("CreateFolder_err() ", Filename)
	//}

	//model
	Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_MODEL
	ok, err = micro.FileExists(Filename)
	if ok == false || err != nil {
		err = CreateFolder_err(Filename, 0777)
		if err != nil {
			log.Panic("CreateFolder_err() ", Filename, " error: ", err)
		}
		log.Info("CreateFolder_err() ", Filename)
	}

	if config.Settings.NEED_CREATE_DB == true {
		//db
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_CRUD
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

		//crud_starter
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_CRUD_STARTER
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

	}

	if config.Settings.NEED_CREATE_GRPC == true {
		//grpc
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

		//grpc_server
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_SERVER
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

		////grpc client
		//Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_CLIENT
		//ok, err = micro.FileExists(Filename)
		//if ok == false || err != nil {
		//	err = CreateFolder_err(Filename, 0777)
		//	if err != nil {
		//		log.Panic("CreateFolder_err() ", Filename, " error: ", err)
		//	}
		//	log.Info("CreateFolder_err() ", Filename)
		//}

		//grpc_proto
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile() + "grpc_proto"
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}
	}

	if config.Settings.NEED_CREATE_NRPC == true {
		//nrpc
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_NRPC
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

		//server_nrpc
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_NRPC_SERVER
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

		//nrpc client
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_NRPC_CLIENT
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}

		//grpc_proto
		Filename = dir + config.Settings.SERVICE_NAME + micro.SeparatorFile() + config.Settings.TEMPLATE_FOLDERNAME_GRPC_PROTO + micro.SeparatorFile() + "grpc_proto"
		ok, err = micro.FileExists(Filename)
		if ok == false || err != nil {
			err = CreateFolder_err(Filename, 0777)
			if err != nil {
				log.Panic("CreateFolder_err() ", Filename, " error: ", err)
			}
			log.Info("CreateFolder_err() ", Filename)
		}
	}

	//return err
}

// CopyAllFiles_Exclude_ - копирует все файлы из src в dest, кроме "*_"
func CopyAllFiles_Exclude_(src, dest string) error {
	var err error

	opt := copy_files.Options{
		Skip: CopyFilesFilterGo,
	}
	err = copy_files.Copy(src, dest, opt)

	return err
}

// filter to all files, insted of "*_"
func CopyFilesFilterGo(info os.FileInfo, src, dest string) (bool, error) {
	var err error
	var Otvet bool

	Filename := src
	if Filename[len(Filename)-1:] == "_" {
		Otvet = true
		return Otvet, err
	}

	FolderName := filepath.Dir(src)
	if len(FolderName) > 0 && FolderName[len(FolderName)-1:] == "_" {
		Otvet = true
		return Otvet, err
	}

	return Otvet, err
}
