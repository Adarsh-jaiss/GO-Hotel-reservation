package types

type User struct{
	ID string `json:"id,omitempty" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`

}