package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAuthenticate(t *testing.T)  {
	tdb := Setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStorer)
	app.Post("/auth",authHandler.HandleAuthenticate)

	params := AuthParams{
		Email: "adarsh@gmail.com",
		Password: "supersecurepassword",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST","/auth",bytes.NewReader(b))
	req.Header.Add("content-type","application/json")
	res,err := app.Test(req) 
	if err!=nil{
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK{
		t.Fatalf("expected http status of 200 but got %d",res.StatusCode)
	}
	var resp AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err!= nil{
		t.Error(err)
	}
	fmt.Println(res)
}