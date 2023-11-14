package crud_starter

import (
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/db_employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/db_files"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/db_lawsuit_status_states"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/db_organization_casebooks"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/db_organizations"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/grpc_employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/grpc_files"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/grpc_lawsuit_status_states"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/grpc_organization_casebooks"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_client/grpc_organizations"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/nrpc_employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/nrpc_files"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/nrpc_lawsuit_status_states"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/nrpc_organization_casebooks"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/nrpc/nrpc_client/nrpc_organizations"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/employees"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/files"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/lawsuit_status_states"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organization_casebooks"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/object_model/entities/organizations"
)

// initCrudTransport_manual_DB - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_DB() {
	employees.Employee{}.SetCrudManualInterface(db_employees.Crud_DB{})
	files.File{}.SetCrudManualInterface(db_files.Crud_DB{})
	lawsuit_status_states.LawsuitStatusState{}.SetCrudManualInterface(db_lawsuit_status_states.Crud_DB{})
	organization_casebooks.OrganizationCasebook{}.SetCrudManualInterface(db_organization_casebooks.Crud_DB{})
	organizations.Organization{}.SetCrudManualInterface(db_organizations.Crud_DB{})
}

// initCrudTransport_manual_GRPC - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_GRPC() {
	employees.Employee{}.SetCrudManualInterface(grpc_employees.Crud_GRPC{})
	files.File{}.SetCrudManualInterface(grpc_files.Crud_GRPC{})
	lawsuit_status_states.LawsuitStatusState{}.SetCrudManualInterface(grpc_lawsuit_status_states.Crud_GRPC{})
	organization_casebooks.OrganizationCasebook{}.SetCrudManualInterface(grpc_organization_casebooks.Crud_GRPC{})
	organizations.Organization{}.SetCrudManualInterface(grpc_organizations.Crud_GRPC{})
}

// initCrudTransport_manual_NRPC - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_NRPC() {
	employees.Employee{}.SetCrudManualInterface(nrpc_employees.Crud_NRPC{})
	files.File{}.SetCrudManualInterface(nrpc_files.Crud_NRPC{})
	lawsuit_status_states.LawsuitStatusState{}.SetCrudManualInterface(nrpc_lawsuit_status_states.Crud_NRPC{})
	organization_casebooks.OrganizationCasebook{}.SetCrudManualInterface(nrpc_organization_casebooks.Crud_NRPC{})
	organizations.Organization{}.SetCrudManualInterface(nrpc_organizations.Crud_NRPC{})
}
