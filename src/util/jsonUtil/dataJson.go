package jsonUtil

import "github.com/gin-gonic/gin"

func DataJsonDefault(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    msg,
		"data":   data,
	})
}


