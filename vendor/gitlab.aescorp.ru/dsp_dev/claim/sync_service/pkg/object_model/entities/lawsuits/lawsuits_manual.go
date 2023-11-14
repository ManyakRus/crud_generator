package lawsuits

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_states"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
	"time"
)

// NewLawsuit Новый объект дела
func NewLawsuit(contractNumber alias.ContractNumber, claimType string) Lawsuit {
	sf := Lawsuit{
		Number:      alias.LawsuitNumber(fmt.Sprintf("ПФ_%s_%s_%s", time.Now().Format("200601-02"), contractNumber, claimType)),
		NumberClaim: alias.ClaimNumber(fmt.Sprintf("ПР_%s_%s_%s", time.Now().Format("200601-02"), contractNumber, claimType)),
		// TODO ВИ, чёж не добил и третий номер до алиаса?
		NumberTrial: fmt.Sprintf("ПИ_%s_%s_%s", time.Now().Format("200601-02"), contractNumber, claimType),
	}
	return sf
}

func AsLawsuit(b []byte) (Lawsuit, error) {
	c := Lawsuit{}
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return Lawsuit{}, err
	}
	return c, nil
}

// GetID - возвращает ID объекта
func (c Lawsuit) GetID() int64 {
	return c.ID
}

// ClaimNumber -- возвращает номер портфеля
func (sf *Lawsuit) ClaimNumber() alias.ClaimNumber {
	return sf.NumberClaim
}

func (l Lawsuit) GetStatusState(statusID int64) lawsuit_status_states.LawsuitStatusState {
	var currentStatusState lawsuit_status_states.LawsuitStatusState

	for _, statusState := range l.StatusStates {
		if statusState.StatusID == statusID {
			currentStatusState = statusState
			break
		}
	}

	return currentStatusState
}
