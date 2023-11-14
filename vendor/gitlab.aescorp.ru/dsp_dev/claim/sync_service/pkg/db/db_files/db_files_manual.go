package db_files

import (
	"context"
	"errors"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/files"
	"time"
)

// Find_ByFileId - находит запись в БД по File_id
func (crud Crud_DB) Find_ByFileId(f *files.File) error {
	// var Otvet files.File
	var err error

	if f.FileID == "" {
		err = errors.New("Error: file_id = ''")
		return err
	}

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByFileId_ctx(ctx, f)

	return err
}

// Find_ByFileId_ctx - находит запись в БД по File_id
func (crud Crud_DB) Find_ByFileId_ctx(ctx context.Context, f *files.File) error {
	// var Otvet files.File
	var err error

	if f.FileID == "" {
		err = errors.New("Error: file_id = ''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("file_id = ?", f.FileID).First(f)
	err = tx.Error

	return err
}

// Find_ByFull_name - находит запись в БД по File_id
func (crud Crud_DB) Find_ByFullName(f *files.File) error {
	// var Otvet files.File
	var err error

	if f.FullName == "" {
		err = errors.New("Error: full_name =''")
		return err
	}

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByFullName_ctx(ctx, f)

	return err
}

// Find_ByFull_name_ctx - находит запись в БД по File_id
func (crud Crud_DB) Find_ByFullName_ctx(ctx context.Context, f *files.File) error {
	// var Otvet files.File
	var err error

	if f.FullName == "" {
		err = errors.New("Error: full_name = ''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("full_name = ?", f.FullName).First(f)
	err = tx.Error

	return err
}
