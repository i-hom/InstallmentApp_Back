package src

import (
	"installment_back/models"
)

type Controller struct {
	UserModel        *models.User
	CardModel        *models.Card
	InstallmentModel *models.Installment
	DataBaseModel    *DataBase
}
