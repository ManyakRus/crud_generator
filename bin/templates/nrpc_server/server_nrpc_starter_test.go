package server_grpc

import (
	config_main "gitlab.aescorp.ru/dsp_dev/claim/nikitin/config"
	"testing"
)

func TestConnect(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	Connect()
	CloseConnection()
}
