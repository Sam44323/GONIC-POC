package services

import (
	"github.com/Sam44323/gin-POC/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
);

type UserServiceImpl struct{
  usercollection *mongo.Collection
  ctx context.Context
}

func NewUserServiceImpl(usercollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
  return &UserServiceImpl{usercollection, ctx}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error{
  return nil
}

func (u *UserServiceImpl) GetUser(id *string) (*models.User, error){
  return nil, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error){
  return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error{
  return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error{
  return nil
}
