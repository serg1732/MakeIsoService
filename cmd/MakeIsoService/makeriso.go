package main

import (
	"github.com/serg1732/MakeIsoService/pkg/handlers"

	"github.com/serg1732/SkeletService/pkg/loggers"
	"github.com/serg1732/SkeletService/pkg/skeletservice"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	cLogger := loggers.NewConsoleLogger()

	router.GET("/makeiso", handlers.CreateIso(cLogger))

	service := skeletservice.NewService(router, cLogger)
	service.Start()
}
