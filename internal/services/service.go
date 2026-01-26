package services

type storageInterface interface {
	productStorage
	userStorage
	cartStorage
}

type service struct {
	storage storageInterface
}

func NewServices(db storageInterface) *service {
	return &service{storage: db}
}
