package services

import (
	"installment_back/models"
	"installment_back/repositories"
)

type UserService struct {
	userRepository     *repositories.UserRepository
	installmentService *InstallmentService
	cardRepository     *repositories.CardRepository
}

func NewUserService(userRepository *repositories.UserRepository, installmentService *InstallmentService, cardRepository *repositories.CardRepository) *UserService {
	return &UserService{userRepository: userRepository, installmentService: installmentService, cardRepository: cardRepository}
}

func (us *UserService) Get(userAuth models.UserLog) models.User {
	var user models.User

	user, _ = us.userRepository.Get(userAuth)

	user.Installments, _ = us.installmentService.GetAll(user.ID)
	user.Cards, _ = us.cardRepository.GetAll(user.ID)
	return user
}
