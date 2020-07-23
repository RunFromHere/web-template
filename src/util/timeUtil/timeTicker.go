package timeUtil

import (
	"log"
	"time"
)

//定时器
func timeHandlerForMonitorClusterHealth(duration time.Duration) {
	log.Println("timeHandlerForMonitorClusterHealth start... ")
	//定时器时间间隔
	ticker := time.NewTicker(duration)
	var nowTime time.Time
	for {
		//插入一个阻断，新阻断插入的前提是前一个阻断已结束
		nowTime = <-ticker.C
		log.Println("timeHandlerForMonitorClusterHealth Ticket Time: ", nowTime)

		//run sth.
		//go monitorClusterHealth()
	}
}
