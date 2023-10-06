package db_lawsuit_status_types

import (
	"context"
	"errors"
	"fmt"
	model "gitlab.aescorp.ru/dsp_dev/claim/common/object_model"
	"gitlab.aescorp.ru/dsp_dev/claim/common/object_model/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/micro"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// TableName - имя таблицы в БД Postgres
const TableName string = "lawsuit_status_types"

// Crud_DB - объект для CRUD операций через БД
type Crud_DB struct {
}

// Read - находит запись в БД по ID
func (crud Crud_DB) Read(l *model.LawsuitStatusType) error {
	//var Otvet model.LawsuitStatusType
	var err error

	//log.Trace("start Read() ", TableName, " id: ", id)
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Read_ctx(ctx, l)
	return err
}

// Read_ctx - находит запись в БД по ID
func (crud Crud_DB) Read_ctx(ctx context.Context, l *model.LawsuitStatusType) error {
	//var Otvet model.LawsuitStatusType
	var err error

	id := l.ID

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.First(l, id)
	err = tx.Error

	return err
}

// Save - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save(l *model.LawsuitStatusType) error {
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Save_ctx(ctx, l)
	return err
}

// Save_ctx - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save_ctx(ctx context.Context, l *model.LawsuitStatusType) error {
	is_create := !micro.BoolFromInt64(l.ID)
	err := crud.create_update_ctx(ctx, l, is_create)
	return err
}

// Update - записывает существующий объект в базу данных
func (crud Crud_DB) Update(l *model.LawsuitStatusType) error {
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Update_ctx(ctx, l)
	return err
}

// Update_ctx - записывает существующий объект в базу данных
func (crud Crud_DB) Update_ctx(ctx context.Context, l *model.LawsuitStatusType) error {
	err := crud.create_update_ctx(ctx, l, false)
	return err
}

// Create - записывает новый объект в базу данных
func (crud Crud_DB) Create(l *model.LawsuitStatusType) error {
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err := crud.Create_ctx(ctx, l)
	return err
}

// Create_ctx - записывает новый объект в базу данных
func (crud Crud_DB) Create_ctx(ctx context.Context, l *model.LawsuitStatusType) error {
	err := crud.create_update_ctx(ctx, l, true)
	return err
}

// create_update - записывает объект в базу данных
func (crud Crud_DB) create_update(l *model.LawsuitStatusType, is_create bool) error {
	var err error

	//log.Trace("start Save() ", TableName, " id: ", m.ID)

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.create_update_ctx(ctx, l, is_create)
	return err
}

// create_update_ctx - записывает объект в базу данных
func (crud Crud_DB) create_update_ctx(ctx context.Context, l *model.LawsuitStatusType, is_create bool) error {
	var err error

	//log.Trace("start Save() ", TableName, " id: ", m.ID)

	// проверка ID
	if is_create == true {
		if l.ID != 0 {
			TextError := fmt.Sprint("db.Save() ", TableName, " error: id !=0")
			//log.Panic(sError)
			err = errors.New(TextError)
			return err
		}
	} else if l.ID == 0 {
		TextError := fmt.Sprint("db.Save() ", TableName, " error: id =0")
		err = errors.New(TextError)
		//log.Panic(sError)
		return err
	}

	//
	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	//заполним даты
	Now := time.Now()
	l.ModifiedAt = Now
	if l.IsDeleted == true && l.DeletedAt.IsZero() == true {
		l.DeletedAt = Now
	} else if l.IsDeleted == false && l.DeletedAt.IsZero() == false {
		l.DeletedAt = time.Time{}
	}

	//колонки с null
	tx := db
	MassOmit := make([]string, 0)
	var ColumnName string

	ColumnName = "DeletedAt"
	if l.DeletedAt.IsZero() == true {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "ExtID"
	if l.ExtID == 0 {
		MassOmit = append(MassOmit, ColumnName)
	}

	//игнор пустых колонок
	tx = tx.Omit(MassOmit...)

	//запись
	if is_create == true {
		tx = tx.Create(&l)
	} else {
		tx = tx.Save(&l)
	}
	err = tx.Error
	if err != nil {
		return err
	}

	//запишем NULL в пустые колонки
	for f := 0; f < len(MassOmit); f++ {
		ColumnName := MassOmit[f]
		tx = db.First(&l).Update(ColumnName, gorm.Expr("NULL"))

		err = tx.Error
		if err != nil {
			TextError := fmt.Sprint("db.Update() ", TableName, " id: ", l.ID, " error: ", err)
			err = errors.New(TextError)
			return err
			//log.Panic(sError)
		}
	}

	return err
}

// Delete - записывает is_deleted = true
func (crud Crud_DB) Delete(l *model.LawsuitStatusType) error {
	//var Otvet model.LawsuitStatusType
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Delete_ctx(ctx, l)
	return err
}

// Delete_ctx - записывает is_deleted = true
func (crud Crud_DB) Delete_ctx(ctx context.Context, l *model.LawsuitStatusType) error {
	//var Otvet model.LawsuitStatusType
	var err error

	l2 := model.LawsuitStatusType{}
	l2.ID = l.ID
	err = crud.Read_ctx(ctx, &l2)
	if err != nil {
		return err
	}

	l2.IsDeleted = true
	l.IsDeleted = true

	err = crud.Save_ctx(ctx, &l2)

	return err
}

// Restore - записывает is_deleted = true
func (crud Crud_DB) Restore(l *model.LawsuitStatusType) error {
	//var Otvet model.LawsuitStatusType
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Restore_ctx(ctx, l)
	return err
}

// Restore_ctx - записывает is_deleted = true
func (crud Crud_DB) Restore_ctx(ctx context.Context, l *model.LawsuitStatusType) error {
	//var Otvet model.LawsuitStatusType
	var err error

	l2 := model.LawsuitStatusType{}
	l2.ID = l.ID
	err = crud.Read_ctx(ctx, &l2)
	if err != nil {
		return err
	}

	l2.IsDeleted = false
	l.IsDeleted = false

	err = crud.Save_ctx(ctx, &l2)

	return err
}

//// Find_ByExtID - находит запись в БД по ext_id и connection_id
//func Find_ByExtID(ext_id int64, connection_id int64) (model.LawsuitStatusType, error) {
//	var Otvet model.LawsuitStatusType
//	var err error
//
//	if ext_id <= 0 {
//		err = errors.New("Error: ext_id <=0")
//		return Otvet, err
//	}
//
//	//
//	ctxMain := context.Background()
//	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
//	defer ctxCancelFunc()
//
//	Otvet, err = Find_ByExtID_ctx(ctx, ext_id, connection_id)
//
//	return Otvet, err
//}
//
//// Find_ByExtID_ctx - находит запись в БД по ext_id и connection_id
//func Find_ByExtID_ctx(ctx context.Context, ext_id int64, connection_id int64) (model.LawsuitStatusType, error) {
//	var Otvet model.LawsuitStatusType
//	var err error
//	//log.Trace("start Find_ByExtID() ", TableName, " ext_id: ", ext_id)
//
//	if ext_id <= 0 {
//		err = errors.New("Error: ext_id <=0")
//		return Otvet, err
//	}
//
//	db := postgres_gorm.GetConnection()
//	db.WithContext(ctx)
//
//	tx := db.Where("ext_id = ?", ext_id).Where("connection_id = ?", connection_id).First(&Otvet)
//	err = tx.Error
//
//	return Otvet, err
//}
