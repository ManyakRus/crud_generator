package lawsuit_payments

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/format_date"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/functions/format_time"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/types/alias"
)

// RegisteredAt -- омент регистрации платежа в системе
func (sf *LawsuitPayment) RegisteredAt() alias.PaymentRegisteredAt {
	strDate := format_time.FormatTime(sf.CreatedAt)
	return alias.PaymentRegisteredAt(strDate)
}

// DatePayAt -- возвращает момент оплаты
func (sf *LawsuitPayment) DatePayAt() alias.FrontDate {
	frontDate := format_date.FormatDate(sf.Document.DocumentAt)
	return alias.FrontDate(frontDate)
}

// InvoiceId -- возвращает ID привязанной С/Ф
func (sf *LawsuitPayment) InvoiceId() alias.InvoiceId {
	return sf.InvoiceID
}

// Id -- возвращает ID платёжки
func (sf *LawsuitPayment) Id() alias.PaymentId {
	return sf.ID
}
