package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"transactions/config"
	"transactions/internal/db"
	"transactions/internal/handler"
	"transactions/internal/repository"
	"transactions/internal/service"
)

func main() {
	config.ReadConfig("./config/config.json")

	dbc := db.InitToDatabase()

	repositoryConnection := repository.NewRepo(dbc)

	serviceConnection := service.NewService(repositoryConnection)

	r := gin.Default()

	handle := handler.NewHandler(serviceConnection, r)

	handle.AllRoutes()

	err := r.Run(config.Configure.PortRun)
	if err != nil {
		log.Fatal("router failed to start")
		return
	}
}
