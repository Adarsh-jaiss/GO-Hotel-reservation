package db

import (
	"context"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type BookingStore interface{
	InsertBooking(context.Context, *types.Booking) (*types.Booking,error)
	GetBookings(context.Context, bson.M) ([]*types.Booking,error)
	GetBookingsByID(context.Context, primitive.ObjectID) (*types.Booking,error)
	UpdateBooking(context.Context, string, bson.M) (error)
}

type MongoBookingStore struct{
	client      *mongo.Client
	coll        *mongo.Collection
	BookingStore
	
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	dbname := os.Getenv(MongoDbEnvName)
	return &MongoBookingStore{
		client: client,
		coll: client.Database(dbname).Collection("bookings"),
	}
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking,error) {
	res,err:= s.coll.InsertOne(ctx,booking)
	if err!= nil{
		return nil,err
	}
	booking.ID = res.InsertedID.(primitive.ObjectID)
	return booking, nil

}

func(s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M ) ([]*types.Booking,error) {
	cur,err := s.coll.Find(ctx,filter)
	if err!=nil{
		return nil,err
	}

	var bookings []*types.Booking
	if err := cur.All(ctx,&bookings); err!= nil{
		return []*types.Booking{},nil
	}

	return bookings,nil
}

func(s *MongoBookingStore) GetBookingsByID(ctx context.Context, id primitive.ObjectID) (*types.Booking,error) {
	var booking types.Booking
	if err := s.coll.FindOne(ctx, bson.M{"_id":id}).Decode(&booking); err!= nil{
		return nil,err
	}
	return &booking,nil
}

func (s *MongoBookingStore) UpdateBooking(ctx context.Context, id string, update bson.M) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	m := bson.M{"$set": update}
	_, err = s.coll.UpdateByID(ctx, oid, m)
	return err
}
