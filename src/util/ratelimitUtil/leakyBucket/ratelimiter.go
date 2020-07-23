package leakyBucket

import (
	"bces-loadbalancing/src/log"
	"bces-loadbalancing/src/util/appcfg"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"time"
)

var rateLimiter ratelimit.Limiter
var prev time.Time

func InitRateLimiterOfLeakyBucket() {
	var rate = appcfg.GetRequestRateLimit()
	log.Info("rate init: ", zap.Int("rate", rate))
	rateLimiter = ratelimit.New(rate) // per second
	prev = time.Now()
}

func Take() {
	now := rateLimiter.Take()
	log.Debug("test", zap.Duration("timeSub", now.Sub(prev)))
	prev = now
}

func SetRateOfLeakyBucket(rate int) {
	rateLimiter = ratelimit.New(rate)
}
