package user

import (
	"context"
	"github.com/anisbouzahar/portfolio-api/internal/app/database"
	"github.com/anisbouzahar/portfolio-api/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	SaveSubscriber(subscriber *models.Subscriber, ctx context.Context) (models.Subscriber, error)
	IsSubscribed(email string, ctx context.Context) (bool, error)
}

type RepositoryImpl struct {
	Db         *database.MongoDb
	collection *mongo.Collection
}

func (r RepositoryImpl) IsSubscribed(email string, ctx context.Context) (bool, error) {
	count, err := r.collection.CountDocuments(ctx, bson.M{"email": email})
	println(count)
	if err != nil {
		return false, err
	}
	return count > 0, err
}

func (r RepositoryImpl) SaveSubscriber(subscriber *models.Subscriber, ctx context.Context) (models.Subscriber, error) {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

	res, err := r.collection.InsertOne(ctx, subscriber)
	if err != nil {

	}
	id := res.InsertedID
	var entry models.Subscriber
	err = r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entry)

	return entry, err
}

func NewUserRepository(db *database.MongoDb) Repository {
	return RepositoryImpl{Db: db, collection: db.Client.Collection("subscribers")}
}
