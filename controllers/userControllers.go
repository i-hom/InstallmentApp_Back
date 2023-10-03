package controllers

import (
	"encoding/json"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Get(params interface{}) models.RPCResponse {
	var userAuth models.UserLog
	if err := json.Unmarshal(models.GetRaw(params), &userAuth); err != nil {
		return errors.Invalid_parameter
	}

	if userAuth.PhoneNumber == "" || userAuth.Password == "" {
		return errors.Missing_parameter
	}

	user := uc.userService.Get(userAuth)

	if user.ID.IsZero() {
		return errors.User_not_found
	}

	return models.RPCResponse{Data: user}
}
