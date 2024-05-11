package router

import (
	"auth-service/internal/server/rest/handler"
	"auth-service/internal/server/rest/middleware"
	"auth-service/internal/service"
	"auth-service/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(port int, services *service.Manager, store *storage.Storage) error {
	authHandler := handler.NewAuthHandler(services.Auth, services.Logger)

	g := gin.New()
	g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := g.Group("/api")
	api.Use(middleware.NewLogMiddleware(services.Logger))
	api.POST("sign-up", authHandler.SignUp)
	api.POST("sign-in", authHandler.SignIn)

	err := g.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}
