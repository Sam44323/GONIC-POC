package services

import (
	"errors"

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
	querym := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, querym).Decode(&user)
	return &user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	// var users []*models.User
	// cursor, err := u.usercollection.Find(u.ctx, bson.D{})

	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount == 0 {
		return errors.New("no user found")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	return nil
}
