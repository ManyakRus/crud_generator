package db_messages

import (
	"context"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/messages"
	"time"
)

// findBy_LawsuitID_MessageTypeID - находит запись в БД по lawsuit_id + message_type_id
func (crud Crud_DB) FindBy_LawsuitID_MessageTypeID(m *messages.Message) error {
	var err error

	// log.Trace("start Read() ", TableName, " id: ", id)
	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	err = crud.FindBy_LawsuitID_MessageTypeID_ctx(ctx, m)
	return err
}

// findBy_LawsuitID_MessageTypeID_ctx - находит запись в БД по lawsuit_id + message_type_id
func (crud Crud_DB) FindBy_LawsuitID_MessageTypeID_ctx(ctx context.Context, m *messages.Message) error {
	// var m messages.Message
	var err error

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	tx := db.Where("lawsuit_id = ?", m.LawsuitID).Where("message_type_id = ?", m.MessageTypeID).Order("created_at desc").First(m)
	err = tx.Error

	return err
}
