package api

import (
	"accounts-api/pkg/auth"
	"accounts-api/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router, serv user.Service) {
	router.Get("/users", FindAll(serv))
	router.Get("/users/:id", FindById(serv))
	router.Post("/users", Save(serv))
	router.Put("/users/:id", Update(serv))
	router.Delete("/users/:id", DeleteById(serv))
}

func AuthRouter(router fiber.Router, serv auth.Service) {
	router.Post("/auth/login", Login(serv))
	router.Post("/auth/signup", Signup(serv))
}
