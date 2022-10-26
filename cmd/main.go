package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kirill0909/resource-keeper-api"
	"github.com/kirill0909/resource-keeper-api/pkg/handler"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
	"github.com/kirill0909/resource-keeper-api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		logrus.Fatalf("faild to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	exitChannel := make(chan int)
	go func() {
		for {
			s := <-signalChannel
			switch s {
			case syscall.SIGINT:
				logrus.Println("The interrupt signal has been triggered")
				exitChannel <- 0
			case syscall.SIGTERM:
				logrus.Println("The terminte signal has been triggered")
				exitChannel <- 0
			case syscall.SIGHUP:
				logrus.Println("The hung up signal has been triggered")
				exitChannel <- 0
			default:
				logrus.Println("Unknown signal")
				exitChannel <- 1
			}
		}
	}()

	exitCode := <-exitChannel
	defer os.Exit(exitCode)

	srv := new(server.Server)
	// Method InitRoutes returns object *gin.Engine and we can use this object
	// as second parametr in method Run because gin.Engine
	// implements method ServeHTTP(ResponseWriter, *Request) from Handler struct
	// from http package
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured while on server shuting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func initDB() (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
