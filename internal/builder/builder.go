package builder

import (
	"Evermos-Virtual-Intern/config"
	"Evermos-Virtual-Intern/internal/http/handler"
	"Evermos-Virtual-Intern/internal/http/router"
	"Evermos-Virtual-Intern/internal/repository"
	"Evermos-Virtual-Intern/internal/service"

	"gorm.io/gorm"
)

func BuildPrivateRoutes(db *gorm.DB, cfg *config.Config) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg, userRepository)
	userHandler := handler.NewUserHandler(userService)


	alamatRepository := repository.NewAlamatRepository(db)
	alamatService := service.NewAlamatService(cfg, alamatRepository)
	alamatHandler := handler.NewAlamatHandler(alamatService)


	return router.PrivateRoutes(userHandler, alamatHandler)
}

func BuildPublicRoutes(db *gorm.DB, cfg *config.Config) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	

	// Initialize services
	
	authService := service.NewAuthService(cfg, userRepository)
	userService := service.NewUserService(cfg, userRepository)


	// Initialize handlers

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	return router.PublicRoutes(authHandler, userHandler)
}