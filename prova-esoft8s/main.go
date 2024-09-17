package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"prova-esoft8s/config"
	"prova-esoft8s/controller"
	"prova-esoft8s/helper"

	"prova-esoft8s/model"
	"prova-esoft8s/repository"
	"prova-esoft8s/router"
	"prova-esoft8s/service"
	"time"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("viagens").AutoMigrate(&model.Viagens{})

	//Init Repository
	viagenRepository := repository.NewViagensRepositoryImpl(db)

	//Init Service
	viagenService := service.NewViagemServiceImpl(viagenRepository, validate)

	//Init controller
	viagenController := controller.NewViagemController(viagenService)

	//Router
	routes := router.NewRouter(viagenController)

	server := &http.Server{
		Addr:           ":6666",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
