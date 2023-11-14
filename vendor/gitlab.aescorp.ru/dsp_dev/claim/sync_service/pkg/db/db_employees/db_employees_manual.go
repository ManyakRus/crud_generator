package db_employees

import (
	"context"
	"errors"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
	"strings"
	"time"
)

// Find_ByLogin - находит запись в БД по Login
func (crud Crud_DB) Find_ByLogin(e *employees.Employee) error {
	// var Otvet employees.Employee
	var err error

	if e.Login == "" {
		err = errors.New("Error: login =''")
		return err
	}

	//
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByLogin_ctx(ctx, e)

	return err
}

// Find_ByLogin_ctx - находит запись в БД по Login
func (crud Crud_DB) Find_ByLogin_ctx(ctx context.Context, e *employees.Employee) error {
	// var Otvet employees.Employee
	var err error

	if e.Login == "" {
		err = errors.New("Error: login =''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("UPPER(login) LIKE UPPER(?)", strings.ToUpper(e.Login)).First(e)
	//	tx := db.Where("login = ?", e.Login).First(e)
	err = tx.Error

	return err
}

// Find_ByEMail - находит запись в БД по EMail
func (crud Crud_DB) Find_ByEMail(e *employees.Employee) error {
	// var Otvet employees.Employee
	var err error

	if e.Email == "" {
		err = errors.New("Error: email =''")
		return err
	}

	//
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByEMail_ctx(ctx, e)

	return err
}

// Find_ByEMail_ctx - находит запись в БД по EMail
func (crud Crud_DB) Find_ByEMail_ctx(ctx context.Context, e *employees.Employee) error {
	// var Otvet employees.Employee
	var err error
	// log.Trace("start Find_ByExtID() ", TableName, " ext_id: ", ext_id)

	if e.Email == "" {
		err = errors.New("Error: email =''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("email = ?", e.Email).First(e)
	err = tx.Error

	return err
}

// Find_ByFIO - находит запись в БД по name + second_name + parent_name
func (crud Crud_DB) Find_ByFIO(e *employees.Employee) error {
	// var Otvet employees.Employee
	var err error

	if e.Name == "" && e.SecondName == "" && e.ParentName == "" {
		err = errors.New("Error: name + second_name + parent_name = ''")
		return err
	}

	//
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByFIO_ctx(ctx, e)

	return err
}

// Find_ByFIO_ctx - находит запись в БД по name + second_name + parent_name
func (crud Crud_DB) Find_ByFIO_ctx(ctx context.Context, e *employees.Employee) error {
	// var Otvet employees.Employee
	var err error

	if e.Name == "" && e.SecondName == "" && e.ParentName == "" {
		err = errors.New("Error: name + second_name + parent_name = ''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("name = ?", e.Name).Where("second_name = ?", e.SecondName).Where("parent_name = ?", e.ParentName).First(e)
	err = tx.Error

	return err
}
