package models

import "time"

type Booking struct {
	ID          int       `json:"id"`
	SeatID      int       `json:"seat_id"`
	BookingID   string    `json:"booking_id" gorm:"unique"`
	UserID      int       `json:"user_id"`
	DateCreated time.Time `json:"date_created"`
}
