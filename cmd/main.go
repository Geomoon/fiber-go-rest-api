package main

import (
	"accounts-api/api"
	"accounts-api/datasource"
	"accounts-api/pkg/auth"
	"accounts-api/pkg/user"
	"accounts-api/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	fmt.Println("Hello")

	var repository user.Repository
	var service user.Service
	var encryptService *services.PasswordEncrypt
	var authService auth.Service

	repository = datasource.NewMKVDatabase()
	service = services.NewUserService(repository)
	encryptService = services.NewPasswordEncrypt(14)
	authService = services.NewAuthService(&repository, encryptService)

	app := fiber.New()

	apipath := app.Group("/api")

	v1 := apipath.Group("/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})

	api.UserRouter(v1, service)
	api.AuthRouter(v1, authService)

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal("Error listen")
		return
	}
	log.Println("Ready")

}
