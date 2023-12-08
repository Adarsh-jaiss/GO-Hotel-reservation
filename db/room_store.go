package db

import (
	"context"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
	// UpdateHotels(context.Context, bson.M, bson.M) error
}

type MongoRoomStore struct {
	client      *mongo.Client
	coll        *mongo.Collection
	HotelStore  *MongoHotelStore // Embed an instance of MongoHotelStore
}

func NewMongoRoomStore(client *mongo.Client, dbname string, hotelStore *MongoHotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:      client,
		coll:        client.Database(dbname).Collection("rooms"),
		HotelStore:  hotelStore,
	}
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
    // Exclude rooms with the placeholder hotel ID
    filter["hotelid"] = bson.M{"$ne": "000000000000000000000000"}

    res, err := s.coll.Find(ctx, filter)
    if err != nil {
        return nil, err
    }

    var rooms []*types.Room
    if err := res.All(ctx, &rooms); err != nil {
        return nil, err
    }
    return rooms, nil
}



func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)

	// Update the hotels with room ID
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}
	if err := s.HotelStore.UpdateHotels(ctx, filter, update); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *MongoRoomStore) UpdateHotels(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
