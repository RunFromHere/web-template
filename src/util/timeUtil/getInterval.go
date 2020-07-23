package timeUtil

import "time"

func GetIntervalFromNow(unixTime int64) int64 {
	//获取时区
	loc, _ := time.LoadLocation("Local")
	time.Now().In(loc).Unix()

	rest := time.Now().In(loc).Unix() - unixTime

	return rest
}
