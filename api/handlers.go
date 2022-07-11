package api

import (
	"accounts-api/pkg/auth"
	"accounts-api/pkg/user"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Save(serv user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(SignupRequest)
		if err := ctx.BodyParser(request); err != nil {
			// TODO improve errors handlers
			return ctx.Status(fiber.StatusBadRequest).BodyParser(request)
		}
		entity := toUser(request)
		saved, err := serv.Create(entity)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		response := toResponse(saved)

		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func Update(serv user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := ctx.ParamsInt("id", -1)

		if id == -1 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": "Path param 'id' is required: /users/:id",
			})
		}

		request := new(SignupRequest)

		if err := ctx.BodyParser(request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		entity := toUser(request)
		entity.Id = id
		updated, err := serv.Update(entity)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(toResponse(updated))
	}
}

func FindAll(serv user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		list := serv.FindAll()

		var responseList []UserResponse
		for _, v := range *list {
			item := toResponse(&v)
			responseList = append(responseList, *item)
		}
		return ctx.Status(fiber.StatusOK).JSON(responseList)
	}
}

func FindById(serv user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := ctx.ParamsInt("id", -1)
		if id == -1 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": "Path param 'id' is required: /users/:id",
			})
		}

		foundUser, err := serv.FindById(id)
		if err != nil {
			msg := "User not found with id: " + strconv.Itoa(id)
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": msg,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(toResponse(foundUser))
	}
}

func DeleteById(serv user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := ctx.ParamsInt("id", -1)

		if id == -1 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": "Path param 'id' is required: /users/:id",
			})
		}

		if err := serv.DeleteById(id); err != nil {
			msg := "User not exists with id: " + strconv.Itoa(id)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": msg,
			})
		}
		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

func Login(serv auth.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(LoginRequest)

		if errParse := ctx.BodyParser(request); errParse != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"msg": "Error parse body",
				})
		}

		login, err := serv.Login(request.Email, request.Password)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"msg": "Incorrect Password or Email",
				})
		}

		response := toResponse(login)
		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}

func Signup(serv auth.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(SignupRequest)
		if err := ctx.BodyParser(request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).BodyParser(request)
		}
		entity := toUser(request)
		saved, err := serv.Signup(entity)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		response := toResponse(saved)

		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}
