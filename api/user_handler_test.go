package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	
	"log"
	"net/http/httptest"
	"testing"

	"github.com/adarsh-jaiss/GO-Hotel-reservation/db"
	"github.com/adarsh-jaiss/GO-Hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Testdburi = "mongodb+srv://adarsh_jaiss:baburaokijai@baburao.dg1eflt.mongodb.net/"
	dbname = "Hotel-reservation-test"
)

type testdb struct{
	db.UserStorer
}

func (tdb *testdb) teardown(t *testing.T)  {
	err:= tdb.UserStorer.Drop(context.TODO()); 
	if err!=nil{
		return 
	}
	
}

func Setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Testdburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStorer: db.NewMongoUserStore(client),
	}
	
}

func TestPostUser(t *testing.T)  {
	tdb := Setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	UserHandler := NewUserhandler(tdb.UserStorer)
	app.Post("/",UserHandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "Testing",
		LastName: "kumar",
		Email: "testing@gmail.com",
		Password: "12345",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST","/",bytes.NewReader(b))
	req.Header.Add("content-type","application/json")
	res,err := app.Test(req) 
	if err!=nil{
		t.Error(err)
	}
	
	var user types.User
	json.NewDecoder(res.Body).Decode(&user)
	fmt.Printf("Received User: %+v\n", user)

	if len(user.ID) == 0 {
		t.Error("expecting a user id to be set in database")
	}

	if len(user.EncryptedPassword) > 0 {
		t.Error("expected the encrypted password not to be included in the json response")
	}

	if user.FirstName != params.FirstName{
		t.Errorf("expected firstname %s but got %s",params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName{
		t.Errorf("expected lastname %s but got %s",params.LastName, user.LastName)
	}

	if user.Email != params.Email{
		t.Errorf("expected Email %s but got %s",params.Email, user.Email)
	}
	

}