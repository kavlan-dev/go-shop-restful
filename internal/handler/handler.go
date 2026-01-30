package handler

import "go.uber.org/zap"

type serviceInterface interface {
	productService
	userService
	cartService
}

type handler struct {
	service serviceInterface
	log     *zap.SugaredLogger
}

func NewHandler(service serviceInterface, log *zap.SugaredLogger) *handler {
	return &handler{service: service, log: log}
}
