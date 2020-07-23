package jsonUtil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResult struct {
	Status  int      `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

//NoResponse 请求的url不存在，返回404
func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"msg":    "404, page not exists!",
	})
}

//报错的时候快速拼凑json响应请求
func ErrJsonDefault(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status": 1,
		"msg": msg,
		"data": nil,
	})
}
