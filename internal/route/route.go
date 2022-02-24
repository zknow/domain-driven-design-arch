package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zknow/my-arch/internal/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", api.Home)
	return r
}
