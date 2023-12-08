package main

import (
	"context"
	"os"
	"os/signal"
	"sportnotes"
	"sportnotes/configs"
	"sportnotes/pkg/handler"
	"sportnotes/pkg/repository"
	"sportnotes/pkg/service"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	var cfg configs.Config
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := cleanenv.ReadConfig("configs/config.yml", &cfg); err != nil {
		logrus.Fatalf("error loading config file:%s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables:%s", err.Error())
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("database initializing failed:%s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(sportnotes.Server)
	go func() {
		if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Server start error: %s", err.Error())
		}
	}()

	logrus.Print("SportNotes started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("SportNotes sutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down:%s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close:%s", err.Error())
	}
}
