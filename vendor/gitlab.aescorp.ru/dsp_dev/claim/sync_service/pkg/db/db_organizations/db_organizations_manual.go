package db_organizations

import (
	"context"
	"errors"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organizations"
	"time"
)

// Find_ByInnKpp - находит запись в БД по ИНН и КПП
func (crud Crud_DB) Find_ByInnKpp(o *organizations.Organization) error {
	// var Otvet organizations.Organization
	var err error

	if o.INN == "" {
		err = errors.New("Error: INN = ''")
		return err
	}

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.Find_ByInnKpp_ctx(ctx, o)

	return err
}

// Find_ByInnKpp_ctx - находит запись в БД по ИНН и КПП
func (crud Crud_DB) Find_ByInnKpp_ctx(ctx context.Context, o *organizations.Organization) error {
	// var Otvet organizations.Organization
	var err error

	if o.INN == "" {
		err = errors.New("Error: INN =''")
		return err
	}

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("INN = ?", o.INN)
	if o.KPP != "" {
		tx = tx.Where("KPP = ?", o.KPP)
	}
	tx = tx.First(&o)

	err = tx.Error

	return err
}
