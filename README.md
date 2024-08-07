# GO-Hotel-Reservation

GO-Hotel-Reservation is a Go-based web application for hotel reservations. It is built using the Fiber web framework and MongoDB as the database. This project provides a simple and scalable solution for managing hotel information, room bookings, user authentication, and more.

## Table of Contents

- [GO-Hotel-Reservation](#go-hotel-reservation)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Getting Started](#getting-started)
  - [Using GO-Hotel-Reservation with Docker](#using-go-hotel-reservation-with-docker)
    - [Prerequisites](#prerequisites-1)
    - [Clone the Repository and run these commands in your terminal](#clone-the-repository-and-run-these-commands-in-your-terminal)
  - [Configuration](#configuration)
  - [Endpoints](#endpoints)
  - [Testing](#testing)
  - [Author's Note](#authors-note)

## Features

- User authentication with JWT
- Hotel and room management
- Room booking and reservation handling
- Role-based access control (admin and user roles)
- API versioning with Fiber
- MongoDB integration for data storage

## Prerequisites

Before you begin, make sure you have the following prerequisites installed:

- [Go](https://golang.org/dl/)
- [MongoDB](https://www.mongodb.com/try/download/community)
- [Docker](https://www.docker.com/get-started)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/Adarsh-jaiss/GO-Hotel-reservation.git
    cd GO-Hotel-reservation
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up the environment variables:

    Create a `.env` file in the project root and add the following:

    ```env
    HTTP_LISTEN_ADDRESS=:3000
    JWT_SECRET=your_secret_key
    MONGO_DB_NAME=hotel-reservation
    MONGO_DB_URL=mongodb://your_mongodb_host:your_mongodb_port/
    ```
4. **Seed the database with some initial values, so that you can test the API's using those values** 

    ```
    make seed
    ```


5. Run the application:

    ```bash
    go run main.go
    ```

The application should now be accessible at `http://localhost:3000`.

## Using GO-Hotel-Reservation with Docker

This guide provides step-by-step instructions on how to build and run the GO-Hotel-Reservation project using Docker.

### Prerequisites

Make sure you have the following prerequisites installed on your system:

- [Docker](https://www.docker.com/get-started)
  
  You can install it via this simple command
  ```
  docker run --name mongob -d mongo:latest -p 27017:27017
  ```

  and run a docker container via 
  ```
  docker run --rm --name my_mongo_container -d mongo:latest
  ```

### Clone the Repository and run these commands in your terminal

Note : Don't forget to configure the enviroment variables before running the application,otherwise it will not start and throw an error

```bash
git clone https://github.com/YourUsername/GO-Hotel-reservation.git
cd GO-Hotel-reservation
go mod download
make seed
make docker or make run
```
If you have exited out of docker container, use this command to create another container out of the docker image
```
docker run -p 3000:3000 --rm --name hotel-reservation 4fde40842494

```

- Now you can test all of the API's using their routes :
  
0. Authentication

```bash

http://localhost:3000/api/auth/signup  -> POST (create a new user)

{
  "firstName": "rohan",
  "lastName": "nanda",
  "email": "rohan@jaiswal.com",
  "password": "12345678",
  "isAdmin": true
}


http://localhost:3000/api/auth/signin -> POST (sign in already existing user)

req body :

  {
  "email": "rohan@jaiswal.com",
  "password": "12345678"
  }


```

1. USER ROUTES :

Note : ISAdmin can only be modified in seed, we can not create a new user with admin value true, we can only seed admin.
   
```bash

http://localhost:3000/api/v1/user -> GET all Users

http://localhost:3000/api/v1/user/:id -> GET

http://localhost:3000/api/v1/user/:id -> DELETE


http://localhost:3000/api/v1/user/:id -> PUT (NOT updating correctly) -> can only update first and last name.

body :

{
  "firstName": "rohan_baba"
}

```

2. Hotels :

```bash
http://localhost:3000/api/v1/hotel  -> GET all hotels

http://localhost:3000/api/v1/hotel/66a91ae3401a46473b599859 -> GET hotels by ID

http://localhost:3000/api/v1/hotel/66a91ae3401a46473b599859/rooms -> GET all rooms of a hotel

```
3. Looking for rooms and booking it

```bash
http://localhost:3000/api/v1/room -> GET all the available rooms


http://localhost:3000/api/v1/room/66a91ae3401a46473b59985a/book -> POST (book rooms)(working only after authentication)

req body :
 { 
  "fromdate":"2024-12-12T00:00:00.0Z",
  "tilldate":"2024-12-25T00:00:00.0Z",
  "numPersons":4
 }
```

4. CHECKING AND CANCELLING BOOKINGS BY USER

```bash
http://localhost:3000/api/v1/booking/id -> GET bookings by a user

http://localhost:3000/api/v1/booking/id/cancel -> GET (Cancel booking)

```

5. ADMIN panel

```bash
http://localhost:3000/api/v1/admin/booking -> GET (all bookings) NOT WORKING 


```


## Configuration

You can configure the application by modifying the values in the `.env` file. Make sure to update the MongoDB connection details and set a secure JWT secret.

## Endpoints

The API endpoints are documented in the [API Documentation](API_DOCUMENTATION.md) file.

## Testing

Run tests using the following command:

```bash
make test
```


## Author's Note

**I haven't implemented few things such as Database migration, adding up another database layer (maybe postgress), Pagination, Custom Test Cases for all the handlers, and some other API functionalities like creating admin directly via post method but in future maybe I'll apply those as well.**

<h4>If you like this project, do give it a star and I am open for suggestions and feedbacks. You can create an issue if you find bugs or got some better way to implement stuff and we can discuss over it.</h4>