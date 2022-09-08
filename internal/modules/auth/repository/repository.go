package repository

import (
	"context"

	"github.com/quankori/go-clear-architecture/internal/models"
	"github.com/quankori/go-clear-architecture/internal/modules/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) auth.UserRepository {
	return &userRepository{
		db:         db,
		collection: collection,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := ur.db.Collection(ur.collection).InsertOne(context.TODO(), &user)

	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	projection := bson.M{"password": 0}
	err := ur.db.Collection(ur.collection).FindOne(
		context.TODO(),
		bson.M{"username": username},
		options.FindOne().SetProjection(projection)).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	var user models.User
	projection := bson.M{"password": 0}
	err := ur.db.Collection(ur.collection).FindOne(
		context.TODO(),
		bson.M{"_id": userId},
		options.FindOne().SetProjection(projection)).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
