package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kdomanski/iso9660"
	"github.com/serg1732/SkeletService/pkg/loggers"
)

func CreateIso(logger loggers.ILogger) func(c *gin.Context) {
	return func(c *gin.Context) {
		uuid := c.Request.URL.Query().Get("uuid")
		paths := c.Request.URL.Query().Get("paths")
		logger.Progress("Задача запущена!", uuid)
		writer, err := iso9660.NewWriter()
		if err != nil {
			logger.Error(fmt.Sprintf("failed to create writer: %s", err), uuid)
		}
		defer writer.Cleanup()
		for _, path := range strings.Split(paths, ",") {
			f, err := os.Open(path)
			if err != nil {
				logger.Error(fmt.Sprintf("failed to open file: %s", err), uuid)
			}
			defer f.Close()

			err = writer.AddFile(f, path)
			if err != nil {
				logger.Error(fmt.Sprintf("failed to add file: %s", err), uuid)
			}
		}

		outputFile, err := os.OpenFile("/home/zed/output.iso", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to create file: %s", err), uuid)
		}

		err = writer.WriteTo(outputFile, "testvol")
		if err != nil {
			logger.Error(fmt.Sprintf("failed to write ISO image: %s", err), uuid)
		}

		err = outputFile.Close()
		if err != nil {
			logger.Error(fmt.Sprintf("failed to close output file: %s", err), uuid)
		}

		logger.Progress("Задача завершена!", uuid)
	}
}
