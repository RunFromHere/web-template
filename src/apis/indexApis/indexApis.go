package indexApis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultIndexApi(c *gin.Context) {
	//something to do

	//return json
	c.String(http.StatusOK, "hello! welcome to koala.")
}
