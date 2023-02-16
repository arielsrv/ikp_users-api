package services

import (
	"github.com/src/main/app/model"
	"github.com/src/main/app/producer"
)

type UserService struct {
	userProducer *producer.UserProducer
}

func NewUserService(userProducer *producer.UserProducer) *UserService {
	return &UserService{
		userProducer: userProducer,
	}
}

func (userService *UserService) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	response := new(model.CreateUserResponse)

	// userRepository result -> id
	userService.userProducer.Send(1)

	return response, nil
}
