package format_time

import (
	"time"
)

func FormatTime(date time.Time) string {
	return date.Format("02.01.2006 15:04:05")
}
