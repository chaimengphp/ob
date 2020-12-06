package Tools

import (
	"fmt"
	"time"
)

func DateFormat(timestamp int64) string {
	tm := time.Unix(timestamp,0)
	dates := tm.Format("2006-01-02 15:04:05")
	return dates
}

func YmDateFormat(timestamp int64) string{
	tm := time.Unix(timestamp,0)
	return fmt.Sprintf("%d年%d日",tm.Year(),tm.Month())
}

func MdTimeFormat(timestamp int64) string{
	tm := time.Unix(timestamp,0)
	return tm.Format("01月02 15:04")
}
