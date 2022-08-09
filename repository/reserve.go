package repository

import (
	"context"
	"go-project/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ReserveRepo struct {
		//logger        logrus.Logger
		db         *mongo.Client
		collection *mongo.Collection
	}

	ReserveRepoImpl interface {
		GetAllTable(ctx context.Context) ([]*model.Reserve, error)
		CreateNewBooking(ctx context.Context, reserve model.Reserve) error
		CancelBooking(ctx context.Context, reserve model.Reserve) error
		UpdateBooking(ctx context.Context, reserve model.Reserve) error
	}
)

func NewReserveRepo(db *mongo.Client) *ReserveRepo {
	collection := db.Database("Cluster0").Collection("reserves")
	return &ReserveRepo{
		collection: collection,
		db:         db,
	}
}

func (r *ReserveRepo) GetAllTable(ctx context.Context, date string) ([]*model.Reserve, error) {
	var reserves []*model.Reserve

	filter := bson.M{
		"datetime": date,
	}

	c, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for c.Next(ctx) {
		var reserve model.Reserve
		err = c.Decode(&reserve)
		if err != nil {
			logrus.Errorf("Find all got error: %v", err)
			continue
		}
		reserves = append(reserves, &reserve)
	}

	return reserves, nil
}

func (r *ReserveRepo) CreateNewBooking(ctx context.Context, reserve model.Reserve) error {
	_, err := r.collection.InsertOne(ctx, reserve)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReserveRepo) CancelBooking(ctx context.Context, phoneNumber string) error {
	filter := bson.M{
		"phone_number": phoneNumber,
	}
	result := r.collection.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *ReserveRepo) UpdateBooking(ctx context.Context, reserve model.Reserve) error {
	filter := bson.M{
		"phone_number": reserve.PhoneNumber,
	}

	data := bson.M{
		"datetime":     reserve.DateTime,
		"total_people": reserve.TotalPeople,
		"table_id":     reserve.TableId,
	}

	update := bson.M{
		"$set": data,
	}

	result := r.collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
