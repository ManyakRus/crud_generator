The service implements synchronous data exchange.

Service for exchanging data with Postgres SQL database
Data exchange is done using different methods:
1. Commands in NATS for reading, changing, etc.
2. DB CRUD operations - direct exchange with the database
    (each table model has methods Read(), Update(), Create(), Save(), Delete(), Restore())
3. GRPC - exchange with the database using the GRPC protocol
    (the client service that needs to exchange with the database connects to the sync_exchange server service, the latter exchanges with the database)
4. NRPC - exchange with the database using the NRPC protocol
    (the client service that needs to exchange with the database connects to the NATS service, which sends commands to the sync_exchange server, the latter exchanges with the database)

Before starting CRUD operations, you must specify the transport over which the exchange will take place (CRUD, GRPC, NRPC)
using one of the commands:
InitCrudTransport_DB()
InitCrudTransport_GRPC()
InitCrudTransport_NRPC()
from module
"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/crud_starter"

Also, to get started, the environment variables must be filled in:

1) for DB CRUD:
    DB_HOST="10.1.9.23"
    DB_NAME="claim"
    DB_SCHEME="public"
    DB_PORT="5432"
    DB_USER=""
    DB_PASSWORD=""

2) for GRPC:
    SYNC_SERVICE_HOST=10.1.9.150
    SYNC_SERVICE_PORT=30031

3) for NRPC:
    BUS_LOCAL_HOST="10.1.9.150"
    BUS_LOCAL_PORT=30222

For NRPC (GRPC) it is advisable to connect there first and disconnect at the end
nrpc_client.Connect()
defer nrpc_client.CloseConnection()
otherwise, the code will still connect there and will not disconnect at the end of the microservice.