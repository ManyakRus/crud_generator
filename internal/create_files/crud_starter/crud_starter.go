package crud_starter

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
)

// CreateAllFiles - создаёт все файлы в папке crud_starter
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	//crud_starter.go
	err = CreateFile_CrudStarter(MapAll)
	if err != nil {
		log.Error("CreateFile_CrudStarter() error: ", err)
		return err
	}

	//crud_starter_manual.go
	//if config.Settings.NEED_CREATE_MANUAL_FILES == true {
	err = CreateFileCrudStarter_manual(MapAll)
	if err != nil {
		log.Error("CreateFileCrudStarter_manual() error: ", err)
		return err
	}

	//ReadObject
	if config.Settings.NEED_CREATE_READOBJECT == true {
		//crud_starter_readobject.go
		err = CreateFile_CrudStarter_ReadObject(MapAll)
		if err != nil {
			log.Error("CreateFile_CrudStarter_ReadObject() error: ", err)
			return err
		}
	}

	return err
}
