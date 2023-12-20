package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct{
	ID primitive.ObjectID 		`json:"id,omitempty" bson:"_id,omitempty"`
	Name string 				`josn:"name" bson:"name"`
	Location string				`json:"location" bson:"location"`
	Rooms []primitive.ObjectID  `json:"rooms" bson:"rooms"`
	Rating int 					`json:"rating" bson:"rating"`
}

type RoomType int

// const (
// 	_ RoomType = iota
// 	SingleRoomType
// 	DoubleRoomType
// 	SeaSideRoomType
// 	DeluxeRoomType
// )

type Room struct{
	ID primitive.ObjectID  			`json:"id,omitempty" bson:"_id,omitempty"`
	Size string 					`json:"size" bson:"size"`   				// size ,normal , kingszie
	Price float64 		  			`json:"price" bson:"price"`
	Seaside bool					`json:"seaside bson:"seaside"`
	HotelID primitive.ObjectID 		`json:"hotelID" bson:"hotelID"`
	// IsAvailable bool 				`json:"available" bson:"-"`

}