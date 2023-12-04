package db

import (
	"context"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ( 
	Usercoll = "users"
)

type UserStorer interface{
	GetUserByID(context.Context,string) (*types.User,error)
	GetUsers(context.Context) ([]*types.User,error)
	InsertUsers(context.Context,*types.User) (*types.User,error)
}

type MongoUserStore struct{
	client *mongo.Client
	coll *mongo.Collection
	
}

func (s *MongoUserStore) InsertUsers(ctx context.Context,user *types.User) (*types.User,error) {
	res,err := s.coll.InsertOne(ctx,user)
	if err!= nil{
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user,err
}

func NewMongoUserStore(c *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll: c.Database(DBNAME).Collection(Usercoll),
	}
	
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