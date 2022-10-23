package main

import (
	"log"

	"github.com/kirill0909/resource-keeper-api"
	"github.com/kirill0909/resource-keeper-api/pkg/handler"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
	"github.com/kirill0909/resource-keeper-api/pkg/service"
)

func main() {

	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	// Method InitRoutes returns object *gin.Engine and we can use this object
	// as second parametr in method Run because gin.Engine
	// implements method ServeHTTP(ResponseWriter, *Request) from Handler struct
	// from http package
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
