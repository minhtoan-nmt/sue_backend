package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"sue_backend/config"
	"sue_backend/internal/common/middleware"
	"sue_backend/internal/domain/repository"
	"sue_backend/internal/domain/service"
	"sue_backend/internal/infra/auth"
	"sue_backend/internal/infra/cache"
	"sue_backend/internal/infra/db"
	"sue_backend/internal/transport/http/route"
	"sue_backend/pkg/logger"
)

func main() {
	ctx := context.Background()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg, err := config.LoadAllConfigs("config")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.New(cfg.Log)
	logger.Info("⇢ initializing backend...")

	dbConn, err := db.InitDB(ctx, cfg.DB)
	if err != nil {
		logger.Fatalf("failed to connect to Postgres: %v", err)
	}
	defer dbConn.Close()
	pgStore := db.NewPostgresStore(dbConn)

	redisClient, err := cache.InitRedis(cfg.Redis)
	if err != nil {
		logger.Fatalf("failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()
	redisStore := cache.NewRedisStore(redisClient)

	// Init Repositories & Services
	userRepo := repository.NewUserRepository(pgStore, redisStore)
	// print cfg.Auth.JWTSecret
	logger.Infof("⇢ JWT Secret: %s", cfg.Auth.JWTSecret, cfg.Auth.JWTExpire)
	jwtManager := auth.NewJWTManager(cfg.Auth.JWTSecret, time.Duration(cfg.Auth.JWTExpire))
	authService := service.NewAuthService(userRepo, jwtManager)

	// Init Gin engine & middleware
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.CORS(), middleware.RequestID())

	api := r.Group("/api/v0")

	// public
	rPublic := api.Group("")
	route.RegisterAuthRoutes(rPublic, authService)

	// protected
	authGroup := api.Group("")
	authGroup.Use(middleware.JWTAuth(cfg.Auth.JWTSecret))

	userService := service.NewUserService(userRepo)
	route.RegisterUserRoutes(authGroup, userService)

	courseRepo := repository.NewCourseRepository(pgStore, redisStore)
	courseService := service.NewCourseService(courseRepo)
	route.RegisterCourseRoutes(authGroup, courseService)

	courseTemplateRepo := repository.NewCourseTemplateRepository(pgStore, redisStore)
	courseTemplateService := service.NewCourseTemplateService(courseTemplateRepo)
	route.RegisterCourseTemplateRoutes(authGroup, courseTemplateService)
	// Start server
	logger.Infof("⇢ starting HTTP server on %s", cfg.App.HostPort)
	if err := r.Run(cfg.App.HostPort); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
