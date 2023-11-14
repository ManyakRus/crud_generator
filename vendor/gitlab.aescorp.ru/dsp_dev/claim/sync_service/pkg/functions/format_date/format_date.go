package format_date

import (
	"time"
)

func FormatDate(date time.Time) string {
	return date.Format("02.01.2006")
}
