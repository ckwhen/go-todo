package main

import (
	"database/sql"
	"fmt"

	_todoHandlerHttpDelivery "github.com/ckwhen/go-todo/internal/todo/delivery/http"
	_todoRepo "github.com/ckwhen/go-todo/internal/todo/repository/postgresql"
	_todoUsecase "github.com/ckwhen/go-todo/internal/todo/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config/env.yaml")
	viper.SetConfigType("yaml")

	if viper.GetBool("debug") {
		logrus.Info("Service RUN on DEBUG mode")
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Fatal error config file: %v\n", err)
	}
}

func main() {
	logrus.Info("HTTP server started")

	apiHost := viper.GetString("server.address")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt64("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Taipei",
			dbHost, dbPort, dbUser, dbPass, dbName,
		),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()
	api := r.Group("/api")

	v1 := api.Group("/v1")

	todoRepo := _todoRepo.NewPostgresqlTodoRepository(db)
	todoUsecase := _todoUsecase.NewTodoUsecase(todoRepo)

	_todoHandlerHttpDelivery.NewTodoHandler(v1.Group("/todos"), todoUsecase)

	fmt.Printf("Starting server at port 8000\n")
	logrus.Fatal(
		r.Run(apiHost),
	)
}
