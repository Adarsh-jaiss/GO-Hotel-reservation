package db

import (
	"context"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type HotelStore interface{
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel,error)
	UpdateHotels(context.Context, bson.M,bson.M) error
	GetHotels(context.Context, bson.M) ([]*types.Hotel,error)
	GetHotelByID(context.Context, primitive.ObjectID) (*types.Hotel,error)
}

type MongoHotelStore struct{
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	dbname := os.Getenv(MongoDbEnvName)
	return &MongoHotelStore{
		client: client,
		coll: client.Database(dbname).Collection("hotels"),
		
	}
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel,error) {
	res,err := s.coll.Find(ctx,filter)
	if err!= nil{
		return nil,err
	}

	var Hotels []*types.Hotel
	if err := res.All(ctx,&Hotels); err!= nil{
		return nil,err
	}

	return Hotels,nil
}

func (s *MongoHotelStore) GetHotelByID(ctx context.Context, id primitive.ObjectID) (*types.Hotel,error) {
	var hotel types.Hotel
	if err := s.coll.FindOne(ctx,bson.M{"_id":id}).Decode(&hotel); err!= nil{
		return nil,err
	}


	return &hotel,nil
}






func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel,error) {
	res,err:= s.coll.InsertOne(ctx,hotel)
	if err!= nil{
		return nil,err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotels(ctx context.Context, filter bson.M, update bson.M) (error) {
	_,err := s.coll.UpdateOne(ctx,filter,update)
	if err!= nil{
		return err
	}
	return nil
}