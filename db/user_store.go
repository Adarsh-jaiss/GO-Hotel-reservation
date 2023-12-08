package db

import (
	"context"
	"fmt"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ( 
	Usercoll = "users"
)

type Dropper interface{
	Drop (context.Context) error

}

type UserStorer interface{
	Dropper
	GetUserByEmail(context.Context,string) (*types.User,error)
	GetUserByID(context.Context,string) (*types.User,error)
	GetUsers(context.Context) ([]*types.User,error)
	InsertUsers(context.Context,*types.User) (*types.User,error)
	DeleteUsers(context.Context,string) (error)
	UpdateUsers(ctx context.Context,filter bson.M , params types.UpdateUserParams) (error)
	
}

type MongoUserStore struct{
	client *mongo.Client
	coll *mongo.Collection
	
}
func NewMongoUserStore(c *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll: c.Database(DBNAME).Collection(Usercoll),
	}
	
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("----------------Dropping user collection -------------")
	return s.coll.Drop(ctx)
}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context,email string) (*types.User,error) {
	
	var user types.User
	if err := s.coll.FindOne(ctx,bson.M{"email":email}).Decode(&user); err!=nil{
		return nil,err
	}
	return &user,nil
	
}

func (s *MongoUserStore) GetUserByID(ctx context.Context,id string) (*types.User,error) {
	// validiate the correctness of the ID
	oid,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil,err
	}

	var user types.User
	if err := s.coll.FindOne(ctx,bson.M{"_id":oid}).Decode(&user); err!=nil{
		return nil,err
	}
	return &user,nil
	
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User,error) {
	cur,err := s.coll.Find(ctx,bson.M{})
	if err!= nil{
		return nil,err
	}
	var users []*types.User
	if err:= cur.All(ctx,&users); err!= nil{
		return []*types.User{},nil
	}
	return users,nil

}
func (s *MongoUserStore) InsertUsers(ctx context.Context,user *types.User) (*types.User,error) {
	res,err := s.coll.InsertOne(ctx,user)
	if err!= nil{
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user,err
}

func (s *MongoUserStore) DeleteUsers(ctx context.Context, id string) error {
	// validate the correctness of the ID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStore) UpdateUsers(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
    update := bson.D{
        {"$set", bson.D{
            {"firstName", params.FirstName},
            {"lastName", params.LastName},
        }},
    }

    _, err := s.coll.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    return nil
}
