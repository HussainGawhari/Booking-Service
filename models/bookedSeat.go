package models

type BookedSeat struct {
	ID        int `json:"id"`
	BookingID int `json:"booking_id"`
	SeatID    int `json:"seat_id"`
}
