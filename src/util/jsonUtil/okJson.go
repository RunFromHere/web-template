package jsonUtil

import "github.com/gin-gonic/gin"

//操作正确的时候快速拼凑json响应请求
func OkJsonDefault(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    msg,
		"data":   nil,
	})
}
