package services

import (
	"github.com/Sam44323/gin-POC/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserServiceImpl(usercollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{usercollection, ctx} // creating a new struct
}

// implementation methods for the interface (automatically inferred)

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user models.User
	querym := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, querym).Decode(&user)
	return &user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	// var users []*models.User
	// cursor, err := u.usercollection.Find(u.ctx, bson.D{})

	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {

	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	return nil
}
