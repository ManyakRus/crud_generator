package grpc_client

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/stopapp"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strings"
)

type SettingsINI struct {
	SYNC_SERVICE_HOST string
	SYNC_SERVICE_PORT string
}

var Settings SettingsINI

var Conn *grpc.ClientConn
var Client grpc_proto.SyncServiceClient

func Connect() {
	var err error

	if Settings.SYNC_SERVICE_HOST == "" {
		FillSettings()
	}

	addr := Settings.SYNC_SERVICE_HOST + ":" + Settings.SYNC_SERVICE_PORT
	Conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Info("GRPC client connected. Address: ", addr)

	Client = grpc_proto.NewSyncServiceClient(Conn)
}

func FillSettings() {
	Settings = SettingsINI{}
	Settings.SYNC_SERVICE_HOST = os.Getenv("SYNC_SERVICE_HOST")
	Settings.SYNC_SERVICE_PORT = os.Getenv("SYNC_SERVICE_PORT")

	if Settings.SYNC_SERVICE_HOST == "" {
		log.Panic("Need fill SYNC_SERVICE_HOST ! in OS Environment ")
	}

	if Settings.SYNC_SERVICE_PORT == "" {
		log.Panic("Need fill SYNC_SERVICE_PORT ! in OS Environment ")
	}
}

// WaitStop - ожидает отмену глобального контекста
func WaitStop() {

	select {
	case <-contextmain.GetContext().Done():
		log.Warn("Context app is canceled. grpc_connect")
	}

	// ждём пока отправляемых сейчас сообщений будет =0
	stopapp.WaitTotalMessagesSendingNow("sync_service_client")

	// закрываем соединение
	CloseConnection()
	stopapp.GetWaitGroup_Main().Done()
}

// Start - необходимые процедуры для запуска сервера GRPC
func Start() {
	Connect()

	stopapp.GetWaitGroup_Main().Add(1)
	go WaitStop()

}

func CloseConnection() {
	err := Conn.Close()
	if err != nil {
		log.Panic("GRPC client CloseConnection() error: ", err)
	} else {
		log.Info("GRPC client connection closed")
	}
}

// IsRecordNotFound - возвращает true если ошибка = "record not found"
func IsRecordNotFound(err error) bool {
	Otvet := false

	if err == nil {
		return Otvet
	}

	TextErr := err.Error()
	pos1 := strings.Index(TextErr, constants.TEXT_RECORD_NOT_FOUND)
	if pos1 >= 0 {
		Otvet = true
	}

	return Otvet
}
