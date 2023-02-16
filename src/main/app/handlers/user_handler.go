package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/src/main/app/model"
	"github.com/src/main/app/server"
	"github.com/src/main/app/server/errors"
	"github.com/src/main/app/services"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: *userService}
}

func (handler UserHandler) CreateUser(ctx *fiber.Ctx) error {
	request := new(model.CreateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		return errors.NewError(http.StatusBadRequest, "bad request error, missing key and value properties")
	}

	err := server.NotEmpty(request.Email, "email is required")
	if err != nil {
		return err
	}

	err = server.NotEmpty(request.Nickname, "nickname is required")
	if err != nil {
		return err
	}

	result, err := handler.userService.CreateUser(request)
	if err != nil {
		return err
	}

	return server.SendCreated(ctx, result)
}
