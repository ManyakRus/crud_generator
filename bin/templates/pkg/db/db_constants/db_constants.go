package db_constants

import (
	"errors"
)

const CONNECTION_ID_TEST = 3

const TIMEOUT_DB_SECONDS = 30

// TEXT_RECORD_NOT_FOUND - текст ошибки "record not found" для gorm
const TEXT_RECORD_NOT_FOUND = "record not found"

// TEXT_NO_ROWS - текст ошибки "no rows in result set" для pgx
const TEXT_NO_ROWS = "no rows in result set"

const TextCrudIsNotInit = "Need initializate crud with InitCrudTransport_DB() function at first."

var ErrorCrudIsNotInit error = errors.New(TextCrudIsNotInit)
