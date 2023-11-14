// File generated automatic with crud_generator app
// Do not change anything here.
package db_file_changes

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/file_changes"
	"gorm.io/gorm"
	"time"
)

// TableName - имя таблицы в БД Postgres
const TableName string = "file_changes"

// Crud_DB - объект для CRUD операций через БД
type Crud_DB struct {
}

// Read - находит запись в БД по ID
func (crud Crud_DB) Read(m *file_changes.FileChange) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Read_ctx(ctx, m)
	return err
}

// Read_ctx - находит запись в БД по ID
func (crud Crud_DB) Read_ctx(ctx context.Context, m *file_changes.FileChange) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	id := int64(m.ID)

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.First(m, id)
	err = tx.Error

	return err
}

// Save - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save(m *file_changes.FileChange) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Save_ctx(ctx, m)
	return err
}

// Save_ctx - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save_ctx(ctx context.Context, m *file_changes.FileChange) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	is_create := !micro.BoolFromInt64(int64(m.ID))
	err = crud.create_update_ctx(ctx, m, is_create)
	return err
}

// Update - записывает существующий объект в базу данных
func (crud Crud_DB) Update(m *file_changes.FileChange) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Update_ctx(ctx, m)
	return err
}

// Update_ctx - записывает существующий объект в базу данных
func (crud Crud_DB) Update_ctx(ctx context.Context, m *file_changes.FileChange) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	err = crud.create_update_ctx(ctx, m, false)
	return err
}

// Create - записывает новый объект в базу данных
func (crud Crud_DB) Create(m *file_changes.FileChange) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Create_ctx(ctx, m)
	return err
}

// Create_ctx - записывает новый объект в базу данных
func (crud Crud_DB) Create_ctx(ctx context.Context, m *file_changes.FileChange) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	err = crud.create_update_ctx(ctx, m, true)
	return err
}

// create_update - записывает объект в базу данных
func (crud Crud_DB) create_update(m *file_changes.FileChange, is_create bool) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.create_update_ctx(ctx, m, is_create)
	return err
}

// create_update_ctx - записывает объект в базу данных
func (crud Crud_DB) create_update_ctx(ctx context.Context, m *file_changes.FileChange, is_create bool) error {
	var err error

	// log.Trace("start Save() ", TableName, " id: ", int64(m.ID))

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	// проверка ID
	if is_create == true {
		if int64(m.ID) != 0 {
			TextError := fmt.Sprint("db.Save() ", TableName, " error: id !=0")
			// log.Panic(sError)
			err = errors.New(TextError)
			return err
		}
	} else if int64(m.ID) == 0 {
		TextError := fmt.Sprint("db.Save() ", TableName, " error: id =0")
		err = errors.New(TextError)
		// log.Panic(sError)
		return err
	}

	//
	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	// заполним даты
	m.ModifiedAt = time.Now()

	//колонки с null
	tx := db
	MassOmit := make([]string, 0)
	var ColumnName string

	ColumnName = "EmployeeID"
	if m.EmployeeID == 0 {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "FileID"
	if m.FileID == 0 {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "ModifiedAt"
	if m.ModifiedAt.IsZero() == true {
		MassOmit = append(MassOmit, ColumnName)
	}

	//игнор пустых колонок
	tx = tx.Omit(MassOmit...)

	// запись
	if is_create == true {
		tx = tx.Create(&m)
	} else {
		tx = tx.Save(&m)
	}
	err = tx.Error
	if err != nil {
		return err
	}

	// запишем NULL в пустые колонки
	for f := 0; f < len(MassOmit); f++ {
		ColumnName := MassOmit[f]
		if ColumnName == "CreatedAt" {
			continue
		}
		tx = db.Model(&m).Update(ColumnName, gorm.Expr("NULL"))

		err = tx.Error
		if err != nil {
			TextError := fmt.Sprint("db.Update() ", TableName, " id: ", m.ID, " error: ", err)
			err = errors.New(TextError)
			return err
			// log.Panic(sError)
		}
	}

	return err
}
