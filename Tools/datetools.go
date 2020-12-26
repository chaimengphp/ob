package Tools

import (
	"fmt"
	"math"
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

func DFormat(timestamp int64) string{
	tm := time.Unix(timestamp,0)
	//今天最大时间
	now_timestamp := time.Now().Unix()
	now_tm := time.Unix(now_timestamp,0)
	todayLast := now_tm.Format("2006-01-02")+" 23:59:59"
	todayLast_time,_ := time.Parse("2006-01-02 15:04:05",todayLast)
	todayLast_timestamp := todayLast_time.Unix()

	agoTimeTrue := now_timestamp-timestamp
	agoTime := todayLast_timestamp-timestamp

	agoTimeTrue_float := agoTime/86400
	agoDay := math.Floor(float64(agoTimeTrue_float))
	//fmt.Println(now_timestamp,"----",timestamp)
	if agoTimeTrue <60 {
		return "刚刚"
	}else if agoTimeTrue<3600 {
		itime:=agoTimeTrue/60
		return fmt.Sprintf("%v分钟前",math.Ceil(float64(itime)))
	}else if agoTimeTrue < 3600 * 12 {
		htime:=agoTimeTrue/3600
		return fmt.Sprintf("%v小时前",math.Ceil(float64(htime)))
	}else if agoDay == 0 {
		return fmt.Sprintf("今天%v",tm.Format("15:04"))
	}else if agoDay == 1 {
		return fmt.Sprintf("昨天%v",tm.Format("15:04"))
	}else if agoDay == 2 {
		return fmt.Sprintf("前天%v",tm.Format("15:04"))
	}else if agoDay > 2 && agoDay < 16 {
		return fmt.Sprintf("%v天前 %s",agoDay,tm.Format("15:04"))
	}else {
		return fmt.Sprintf("%s",tm.Format("2006-01-02 15:04"))
	}
}
