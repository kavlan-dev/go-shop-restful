package main

import (
	"context"
	"go-shop-restful/internal/app"
	"go-shop-restful/internal/config"
	"go-shop-restful/internal/handlers"
	"go-shop-restful/internal/services"
	"go-shop-restful/internal/storage/postgres"
	"go-shop-restful/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	server, err := app.Router(cfg, handler)
	if err != nil {
		logger.Fatalln("Ошибка создания сервера:", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalln("Ошибка запуска сервер:", err)
		}
	}()

	<-sigCh
	logger.Infoln("Получен сигнал завершения, начинаем graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("Ошибка graceful shutdown: %v", err)
		return
	}

	logger.Infoln("Сервер успешно завершил работу")
}
