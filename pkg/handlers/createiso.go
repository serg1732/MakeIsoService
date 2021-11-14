package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/serg1732/SkeletService/pkg/loggers"
)

func CreateIso(logger loggers.ILogger) func(c *gin.Context) {
	return func(c *gin.Context) {
		uuid := c.Request.URL.Query().Get("uuid")
		logger.Progress("Задача запущена!", uuid)
	}
}
