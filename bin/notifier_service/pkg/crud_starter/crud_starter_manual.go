//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_attachament"

	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_attachament"

	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_attachament"
)

// initCrudTransport_manual_DB - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_DB() {
	crud_starter_attachament.SetCrudManualInterface(crud_attachament.Crud_DB{})
}

// initCrudTransport_manual_GRPC - заполняет объекты crud для работы с БД через протокол GRPC
func initCrudTransport_manual_GRPC() {
	crud_starter_attachament.SetCrudManualInterface(grpc_attachament.Crud_GRPC{})
}

// initCrudTransport_manual_NRPC - заполняет объекты crud для работы с БД через протокол NRPC
func initCrudTransport_manual_NRPC() {
	crud_starter_attachament.SetCrudManualInterface(grpc_attachament.Crud_GRPC{})
}