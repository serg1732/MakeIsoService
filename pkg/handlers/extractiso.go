package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kdomanski/iso9660/util"
	"github.com/serg1732/SkeletService/pkg/loggers"
)

func ExtractIso(logger loggers.ILogger) func(c *gin.Context) {
	return func(c *gin.Context) {
		uuid := c.Request.URL.Query().Get("uuid")
		paths := strings.Split(c.Request.URL.Query().Get("paths"), ",")
		logger.Progress("Задача запущена!", uuid)
		f, err := os.Open(paths[0])
		if err != nil {
			logger.Error(fmt.Sprintf("failed to open file: %s", err), uuid)
		}
		defer f.Close()

		if err = util.ExtractImageToDirectory(f, paths[1]); err != nil {
			logger.Error(fmt.Sprintf("failed to extract image: %s", err), uuid)
		}
		logger.Progress("Задача завершена!", uuid)
	}
}
