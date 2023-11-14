package nrpc_client

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/nats-io/nats.go"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/db/constants"
	"gitlab.aescorp.ru/dsp_dev/claim/sync_service/pkg/network/grpc/grpc_proto"

	"os"
	"strings"
	"time"
)

type SettingsINI struct {
	NATS_HOST string
	NATS_PORT string
}

var Settings SettingsINI

// Conn - подключение к NATS
var Conn *nats.Conn

// Client - подключение к клиенту NRPC
var Client *grpc_proto.Sync_serviceClient

// Connect - подключается к NATS
func Connect() {
	var err error

	if Settings.NATS_HOST == "" {
		FillSettings()
	}

	NatsURL := "nats://" + Settings.NATS_HOST + ":" + Settings.NATS_PORT
	// Connect to the NATS server.
	Conn, err = nats.Connect(NatsURL, nats.Timeout(5*time.Second))
	if err != nil {
		log.Panic(err)
	}
	// defer Conn.Close()

	// This is our generated client.
	Client = grpc_proto.NewSync_serviceClient(Conn)

	log.Info("Client NRPC connected: ", NatsURL)
}

// FillSettings - заполняет настройки из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.NATS_HOST = os.Getenv("BUS_LOCAL_HOST")
	Settings.NATS_PORT = os.Getenv("BUS_LOCAL_PORT")

	if Settings.NATS_HOST == "" {
		log.Panic("Need fill BUS_LOCAL_HOST ! in OS Environment ")
	}

	if Settings.NATS_PORT == "" {
		log.Panic("Need fill BUS_LOCAL_PORT ! in OS Environment ")
	}
}

// WaitStop - ожидает отмену глобального контекста
func WaitStop() {

	select {
	case <-contextmain.GetContext().Done():
		log.Warn("Context app is canceled. nrpc client connect")
	}

	// ждём пока отправляемых сейчас сообщений будет =0
	stopapp.WaitTotalMessagesSendingNow("sync_service_client")

	// закрываем соединение
	CloseConnection()
	stopapp.GetWaitGroup_Main().Done()
}

// Start - необходимые процедуры для запуска сервера NRPC
func Start() {
	Connect()

	stopapp.GetWaitGroup_Main().Add(1)
	go WaitStop()

}

// CloseConnection - закрывает подключение к NATS
func CloseConnection() {
	Conn.Close()
	log.Info("NRPC client connection closed")
}

// IsRecordNotFound - возвращает true если ошибка = "record not found"
func IsRecordNotFound(err error) bool {
	Otvet := false

	if err == nil {
		return Otvet
	}

	// len1 := len(constants.TEXT_RECORD_NOT_FOUND)
	TextErr := err.Error()
	pos1 := strings.Index(TextErr, constants.TEXT_RECORD_NOT_FOUND)
	// if TextErr[0:len1] == constants.TEXT_RECORD_NOT_FOUND {
	if pos1 >= 0 {
		Otvet = true
	}

	return Otvet
}
