package httpUtil

import "net/url"

func encodeUrl(urlR string) string {
	urlR = "http://10.20.20.20:9200/webapi/agent?Action=beat&a=2&c=a"
	return url.QueryEscape(urlR)
}
