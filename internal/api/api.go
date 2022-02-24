package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zknow/my-arch/internal/errmsg"
	"github.com/zknow/my-arch/internal/response"
)

func Home(c *gin.Context) {
	response.Response(c, errmsg.SUCCESS)
}
