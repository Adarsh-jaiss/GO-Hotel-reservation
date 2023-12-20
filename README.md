# Hotel reservation backend

# Project outline
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
