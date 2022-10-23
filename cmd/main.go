package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kirill0909/resource-keeper-api"
	"github.com/kirill0909/resource-keeper-api/pkg/handler"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
	"github.com/kirill0909/resource-keeper-api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("faild to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	// Method InitRoutes returns object *gin.Engine and we can use this object
	// as second parametr in method Run because gin.Engine
	// implements method ServeHTTP(ResponseWriter, *Request) from Handler struct
	// from http package
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
