package handlers

type ServicesInterface interface {
	ProductService
	UserService
	CartService
}

type Handler struct {
	service ServicesInterface
}

func NewHandler(service ServicesInterface) *Handler {
	return &Handler{service: service}
}
