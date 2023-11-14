package db_lawsuit_status_states

import (
	"context"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_states"
	"strconv"
	"strings"
	"time"
)

func (crud Crud_DB) Fill_from_Lawsuit(Lawsuit_id int64, Status_id int64) error {
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	TextSQL := `
INSERT INTO public.lawsuit_status_states as lss (
	lawsuit_id,
	status_at,
	modified_at,
	total_debt,
	invoice_sum,
	pay_sum,
	main_sum,
	penalty_sum,
	penny_sum,
	restrict_sum,
	status_id,
	created_at,
	state_duty_sum
	)
SELECT 
	l.id, 
	Now(),
	Now(),
	debt_sum,
	invoice_sum,
	pay_sum,
	main_sum,
	penalty as penalty_sum,
	penny as penny_sum,
	restrict_sum,
	:status_id as status_id,
	Now(),
	state_duty as state_duty_sum
from 
    public.lawsuits as l 
WHERE 1=1
	and l.id = :lawsuit_id
	
ON CONFLICT (lawsuit_id, status_id) DO UPDATE
SET
	lawsuit_id = EXCLUDED.lawsuit_id,
	status_at = Now(),
	modified_at = Now(),
	total_debt = EXCLUDED.total_debt,
	invoice_sum = EXCLUDED.invoice_sum,
	pay_sum = EXCLUDED.pay_sum,
	main_sum = EXCLUDED.main_sum,
	penalty_sum = EXCLUDED.penalty_sum,
	penny_sum = EXCLUDED.penny_sum,
	restrict_sum = EXCLUDED.restrict_sum,
	status_id = EXCLUDED.status_id,
	state_duty_sum = EXCLUDED.state_duty_sum
`

	TextSQL = strings.ReplaceAll(TextSQL, ":lawsuit_id", strconv.Itoa(int(Lawsuit_id)))
	TextSQL = strings.ReplaceAll(TextSQL, ":status_id", strconv.Itoa(int(Status_id)))

	tx := db.Exec(TextSQL)
	err = tx.Error

	return err
}

func (crud Crud_DB) FindDebtSum(lawsuit_id, status_id int64) (float64, error) {
	var Otvet float64
	var err error

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	lawsuitStatusState := lawsuit_status_states.LawsuitStatusState{}
	tx := db.Model(lawsuitStatusState).Where("lawsuit_id = ?", lawsuit_id).Where("status_id = ?", status_id).First(&lawsuitStatusState)
	err = tx.Error
	Otvet = lawsuitStatusState.TotalDebt

	return Otvet, err
}
