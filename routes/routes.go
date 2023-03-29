package routes

import (
	"gin_project_template/config"
	"gin_project_template/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,gin!!!")
	})
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, config.Conf.Version)
	})
	return r
}