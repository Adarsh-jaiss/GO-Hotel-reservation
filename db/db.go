package db

const (
	DBNAME = "Hotel-reservation"
	TestDBNAME = "Hotel-reservation-test"
	DBURI = "mongodb+srv://adarsh_jaiss:baburaokijai@baburao.dg1eflt.mongodb.net/"

)

type Store struct{
	User UserStorer
	Hotel HotelStore
	Room RoomStore
	Booking BookingStore
}


