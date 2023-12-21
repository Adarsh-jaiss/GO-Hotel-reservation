# Hotel reservation backend

## Project Enviroment variables
```
HTTP_LISTEN_ADDRESS:=3000
JWT_SECRET= somethingsupersecretTHATKNOWBODYKNOWS
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://172.17.0.2:27017/
```

## Project outline
- users -> Book room in a hotel
- admins -> Going to check bookings/reservations
- Authentation and Authorization -> JWT Tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> DB managemnet -> Seeding , migration

## Resources
### Mongo DB driver
Documentation

```
https://www.mongodb.com/docs/drivers/go/current/quick-start/#std-label-golang-quickstart
```

Installing MongoDB client

```
go get go.mongodb.org/mongo-driver/mongo
```

### GO Fiber
Documentation

```
https://gofiber.io
```

Installing Gofiber

```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongo DB as docker container

```
docker run --name mongob -d mongo:latest -p 27017:27017
```

running the docker container
```
docker run --rm --name my_mongo_container -d mongo:latest
```


# booking format
{
  "fromdate":"2023-12-22T00:00:00.0Z",
  "tilldate":"2023-12-25T00:00:00.0Z",
  "numPersons":4
  
}