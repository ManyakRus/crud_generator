package grpc_client_tables

import (
	"github.com/ManyakRus/crud_generator/internal/config"
	"github.com/ManyakRus/crud_generator/internal/create_files"
	"github.com/ManyakRus/crud_generator/internal/types"
	"github.com/ManyakRus/starter/log"
)

// CreateAllFiles - создаёт все файлы в папке grpc_client
func CreateAllFiles(MapAll map[string]*types.Table) error {
	var err error

	for _, Table1 := range MapAll {
		//проверка что таблица нормальная
		err1 := create_files.IsGood_Table(Table1)
		if err1 != nil {
			log.Warn(err1)
			continue
		}

		//файлы grpc_client
		err = CreateFiles(Table1)
		if err != nil {
			log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
			return err
		}

		//тестовые файлы grpc_client
		if config.Settings.NEED_CREATE_GRPC_CLIENT_TEST == true {
			err = CreateFilesTest(Table1)
			if err != nil {
				log.Error("CreateFilesTest() table: ", Table1.Name, " error: ", err)
				return err
			}
		}

		//
		if config.Settings.NEED_CREATE_UPDATE_EVERY_COLUMN == true {
			//файлы grpc_client update
			err = CreateFilesUpdateEveryColumn(Table1)
			if err != nil {
				log.Error("CreateFilesUpdateEveryColumn() table: ", Table1.Name, " error: ", err)
				return err
			}

			//тестовые файлы grpc_client update
			err = CreateFilesUpdateEveryColumnTest(Table1)
			if err != nil {
				log.Error("CreateFilesUpdateEveryColumnTest() table: ", Table1.Name, " error: ", err)
				return err
			}

		}
		//NEED_CREATE_CACHE_API
		if config.Settings.NEED_CREATE_CACHE_API == true {
			//файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_FILES == true {
				err = CreateFiles_GRPC_Client_Cache(Table1)
				if err != nil {
					log.Error("CreateFiles() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

			//тестовые файлы grpc_server cache
			if config.Settings.NEED_CREATE_CACHE_TEST_FILES == true {
				err = CreateFiles_GRPC_Client_Cache_Test(Table1)
				if err != nil {
					log.Error("CreateFilesTest() table: ", Table1.Name, " error: ", err)
					return err
				}
			}

		}

		//
		err = CreateFilesFindBy(Table1)
		if err != nil {
			log.Error("CreateFilesFindBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFilesFindByTest(Table1)
		if err != nil {
			log.Error("CreateFilesFindByTest() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFilesFindMassBy(Table1)
		if err != nil {
			log.Error("CreateFilesFindMassBy() table: ", Table1.Name, " error: ", err)
			return err
		}

		//
		err = CreateFilesFindMassByTest(Table1)
		if err != nil {
			log.Error("CreateFilesFindMassByTest() table: ", Table1.Name, " error: ", err)
			return err
		}

	}

	return err
}
