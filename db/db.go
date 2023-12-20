package db

const (
	DBNAME = "Hotel-reservation"
	TestDBNAME = "Hotel-reservation-test"
	DBURI = "mongodb://172.17.0.2:27017/"

)

type Store struct{
	User UserStorer
	Hotel HotelStore
	Room RoomStore
	Booking BookingStore
}


