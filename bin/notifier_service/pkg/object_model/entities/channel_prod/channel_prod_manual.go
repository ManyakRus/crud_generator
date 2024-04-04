package channel_prod

// Crud_manual_ChannelProd - объект контроллер crud операций
var Crud_manual_ChannelProd ICrud_manual_ChannelProd

// интерфейс CRUD операций сделанных вручную, для использования в DB или GRPC или NRPC
type ICrud_manual_ChannelProd interface {
}

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m ChannelProd) SetCrudManualInterface(crud ICrud_manual_ChannelProd) {
	Crud_manual_ChannelProd = crud

	return
}
