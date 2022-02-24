package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zknow/my-arch/internal/errmsg"
)

func Response(c *gin.Context, status int) {
	var context = gin.H{
		"status":  status,
		"message": errmsg.GetErrMsg(status),
	}
	c.JSON(
		http.StatusOK,
		context,
	)
}

func ResponseWithData(c *gin.Context, status int, data map[string]interface{}) {
	var context = gin.H{
		"status":  status,
		"message": errmsg.GetErrMsg(status),
	}
	for key, value := range data {
		context[key] = value
	}
	c.JSON(
		http.StatusOK,
		context,
	)
}
