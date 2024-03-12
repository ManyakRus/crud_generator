package constants

const SERVICE_NAME = "crud_generator"

const TEXT_HELP = `
Need fill settings in settings.txt file
`

//const FolderTemplates string = "templates"
//
//const FolderReady string = "ready"

const FILE_PERMISSIONS = 0666

var TEMPLATES_FOLDER_NAME = "templates"

const CONFIG_FOLDER_NAME = "configs_"

const TemplateFilenameCrudGo = "crud.go_"
const TemplateFilenameCrudGoTest = "crud_test.go_"

//const SETTINGS_FOLDER_NAME = "templates/configs_"

const GENERATION_PROTO_FILENAME = "generation_code.sh"

const GRPC_CLIENT_FILENAME = "grpc_client.go"
const GRPC_CLIENT_TEST_FILENAME = "grpc_client_test.go"

const NRPC_CLIENT_FILENAME = "nrpc_client.go"
const NRPC_CLIENT_TEST_FILENAME = "nrpc_client_test.go"

const NRPC_CLIENT_TABLE_FILENAME = "nrpc_client_table.go"
const NRPC_CLIENT_TABLE_TEST_FILENAME = "nrpc_client_table_test.go"

const SERVER_GRPC_STARTER_FILENAME = "server_grpc_starter.go"
const SERVER_GRPC_FUNC_FILENAME = "server_grpc_func.go"

const MAKEFILE_FILENAME = "Makefile"
const ENV_FILENAME = ".env"

const STARTER_TABLES_FILENAME = "starter_tables.go_"
const STARTER_TABLES_TEST_FILENAME = "starter_tables_test.go_"
const STARTER_TABLES_MANUAL_FILENAME = "starter_tables_manual.go_"
const STARTER_TABLES_TEST_MANUAL_FILENAME = "starter_tables_manual_test.go_"
const STARTER_TABLES_PREFIX = "crud_starter_"
const CRUD_TABLES_FREFIX = "crud_"

const MODEL_TABLE_MANUAL_FILENAME = "model_table_manual.go_"
const MODEL_TABLE_UPDATE_FILENAME = "model_table_update.go_"

const SERVER_GRPC_TABLE_UPDATE_FUNC_FILENAME = "server_grpc_table_update_func.go_"
const SERVER_GRPC_TABLE_UPDATE_FUNC_TEST_FILENAME = "server_grpc_table_update_func_test.go_"

const GRPC_CLIENT_TABLE_UPDATE_FUNC_FILENAME = "grpc_client_table_update_func.go_"
const GRPC_CLIENT_TABLE_UPDATE_FUNC_TEST_FILENAME = "grpc_client_table_update_func_test.go_"

const CRUD_TABLES_CACHE_FILENAME = "crud_table_cache.go_"
const CRUD_TABLES_CACHE_TEST_FILENAME = "crud_table_cache_test.go_"
const TEXT_CACHE_REMOVE = "cache.Remove(int64(m.ID))"

const SERVER_GRPC_TABLE_CACHE_FILENAME = "server_grpc_table_cache.go_"
const SERVER_GRPC_TABLE_CACHE_TEST_FILENAME = "server_grpc_table_cache_test.go_"

const TEXT_OTVET_ID_ALIAS = "Otvet.ID = ID"
