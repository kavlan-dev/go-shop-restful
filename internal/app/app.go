package app

import (
	"context"
	"go-shop-restful/internal/config"
	"go-shop-restful/internal/handler"
	"go-shop-restful/internal/router"
	"go-shop-restful/internal/service"
	"go-shop-restful/internal/storage/postgres"
	"go-shop-restful/internal/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := util.NewLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	defer logger.Sync()

	util.InitJWT(cfg.JWTSecret)

	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	storage := postgres.NewStorage(db)
	service := service.NewService(storage)
	handler := handler.NewHandler(service, logger)

	if err := service.CreateAdminIfNotExists(cfg.AdminName, cfg.AdminEmail, cfg.AdminPassword); err != nil {
		logger.Errorf("Ошибка создания аккаунта администратора: %v", err)
	}

	server, err := router.Router(cfg, handler)
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
