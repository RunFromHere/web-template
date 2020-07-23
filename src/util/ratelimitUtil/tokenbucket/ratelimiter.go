package tokenbucket

import (
	"bces-loadbalancing/src/log"
	"bces-loadbalancing/src/util/appcfg"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

var RateLimiter *rate.Limiter
var AllowAllGetRequest bool

func InitRateLimiterOfTokenBucket() {
	var requestRateLimit = appcfg.GetRequestRateLimit()
	var requestBucketNum = appcfg.GetRequestBucketNum()
	var rateLimitDeveloperLevel = appcfg.GetRateLimitDeveloperLevel()
	AllowAllGetRequest = appcfg.GetAllowAllGetRequest()
	if !rateLimitDeveloperLevel && (requestBucketNum != requestRateLimit) {
		requestBucketNum = requestRateLimit
	}
	log.Info("InitRateLimiter: ", zap.Int("requestRateLimit", requestRateLimit),
		zap.Int("requestBucketNum", requestBucketNum), zap.Bool("AllowAllGetRequest", AllowAllGetRequest),
		zap.Bool("rateLimitDeveloperLevel", rateLimitDeveloperLevel))
	RateLimiter = rate.NewLimiter(rate.Limit(requestRateLimit), requestBucketNum)
}

//func Wait()  {
//	//RateLimiter.Wait()
//}
//
//func Allow() {
//	if RateLimiter.Allow() {}
//}

//func SetLimitOfTokenBucket(limit float64) {
//	RateLimiter.SetLimit(rate.Limit(limit))
//}
//
//func SetBurstOfTokenBucket(burst int) {
//	RateLimiter.SetBurst(burst)
//}
