package server_nrpc

import (
	config_main "github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestConnect(t *testing.T) {
	config_main.LoadEnv()
	FillSettings()

	Connect()
	CloseConnection()
}
