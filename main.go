package main

import (
	exampleDelivery "arch/example/delivery"
	exampleRepo "arch/example/repository"
	exampleUsecase "arch/example/usecase"

	"arch/pkg/shutdown"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Fatal error config file: %v\n", err)
	}
}

func main() {
	serverPort := viper.GetString("Server_PORT")
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)
	checkError(err)

	err = db.Ping()
	checkError(err)

	route := gin.Default()

	repo := exampleRepo.NewExampleRepository(db)
	exUsecase := exampleUsecase.NewExampleRepoUsecase(repo)
	exampleDelivery.SetExampleHandler(route, exUsecase)

	srv := &http.Server{
		Addr:    ":" + serverPort,
		Handler: route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			srv.Shutdown(ctx)
		},
	)
}

func checkError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
