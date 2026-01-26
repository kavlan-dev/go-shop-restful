package handlers

import "go.uber.org/zap"

type serviceInterface interface {
	productService
	userService
	cartService
}

type Handler struct {
	service serviceInterface
	log     *zap.SugaredLogger
}

func NewHandler(service serviceInterface, log *zap.SugaredLogger) *Handler {
	return &Handler{service: service, log: log}
}
