# GO-Hotel-Reservation API Documentation

## Authentication

### Authenticate User

**Endpoint:** `/api/auth` (POST)

- **Description:** Authenticate a user and receive a JWT token.
- **Request:**
  - Body:
    - `email` (string): User's email.
    - `password` (string): User's password.
- **Response:**
  - Success (200 OK):
    - `token` (string): JWT token for authentication.
    - Manually add the token into the headers with the field name **X-Api-Token**
    - Also don't forget to include content-type in the headers as **application/json**

## User Endpoints

### Get All Users

**Endpoint:** `/api/v1/user` (GET)

- **Description:** Retrieve a list of all users.
- **Request:** None.
- **Response:**
  - Success (200 OK):
    - List of users.

### Get User by ID

**Endpoint:** `/api/v1/user/:id` (GET)

- **Description:** Retrieve details of a specific user.
- **Request:**
  - Path parameter:
    - `id` (string): User ID.
- **Response:**
  - Success (200 OK):
    - User details.

### Create User

**Endpoint:** `/api/v1/user` (POST)

- **Description:** Create a new user.
- **Request:**
  - Body:
    - `firstname` (string): User's username.
    - `lastname` (string): User's lastname.
    - `email` (string): User's email.
    - `password` (string): User's password.
  
- **Response:**
  - Success (201 Created):
  - User created successfully.

Here's an example :
```
{
  "firstname":"rohan",
  "lastname":"jaiswal",
  "email":"rohan@gmail.com",
  "password":"supersecurepassword"
  
}
```

### Delete User

**Endpoint:** `/api/v1/user/:id` (DELETE)

- **Description:** Delete a user.
- **Request:**
  - Path parameter:
    - `id` (string): User ID.
- **Response:**
  - Success (204 No Content):
    - User deleted successfully.

### Update User

**Endpoint:** `/api/v1/user/:id` (PUT)

- **Description:** Update user details.
- **Request:**
  - Path parameter:
    - `id` (string): User ID.
  - Body: Updated User details
    - `firstname` (string): User's username.
    - `lastname` (string): User's lastname.
    - `email` (string): User's email.
    - `password` (string): User's password.
    
    Note : Only firstName and lastName can be updated
- **Response:**
  - Success (200 OK):
    - User details updated successfully.

## Hotel Endpoints

### Get All Hotels

**Endpoint:** `/api/v1/hotel` (GET)

- **Description:** Retrieve a list of all hotels.
- **Request:** None.
- **Response:**
  - Success (200 OK):
    - List of hotels.

### Get Rooms for a Hotel

**Endpoint:** `/api/v1/hotel/:id/rooms` (GET)

- **Description:** Retrieve rooms for a specific hotel.
- **Request:**
  - Path parameter:
    - `id` (string): Hotel ID.
- **Response:**
  - Success (200 OK):
    - List of rooms for the hotel.

### Get Hotel by ID

**Endpoint:** `/api/v1/hotel/:id` (GET)

- **Description:** Retrieve details of a specific hotel.
- **Request:**
  - Path parameter:
    - `id` (string): Hotel ID.
- **Response:**
  - Success (200 OK):
    - Hotel details.

## Room Endpoints

### Get All Rooms

**Endpoint:** `/api/v1/room` (GET)

- **Description:** Retrieve a list of all rooms.
- **Request:** 
    - 
- **Response:**
  - Success (200 OK):
    - List of rooms.

### Book a Room

**Endpoint:** `/api/v1/room/:id/book` (POST)

- **Description:** Book a room.
- **Request:**
  - Path parameter:
    - `id` (string): Room ID.

  - Body:
    - `fromdate` 
    - `tilldate`
    - `numPersons`
- **Response:**
  - Success (201 Created):
    - Room booked successfully.
  
  Here's an example :
  ```
  {
  "fromdate":"2023-12-22T00:00:00.0Z",
  "tilldate":"2023-12-25T00:00:00.0Z",
  "numPersons":4
  }
  ```

## Booking Endpoints

### Get Booking by ID

**Endpoint:** `/api/v1/booking/:id` (GET)

- **Description:** Retrieve details of a specific booking.
- **Request:**
  - Path parameter:
    - `id` (string): Booking ID.
- **Response:**
  - Success (200 OK):
    - Booking details.

### Cancel Booking

**Endpoint:** `/api/v1/booking/:id/cancel` (GET)

- **Description:** Cancel a booking.
- **Request:**
  - Path parameter:
    - `id` (string): Booking ID.
- **Response:**
  - Success (200 OK):
    - Booking canceled successfully.

## Admin Endpoints

### Get All Bookings (Admin Access Required)

**Endpoint:** `/api/v1/admin/booking` (GET)

- **Description:** Retrieve a list of all bookings (Admin access required).
- **Request:** None.
- **Response:**
  - Success (200 OK):
  - List of bookings.

Note : Admins can only be added through seed. I haven't implemented any direct way to add admin, because the project is itself quite big. But maybe I can add that in future.

## Error Handling

The API handles errors gracefully and returns appropriate HTTP status codes with error details in the response body.
