package app

import (
	"context"
	"go-shop-restful/internal/config"
	"go-shop-restful/internal/handler"
	"go-shop-restful/internal/repository"
	"go-shop-restful/internal/repository/postgres"
	"go-shop-restful/internal/router"
	"go-shop-restful/internal/service"
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
		log.Fatal("Ошибка инициализации конфигурации:", err)
	}

	logger, err := util.NewLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	defer logger.Sync()

	util.InitJWT(cfg.JWTSecret)

	db, err := postgres.NewStorage(cfg)
	if err != nil {
		logger.Fatalf("Ошибка инициализации хранилища: %v", err)
	}

	cartRepo := repository.NewCartRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	cartService := service.NewCartService(cartRepo, userRepo, productRepo)
	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo, cartService)

	cartHandler := handler.NewCartHandler(cartService, logger)
	productHandler := handler.NewProductHandler(productService, logger)
	userHandler := handler.NewUserHandler(userService, logger)

	if err := userService.CreateAdminIfNotExists(cfg.Admin.Name, cfg.Admin.Email, cfg.Admin.Password); err != nil {
		logger.Errorf("Ошибка создания аккаунта администратора: %v", err)
	}

	r, err := router.NewRouter(cfg, productHandler, userHandler, cartHandler)
	if err != nil {
		logger.Fatalln("Ошибка создания сервера:", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		logger.Infoln("Остановка сервера...")
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf("Ошибка graceful shutdown: %v", err)
			return
		}
		logger.Infoln("Сервер успешно завершил работу")
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalln("Ошибка запуска сервера:", err)
	}
}
