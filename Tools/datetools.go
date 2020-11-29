package Tools

import "time"

func DateFormat(timestamp int64) string {
	tm := time.Unix(timestamp,0)
	dates := tm.Format("2006-01-02 15:04:05")
	return dates
}