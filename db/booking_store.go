package db

import (
	"context"
	"errors"
	"os"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	   // Validate and convert ObjectID fields in the filter
	   for key, value := range filter {
        if str, ok := value.(string); ok {
            objID, err := primitive.ObjectIDFromHex(str)
            if err != nil {
                return nil, errors.New("the provided hex string is not a valid ObjectID, error found")
            }
            filter[key] = objID
        }
    }

	
	cur,err := s.coll.Find(ctx,filter)
	if err!=nil{
		return nil,err
	}
	defer cur.Close(ctx)

	var bookings []*types.Booking
	if err := cur.All(ctx,&bookings); err!= nil{
		return nil,err
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
