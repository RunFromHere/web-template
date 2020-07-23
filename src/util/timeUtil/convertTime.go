package timeUtil

import (
	"log"
	"strings"
	"time"
)

func ConvertMySqlTimeToUnixTime(datetime2 string) (*time.Time, error) {
	datetime2 = strings.Replace(datetime2[:len(datetime2)-6], "T", " ", 10)
	println(datetime2)
	//转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	//获取时区
	loc, err := time.LoadLocation("Local")
	if err != nil {
		log.Println("convertMySqlTimeToUnixTime method:", err)
		return nil, err
	}
	tmp, err := time.ParseInLocation(timeLayout, datetime2, loc)
	if err != nil {
		log.Println("convertMySqlTimeToUnixTime method:", err)
		return nil, err
	}

	return &tmp, nil
}
