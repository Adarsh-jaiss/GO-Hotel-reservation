package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct{
	ID primitive.ObjectID 	`json:"id,omitempty" bson:"_id,omitempty"`
	Name string 			`josn:"name" bson:"name"`
	Location string			`json:"location" bson:"location"`
	Rooms []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating int 				`json:"rating" bson:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct{
	ID primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Type RoomType 		   `json:"type" bson:"type"`
	BasePrice float64 	   `json:"basePrice" bson:"basePrice"`
	Price float64 		   `json:"price" bson:"price"`
	HotelID primitive.ObjectID `json:"hotelID" bson:"hotelID"`

}