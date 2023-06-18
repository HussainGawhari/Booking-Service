package models

// Seat represents a single seat
type Seat struct {
	ID        int    `json:"id"`
	SeatID    int    `json:"seat_id"`
	SeatClass string `json:"seat_class"`
}

var seats []Seat
