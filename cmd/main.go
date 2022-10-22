package main

import (
	"log"

	"github.com/kirill0909/resource-keeper-api"
	"github.com/kirill0909/resource-keeper-api/pkg/handler"
)

func main() {

	srv := new(server.Server)
	handler := new(handler.Handler)

	// Method InitRoutes returns object *gin.Engine and we can use this object
	// as second parametr in method Run because gin.Engine
	// implements method ServeHTTP(ResponseWriter, *Request) from Handler struct
	// from http package
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
