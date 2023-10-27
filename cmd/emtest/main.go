package main

import (
	emtest "em-test"
	"em-test/internal/handler"
	"em-test/internal/repository"
	"em-test/internal/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		logrus.Fatalf("Config init error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Loading env variables error: %s", err.Error())
	}
	logrus.Debug("Config init complete")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Usename:  viper.GetString("db.usename"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("DB init fail: %s", err.Error())
	}
	logrus.Info("DB init complete")

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(emtest.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Running error: %s", err.Error())
	}
	logrus.Info("Server start")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
