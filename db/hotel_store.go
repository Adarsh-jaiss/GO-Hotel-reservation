package db

import (
	"context"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface{
	InsertHotels(context.Context, *types.Hotel) (*types.Hotel,error)
	UpdateHotels(context.Context, bson.M,bson.M) error
}

type MongoHotelStore struct{
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll: client.Database(DBNAME).Collection("hotels"),
		
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel,error) {
	res,err:= s.coll.InsertOne(ctx,hotel)
	if err!= nil{
		return nil,err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, update bson.M) (error) {
	_,err := s.coll.UpdateOne(ctx,filter,update)
	if err!= nil{
		return err
	}
	return nil
}