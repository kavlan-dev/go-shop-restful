package service

type storageInterface interface {
	productStorage
	userStorage
	cartStorage
}

type service struct {
	storage storageInterface
}

func NewService(db storageInterface) *service {
	return &service{storage: db}
}
