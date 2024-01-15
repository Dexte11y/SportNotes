package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sportnotes/configs"
	sportnotes "sportnotes/internal/app"
	"sportnotes/internal/handler"
	"sportnotes/internal/repository"
	"sportnotes/internal/service"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus.SetFormatter(new(logrus.JSONFormatter))
	err := configs.LoadEnvironment()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := configs.ReadConfig()

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
