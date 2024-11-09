package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler interface {
	InitRouter()
}

type handler struct {
	fiberService *fiber.App
	logger       *zap.Logger
}

func New(fiber *fiber.App, logger *zap.Logger) Handler {
	return &handler{
		fiberService: fiber,
		logger:       logger,
	}
}
