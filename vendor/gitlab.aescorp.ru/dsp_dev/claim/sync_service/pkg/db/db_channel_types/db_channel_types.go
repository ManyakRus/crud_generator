// File generated automatic with crud_generator app
// Do not change anything here.
package db_channel_types

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/channel_types"
	"gorm.io/gorm"
	"time"
)

// TableName - имя таблицы в БД Postgres
const TableName string = "channel_types"

// Crud_DB - объект для CRUD операций через БД
type Crud_DB struct {
}

// Read - находит запись в БД по ID
func (crud Crud_DB) Read(m *channel_types.ChannelType) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Read_ctx(ctx, m)
	return err
}

// Read_ctx - находит запись в БД по ID
func (crud Crud_DB) Read_ctx(ctx context.Context, m *channel_types.ChannelType) error {
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
func (crud Crud_DB) Save(m *channel_types.ChannelType) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Save_ctx(ctx, m)
	return err
}

// Save_ctx - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save_ctx(ctx context.Context, m *channel_types.ChannelType) error {
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
func (crud Crud_DB) Update(m *channel_types.ChannelType) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Update_ctx(ctx, m)
	return err
}

// Update_ctx - записывает существующий объект в базу данных
func (crud Crud_DB) Update_ctx(ctx context.Context, m *channel_types.ChannelType) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	err = crud.create_update_ctx(ctx, m, false)
	return err
}

// Create - записывает новый объект в базу данных
func (crud Crud_DB) Create(m *channel_types.ChannelType) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Create_ctx(ctx, m)
	return err
}

// Create_ctx - записывает новый объект в базу данных
func (crud Crud_DB) Create_ctx(ctx context.Context, m *channel_types.ChannelType) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	err = crud.create_update_ctx(ctx, m, true)
	return err
}

// create_update - записывает объект в базу данных
func (crud Crud_DB) create_update(m *channel_types.ChannelType, is_create bool) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.create_update_ctx(ctx, m, is_create)
	return err
}

// create_update_ctx - записывает объект в базу данных
func (crud Crud_DB) create_update_ctx(ctx context.Context, m *channel_types.ChannelType, is_create bool) error {
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
	if m.IsDeleted == true && m.DeletedAt.IsZero() == true {
		m.DeletedAt = time.Now()
	} else if m.IsDeleted == false && m.DeletedAt.IsZero() == false {
		m.DeletedAt = time.Time{}
	}

	//колонки с null
	tx := db
	MassOmit := make([]string, 0)
	var ColumnName string

	ColumnName = "ModifiedAt"
	if m.ModifiedAt.IsZero() == true {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "DeletedAt"
	if m.DeletedAt.IsZero() == true {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "ExtID"
	if m.ExtID == 0 {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "CreatedAt"
	if m.CreatedAt.IsZero() == true {
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

// Delete - записывает is_deleted = true
func (crud Crud_DB) Delete(m *channel_types.ChannelType) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Delete_ctx(ctx, m)
	return err
}

// Delete_ctx - записывает is_deleted = true
func (crud Crud_DB) Delete_ctx(ctx context.Context, m *channel_types.ChannelType) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	m2 := channel_types.ChannelType{}
	m2.ID = m.ID
	err = crud.Read_ctx(ctx, &m2)
	if err != nil {
		return err
	}

	m2.IsDeleted = true
	m.IsDeleted = true

	err = crud.Save_ctx(ctx, &m2)

	return err
}

// Restore - записывает is_deleted = true
func (crud Crud_DB) Restore(m *channel_types.ChannelType) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Restore_ctx(ctx, m)
	return err
}

// Restore_ctx - записывает is_deleted = true
func (crud Crud_DB) Restore_ctx(ctx context.Context, m *channel_types.ChannelType) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	m2 := channel_types.ChannelType{}
	m2.ID = m.ID
	err = crud.Read_ctx(ctx, &m2)
	if err != nil {
		return err
	}

	m2.IsDeleted = false
	m.IsDeleted = false

	err = crud.Save_ctx(ctx, &m2)

	return err
}
