package grpc_constants

import "sync"

// timeout_seconds - время ожидания ответа
var timeout_seconds int = 30

// TEXT_ERROR_MODEL_VERSION - текст ошибки версии модели
const TEXT_ERROR_MODEL_VERSION = "Error: wrong version object model"

// mutex_TIMEOUT_SECONDS - защита от многопоточности GetTimeoutSeconds()
var mutex_TIMEOUT_SECONDS sync.RWMutex

// GetTimeoutSeconds - возвращает время ожидания ответа
func GetTimeoutSeconds() int {
	mutex_TIMEOUT_SECONDS.RLock()
	defer mutex_TIMEOUT_SECONDS.RUnlock()

	return timeout_seconds
}

// SetTimeoutSeconds - устанавливает время ожидания ответа
func SetTimeoutSeconds(seconds int) {
	mutex_TIMEOUT_SECONDS.Lock()
	defer mutex_TIMEOUT_SECONDS.Unlock()

	timeout_seconds = seconds
}
