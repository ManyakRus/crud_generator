package db_organization_casebooks

import (
	"context"
	"errors"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_casebooks"
	"time"
)

// Find_ByInnKpp - находит запись в БД по ИНН и КПП
func (crud Crud_DB) Find_ByInn(o *organization_casebooks.OrganizationCasebook) error {
	// var Otvet organization_casebooks.OrganizationCasebook
	var err error

	if o.INN == "" {
		err = errors.New("Error: INN =''")
		return err
	}

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByInn_ctx(ctx, o)
	return err
}

// Find_ByInn_ctx - находит запись в БД по ИНН и КПП
func (crud Crud_DB) Find_ByInn_ctx(ctx context.Context, o *organization_casebooks.OrganizationCasebook) error {
	// var Otvet organization_casebooks.OrganizationCasebook
	var err error

	if o.INN == "" {
		err = errors.New("Error: INN =''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("INN = ?", o.INN).First(&o)
	tx = tx.First(&o)
	err = tx.Error

	return err
}

// Find_ByOrganizationId - находит запись в БД по organization_id
func (crud Crud_DB) Find_ByOrganizationId(o *organization_casebooks.OrganizationCasebook) error {
	// var Otvet organization_casebooks.OrganizationCasebook
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByOrganizationId_ctx(ctx, o)

	return err
}

// Find_ByOrganizationId_ctx - находит запись в БД по organization_id
func (crud Crud_DB) Find_ByOrganizationId_ctx(ctx context.Context, o *organization_casebooks.OrganizationCasebook) error {
	// var Otvet organization_casebooks.OrganizationCasebook
	var err error

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("organization_id = ?", o.OrganizationID).First(&o)
	err = tx.Error

	return err
}
