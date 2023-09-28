package object_model

import (
	"context"
	"errors"
	"fmt"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/micro"
	"gitlab.aescorp.ru/dsp_dev/claim/nikitin/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

type crud_MessageType struct {
}

// Read - находит запись в БД по ID
func (crud crud_MessageType) read(m *MessageType) error {
	var err error

	//log.Trace("start Read() ", TableName, " id: ", id)
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.read_ctx(ctx, m)
	return err
}

// Read_ctx - находит запись в БД по ID
func (crud crud_MessageType) read_ctx(ctx context.Context, m *MessageType) error {
	var err error

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.First(m, m.ID)
	err = tx.Error

	return err
}

// Save - записывает новый или существующий объект в базу данных
func (crud crud_MessageType) save(m *MessageType) error {
	err := crud.create_update(m, false)
	return err
}

// Save_ctx - записывает новый или существующий объект в базу данных
func (crud crud_MessageType) save_ctx(ctx context.Context, m *MessageType) error {
	is_create := !micro.BoolFromInt64(m.ID)
	err := crud.create_update_ctx(ctx, m, is_create)
	return err
}

// Update - записывает существующий объект в базу данных
func (crud crud_MessageType) update(m *MessageType) error {
	err := crud.create_update(m, false)
	return err
}

// Update_ctx - записывает существующий объект в базу данных
func (crud crud_MessageType) update_ctx(ctx context.Context, m *MessageType) error {
	err := crud.create_update_ctx(ctx, m, false)
	return err
}

// Create - записывает новый объект в базу данных
func (crud crud_MessageType) create(m *MessageType) error {
	err := crud.create_update(m, true)
	return err
}

// Create_ctx - записывает новый объект в базу данных
func (crud crud_MessageType) create_ctx(ctx context.Context, m *MessageType) error {
	err := crud.create_update_ctx(ctx, m, true)
	return err
}

// create_update - записывает объект в базу данных
func (crud crud_MessageType) create_update(m *MessageType, is_create bool) error {
	var err error

	//log.Trace("start Save() ", TableName, " id: ", m.ID)

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.create_update_ctx(ctx, m, is_create)
	return err
}

// create_update_ctx - записывает объект в базу данных
func (crud crud_MessageType) create_update_ctx(ctx context.Context, m *MessageType, is_create bool) error {
	var err error

	//log.Trace("start Save() ", TableName, " id: ", m.ID)

	// проверка ID
	if is_create == true {
		if m.ID != 0 {
			TextError := fmt.Sprint("db.Save() ", m.TableName(), " error: id !=0")
			//log.Panic(sError)
			err = errors.New(TextError)
			return err
		}
	} else if m.ID == 0 {
		TextError := fmt.Sprint("db.Save() ", m.TableName(), " error: id =0")
		err = errors.New(TextError)
		//log.Panic(sError)
		return err
	}

	//
	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	//заполним даты
	Now := time.Now()
	m.ModifiedAt = Now
	if m.IsDeleted == true && m.DeletedAt.IsZero() == true {
		m.DeletedAt = Now
	} else if m.IsDeleted == false && m.DeletedAt.IsZero() == false {
		m.DeletedAt = time.Time{}
	}

	//колонки с null
	tx := db
	MassOmit := make([]string, 0)
	var ColumnName string

	ColumnName = "DeletedAt"
	if m.DeletedAt.IsZero() == true {
		MassOmit = append(MassOmit, ColumnName)
	}

	ColumnName = "ExtID"
	if m.ExtID == 0 {
		MassOmit = append(MassOmit, ColumnName)
	}

	//игнор пустых колонок
	tx = tx.Omit(MassOmit...)

	//запись
	if is_create == true {
		tx = tx.Create(m)
	} else {
		tx = tx.Save(m)
	}
	err = tx.Error
	if err != nil {
		return err
	}

	//запишем NULL в пустые колонки
	for f := 0; f < len(MassOmit); f++ {
		ColumnName := MassOmit[f]
		tx = db.First(m).Update(ColumnName, gorm.Expr("NULL"))

		err = tx.Error
		if err != nil {
			TextError := fmt.Sprint("db.Update() ", m.TableName(), " id: ", m.ID, " error: ", err)
			err = errors.New(TextError)
			return err
			//log.Panic(sError)
		}
	}

	return err
}

// Delete - записывает is_deleted = true
func (crud crud_MessageType) delete(m *MessageType) error {
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.delete_ctx(ctx, m)
	return err
}

// Delete_ctx - записывает is_deleted = true
func (crud crud_MessageType) delete_ctx(ctx context.Context, m *MessageType) error {
	var err error

	var m2 *MessageType
	m2.ID = m.ID
	err = crud.read_ctx(ctx, m2)
	if err != nil {
		return err
	}

	m.IsDeleted = true
	m2.IsDeleted = true

	err = crud.save_ctx(ctx, m2)

	return err
}

// Restore - записывает is_deleted = true
func (crud crud_MessageType) restore(m *MessageType) error {
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.restore_ctx(ctx, m)
	return err
}

// Restore_ctx - записывает is_deleted = true
func (crud crud_MessageType) restore_ctx(ctx context.Context, m *MessageType) error {
	var err error

	var m2 *MessageType
	m2.ID = m.ID
	err = crud.read_ctx(ctx, m2)
	if err != nil {
		return err
	}

	m.IsDeleted = false
	m2.IsDeleted = false

	err = crud.save_ctx(ctx, m2)

	return err
}

//// Find_ByExtID - находит запись в БД по ext_id и connection_id
//func Find_ByExtID(ext_id int64, connection_id int64) (MessageType, error) {
//	var Otvet MessageType
//	var err error
//
//	if ext_id <= 0 {
//		err = errors.New("Error: ext_id <=0")
//		return Otvet, err
//	}
//
//	//
//	ctxMain := context.Background()
//	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(TIMEOUT_DB_SECONDS))
//	defer ctxCancelFunc()
//
//	Otvet, err = Find_ByExtID_ctx(ctx, ext_id, connection_id)
//
//	return Otvet, err
//}

//// Find_ByExtID_ctx - находит запись в БД по ext_id и connection_id
//func Find_ByExtID_ctx(ctx context.Context, ext_id int64, connection_id int64) (MessageType, error) {
//	var Otvet MessageType
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
