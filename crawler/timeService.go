package crawler

import (
	"time"
)

func TimeStrToTime(timeStr string) time.Time {
	//timeStr := "2018-01-01"
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}
	}
	t, err := time.ParseInLocation("2006-01-02", timeStr, location)
	if err != nil {
		panic("[error] 时间转换失败")
	}
	return t
}

func BeijingTime() time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = local
	return time.Now()
}

func StampToTime(timestamp int64) time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = local
	ts := time.Unix(timestamp, 0)
	return ts
}
