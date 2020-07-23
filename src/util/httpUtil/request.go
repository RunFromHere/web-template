package httpUtil

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"net/url"
)

//post请求
//json数据
func PostRequestByJson(requestUrl string, data interface{}) (*http.Response, error) {
	client := http.Client{}
	jsonData, err := jsoniter.Marshal(data)
	if err != nil {
		log.Println("requestByJson err:", err)
		return nil, err
	}
	//请求内容
	request, err := http.NewRequest("POST", requestUrl, bytes.NewReader(jsonData))
	if err != nil {
		log.Println("requestByJson err:", err)
		return nil, err
	}
	//请求内容类型为json
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	log.Println("requestByJson request:", request)
	//发送请求，接收返回内容
	response, err := client.Do(request)
	if err != nil {
		log.Println("requestByJson err:", err)
		return nil, err
	}
	return response, nil
}

//delete请求
//json数据
func DeleteRequestByJson(requestUrl string, data interface{}) (*http.Response, error) {
	client := http.Client{}
	jsonData, err := jsoniter.Marshal(data)
	if err != nil {
		log.Println("DeleteRequestByJson err:", err)
		return nil, err
	}
	//请求内容
	request, err := http.NewRequest("DELETE", url.QueryEscape(requestUrl), bytes.NewReader(jsonData))
	if err != nil {
		log.Println("DeleteRequestByJson err:", err)
		return nil, err
	}
	//请求内容类型为json
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	log.Println("DeleteRequestByJson request:", request)
	//发送请求，接收返回内容
	response, err := client.Do(request)
	if err != nil {
		log.Println("DeleteRequestByJson err:", err)
		return nil, err
	}
	return response, nil
}

//普通请求
func GetRequest(requestUrl string) (*http.Response, error) {
	client := http.Client{}
	//请求内容
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println("GetRequest err:", err)
		return nil, err
	}
	log.Println("GetRequest request:", request)
	//发送请求，接收返回内容
	response, err := client.Do(request)
	if err != nil {
		log.Println("GetRequest err:", err)
		return nil, err
	}
	return response, nil
}

