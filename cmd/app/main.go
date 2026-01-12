package main

import (
	"golang-shop-restful/internal/app"
	"golang-shop-restful/internal/config"
	"golang-shop-restful/internal/handlers"
	"golang-shop-restful/internal/services"
	"golang-shop-restful/internal/storage/postgres"
	"golang-shop-restful/internal/utils"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := utils.InitLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	defer logger.Sync()

	utils.InitJWT(cfg.JWTSecret)

	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	storage := postgres.NewStorage(db)
	service := services.NewServices(storage)
	handler := handlers.NewHandler(service, logger)

	if err := service.CreateAdminIfNotExists(cfg.AdminName, cfg.AdminEmail, cfg.AdminPassword); err != nil {
		logger.Errorf("Ошибка создания аккаунта администратора: %v", err)
	}

	log.Fatal(app.Router(cfg, handler))
}
