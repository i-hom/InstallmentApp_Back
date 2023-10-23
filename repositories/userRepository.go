package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"installment_back/models"
	"installment_back/storage"
)

type UserRepository struct {
	db *storage.DataBase
}

func NewUserRepository(db *storage.DataBase) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(userAuth models.UserLog) (models.User, error) {
	var user models.BUser
	err := ur.db.FindOne("Users", bson.M{"phonenumber": userAuth.PhoneNumber, "password": userAuth.Password}, &user)
	return user.ToJUser(), err
}
