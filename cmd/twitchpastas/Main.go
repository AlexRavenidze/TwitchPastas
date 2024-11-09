package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexRavenidze/TwitchPastas/internal/config"
	"github.com/AlexRavenidze/TwitchPastas/internal/controller"
	"github.com/AlexRavenidze/TwitchPastas/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	var cfg config.Configuration
	if err = envconfig.Process("", &cfg); err != nil {
		log.Println(err)
	}
	logger, err := logger.InitLogger(cfg.AppConfig.LogLevel)
	if err != nil {
		log.Println(err)
	}

	fiberConfig := fiber.Config{
		AppName:      cfg.AppConfig.AppName,
		BodyLimit:    cfg.AppConfig.BodyLimit,
		WriteTimeout: cfg.AppConfig.WriteTimeout,
		ReadTimeout:  cfg.AppConfig.ReadTimeout,
		IdleTimeout:  cfg.AppConfig.IdleTimeout,
	}
	apiServer := fiber.New(fiberConfig)
	handler := controller.New(apiServer, logger)
	handler.InitRouter()

	go func() {
		if err = apiServer.Listen(":" + cfg.AppConfig.Port); err != nil {
			logger.Error("server listen error", zap.Error(err))
			return
		}
	}()
	defer func() {
		err = apiServer.Shutdown()
		if err != nil {
			logger.Error("server shut down error", zap.Error(err))
			return
		}
	}()
	logger.Info("Application started", zap.String("name", cfg.AppConfig.AppName),
		zap.String("listen", cfg.AppConfig.Port))

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, syscall.SIGTERM, syscall.SIGINT)
	<-shutdownSignals

}
