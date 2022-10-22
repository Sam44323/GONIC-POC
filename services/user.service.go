package services

import (
  "github.com/Sam44323/gin-POC/models"
  );

type UserService interface{
  CreateUser(*models.User) error
  GetUser(*string) (*models.User, error) // either the user or any error will be returned
  GetAll() ([]*models.User, error)
  UpdateUser(*models.User) error
  DeleteUser(*string) error
}
